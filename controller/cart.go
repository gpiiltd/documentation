package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddToCart function handles adding new data into the cart
// @Summary Endpoint to add new product to user cart information
// @Description Add new product information to cart system
// @Tags User Cart
// @Accept json
// @Produce json
// @Param Cart body Cart true  "The product data in the cart, as well as the user information"
// @Success 200 {object} []Cart "Success cart data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /cart/add [post]
func (c *Controller) AddToCart(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// RemoveItemFromCart removes a product from the cart
// @Summary Endpoint to remove product from cart item
// @Description This is used when a user wants to remove a particular data from cart
// @Tags User Cart
// @Accept json
// @Produce json
// @Param Cart body Cart true  "The product data in the cart, as well as the user information"
// @Success 200 {object} []Cart "Success data "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /cart/remove/{productid} [delete]
func (c *Controller) RemoveItemFromCart(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// EmptyCart removes all product from the cart
// @Summary Endpoint to remove all product from cart item
// @Description This is used when a user wants to empty user cart
// @Tags User Cart
// @Accept json
// @Produce json
// @Param Cart body Cart true  "The product data in the cart, as well as the user information"
// @Success 200 {object} []Cart "Success data "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /cart/empty [delete]
func (c *Controller) EmptyCart(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
