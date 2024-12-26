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

	baseUrl := "/app"
	app := echo.New()

	conn := redis.GetConnection()
	redisRepo := redis.NewRepository(conn)
	service := service.NewPbinService(redisRepo)
	handler := handler.NewHandler(service)

	webapp := app.Group(baseUrl)
	webapp.File("/style.css", "./src/view/static/style/style.css")
	webapp.File("/index.js","./src/view/static/script/index.js")
	// webapp.Static("/static","./src/view/static")

	webapp.GET("", func(c echo.Context) error {
		home := view.Home()
		return home.Render(c.Request().Context(), c.Response())
	})

	webapp.POST("/pasteData", handler.PasteData)
	
	webapp.GET("/:id",func( c echo.Context) error {
		id := c.Param("id")
		if handler.IsDataPresent(c,id) {
		result := view.ResultPage(id)
		return result.Render(c.Request().Context(),c.Response())
		}
		notFound := view.NotFoundPage()
		return notFound.Render(c.Request().Context(),c.Response())
	})

	webapp.GET("/getContent/:id", handler.GetData)

	app.Logger.Fatal(app.Start(":" + constants.Env.AppPort))
}
