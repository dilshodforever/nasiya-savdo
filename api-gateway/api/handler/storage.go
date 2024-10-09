package handler

import (
	"context"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// // CreateStorage handles creating a new storage
// // @Summary      Create a new Storage
// // @Description  Create a new storage record with the specified details
// // @Tags         Storage
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        storage body pb.CreateStorageRequest true "Storage details"
// // @Success      200 {object} pb.StorageResponse "Storage created successfully"
// // @Failure      400 {string} string "Invalid input"
// // @Failure      500 {string} string "Error while creating storage"
// // @Router       /storage/create [post]
// func (h *Handler) CreateStorage(ctx *gin.Context) {
// 	var req pb.CreateStorageRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(400, gin.H{"error": "Invalid input"})
// 		return
// 	}
// 	res, err := h.StorageService.CreateStorage(context.Background(), &req)
// 	if err != nil {
// 		log.Print(err)
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(200, res)
// }


// GetStorage retrieves a storage by its ID
// @Summary      Get Storage by ID
// @Description  Retrieve a specific storage record using its ID
// @Tags         Storage
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Storage ID"
// @Success      200 {object} pb.GetStorageResponse "Storage retrieved successfully"
// @Failure      404 {string} string "Storage not found"
// @Failure      500 {string} string "Error while retrieving storage"
// @Router       /storage/get/{id} [get]
func (h *Handler) GetStorage(ctx *gin.Context) {
	var req pb.StorageIdRequest
	req.Id = ctx.Param("id")
	res, err := h.StorageService.GetStorage(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// UpdateStorage handles updating an existing storage
// @Summary      Update a Storage
// @Description  Update an existing storage record with new details
// @Tags         Storage
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Storage ID"
// @Param        storage body pb.UpdateStorageRequest true "Storage details"
// @Success      200 {object} pb.StorageResponse "Storage updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      404 {string} string "Storage not found"
// @Failure      500 {string} string "Error while updating storage"
// @Router       /storage/update/{id} [put]
func (h *Handler) UpdateStorage(ctx *gin.Context) {
	var req pb.UpdateStorageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	req.Id = ctx.Param("id")
	res, err := h.StorageService.UpdateStorage(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// DeleteStorage handles deleting a storage by its ID
// @Summary      Delete Storage by ID
// @Description  Delete a specific storage record using its ID
// @Tags         Storage
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Storage ID"
// @Success      200 {object} pb.StorageResponse "Storage deleted successfully"
// @Failure      404 {string} string "Storage not found"
// @Failure      500 {string} string "Error while deleting storage"
// @Router       /storage/delete [delete]
func (h *Handler) DeleteStorage(ctx *gin.Context) {
	var req pb.StorageIdRequest
	req.Id = ctx.Param("id")
	res, err := h.StorageService.DeleteStorage(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// ListStorages handles listing all storages with optional filters
// @Summary      List all Storages
// @Description  Retrieve a list of all storage records with optional filtering by name or user ID
// @Tags         Storage
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        name query string false "Storage name"
// @Param        user_id query string false "User ID"
// @Param        limit query string false "Limit"
// @Param        offset query string false "Offset"
// @Success      200 {object} pb.GetAllStorageResponse "List of storages retrieved successfully"
// @Failure      500 {string} string "Error while listing storages"
// @Router       /storage/list [get]
func (h *Handler) ListStorages(ctx *gin.Context) {
	var req pb.GetAllStorageRequest
	req.Name = ctx.Query("name")
	req.UserId = ctx.Query("user_id")
	req.Limit = ParseQueryInt32(ctx, "limit", 10) // Default limit 10
	req.Offset = ParseQueryInt32(ctx, "offset", 0) // Default offset 0
	res, err := h.StorageService.ListStorages(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
