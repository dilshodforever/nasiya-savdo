package handler

import (
	"bytes"
	"fmt"
	"strconv"

	"log"
	"net/http"

	"github.com/dilshodforever/nasiya-savdo/api/middleware"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf/v2"
)

// CreateContract handles the creation of a new contract
// @Summary      Create Contract
// @Description  Create a new contract
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        contract body pb.CreateContractRequestSwagger true "Contract details"
// @Success      200 {object} pb.ContractResponse "Contract created successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while creating contract"
// @Router       /contract/create [post]
func (h *Handler) CreateContract(ctx *gin.Context) {
	var req pb.CreateContractRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	storageid := middleware.GetStorageId(ctx)
	req.StorageId = storageid
	res, err := h.ContractService.CreateContract(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// GetContract retrieves a specific contract by ID
// @Summary      Get Contract
// @Description  Retrieve a contract by ID
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Contract ID"
// @Success      200 {object} pb.GetContractResponse "Contract details"
// @Failure      404 {string} string "Contract not found"
// @Failure      500 {string} string "Error while retrieving contract"
// @Router       /contract/get/{id} [get]
func (h *Handler) GetContract(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.ContractIdRequest{Id: id}
	res, err := h.ContractService.GetContract(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(404, gin.H{"error": "Contract not found"})
		return
	}
	ctx.JSON(200, res)
}

// UpdateContract handles updating an existing contract
// @Summary      Update Contract
// @Description  Update an existing contract
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Contract ID"
// @Param        contract body pb.UpdateContractRequest true "Updated contract details"
// @Success      200 {object} pb.ContractResponse "Contract updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while updating contract"
// @Router       /contract/update/{id} [put]
func (h *Handler) UpdateContract(ctx *gin.Context) {
	var req pb.UpdateContractRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	req.Id = ctx.Param("id")
	res, err := h.ContractService.UpdateContract(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// DeleteContract handles deleting a contract by ID
// @Summary      Delete Contract
// @Description  Delete a contract by ID
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Contract ID"
// @Success      200 {object} pb.ContractResponse "Contract deleted successfully"
// @Failure      404 {string} string "Contract not found"
// @Failure      500 {string} string "Error while deleting contract"
// @Router       /contract/delete/{id} [delete]
func (h *Handler) DeleteContract(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.ContractIdRequest{Id: id}
	res, err := h.ContractService.DeleteContract(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(404, gin.H{"error": "Contract not found"})
		return
	}
	ctx.JSON(200, res)
}

// ListContracts retrieves a list of all contracts with optional filters
// @Summary      List Contracts
// @Description  Retrieve a list of contracts with optional filters
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        search query string false "Filter by consumer name"
// @Param        limit query string false "Limit"
// @Param        page query string false "Offset"
// @Success      200 {object} pb.GetAllContractResponse "List of contracts"
// @Failure      500 {string} string "Error while retrieving contracts"
// @Router       /contract/list [get]
func (h *Handler) ListContracts(ctx *gin.Context) {
	req := &pb.GetAllContractRequest{
		ConsumerName: ctx.Query("search"),
		//Status:       ctx.Query("status"),
		//PasportSeria: ctx.Query("pasport_seria"),
		Limit:        ParseQueryInt32(ctx, "limit", 10), // Default limit 10
		Page:       ParseQueryInt32(ctx, "page", 1), // Default offset 0
	}

	//req.StorageId = middleware.GetStorageId(ctx)
	res, err := h.ContractService.ListContracts(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

func ParseQueryInt32(ctx *gin.Context, key string, defaultValue int32) int32 {
	valueStr := ctx.Query(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return int32(value)
}

// GetContract retrieves a specific contract by ID
// @Summary      Get Contract
// @Description  Retrieve a contract by ID and return it in PDF format
// @Tags         Contract
// @Accept       json
// @Produce      application/pdf
// @Security     BearerAuth
// @Param        id path string true "Contract ID"
// @Success      200 {file} application/pdf "Contract details in PDF format"
// @Failure      404 {string} string "Contract not found"
// @Failure      500 {string} string "Error while retrieving contract"
// @Router       /contract/getpdf/{id} [get]
func (h *Handler) GetContractPdf(ctx *gin.Context) {
	id := ctx.Param("id")
	req := &pb.ContractIdRequest{Id: id}
	res, err := h.ContractService.GetContract(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(404, gin.H{"error": "Contract not found"})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Contract Details")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Contract ID: %s", res.Id))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Consumer Name: %s", res.ConsumerName))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Consumer Passport Serial: %s", res.ConsumerPassportSerial))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Consumer Address: %s", res.ConsumerAddress))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Consumer Phone Number: %s", res.ConsumerPhoneNumber))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Passport Image: %s", res.PassportImage))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Status: %s", res.Status))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Duration: %d", res.Duration))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Created At: %s", res.CreatedAt))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Deleted At: %s", res.DeletedAt))

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		log.Printf("Error generating PDF: %v", err)
		ctx.JSON(500, gin.H{"error": "Error generating PDF"})
		return
	}

	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", "attachment; filename=contract_details.pdf")
	ctx.Data(http.StatusOK, "application/pdf", buf.Bytes())
}
