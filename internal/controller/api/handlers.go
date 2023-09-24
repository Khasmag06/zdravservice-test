package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	defaultPaginationLimit = 2
	defaultPageNumber      = 1
)

// @Tags Products
// @Summary get list of product
// @Description get a list of product with pagination and sorting
// @ID getProducts
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param limit query int false "Number of items per page (default is 2)"
// @Param sortOrder query string false "Sorting order (default is 'asc')"
// @Success 200 {array} models.Product "List of product"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /products/get [get]
func (h *Handler) getProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = defaultPageNumber
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		limit = defaultPaginationLimit
	}
	sortOrder := c.Query("sortOrder")

	ctx := context.Background()

	products, err := h.productService.GetProducts(ctx, page, limit, sortOrder)
	if err != nil {
		h.logger.Errorf("failed to fetch products: %v", err.Error())
		writeErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, products)
}
