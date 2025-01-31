package handler

import (
    "contact-api/internal/model"
    "contact-api/internal/repository"
    "contact-api/internal/validator"
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
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

// CreateContact - Criar novo contato
func (h *ContactHandler) CreateContact(c *gin.Context) {
    var contact model.Contact

    if err := c.ShouldBindJSON(&contact); err != nil {
        if err.Error() == "EOF" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Body da requisição está vazio"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
        return
    }

    if errors := h.validator.ValidateContact(&contact); len(errors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
        return
    }

    if err := h.repo.Create(&contact); err != nil {
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            c.JSON(http.StatusConflict, gin.H{"error": "Email ou telefone já cadastrado"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar contato"})
        return
    }

    c.JSON(http.StatusCreated, contact)
}

// GetAllContacts - Listar todos os contatos
func (h *ContactHandler) GetAllContacts(c *gin.Context) {
    contacts, err := h.repo.FindAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contatos"})
        return
    }

    c.JSON(http.StatusOK, contacts)
}

// GetContactByID - Buscar contato por ID
func (h *ContactHandler) GetContactByID(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    contact, err := h.repo.FindByID(uint(id))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Contato não encontrado"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contato"})
        return
    }

    c.JSON(http.StatusOK, contact)
}

// UpdateContact - Atualizar contato existente
func (h *ContactHandler) UpdateContact(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    existingContact, err := h.repo.FindByID(uint(id))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Contato não encontrado"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contato"})
        return
    }

    var updatedContact model.Contact
    if err := c.ShouldBindJSON(&updatedContact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }

    updatedContact.ID = existingContact.ID
    
    if errors := h.validator.ValidateContact(&updatedContact); len(errors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
        return
    }

    if err := h.repo.Update(&updatedContact); err != nil {
        if err.Error() == "duplicated key value violates unique constraint" {
            c.JSON(http.StatusConflict, gin.H{"error": "Email ou telefone já cadastrado"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar contato"})
        return
    }

    c.JSON(http.StatusOK, updatedContact)
}

// DeleteContact - Deletar contato
func (h *ContactHandler) DeleteContact(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if err := h.repo.Delete(uint(id)); err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Contato não encontrado"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar contato"})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}