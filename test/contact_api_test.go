package test

import (
	"bytes"
	"contact-api/internal/handler"
	"contact-api/internal/model"
	"contact-api/internal/repository"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockContactRepository struct {
	mock.Mock
}

func (m *MockContactRepository) Create(contact *model.Contact) error {
	args := m.Called(contact)
	return args.Error(0)
}

func setupRouter(repo repository.ContactRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/contacts", func(c *gin.Context) {
		handler := handler.NewContactHandler(repo)
		handler.CreateContact(c)
	})

	return r
}

func TestCreateContact(t *testing.T) {
	mockRepo := new(MockContactRepository)
	contact := &model.Contact{Name: "Paulo Silva", Email: "paulo@exemplo.com"}
	mockRepo.On("Create", contact).Return(nil)

	router := setupRouter(mockRepo)
	body, _ := json.Marshal(contact)
	req, _ := http.NewRequest("POST", "/contact", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockRepo.AssertExpectations(t)
}