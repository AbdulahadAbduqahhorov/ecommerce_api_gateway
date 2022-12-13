package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/auth_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID          create_user
// @Summary     Create User
// @Description Create User
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user body     auth_service.CreateUserRequest               true "user"
// @Success     200  {object} models.ResponseModel{data=auth_service.User} "desc"
// @Response    400  {object} models.ResponseModel{error=string}           "Bad Request"
// @Failure     500  {object} models.ResponseModel{error=string}           "Server Error"
// @Router      /user [POST]
func (h *handler) CreateUser(c *gin.Context) {
	var user auth_service.CreateUserRequest

	err := c.BindJSON(&user)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	res, err := h.services.AuthService().CreateUser(
		context.Background(),
		&user,
	)

	if !handleError(h.log, c, err, "error while creating user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", res)
}

// Get All User godoc
// @ID          get all user
// @Summary     Get All User
// @Description Get All User
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       offset   query    string                                                      false "offset"
// @Param       limit    query    string                                                      false "limit"
// @Param       search   query    string                                                      false "search"
// @Param       userType query    string                                                      false "sort by user type" Enums(ADMIN,USER)
// @Success     200      {object} models.ResponseModel{data=auth_service.GetUserListResponse} "desc"
// @Response    400      {object} models.ResponseModel{error=string}                          "Bad Request"
// @Failure     500      {object} models.ResponseModel{error=string}                          "Server Error"
// @Router      /user [GET]
func (h *handler) GetAllUser(c *gin.Context) {
	
	offset, err := h.ParseQueryParam(c, "offset", h.cfg.DefaultOffset)
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", h.cfg.DefaultLimit)
	if err != nil {
		return
	}

	resp, err := h.services.AuthService().GetUserList(
		context.Background(),
		&auth_service.GetUserListRequest{
			Offset:     int32(offset),
			Limit:      int32(limit),
			Search:     c.Query("search"),
			Type: c.Query("userType"),
		},
	)
	if !handleError(h.log, c, err, "error while getting all users") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// GetUserById godoc
// @ID          get_user
// @Router      /user/{user_id} [GET]
// @Summary     Get User
// @Description Get User
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user_id path     string                                       true "user_id"
// @Success     200     {object} models.ResponseModel{data=auth_service.User} "desc"
// @Response    400     {object} models.ResponseModel{error=string}           "Bad Request"
// @Failure     500     {object} models.ResponseModel{error=string}           "Server Error"
func (h *handler) GetUserById(c *gin.Context) {
	user_id := c.Param("user_id")
	
	if !util.IsValidUUID(user_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid user id", errors.New("user id is not valid"))
		return
	}

	resp, err := h.services.AuthService().GetUserById(
		context.Background(),
		&auth_service.GetUserByIdRequest{
			Id: user_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Update User godoc
// @ID          update_user
// @Router      /user [PUT]
// @Summary     Update User
// @Description Update User by ID
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user body     auth_service.UpdateUserRequest               true "UpdateUserRequestBody"
// @Success     200  {object} models.ResponseModel{data=auth_service.User} "User data"
// @Response    400  {object} models.ResponseModel{error=string}           "Bad Request"
// @Failure     500  {object} models.ResponseModel{error=string}           "Server Error"
func (h *handler) UpdateUser(c *gin.Context) {
	var user auth_service.UpdateUserRequest
	err := c.BindJSON(&user)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	res, err := h.services.AuthService().UpdateUser(
		context.Background(),
		&user,
	)

	if !handleError(h.log, c, err, "error while updated user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", res)
}

// Delete User godoc
// @ID          delete_user
// @Router      /user/{user_id} [DELETE]
// @Summary     Delete User
// @Description Delete User by given ID
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user_id path string true "user_id"
// @Success     204
// @Response    400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure     500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeleteUser(c *gin.Context) {
	user_id := c.Param("user_id")

	if !util.IsValidUUID(user_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid user id", errors.New("user id is not valid"))
		return
	}

	resp, err := h.services.AuthService().DeleteUser(
		context.Background(),
		&auth_service.DeleteUserRequest{
			Id: user_id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok",resp )
}

// Register User godoc
// @ID          register_user
// @Summary     Register User
// @Description Register User
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user body     auth_service.RegisterUserRequest             true "user"
// @Success     200  {object} models.ResponseModel{data=auth_service.User} "desc"
// @Response    400  {object} models.ResponseModel{error=string}           "Bad Request"
// @Failure     500  {object} models.ResponseModel{error=string}           "Server Error"
// @Router      /register [POST]
func (h *handler) Register(c *gin.Context) {
	var user auth_service.RegisterUserRequest

	err := c.BindJSON(&user)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	res, err := h.services.AuthService().Register(
		context.Background(),
		&user,
	)

	if !handleError(h.log, c, err, "error while register user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", res)
}

// Login User godoc
// @ID          login_user
// @Summary     Login User
// @Description Login User
// @Tags        user
// @Accept      json
// @Produce     json
// @Param       user body     auth_service.LoginRequest                             true "user"
// @Success     200  {object} models.ResponseModel{data=auth_service.TokenResponse} "desc"
// @Response    400  {object} models.ResponseModel{error=string}                    "Bad Request"
// @Failure     500  {object} models.ResponseModel{error=string}                    "Server Error"
// @Router      /login [POST]
func (h *handler) Login(c *gin.Context) {
	var login auth_service.LoginRequest

	err := c.BindJSON(&login)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	res, err := h.services.AuthService().Login(
		context.Background(),
		&login,
	)

	if !handleError(h.log, c, err, "error while login user") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", res)
}
