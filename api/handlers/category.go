package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/api/models"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/category_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

// Create Category godoc
// @ID          create_category
// @Summary     Create Category
// @Description Create Category
// @Tags        category
// @Accept      json
// @Produce     json
// @Param       category      body     models.CreateCategoryModel         true  "category"
// @Param       Authorization header   string                             false "Authorization"
// @Success     200           {object} models.ResponseModel{data=string}  "desc"
// @Response    400           {object} models.ResponseModel{error=string} "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string} "Server Error"
// @Router      /category [POST]
func (h *handler) CreateCategory(c *gin.Context) {
	var category models.CreateCategoryModel

	err := c.BindJSON(&category)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CategoryService().CreateCategory(
		context.Background(),
		&category_service.CreateCategoryRequest{
			Title:      category.Title,
		},
	)

	if !handleError(h.log, c, err, "error while creating category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp.Id)
}

// Get All Category godoc
// @ID          get all category
// @Summary     Get All Category
// @Description Get All Category
// @Tags        category
// @Accept      json
// @Produce     json
// @Param       offset        query    string                                                false "offset"
// @Param       limit         query    string                                                false "limit"
// @Param       search        query    string                                                false "search"
// @Param       Authorization header   string                                                false "Authorization"
// @Success     200           {object} models.ResponseModel{data=models.GetAllCategoryModel} "desc"
// @Response    400           {object} models.ResponseModel{error=string}                    "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string}                    "Server Error"
// @Router      /category [GET]
func (h *handler) GetAllCategory(c *gin.Context) {
	// var categorys models.GetAllCategoryModel
	offset, err := h.ParseQueryParam(c, "offset", h.cfg.DefaultOffset)
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", h.cfg.DefaultLimit)
	if err != nil {
		return
	}

	resp, err := h.services.CategoryService().GetCategoryList(
		context.Background(),
		&category_service.GetCategoryListRequest{
			Offset:     int32(offset),
			Limit:      int32(limit),
			Search:     c.Query("search"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all categories") {
		return
	}
	// err = ParseToStruct(&categorys, resp)

	// if !handleError(h.log, c, err, "error while parse to struct") {
	// 	return
	// }

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// GetCategoryById godoc
// @ID          get_category
// @Router      /category/{category_id} [GET]
// @Summary     Get Category
// @Description Get Category
// @Tags        category
// @Accept      json
// @Produce     json
// @Param       category_id   path     string                                                true  "category_id"
// @Param       Authorization header   string                                                false "Authorization"
// @Success     200           {object} models.ResponseModel{data=models.GetAllCategoryModel} "desc"
// @Response    400           {object} models.ResponseModel{error=string}                    "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string}                    "Server Error"
func (h *handler) GetCategoryById(c *gin.Context) {
	// var category models.GetCategoryByIdModel
	category_id := c.Param("category_id")

	if !util.IsValidUUID(category_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid category id", errors.New("category id is not valid"))
		return
	}

	resp, err := h.services.CategoryService().GetCategoryById(
		context.Background(),
		&category_service.GetCategoryByIdRequest{
			Id: category_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting category") {
		return
	}

	// err = ParseToStruct(&category, resp)
	// if !handleError(h.log, c, err, "error while parsing to struct") {
	// 	return
	// }

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Update Category godoc
// @ID          update_category
// @Router      /category/{category_id} [PUT]
// @Summary     Update Category
// @Description Update Category by ID
// @Tags        category
// @Accept      json
// @Produce     json
// @Param       category_id   path     string                                     true  "category_id"
// @Param       category      body     models.UpdateCategoryModel                 true  "category"
// @Param       Authorization header   string                                     false "Authorization"
// @Success     200           {object} models.ResponseModel{data=models.Category} "desc"
// @Response    400           {object} models.ResponseModel{error=string}         "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string}         "Server Error"
func (h *handler) UpdateCategory(c *gin.Context) {
	var category models.UpdateCategoryModel
	// var response models.Category
	category_id := c.Param("category_id")

	if !util.IsValidUUID(category_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid category id", errors.New("category id is not valid"))
		return
	}

	err := c.BindJSON(&category)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	res, err := h.services.CategoryService().UpdateCategory(
		context.Background(),
		&category_service.UpdateCategoryRequest{
			Id:   category_id,
			Title: category.Title,
		},
	)

	if !handleError(h.log, c, err, "error while updated category") {
		return
	}

	// err = ParseToStruct(&response, res)
	// if !handleError(h.log, c, err, "error while parsing to struct") {
	// 	return
	// }

	h.handleSuccessResponse(c, http.StatusOK, "ok", res)
}


// Delete Category godoc
// @ID          delete_category
// @Router      /category/{category_id} [DELETE]
// @Summary     Delete Category
// @Description Delete Category by given ID
// @Tags        category
// @Accept      json
// @Produce     json
// @Param       category_id   path     string                             true  "category_id"
// @Param       Authorization header   string                             false "Authorization"
// @Success     200           {object} models.ResponseModel{data=string}  "desc"
// @Response    400           {object} models.ResponseModel{error=string} "Bad Request"
// @Failure     500           {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeleteCategory(c *gin.Context) {
	category_id := c.Param("category_id")

	if !util.IsValidUUID(category_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid category id", errors.New("category id is not valid"))
		return
	}

	_, err := h.services.CategoryService().DeleteCategory(
		context.Background(),
		&category_service.DeleteCategoryRequest{
			Id: category_id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Deleted")
}
