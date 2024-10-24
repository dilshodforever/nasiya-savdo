package handler

import (
	"context"
	"log"

	//"github.com/dilshodforever/nasiya-savdo/api/middleware"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// CreateTransaction handles creating a new transaction
// @Summary      Create a new Transaction
// @Description  Create a new transaction record with the specified details
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        transaction body pb.CreateTransactionRequestSwagger true "Transaction details"
// @Success      200 {object} pb.TransactionResponse "Transaction created successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while creating transaction"
// @Router       /transaction/create [post]
func (h *Handler) CreateTransaction(ctx *gin.Context) {
	var req pb.CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.TransactionService.CreateTransaction(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// GetTransaction handles retrieving a transaction by ID
// @Summary      Get Transaction by ID
// @Description  Retrieve transaction details by transaction ID
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Transaction ID"
// @Success      200 {object} pb.GetTransactionResponse "Transaction details"
// @Failure      400 {string} string "Invalid input"
// @Failure      404 {string} string "Transaction not found"
// @Failure      500 {string} string "Error while retrieving transaction"
// @Router       /transaction/get/{id} [get]
func (h *Handler) GetTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.TransactionIdRequest{Id: id}

	res, err := h.TransactionService.GetTransaction(context.Background(), req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// UpdateTransaction handles updating an existing transaction
// @Summary      Update Transaction
// @Description  Update an existing transaction record by ID
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Transaction ID"
// @Param        transaction body pb.UpdateTransactionRequest true "Updated transaction details"
// @Success      200 {object} pb.TransactionResponse "Transaction updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      404 {string} string "Transaction not found"
// @Failure      500 {string} string "Error while updating transaction"
// @Router       /transaction/{id} [put]
func (h *Handler) UpdateTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	var req pb.UpdateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	req.Id = id

	res, err := h.TransactionService.UpdateTransaction(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// DeleteTransaction handles deleting a transaction by ID
// @Summary      Delete Transaction by ID
// @Description  Delete a transaction record by ID
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Transaction ID"
// @Success      200 {object} pb.TransactionResponse "Transaction deleted successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      404 {string} string "Transaction not found"
// @Failure      500 {string} string "Error while deleting transaction"
// @Router       /transaction/{id} [delete]
func (h *Handler) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.TransactionIdRequest{Id: id}

	res, err := h.TransactionService.DeleteTransaction(context.Background(), req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// ListTransactions handles listing all transactions with optional filtering by contract_id
// @Summary      List Transactions
// @Description  List all transactions, optionally filtered by contract ID
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        contract_id query string false "Contract ID"
// @Param        limit query string false "Limit"
// @Param        offset query string false "Offset"
// @Success      200 {object} pb.GetAllTransactionResponse "List of transactions"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while listing transactions"
// @Router       /transaction/list [get]
func (h *Handler) ListTransactions(ctx *gin.Context) {
	contractID := ctx.Query("contract_id")
	req := &pb.GetAllTransactionRequest{ContractId: contractID}
	req.Limit = ParseQueryInt32(ctx, "limit", 10) // Default limit 10
	req.Offset = ParseQueryInt32(ctx, "offset", 0) // Default offset 0
	res, err := h.TransactionService.ListTransactions(context.Background(), req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// // CheckTransactions handles checking for due payments
// // @Summary      Test Transactions
// // @Description  Check all pending transactions and return a message if any payments are due this month
// // @Tags         Transactions
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Success      200 {object} pb.CheckResponse "Payments due this month"
// // @Failure      500 {string} string "Error while checking transactions"
// // @Router       /transaction/check [post]
// func (h *Handler) CheckTransactions(ctx *gin.Context) {
// 	req := &pb.CheckRequest{}
// 	req.StorageId = middleware.GetStorageId(ctx)
// 	req.UserId=middleware.GetUserId(ctx)
// 	res, err := h.TransactionService.CheckTransactions(ctx.Copy(), req)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, res)
// }

// // TestNotification handles checking for due payments
// // @Summary      Test Notification
// // @Description  Check all pending transactions and return a message if any payments are due this month
// // @Tags         Transactions
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Success      200 {object} pb.Testresponse "Payments due this month"
// // @Failure      500 {string} string "Error while checking transactions"
// // @Router       /transaction/test [post]
// func (h *Handler) TestNotification(ctx *gin.Context) {
// 	req := &pb.Testresponse{}
// 	res, err := h.TransactionService.TestNotification(ctx.Copy(), req)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, res)
// }
