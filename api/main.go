package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not read .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not set in .env file")
	}

	InitAuth(port)

	e := echo.New()

	e.GET("/", indexHandler)
	e.GET("/login", LoginHandler)
	e.GET("/auth/callback", AuthCallbackHandler)

	e.Logger.Fatal(e.Start(":" + port))
}

func indexHandler(c echo.Context) error {
	indexHTML := `<h2>login with <a href="/login">GitHub</a></h2>`
	return c.HTML(http.StatusOK, indexHTML)
}
