package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Business Registration Documentation godoc
// @Summary Endpoint to send user data for registration
// @Description Send user data to the main authentication endpoint.
// @Tags APICORE Business
// @Accept json
// @Produce json
// @Param BusinessRegistration body BusinessRegistration true  "The Business Registration Data"
// @Success 200 {object} BusinessProfile "Success data "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/business/ [post]
func (c *Controller) BusinessRegistration(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// UpdateBusinessData
// @Summary Endpoint to update user business data
// @Description Perform Update endpoint for Business Registration
// @Tags APICORE Business
// @Accept json
// @Produce json
// @Param BusinessProfile body BusinessProfile true  "The Business Registration Data"
// @Success 200 {object} BusinessProfile "Success data "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/business/update [post]
func (c *Controller) UpdateBusinessData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// DeleteBuiness removes a business data
// @Summary Endpoint to delete a user business information
// @Description Perform Delete endpoint for Business Registration
// @Tags APICORE Business
// @Accept json
// @Produce json
// @Param BusinessProfile body BusinessProfile true  "The Business Registration Data"
// @Success 200 {object} BusinessProfile "Success data "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/business/ [delete]
func (c *Controller) DeleteBuiness(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetBusinessByID
// @Summary Endpoint to get business information by id
// @Description Get an array of all business and business data by business unique ids
// @Tags APICORE Business
// @Accept json
// @Produce json
// @Param businessid body string true  "The Business unique id string"
// @Success 200 {object} BusinessResponse "Success data "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/business/{buid} [get]
func (c *Controller) GetBusinessByID(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetAllBusiness retrieves the list of all businesses in barafiri
// @Summary Endpoint to get all businesses on BARAFIRI
// @Description Get an array of all business and business data
// @Tags APICORE Business
// @Accept json
// @Produce json
// @Success 200 {object} []BusinessResponse "Success data "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/business/all [get]
func (c *Controller) GetAllBusiness(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
