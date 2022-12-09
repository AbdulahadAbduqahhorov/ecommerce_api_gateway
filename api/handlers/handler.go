package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/api/models"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/clients"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"
)

var (
	ErrAlreadyExists       = "ALREADY_EXISTS"
	ErrNotFound            = "NOT_FOUND"
	ErrInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrServiceUnavailable  = "SERVICE_UNAVAILABLE"
	SigningKey             = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
	SuperAdminUserType     = "superadmin"
	SystemUserType         = "admin"
)

type handler struct {
	log      logger.Logger
	cfg      config.Config
	services clients.ServiceManagerI
}

type HandlerOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services clients.ServiceManagerI
}

func New(options *HandlerOptions) *handler {
	return &handler{
		log:      options.Log,
		cfg:      options.Cfg,
		services: options.Services,
	}
}

func handleError(log logger.Logger, c *gin.Context, err error, message string) (hasError bool) {
	st, ok := status.FromError(err)
	if st.Code() == codes.Canceled {
		log.Error(message+", canceled ", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   st.Message(),
		})
		return
	} else if st.Code() == codes.AlreadyExists || st.Code() == codes.InvalidArgument {
		log.Error(message+", already exists", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   ErrAlreadyExists,
		})
		return
	} else if st.Code() == codes.NotFound {
		log.Error(message+", not found", logger.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   ErrNotFound,
		})
		return
	} else if st.Code() == codes.Unavailable {
		log.Error(message+", service unavailable", logger.Error(err))
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"error":   ErrServiceUnavailable,
		})
		return
	} else if !ok || st.Code() == codes.Internal || st.Code() == codes.Unknown || err != nil {
		log.Error(message+", internal server error", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   ErrInternalServerError,
		})
		return
	}
	return true
}

func (h *handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, models.ResponseModel{
		Code:    code,
		Message: message,
		Error:   err,
	})
}

func (h *handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, models.ResponseModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (h *handler) ParseQueryParam(c *gin.Context, key string, defaultValue string) (int, error) {
	valueStr := c.DefaultQuery(key, defaultValue)

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		h.log.Error("error while parsing query param"+", canceled ", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return 0, err
	}

	return value, nil
}

func (h *handler) BadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func ParseToStruct(data interface{}, m protoiface.MessageV1) error {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	js, err := jspbMarshal.MarshalToString(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(js), data)
	return err
}
