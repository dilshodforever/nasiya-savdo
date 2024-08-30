package handler

import (
	"fmt"
	"io"
	"net/http"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// UploadFile handles uploading a file to MinIO
// @Summary      Upload File
// @Description  Upload a file to MinIO
// @Tags         MinIO
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        file formData file true "File to upload"
// @Param        filename formData string true "Filename"
// @Success      200 {object} pb.UploadFileResponse "File uploaded successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while uploading file"
// @Router       /minio/upload [post]
func (h *Handler) UploadFile(c *gin.Context) {
	// Fileni olish va defer bilan yopish
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close() // Faylni yopishni ta'minlash

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	filename := c.Request.FormValue("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filename is required"})
		return
	}

	// MinIO obyektini to'g'ri initsializatsiya qilinganligini tekshirish
	if h.MinIO == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "MinIO service is not initialized"})
		return
	}

	resp, err := h.MinIO.UploadFile(c, &pb.UploadFileRequest{
		File:     fileBytes,
		Filename: filename,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload file %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": resp})
}
