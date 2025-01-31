package test

import (
    "bytes"
    "contact-api/internal/handler"
    "contact-api/internal/model"
    "contact-api/internal/repository"
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "net/http/httptest"
    "os"
    "strconv"
    "testing"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

func setupRouter() *gin.Engine {
    if err := godotenv.Load("../.env"); err != nil {
        log.Println("Aviso: .env não encontrado, usando variáveis de ambiente padrão")
    }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Falha ao conectar ao banco de dados: ", err)
    }

    db.AutoMigrate(&model.Contact{})
    db.Exec("DELETE FROM contacts") // Limpa os contatos antes de cada teste

    contactRepo := repository.NewContactRepository(db)
    contactHandler := handler.NewContactHandler(contactRepo)

    r := gin.Default()
    r.POST("/contacts", contactHandler.CreateContact)
    r.GET("/contacts", contactHandler.GetAllContacts)
    r.GET("/contacts/:id", contactHandler.GetContactByID)
    r.PUT("/contacts/:id", contactHandler.UpdateContact)
    r.DELETE("/contacts/:id", contactHandler.DeleteContact)

    return r
}

func generateUniqueEmail() string {
    return fmt.Sprintf("test%d@example.com", time.Now().UnixNano()+rand.Int63())
}

func TestCreateContact(t *testing.T) {
    r := setupRouter()

    newContact := model.Contact{
        Name:  "Jane Doe",
        Email: generateUniqueEmail(),
        Phone: "0987654321",
    }

    jsonValue, _ := json.Marshal(newContact)
    req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllContacts(t *testing.T) {
    r := setupRouter()
    req, _ := http.NewRequest("GET", "/contacts", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetContactByID(t *testing.T) {
    r := setupRouter()

    newContact := model.Contact{
        Name:  "Jane Doe",
        Email: generateUniqueEmail(),
        Phone: "0987654321",
    }

    jsonValue, _ := json.Marshal(newContact)
    req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var createdContact model.Contact
    json.Unmarshal(w.Body.Bytes(), &createdContact)
    assert.NotZero(t, createdContact.ID)

    req, _ = http.NewRequest("GET", "/contacts/"+strconv.Itoa(int(createdContact.ID)), nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateContact(t *testing.T) {
    r := setupRouter()

    newContact := model.Contact{
        Name:  "Jane Doe",
        Email: generateUniqueEmail(),
        Phone: "0987654321",
    }

    jsonValue, _ := json.Marshal(newContact)
    req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var createdContact model.Contact
    json.Unmarshal(w.Body.Bytes(), &createdContact)
    assert.NotZero(t, createdContact.ID)

    updatedContact := model.Contact{
        Name:  "Jane Smith",
        Email: generateUniqueEmail(),
        Phone: "0987654321",
    }

    jsonValue, _ = json.Marshal(updatedContact)
    req, _ = http.NewRequest("PUT", "/contacts/"+strconv.Itoa(int(createdContact.ID)), bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteContact(t *testing.T) {
    r := setupRouter()

    newContact := model.Contact{
        Name:  "Jane Doe",
        Email: generateUniqueEmail(),
        Phone: "0987654321",
    }

    jsonValue, _ := json.Marshal(newContact)
    req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var createdContact model.Contact
    json.Unmarshal(w.Body.Bytes(), &createdContact)
    assert.NotZero(t, createdContact.ID)

    req, _ = http.NewRequest("DELETE", "/contacts/"+strconv.Itoa(int(createdContact.ID)), nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNoContent, w.Code)
}