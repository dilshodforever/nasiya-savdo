package handler

import (
	"context"
	"log"

	"github.com/dilshodforever/nasiya-savdo/api/middleware"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// CreateProduct handles the creation of a new product
// @Summary      Create Product
// @Description  Create a new product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        product body pb.CreateProductRequestSwagger true "Product details"
// @Success      200 {object} pb.ProductResponse "Product created successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while creating product"
// @Router       /product/create [post]
func (h *Handler) CreateProduct(ctx *gin.Context) {
	var req pb.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	storageid:=middleware.GetStorageId(ctx)
	req.StorageId=storageid
	res, err := h.ProductService.CreateProduct(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// GetProduct handles retrieving a product by ID
// @Summary      Get Product
// @Description  Retrieve a product by its ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Product ID"
// @Success      200 {object} pb.GetProductResponse "Product retrieved successfully"
// @Failure      404 {string} string "Product not found"
// @Failure      500 {string} string "Error while retrieving product"
// @Router       /product/get/{id} [get]
func (h *Handler) GetProduct(ctx *gin.Context) {
	var req pb.ProductIdRequest
	req.Id = ctx.Param("id")
	res, err := h.ProductService.GetProduct(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// UpdateProduct handles updating an existing product
// @Summary      Update Product
// @Description  Update an existing product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Product ID"
// @Param        product body pb.UpdateProductRequest true "Updated product details"
// @Success      200 {object} pb.ProductResponse "Product updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while updating product"
// @Router       /product/update/{id} [put]
func (h *Handler) UpdateProduct(ctx *gin.Context) {
	var req pb.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	req.Id = ctx.Param("id")
	res, err := h.ProductService.UpdateProduct(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// DeleteProduct handles deleting a product by ID
// @Summary      Delete Product
// @Description  Delete a product by its ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Product ID"
// @Success      200 {object} pb.ProductResponse "Product deleted successfully"
// @Failure      404 {string} string "Product not found"
// @Failure      500 {string} string "Error while deleting product"
// @Router       /product/delete/{id} [delete]
func (h *Handler) DeleteProduct(ctx *gin.Context) {
	var req pb.ProductIdRequest
	req.Id = ctx.Param("id")
	res, err := h.ProductService.DeleteProduct(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// ListProducts handles listing all products with optional filters
// @Summary      List Products
// @Description  Retrieve a list of products with optional filters
// @Tags         Product
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        search query string false "Product Name"
// @Param        limit query string false "Limit"
// @Param        offset query string false "Offset"
// @Success      200 {object} pb.GetAllProductResponse "Products retrieved successfully"
// @Failure      500 {string} string "Error while listing products"
// @Router       /product/list [get]
func (h *Handler) ListProducts(ctx *gin.Context) {
	var req pb.GetAllProductRequest
	req.Name = ctx.Query("search")
	// req.Color = ctx.Query("color")
	// req.Model = ctx.Query("model")
	req.Limit = ParseQueryInt32(ctx, "limit", 10) // Default limit 10
	req.Offset = ParseQueryInt32(ctx, "offset", 0) // Default offset 0
	// req.StorageId = middleware.GetStorageId(ctx)
	res, err := h.ProductService.ListProducts(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
