package frontend

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/user"
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/role"

	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/logout"

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

	// Login endpoints.
	loginHandler := login.NewLoginHandlers(ctx)
	s.Echo.GET("/login", loginHandler.GetLoginHandler).Name = "get-login"
	s.Echo.POST("/login", loginHandler.PostLoginHandler).Name = "post-login"

	// Logout endpoints.
	logoutHandler := logout.NewLogoutHandlers(ctx)
	s.Echo.GET("/logout", logoutHandler.GetLogoutHandler, s.SessionMiddleware).Name = "get-logout"

	// Users endpoints.
	usersHandler := user.NewUsersHandlers(ctx)
	usersGrp := s.Echo.Group("/users")
	usersGrp.GET("/profile", usersHandler.GetUserProfileHandler, s.SessionMiddleware).Name = "get-user-profile"
	usersGrp.GET("/profile/edit", usersHandler.GetEditUserProfileHandler, s.SessionMiddleware).Name = "get-edit-user-profile"
	usersGrp.POST("/profile/edit", usersHandler.PostEditUserProfileHandler, s.SessionMiddleware).Name = "post-edit-user-profile"

	// Admin endpoints.
	adminGrp := s.Echo.Group("/admin")

	// Role endpoints.
	rolesHandler := role.NewRolesHandlers(ctx)
	rolesGrp := adminGrp.Group("/roles")
	rolesGrp.GET("/list", rolesHandler.GetRolesListHandler, s.SessionMiddleware).Name = "list-roles"
	/*
		rolesGrp.GET("/new", rolesHandler.GetRoleCreateHandler, s.SessionMiddleware).Name = "new-role"
		rolesGrp.GET("/view/:role_id", rolesHandler.GetRoleViewHandler, s.SessionMiddleware).Name = "view-role"
		rolesGrp.GET("/edit/:role_id", rolesHandler.GetRoleUpdateHandler, s.SessionMiddleware).Name = "edit-role"
		rolesGrp.GET("/delete/:role_id", rolesHandler.GetRoleDeleteHandler, s.SessionMiddleware).Name = "delete-role"
	*/

	// Health endpoints.
	healthGrp := s.Echo.Group("/health")
	metricsHandler := fmetrics.NewMetricsHandlers(ctx)
	healthGrp.GET("/alive", metricsHandler.GetHealthAliveHandler).Name = "get-health-alive"

	// Debug endpoints.
	debugGrp := s.Echo.Group("/debug")
	debugGrp.GET("/vars", echo.WrapHandler(http.DefaultServeMux)).Name = "get-debug-vars"
}
