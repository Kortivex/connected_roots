package frontend

import (
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/login"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/home"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/fmetrics"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/labstack/echo/v4"
)

// registerRoutes method to register all new routes that the http server exposes.
func (s *Service) registerRoutes(ctx *connected_roots.Context) {
	homeHandler := home.NewHomeHandlers(ctx)
	s.Echo.GET("/", homeHandler.GetHomeHandler).Name = "get-home"

	loginHandler := login.NewLoginHandlers(ctx)
	s.Echo.GET("/login", loginHandler.GetLoginHandler).Name = "get-login"
	s.Echo.POST("/login", loginHandler.PostLoginHandler).Name = "post-login"

	// Health endpoints.
	healthGrp := s.Echo.Group("/health")
	metricsHandler := fmetrics.NewMetricsHandlers(ctx)
	healthGrp.GET("/alive", metricsHandler.GetHealthAliveHandler).Name = "get-health-alive"

	// Debug endpoints.
	debugGrp := s.Echo.Group("/debug")
	debugGrp.GET("/vars", echo.WrapHandler(http.DefaultServeMux)).Name = "get-debug-vars"
}
