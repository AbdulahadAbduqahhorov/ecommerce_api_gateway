package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/api/models"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/order_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

// Create Order godoc
// @ID          create_order
// @Summary     Create Order
// @Description Create Order
// @Tags        order
// @Accept      json
// @Produce     json
// @Param       order         body     models.CreateOrderModel            true  "order"
// @Param       Authorization header   string                             false "Authorization"
// @Success     200           {object} models.ResponseModel{data=string}  "desc"
// @Response    400           {object} models.ResponseModel{error=string} "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string} "Server Error"
// @Router      /order [POST]
func (h *handler) CreateOrder(c *gin.Context) {
	var order models.CreateOrderModel
	userIdRaw, ok := c.Get("auth_user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, "Something is worng")
		return
	}
	userId, ok := userIdRaw.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, "Something is worng")
		return
	}
	err := c.BindJSON(&order)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	ordI := make([]*order_service.OrderItem, 0)
	for _, v := range order.OrderItems {
		ordI = append(ordI, &order_service.OrderItem{
			ProductId: v.ProductId,
			Quantity:  int32(v.Quantity),
		})
	}
	resp, err := h.services.OrderService().CreateOrder(
		context.Background(),
		&order_service.CreateOrderRequest{
			UserId:          userId,
			CustomerName:    order.CustomerName,
			CustomerAddress: order.CustomerAddress,
			CustomerPhone:   order.CustomerPhone,
			Orderitems:      ordI,
		},
	)

	if !handleError(h.log, c, err, "error while creating order") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp.Id)
}

// Get All Order godoc
// @ID          get all order
// @Summary     Get All Order
// @Description Get All Order
// @Tags        order
// @Accept      json
// @Produce     json
// @Param       offset        query    string                                             false "offset"
// @Param       limit         query    string                                             false "limit"
// @Param       search        query    string                                             false "search"
// @Param       Authorization header   string                                             false "Authorization"
// @Success     200           {object} models.ResponseModel{data=models.GetAllOrderModel} "desc"
// @Response    400           {object} models.ResponseModel{error=string}                 "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string}                 "Server Error"
// @Router      /order [GET]
func (h *handler) GetAllOrder(c *gin.Context) {
	userTypeRaw, ok := c.Get("auth_user_type")
	if !ok {
		c.JSON(http.StatusUnauthorized, "Something is worng")
		return
	}
	userType, ok := userTypeRaw.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, "Something is worng")
		return
	}
	var userId string
	if userType == "USER" {
		userIdRaw, ok := c.Get("auth_user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, "Something is worng")
			return
		}
		userId, ok = userIdRaw.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, "Something is worng")
			return
		}
	}

	offset, err := h.ParseQueryParam(c, "offset", h.cfg.DefaultOffset)
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", h.cfg.DefaultLimit)
	if err != nil {
		return
	}

	resp, err := h.services.OrderService().GetOrderList(
		context.Background(),
		&order_service.GetOrderListRequest{
			Offset: int32(offset),
			Limit:  int32(limit),
			Search: c.Query("search"),
			UserId: userId,
		},
	)

	if !handleError(h.log, c, err, "error while getting all orders") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// GetOrderById godoc
// @ID          get_order
// @Router      /order/{order_id} [GET]
// @Summary     Get Order
// @Description Get Order
// @Tags        order
// @Accept      json
// @Produce     json
// @Param       order_id      path     string                                             true  "order_id"
// @Param       Authorization header   string                                             false "Authorization"
// @Success     200           {object} models.ResponseModel{data=models.GetAllOrderModel} "desc"
// @Response    400           {object} models.ResponseModel{error=string}                 "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string}                 "Server Error"
func (h *handler) GetOrderById(c *gin.Context) {
	order_id := c.Param("order_id")

	if !util.IsValidUUID(order_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid order id", errors.New("order id is not valid"))
		return
	}
	userTypeRaw, ok := c.Get("auth_user_type")
	if !ok {
		c.JSON(http.StatusUnauthorized, "Something is worng")
		return
	}
	userType, ok := userTypeRaw.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, "Something is worng")
		return
	}
	var userId string
	if userType == "USER" {
		userIdRaw, ok := c.Get("auth_user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, "Something is worng")
			return
		}
		userId, ok = userIdRaw.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, "Something is worng")
			return
		}
	}

	resp, err := h.services.OrderService().GetOrderById(
		context.Background(),
		&order_service.GetOrderByIdRequest{
			Id: order_id,
			UserId: userId,
		},
	)

	if !handleError(h.log, c, err, "error while getting order") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
