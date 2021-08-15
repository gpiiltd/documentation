package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Search database for text match
// @Summary Endpoint to search the system database for text matches in products and services
// @Description Search the system for products and data match
// @Tags Public Search
// @Accept json
// @Produce json
// @Success 200 {object} SearchResult "Success search result "
// @Failure 400 {object} ResponseObject "API response object"
// @Router /search/{text} [get]
func (c *Controller) Search(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// SortSearchByProducts sorts the search result of an api based on the products
// @Summary Endpoint to search the system database for text matches in products
// @Description Search the system for products match
// @Tags Public Search
// @Accept json
// @Produce json
// @Success 200 {object} SearchResult "Success search result from the procuts section"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /search/products/{text} [get]
func (c *Controller) SortSearchByProducts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// SortSearchByProductsPriceRange sorts the search result of an api based on the products price range
// @Summary Endpoint to search the system database for text matches in products and match in prices
// @Description Search the system for products match and price range match
// @Tags Public Search
// @Accept json
// @Produce json
// @Success 200 {object} SearchResult "Success search result from the procuts section"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /search/products/price/{minprice}/{maxprice} [get]
func (c *Controller) SortSearchByProductsPriceRange(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// SortSearchByProductsCategory sorts the search result of an api based on the products category
// @Summary Endpoint to search the system database for text matches in products and category
// @Description Search the system for products match and category match
// @Tags Public Search
// @Accept json
// @Produce json
// @Success 200 {object} SearchResult "Success search result from the procuts section"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /search/products/category/{text} [get]
func (c *Controller) SortSearchByProductsCategory(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// SortSearchByProductsIndustry sorts the search result of an api based on the products industry
// @Summary Endpoint to search the system database for text matches in products and industry
// @Description Search the system for products match and industry match
// @Tags Public Search
// @Accept json
// @Produce json
// @Success 200 {object} SearchResult "Success search result from the procuts section"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /search/products/industry/{text} [get]
func (c *Controller) SortSearchByProductsIndustry(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
