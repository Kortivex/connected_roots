package httpserver

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/role"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/user"

	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/labstack/echo/v4"

	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/metrics"
)

// registerRoutes method to register all new routes that the http server exposes.
func (s *Service) registerRoutes(ctx *connected_roots.Context) {
	// User endpoints.
	usersHandler := user.NewUsersHandlers(ctx)
	usersGrp := s.Echo.Group("/users")
	usersGrp.GET("/:user_id", usersHandler.GetUserHandler).Name = "get-user"
	usersGrp.PATCH("/:user_id", usersHandler.PatchUserPartiallyHandler).Name = "patch-user"
	usersGrp.POST("/:user_id/auth", usersHandler.PostUserAuthHandler).Name = "post-user-auth"

	// Roles endpoints.
	rolesHandler := role.NewRolesHandlers(ctx)
	rolesGrp := s.Echo.Group("/roles")
	rolesGrp.POST("", rolesHandler.PostRolesHandler).Name = "post-role"
	rolesGrp.PUT("/:role_id", rolesHandler.PutRolesHandler).Name = "put-role"
	rolesGrp.GET("/:role_id", rolesHandler.GetRolesHandler).Name = "get-role"
	rolesGrp.GET("", rolesHandler.ListRolesHandler).Name = "list-roles"

	// Health endpoints.
	healthGrp := s.Echo.Group("/health")
	metricsHandler := metrics.NewMetricsHandlers(ctx)
	healthGrp.GET("/alive", metricsHandler.GetHealthAliveHandler).Name = "get-health-alive"

	// Debug endpoints.
	debugGrp := s.Echo.Group("/debug")
	debugGrp.GET("/vars", echo.WrapHandler(http.DefaultServeMux)).Name = "get-debug-vars"
}
