package handler

import (
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/gin-gonic/gin"
)

// CreateContract handles the creation of a new contract
// @Summary      Create Contract
// @Description  Create a new contract
// @Tags         Contract
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        contract body pb.CreateContractRequest true "Contract details"
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
// @Router       /contract/{id} [get]
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
// @Param        contract body pb.UpdateContractRequest true "Updated contract details"
// @Success      200 {object} pb.ContractResponse "Contract updated successfully"
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Error while updating contract"
// @Router       /contract/update [put]
func (h *Handler) UpdateContract(ctx *gin.Context) {
	var req pb.UpdateContractRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
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
// @Param        consumer_name query string false "Filter by consumer name"
// @Param        status query string false "Filter by status"
// @Param        pasport_seria query string false "Filter by passport seria"
// @Success      200 {object} pb.GetAllContractResponse "List of contracts"
// @Failure      500 {string} string "Error while retrieving contracts"
// @Router       /contract/list [get]
func (h *Handler) ListContracts(ctx *gin.Context) {
	req := &pb.GetAllContractRequest{
		ConsumerName: ctx.Query("consumer_name"),
		Status:       ctx.Query("status"),
		PasportSeria: ctx.Query("pasport_seria"),
	}
	res, err := h.ContractService.ListContracts(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
