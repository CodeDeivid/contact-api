package handler

import (
	"contact-api/internal/model"
	"contact-api/internal/repository"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	repo repository.ContactRepository
}

func NewContactHandler(repo repository.ContactRepository) *ContactHandler {
	return &ContactHandler{repo}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var contact model.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contact)
}