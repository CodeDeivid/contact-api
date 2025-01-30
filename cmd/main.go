package main

import (
	"contact-api/internal/handler"
	"contact-api/internal/model"
	"contact-api/internal/repository"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Conectar ao PostgreSQL
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=UTC"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrar schema
	db.AutoMigrate(&model.Contact{})

	// Inicializar repositório e handler
	contactRepo := repository.NewContactRepository(db)
	contactHandler := handler.NewContactHandler(contactRepo)

	// Configurar rotas
	r := gin.Default()
	
	r.POST("/contacts", contactHandler.CreateContact)

	// Executar servidor
	r.Run(":8080")
}