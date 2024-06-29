package server

import (
	"github.com/Montheankul-K/game-microservices/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthCheck struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

func (s *server) healthCheckService(c echo.Context) error {
	return response.SuccessResponse(c, http.StatusOK, &healthCheck{
		App:    s.cfg.App.Name,
		Status: "ok",
	})
}
