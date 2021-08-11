package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCategories Documentation godoc
// @Summary Endpoint to get the list of all categories on the database
// @Description Get the list of all Categories
// @Tags Public API
// @Accept json
// @Produce json
// @Success 200 {object} SubCategoryList "Category list and sub category"
// @Failure 400 {object} ResponseObject "Category list"
// @Router /util/category/ [get]
func (c *Controller) GetAllCategories(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetIndustries Documentation godoc
// @Summary Endpoint to get the list of all industries on the database
// @Description Get the list of all industries
// @Tags Public API
// @Accept json
// @Produce json
// @Success 200 {object} []Industries "Category list and sub category"
// @Failure 400 {object} ResponseObject "Category list"
// @Router /util/industry/ [get]
func (c *Controller) GetIndustries(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
