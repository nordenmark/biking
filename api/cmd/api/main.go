package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nordenmark/biking/api/packages/config"
	"github.com/nordenmark/biking/api/packages/parser"
	"github.com/nordenmark/biking/api/packages/storage"
)

func main() {
	godotenv.Load()
	conf := config.NewServerConfig()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.CorsAllowOrigins,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet},
	}))

	e.GET("/sessions", sessionsHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}

func sessionsHandler(c echo.Context) error {
	files, err := storage.ReadFiles("biking-sessions")
	if err != nil {
		log.Printf("failed reading storage files: %v", err)
		return c.String(http.StatusInternalServerError, "unable to read sessions")
	}

	sessions := parser.ParseFiles(files)

	return c.JSON(http.StatusOK, sessions)
}
