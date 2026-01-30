package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	httpadapter "crudWithDB/internal/adapters/httpadapter"
	"crudWithDB/internal/adapters/repository"
	"crudWithDB/internal/application"
)

func main() {
	logger := slog.Default()
	dsn := getEnv("DB_DSN", "postgres://postgres:kalel@localhost:5433/postgres?sslmode=disable")
	port := getEnv("PORT", "8080")

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("db connection error:", err)
	}

	repo := repository.NewUserRepositoryPostgres(db)
	service := application.NewUserService(repo)
	handler := httpadapter.NewUserHTTPHandler(service)

	r := gin.Default()
	handler.RegisterRoutes(r)

	logger.Info("listening on :" + port)
	err = r.Run(":" + port)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
