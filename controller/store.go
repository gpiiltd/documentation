package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddProductToStore adds a new user product to the store
// @Summary Add product to store create new product appended to the store
// @Description Add products information to store. This makes the product visible to all users
// @Tags Product Store
// @Accept json
// @Produce json
// @Param Products body Products true  "The Product data to be stored"
// @Success 200 {object} Products "The product data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /products/store/ [post]
func (c *Controller) AddProductToStore(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// UpdateProductStoreData updates the information of a product in a store
// @Summary Update product information of the store
// @Description Update product store edits some store data
// @Tags Product Store
// @Accept json
// @Produce json
// @Param Products body Products true  "The Product data to be stored"
// @Success 200 {object} Products "The product data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /products/store/ [patch]
func (c *Controller) UpdateProductStoreData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// DeleteProductStoreData deletes the information of a product in a store
// @Summary Delete product information of the store
// @Description Delete product information of the store
// @Tags Product Store
// @Accept json
// @Produce json
// @Param Products body Products true  "The Product data to be stored"
// @Success 200 {object} Products "The product data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /products/store/ [delete]
func (c *Controller) DeleteProductStoreData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetAllBusinessProduct retreives the list of all products belonging to a business
// @Summary This endpoint retrieves all product data belonging to the business
// @Description Get all product information about a business from business unique id
// @Tags Product Store
// @Accept json
// @Produce json
// @Success 200 {object} UserData "The product data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /products/store/{buid} [get]
func (c *Controller) GetAllBusinessProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
