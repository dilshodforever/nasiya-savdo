package handler

import (
	"fmt"

	"github.com/dilshodforever/nasiya-savdo/api/middleware"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// GetNotification handles retrieving a notification by user_id
// @Summary      Get Notification
// @Description  Retrieve a notification by user_id
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.GetNotificationByidResponse "Notification retrieved successfully"
// @Failure      404 {string} string "Notification not found"
// @Failure      500 {string} string "Error while retrieving notification"
// @Router       /notifications/get [get]
func (h *Handler) GetNotification(ctx *gin.Context) {
	var req pb.GetNotificationByidRequest
	id := middleware.GetUserId(ctx)
	req.UserId = id

	res, err := h.Notification.GetNotification(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	
	ctx.JSON(200, res)
}

// DeleteNotification handles deleting a notification by user_id
// @Summary      Delete Notification
// @Description  Delete a notification by user_id
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.NotificationsResponse "Notification deleted successfully"
// @Failure      404 {string} string "Notification not found"
// @Failure      500 {string} string "Error while deleting notification"
// @Router       /notifications/delete [delete]
func (h *Handler) DeleteNotification(ctx *gin.Context) {
	var req pb.GetNotificationByidRequest
	id:=middleware.GetUserId(ctx)
	req.UserId=id
	res, err := h.Notification.DeleteNotification(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// ListNotification handles listing all notifications
// @Summary      List Notifications
// @Description  Retrieve a list of all notifications
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListNotificationResponse "List of notifications retrieved successfully"
// @Failure      500 {string} string "Error while retrieving notifications"
// @Router       /notifications/getlist [get]
func (h *Handler) ListNotification(ctx *gin.Context) {
	var req pb.GetNotificationByidRequest
	req.UserId = middleware.GetUserId(ctx)
	res, err := h.Notification.ListNotification(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
