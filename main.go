package main

import (
	"html/template"
	"io"
	log "log/slog"
	"net/http"

	"github.com/Abhishekkarunakaran/pbin/src/adapters/handler"
	"github.com/Abhishekkarunakaran/pbin/src/adapters/redis"
	"github.com/Abhishekkarunakaran/pbin/src/core/constants"
	"github.com/Abhishekkarunakaran/pbin/src/core/service"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	if viewContext, isMap := data.(map[string]any); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)

}
func main() {

	baseUrl := "pbin/app"
	apiUrl := "pbin/v1"
	app := echo.New()

	conn := redis.GetConnection()
	redisRepo := redis.NewRepository(conn)
	service := service.NewPbinService(redisRepo)
	handler := handler.NewHandler(service)

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("src/view/*.html")),
	}

	app.Renderer = renderer

	private := app.Group(apiUrl + "/private")

	private.POST("", handler.PasteData)

	webapp := app.Group(baseUrl)
	webapp.GET("", func(e echo.Context) error {
		if err := e.Render(http.StatusOK, "index", map[string]string{
			"name": "Lorem Ipsum",
		}); err != nil {
			log.Error(err.Error())
			return err
		}
		return nil
	})

	app.Logger.Fatal(app.Start(":"+constants.Env.AppPort))
}
