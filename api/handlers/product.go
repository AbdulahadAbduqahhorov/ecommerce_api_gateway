package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/api/models"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/genproto/product_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_api_gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

// Create Product godoc
// @ID create_product
// @Summary Create Product
// @Description Create Product
// @Tags product
// @Accept json
// @Produce json
// @Param product body models.CreateProductModel true "product"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /product [POST]
func (h *handler) CreateProduct(c *gin.Context) {
	var product models.CreateProductModel

	err := c.BindJSON(&product)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.ProductService().CreateProduct(
		context.Background(),
		&product_service.CreateProductRequest{
			Title:      product.Title,
			Desc:       product.Description,
			Quantity:   product.Quantity,
			Price:      product.Price,
			CategoryId: product.Category_id,
		},
	)

	if !handleError(h.log, c, err, "error while creating product") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp.Id)
}

// Get All Product godoc
// @ID get all product
// @Summary Get All Product
// @Description Get All Product
// @Tags product
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Param category_id query string false "category_id"
// @Success 200 {object} models.ResponseModel{data=models.GetAllProductModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
// @Router /product [GET]
func (h *handler) GetAllProduct(c *gin.Context) {
	// var products models.GetAllProductModel
	offset, err := h.ParseQueryParam(c, "offset", h.cfg.DefaultOffset)
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", h.cfg.DefaultLimit)
	if err != nil {
		return
	}

	resp, err := h.services.ProductService().GetProductList(
		context.Background(),
		&product_service.GetProductListRequest{
			Offset:     int32(offset),
			Limit:      int32(limit),
			Search:     c.Query("search"),
			CategoryId: c.Query("category_id"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all products") {
		return
	}
	// err = ParseToStruct(&products, resp)

	// if !handleError(h.log, c, err, "error while parse to struct") {
	// 	return
	// }

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// GetProductById godoc
// @ID get_product
// @Router /product/{product_id} [GET]
// @Summary Get Product
// @Description Get Product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.ResponseModel{data=models.GetAllProductModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetProductById(c *gin.Context) {
	// var product models.GetProductByIdModel
	product_id := c.Param("product_id")

	if !util.IsValidUUID(product_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid product id", errors.New("product id is not valid"))
		return
	}

	resp, err := h.services.ProductService().GetProductById(
		context.Background(),
		&product_service.GetProductByIdRequest{
			Id: product_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting product") {
		return
	}

	// err = ParseToStruct(&product, resp)
	// if !handleError(h.log, c, err, "error while parsing to struct") {
	// 	return
	// }

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Update Product godoc
// @ID update_product
// @Router /product/{product_id} [PUT]
// @Summary Update Product
// @Description Update Product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Param product body models.UpdateProductModel true "product"
// @Success 200 {object} models.ResponseModel{data=models.Product} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) UpdateProduct(c *gin.Context) {
	var product models.UpdateProductModel
	// var response models.Product
	product_id := c.Param("product_id")

	if !util.IsValidUUID(product_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid product id", errors.New("product id is not valid"))
		return
	}

	err := c.BindJSON(&product)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	res, err := h.services.ProductService().UpdateProduct(
		context.Background(),
		&product_service.UpdateProductRequest{
			Id:   product_id,
			Title: product.Title,
			Desc: product.Description,
			Quantity: int32(product.Quantity),
			Price: int32(product.Price),
		},
	)

	if !handleError(h.log, c, err, "error while updated product") {
		return
	}

	// err = ParseToStruct(&response, res)
	// if !handleError(h.log, c, err, "error while parsing to struct") {
	// 	return
	// }

	h.handleSuccessResponse(c, http.StatusOK, "ok", res)
}


// Delete Product godoc
// @ID delete_product
// @Router /product/{product_id} [DELETE]
// @Summary Delete Product
// @Description Delete Product by given ID
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeleteProduct(c *gin.Context) {
	product_id := c.Param("product_id")

	if !util.IsValidUUID(product_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid product id", errors.New("product id is not valid"))
		return
	}

	_, err := h.services.ProductService().DeleteProduct(
		context.Background(),
		&product_service.DeleteProductRequest{
			Id: product_id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting product") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Deleted")
}
