package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransaction creates a new transaction on system
// @Summary Endpoint to create new transaction for businesses
// @Description Endpoint to create new transaction for businesses.
// @Tags Public API
// @Accept json
// @Produce json
// @Param Pay body Pay true "The payment data"
// @Success 200 {object} []ProductResponseObject "Success data. Check response body for data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /payments/transaction/pay [post]
func (c *Controller) CreateTransaction(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
