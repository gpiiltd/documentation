package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPublicProducts retrieves the list of all public products in barafiri
// @Summary Endpoint to get all product information on BARAFIRI
// @Description Get an array of all business and business data
// @Tags Public API
// @Accept json
// @Produce json
// @Success 200 {object} []ProductResponseObject "Success data. Check response body for data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /public/product/all [get]
func (c *Controller) GetAllPublicProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetProductByProductID retrieves the list of product by puid
// @Summary Endpoint to get all product information on BARAFIRI
// @Description Get product by product id
// @Tags Public API
// @Accept json
// @Produce json
// @Success 200 {object} []ProductResponseObject "Success data. Check response body for data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /public/product/{puid} [get]
func (c *Controller) GetProductByProductID(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetPublicBusinessByID get business information by business id
// @Summary Endpoint to get business information based on business unique id
// @Description Get business information and products
// @Tags Public API
// @Accept json
// @Produce json
// @Success 200 {object} []BusinessResponse "Success data. Check response body for data containing business data and all product array"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /public/business/products/all [get]
func (c *Controller) GetPublicBusinessByID(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
