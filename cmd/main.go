package main

import (
	"contact-api/internal/handler"
	"contact-api/internal/model"
	"contact-api/internal/repository"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Erro ao carregar o arquivo .env")
	}

	// Obter as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Definir uma porta padrão caso DB_PORT esteja vazia
	if dbPort == "" {
		dbPort = "5432"
	}

	// Criar a string de conexão DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Exibir o DSN para depuração
	fmt.Println("DSN:", dsn)

	// Conectar ao PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrar schema
	db.AutoMigrate(&model.Contact{})

	// Inicializar repositório e handler
	contactRepo := repository.NewContactRepository(db)
	contactHandler := handler.NewContactHandler(contactRepo)

	// Configurar rotas
	r := gin.Default()
	r.POST("/contact", contactHandler.CreateContact)

	// Executar servidor
	r.Run(":8080")
}
