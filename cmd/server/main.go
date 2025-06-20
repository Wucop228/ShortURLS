package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/Wucop228/ShortURLS/internal/config"
	"github.com/Wucop228/ShortURLS/internal/handler"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s port=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode, cfg.DBPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	e := echo.New()
	h := handler.NewHandlers(db)

	e.GET("/urls/:key", h.Redirect)
	e.POST("/create-url", h.Shorten)

	log.Println("Server running on port", cfg.ServerPort)
	e.Logger.Fatal(e.Start(":" + cfg.ServerPort))
}
