package router

import (
	"github.com/gorilla/mux"
	"github.com/v420v/cloudwatch-logs/internal/controller"
	"github.com/v420v/cloudwatch-logs/internal/middleware"
)

func NewRouter(m *middleware.Middleware, c *controller.Controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", c.HandleHome)
	router.HandleFunc("/about", c.HandleAbout)

	router.Use(m.LoggingMiddleware)

	return router
}
