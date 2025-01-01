package main

import (
	"github.com/Abhishekkarunakaran/pbin/src/adapters/handler"
	"github.com/Abhishekkarunakaran/pbin/src/adapters/redis"
	"github.com/Abhishekkarunakaran/pbin/src/core/constants"
	"github.com/Abhishekkarunakaran/pbin/src/core/service"
	"github.com/Abhishekkarunakaran/pbin/src/view"
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()

	conn := redis.GetConnection()
	redisRepo := redis.NewRepository(conn)
	service := service.NewPbinService(redisRepo)
	handler := handler.NewHandler(service)

	app.File("/style.css", "./src/view/static/style/style.css")
	app.File("/index.js", "./src/view/static/script/index.js")
	app.Static("/images", "./images")

	app.GET("", func(c echo.Context) error {
		home := view.Home()
		return home.Render(c.Request().Context(), c.Response())
	})

	app.POST("/pasteData", handler.PasteData)

	app.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		if handler.IsDataPresent(c, id) {
			result := view.ResultPage(id)
			return result.Render(c.Request().Context(), c.Response())
		}
		notFound := view.NotFoundPage()
		return notFound.Render(c.Request().Context(), c.Response())
	})

	app.GET("/getContent/:id", handler.GetData)

	app.Logger.Fatal(app.Start(":" + constants.Env.AppPort))
}
