package main

import (
	"github.com/habibridho/cetak-webapp/public"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strings"
)

func main() {
	e := echo.New()
	e.Renderer = public.NewRenderer()

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "public",
		Skipper: func(ctx echo.Context) bool {
			if strings.HasSuffix(ctx.Path(), ".html") {
				return true
			}
			return false
		},
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
	e.GET("/presubmit", func(ctx echo.Context) error {
		err := ctx.Render(http.StatusOK, "presubmit", nil)
		if err != nil {
			log.Printf("error: %v", err.Error())
		}
		return err
	})
	e.Logger.Fatal(e.Start(":8080"))
}
