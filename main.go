package main

import (
	"github.com/habibridho/cetak-webapp/public/views"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.Renderer = views.NewRenderer()

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "public/static",
	}))

	e.GET("/", func(ctx echo.Context) error {
		err := ctx.Render(http.StatusOK, "home", nil)
		if err != nil {
			log.Printf("error: %v", err.Error())
		}
		return err
	})
	e.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "pong")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
