package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/controle-estoque/backend/internal/model"
	"github.com/user/controle-estoque/backend/internal/repository"
)

type MovementHandler struct {
	repo repository.MovementRepository
}

func NewMovementHandler(repo repository.MovementRepository) *MovementHandler {
	return &MovementHandler{repo: repo}
}

func (h *MovementHandler) Register(c *gin.Context) {
	var m model.Movement
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Register(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, m)
}

func (h *MovementHandler) List(c *gin.Context) {
	productIDStr := c.Query("product_id")

	var movements []model.Movement
	var err error

	if productIDStr == "" {
		// Listagem global (ex: as últimas 50)
		movements, err = h.repo.ListAll(50)
	} else {
		productID, parseErr := strconv.ParseInt(productIDStr, 10, 64)
		if parseErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product_id inválido"})
			return
		}
		movements, err = h.repo.ListByProduct(productID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar movimentações"})
		return
	}

	c.JSON(http.StatusOK, movements)
}
