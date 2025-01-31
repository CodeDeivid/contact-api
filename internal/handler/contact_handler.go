package handler

import (
    "contact-api/internal/model"
    "contact-api/internal/repository"
    "contact-api/internal/validator"
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
    validatorPkg "github.com/go-playground/validator/v10"
)

type ContactHandler struct {
    repo      repository.ContactRepository
    validator *validator.Validator
}

func NewContactHandler(repo repository.ContactRepository) *ContactHandler {
    return &ContactHandler{
        repo:      repo,
        validator: validator.NewValidator(),
    }
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
    var contact model.Contact

    if err := c.ShouldBindJSON(&contact); err != nil {
        var ve validatorPkg.ValidationErrors
        if err.Error() == "EOF" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Body da requisição está vazio"})
            return
        }
        if errors.As(err, &ve) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
            return
        }
    }

    if errors := h.validator.ValidateContact(&contact); len(errors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
        return
    }

    if err := h.repo.Create(&contact); err != nil {
        if err.Error() == "duplicated key value violates unique constraint" {
            c.JSON(http.StatusConflict, gin.H{"error": "Email ou telefone já cadastrado"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar contato"})
        return
    }

    c.JSON(http.StatusCreated, contact)
}