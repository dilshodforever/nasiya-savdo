package handler

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type File struct {
	File multipart.FileHeader `form:"file" binding:"required"`
}

// uploadFile
// @Summary uploadFile
// @Description Upload a media file
// @Tags media
// @Accept multipart/form-data
// @Param file formData file true "UploadMediaForm"
// @Success 201 {object} string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /minio/media [post]
func (h *Handler) Media(c *gin.Context) {
	var file File
	err := c.ShouldBind(&file)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	fileUrl := filepath.Join("./media", file.File.Filename)

	err = c.SaveUploadedFile(&file.File, fileUrl)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	fileExt := filepath.Ext(file.File.Filename)

	newFile := uuid.NewString() + fileExt

	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
				{
						"Effect": "Allow",
						"Principal": {
								"AWS": ["*"]
						},
						"Action": ["s3:GetObject"],
						"Resource": ["arn:aws:s3:::%s/*"]
				}
		]
}`, "photos")

	info, err := h.MinIO.FPutObject(context.Background(), "photos", newFile, fileUrl, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = h.MinIO.SetBucketPolicy(context.Background(), "photos", policy)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	println("\n Info Bucket:", info.Bucket)

	madeUrl := fmt.Sprintf("https://solihov.uz/photos/%s", newFile)

	c.JSON(201, gin.H{
		"made_url": madeUrl,
	})

}
