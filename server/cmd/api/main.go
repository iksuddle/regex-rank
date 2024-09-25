package main

import (
	"github.com/iksuddle/regex-rank/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := newConfig()
	db := database.NewDB(newMySQLConfig(config))

	app := &app{
		config: config,
		db:     db,
		auth:   initAuth(config),
		store:  database.NewStore(db),
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	e.GET("/login", app.loginHandler)
	e.GET("/login/callback", app.loginCallbackHandler)
	e.GET("/logout", app.logoutHandler)

	// e.GET("/test", handlers.TestAuthRoute, handlers.AuthRoute)
	e.GET("/user", app.getUserHandler, app.authRoute)
	e.GET("/delete", app.deleteUserHandler, app.authRoute)

	e.Logger.Fatal(e.Start(":" + config.port))
}
