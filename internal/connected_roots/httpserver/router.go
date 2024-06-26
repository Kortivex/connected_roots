package httpserver

import (
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/user"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/labstack/echo/v4"

	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/metrics"
)

// registerRoutes method to register all new routes that the http server exposes.
func (s *Service) registerRoutes(ctx *connected_roots.Context) {
	usersHandler := user.NewUsersHandlers(ctx)
	usersGrp := s.Echo.Group("/users")
	usersGrp.GET("/:user_id", usersHandler.GetUserHandler).Name = "get-user"
	usersGrp.POST("/:user_id/auth", usersHandler.PostUserAuthHandler).Name = "post-user-auth"

	// Health endpoints.
	healthGrp := s.Echo.Group("/health")
	metricsHandler := metrics.NewMetricsHandlers(ctx)
	healthGrp.GET("/alive", metricsHandler.GetHealthAliveHandler).Name = "get-health-alive"

	// Debug endpoints.
	debugGrp := s.Echo.Group("/debug")
	debugGrp.GET("/vars", echo.WrapHandler(http.DefaultServeMux)).Name = "get-debug-vars"
}
