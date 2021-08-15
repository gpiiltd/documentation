package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrderFunction creates a new user order
// @Summary Endpoint to create a new user order
// @Description Send create order data to order for a new product.
// @Tags Sales
// @Accept json
// @Produce json
// @Param CreateProductOrder body CreateProductOrder true  "The Product order data"
// @Success 200 {object} OrderResponseBody "Order response body"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /user/order/ [post]
func (c *Controller) CreateOrderFunction(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// Cancel Order cancels user's order
// @Summary Endpoint to create a new user order
// @Description Send create order data to order for a new product.
// @Tags Sales
// @Accept json
// @Produce json
// @Param CreateOrder body CreateOrder true  "The Product order data"
// @Success 200 {object} OrderResponseBody "Order response body"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /user/order/cancel [patch]
func (c *Controller) CancelOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// CancelUserOrder rejects a  user order from an organization
// @Summary Endpoint to cancel or reject a user order
// @Description Businesses can reject or cancel a user order of their endpoint. Frontend can call this API for that purpose
// @Tags Sales
// @Accept json
// @Produce json
// @Param CreateOrder body CreateOrder true  "The Product order data"
// @Success 200 {object} OrderResponseBody "Order response body"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /business/{buid}/order/{orderid}/cancel [patch]
func (c *Controller) CancelUserOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetAllMyOrder gets all orders belonging to an individual user
// @Summary Endpoint to get all orders of an individual user
// @Description Get business information and products
// @Tags Sales
// @Accept json
// @Produce json
// @Success 200 {object} []OrderResponseBody "Success data. Check response body for data containing business data and all product array"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /user/order/all [get]
func (c *Controller) GetAllMyOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetAllMyBusinessOrder gets all orders belonging to an business user
// @Summary Endpoint to get all business orders
// @Description Get orders for a particular business
// @Tags Sales
// @Accept json
// @Produce json
// @Success 200 {object} []OrderResponseBody "Success data. Check response body for data containing business data and all product array"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /business/{buid}/order/all [get]
func (c *Controller) GetAllMyBusinessOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetMyBusinessOrder gets an order belonging to an business user
// @Summary Endpoint to get a business order
// @Description Get order for a particular business
// @Tags Sales
// @Accept json
// @Produce json
// @Success 200 {object} OrderResponseBody "Success data. Check response body for data containing business data and all product array"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /business/{buid}/order/{orderid} [get]
func (c *Controller) GetMyBusinessOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetMyOrder gets an order belonging to an individual user
// @Summary Endpoint to get an order for an individual user
// @Description Get order data
// @Tags Sales
// @Accept json
// @Produce json
// @Success 200 {object} OrderResponseBody "Success data. Check response body for data containing business data and all product array"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /user/order/{orderid} [get]
func (c *Controller) GetMyOrder(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
