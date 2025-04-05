package controller

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.Logger
}

func NewController(logger *zap.Logger) *Controller {
	return &Controller{
		logger: logger,
	}
}

func (c *Controller) HandleHome(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("home_page_accessed")
	w.Write([]byte("Welcome to the homepage!"))
}

func (c *Controller) HandleAbout(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("about_page_accessed")
	fmt.Fprintf(w, "This is the about page!")
}
