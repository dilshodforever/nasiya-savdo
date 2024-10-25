package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// CreateExchange handles creating a new exchange
// @Summary      Create a new Exchange
// @Description  Create a new exchange record with the specified details
// @Tags         Exchange
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        status query string true "Exchange Status" Enums(buy, sell)
// @Param        exchange body []pb.CreateExchangeRequestSwagger true "Exchange details"
// @Success      200 {object} []pb.ExchangeResponse "Exchange created successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while creating exchange"
// @Router       /exchange/create [post]
func (h *Handler) CreateExchange(ctx *gin.Context) {
	var reqs []pb.CreateExchangeRequest // Accept an array of requests
	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	for i := 0; i < len(reqs); i++ {
		reqs[i].Status = ctx.Query("status")

		_, err := h.ExchangeService.CreateExchange(context.Background(), &reqs[i])
		if err != nil {
			log.Print(err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(200, gin.H{"message": "Exchanges created successfully"})
}

// GetExchange retrieves an exchange by its ID
// @Summary      Get Exchange by ID
// @Description  Retrieve a specific exchange record using its ID
// @Tags         Exchange
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Exchange ID"
// @Success      200 {object} pb.GetExchangeResponse "Exchange retrieved successfully"
// @Failure      404 {string} string "Exchange not found"
// @Failure      500 {string} string "Error while retrieving exchange"
// @Router       /exchange/get/{id} [get]
func (h *Handler) GetExchange(ctx *gin.Context) {
	var req pb.ExchangeIdRequest
	req.Id = ctx.Param("id")
	res, err := h.ExchangeService.GetExchange(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// UpdateExchange handles updating an existing exchange
// @Summary      Update an Exchange
// @Description  Update an existing exchange record with new details
// @Tags         Exchange
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Exchange ID"
// @Param        exchange body pb.UpdateExchangeRequest true "Exchange details"
// @Param        status query string true "Exchange Status" Enums(buy, sell)
// @Success      200 {object} pb.ExchangeResponse "Exchange updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      404 {string} string "Exchange not found"
// @Failure      500 {string} string "Error while updating exchange"
// @Router       /exchange/update/{id} [put]
func (h *Handler) UpdateExchange(ctx *gin.Context) {
	var req pb.UpdateExchangeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	req.Id = ctx.Param("id")
	res, err := h.ExchangeService.UpdateExchange(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// DeleteExchange handles deleting an exchange by its ID
// @Summary      Delete Exchange by ID
// @Description  Delete a specific exchange record using its ID
// @Tags         Exchange
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Exchange ID"
// @Success      200 {object} pb.ExchangeResponse "Exchange deleted successfully"
// @Failure      404 {string} string "Exchange not found"
// @Failure      500 {string} string "Error while deleting exchange"
// @Router       /exchange/delete/{id} [delete]
func (h *Handler) DeleteExchange(ctx *gin.Context) {
	var req pb.ExchangeIdRequest
	req.Id = ctx.Param("id")
	res, err := h.ExchangeService.DeleteExchange(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// ListExchanges handles listing all exchanges with optional filters
// @Summary      List all Exchanges
// @Description  Retrieve a list of all exchange records with optional filtering by product ID or status
// @Tags         Exchange
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        status query string false "Exchange Status" Enums(buy, sell)
// @Param        limit query string false "Limit"
// @Param        offset query string false "Offset"
// @Success      200 {object} pb.GetAllExchangeResponse "List of exchanges retrieved successfully"
// @Failure      500 {string} string "Error while listing exchanges"
// @Router       /exchange/list [get]
func (h *Handler) ListExchanges(ctx *gin.Context) {
	var req pb.GetAllExchangeRequest
	// req.ProductId = ctx.Query("product_id")
	status := ctx.Query("status")
	if status != "" {
		req.Status = status // Directly assigning status if no enum conversion is required
	}
	req.Limit = ParseQueryInt32(ctx, "limit", 10)  // Default limit 10
	req.Offset = ParseQueryInt32(ctx, "offset", 0) // Default offset 0
	res, err := h.ExchangeService.ListExchanges(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// GetStatistika retrieves monthly statistika for exchange activities
// @Summary      Retrieve Monthly Statistika
// @Description  Get monthly statistika on total bought, total sold, total spend, total revenue, net amount, and net profit.
// @Tags         Statistika
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        month query int true "Month (1-12)"
// @Param        year  query int true "Year"
// @Success      200 {object} pb.ExchangeStatisticsResponse "Monthly statistika retrieved successfully"
// @Failure      400 {string} string "Invalid input parameters"
// @Failure      500 {string} string "Error while retrieving statistika"
// @Router       /exchange/statistik [get]
func (h *Handler) GetStatistika(ctx *gin.Context) {
	var req pb.ExchangeStatisticsRequest
	fmt.Println(req.Month,"\n", req.Year)
	// Parse and validate month
	month, err := strconv.Atoi(ctx.Query("month"))
	if err != nil || month < 1 || month > 12 {
		ctx.JSON(400, gin.H{"error": "Invalid month; must be between 1 and 12"})
		return
	}
	req.Month = int32(month)

	// Parse and validate year
	year, err := strconv.Atoi(ctx.Query("year"))
	if err != nil || year < 1 {
		ctx.JSON(400, gin.H{"error": "Invalid year; must be a positive integer"})
		return
	}
	req.Year = int32(year)

	// Call the service method
	res, err := h.ExchangeService.GetStatistika(context.Background(), &req)
	if err != nil {
		log.Print(err)
		ctx.JSON(500, gin.H{"error": "Error while retrieving statistika"})
		return
	}

	ctx.JSON(200, res)
}
