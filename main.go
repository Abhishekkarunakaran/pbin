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
	webapp.File("/style.css", "./src/view/style/style.css")
	webapp.File("/index.js","./src/view/script/index.js")

	home := view.Home()

	webapp.GET("", func(c echo.Context) error {
		return home.Render(c.Request().Context(), c.Response())
	})

	webapp.POST("/pasteData", handler.PasteData)
	

	app.Logger.Fatal(app.Start(":" + constants.Env.AppPort))
}
