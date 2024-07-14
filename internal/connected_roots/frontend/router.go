package frontend

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/crop_types"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/orchard"
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

	// Login endpoints:
	loginHandler := login.NewLoginHandlers(ctx)
	s.Echo.GET("/login", loginHandler.GetLoginHandler).Name = "get-login"
	s.Echo.POST("/login", loginHandler.PostLoginHandler).Name = "post-login"

	// Logout endpoints:
	logoutHandler := logout.NewLogoutHandlers(ctx)
	s.Echo.GET("/logout", logoutHandler.GetLogoutHandler, s.SessionMiddleware).Name = "get-logout"

	// Users endpoints.
	usersHandler := user.NewUsersHandlers(ctx)
	usersGrp := s.Echo.Group("/users")
	usersGrp.GET("/profile", usersHandler.GetUserProfileHandler, s.SessionMiddleware).Name = "get-user-profile"
	usersGrp.GET("/profile/edit", usersHandler.GetEditUserProfileHandler, s.SessionMiddleware).Name = "get-edit-user-profile"
	usersGrp.POST("/profile/edit", usersHandler.PostEditUserProfileHandler, s.SessionMiddleware).Name = "post-edit-user-profile"

	// Admin endpoints:
	// |
	// |
	adminGrp := s.Echo.Group("/admin")
	// |
	// |
	// + --- Role endpoints:
	rolesHandler := role.NewRolesHandlers(ctx)
	adminRolesGrp := adminGrp.Group("/roles")
	adminRolesGrp.GET("/list", rolesHandler.GetRolesListHandler, s.SessionMiddleware).Name = "get-list-roles"
	adminRolesGrp.GET("/new", rolesHandler.GetRoleCreateHandler, s.SessionMiddleware).Name = "get-new-role"
	adminRolesGrp.POST("/new", rolesHandler.PostRoleCreateHandler, s.SessionMiddleware).Name = "post-new-role"
	adminRolesGrp.GET("/edit/:role_id", rolesHandler.GetRoleUpdateHandler, s.SessionMiddleware).Name = "get-edit-role"
	adminRolesGrp.POST("/edit/:role_id", rolesHandler.PostRoleUpdateHandler, s.SessionMiddleware).Name = "post-edit-role"
	adminRolesGrp.GET("/view/:role_id", rolesHandler.GetRoleViewHandler, s.SessionMiddleware).Name = "get-view-role"
	adminRolesGrp.GET("/delete/:role_id", rolesHandler.GetRoleDeleteHandler, s.SessionMiddleware).Name = "get-delete-role"
	adminRolesGrp.POST("/delete/:role_id", rolesHandler.PostRoleDeleteHandler, s.SessionMiddleware).Name = "post-delete-role"
	// |
	// |
	// + --- User endpoints:
	adminUsersGrp := adminGrp.Group("/users")
	adminUsersGrp.GET("/list", usersHandler.GetUsersListHandler, s.SessionMiddleware).Name = "get-list-users"
	adminUsersGrp.GET("/new", usersHandler.GetUserCreateHandler, s.SessionMiddleware).Name = "get-new-user"
	adminUsersGrp.POST("/new", usersHandler.PostUserCreateHandler, s.SessionMiddleware).Name = "post-new-user"
	adminUsersGrp.GET("/edit/:user_id", usersHandler.GetUserUpdateHandler, s.SessionMiddleware).Name = "get-edit-user"
	adminUsersGrp.POST("/edit/:user_id", usersHandler.PostUserUpdateHandler, s.SessionMiddleware).Name = "post-edit-user"
	adminUsersGrp.GET("/view/:user_id", usersHandler.GetUserViewHandler, s.SessionMiddleware).Name = "get-view-user"
	adminUsersGrp.GET("/delete/:user_id", usersHandler.GetUserDeleteHandler, s.SessionMiddleware).Name = "get-delete-user"
	adminUsersGrp.POST("/delete/:user_id", usersHandler.PostUserDeleteHandler, s.SessionMiddleware).Name = "post-delete-user"
	// |
	// |
	// + --- Orchard endpoints:
	orchardsHandler := orchard.NewOrchardsHandlers(ctx)
	adminOrchardsGrp := adminGrp.Group("/orchards")
	adminOrchardsGrp.GET("/list", orchardsHandler.GetOrchardsListHandler, s.SessionMiddleware).Name = "get-list-orchards"
	adminOrchardsGrp.GET("/new", orchardsHandler.GetOrchardCreateHandler, s.SessionMiddleware).Name = "get-new-orchard"
	adminOrchardsGrp.POST("/new", orchardsHandler.PostOrchardCreateHandler, s.SessionMiddleware).Name = "post-new-orchard"
	adminOrchardsGrp.GET("/edit/:orchard_id", orchardsHandler.GetOrchardUpdateHandler, s.SessionMiddleware).Name = "get-edit-orchard"
	adminOrchardsGrp.POST("/edit/:orchard_id", orchardsHandler.PostOrchardUpdateHandler, s.SessionMiddleware).Name = "post-edit-orchard"
	adminOrchardsGrp.GET("/view/:orchard_id", orchardsHandler.GetOrchardViewHandler, s.SessionMiddleware).Name = "get-view-orchard"
	adminOrchardsGrp.GET("/delete/:orchard_id", orchardsHandler.GetOrchardDeleteHandler, s.SessionMiddleware).Name = "get-delete-orchard"
	adminOrchardsGrp.POST("/delete/:orchard_id", orchardsHandler.PostOrchardDeleteHandler, s.SessionMiddleware).Name = "post-delete-orchard"
	// |
	// |
	// + --- CropTypes endpoints:
	cropTypesHandler := crop_types.NewCropTypesHandlers(ctx)
	adminCropTypesGrp := adminGrp.Group("/crop-types")
	adminCropTypesGrp.GET("/list", cropTypesHandler.GetCropTypesListHandler, s.SessionMiddleware).Name = "get-list-crop-types"
	adminCropTypesGrp.GET("/new", cropTypesHandler.GetCropTypeCreateHandler, s.SessionMiddleware).Name = "get-new-crop-type"
	adminCropTypesGrp.POST("/new", cropTypesHandler.PostCropTypeCreateHandler, s.SessionMiddleware).Name = "post-new-crop-type"
	adminCropTypesGrp.GET("/edit/:crop_type_id", cropTypesHandler.GetCropTypeUpdateHandler, s.SessionMiddleware).Name = "get-edit-crop-type"
	adminCropTypesGrp.POST("/edit/:crop_type_id", cropTypesHandler.PostCropTypeUpdateHandler, s.SessionMiddleware).Name = "post-edit-crop-type"
	adminCropTypesGrp.GET("/view/:crop_type_id", cropTypesHandler.GetCropTypeViewHandler, s.SessionMiddleware).Name = "get-view-crop-type"
	adminCropTypesGrp.GET("/delete/:crop_type_id", cropTypesHandler.GetCropTypeDeleteHandler, s.SessionMiddleware).Name = "get-delete-crop-type"
	adminCropTypesGrp.POST("/delete/:crop_type_id", cropTypesHandler.PostCropTypeDeleteHandler, s.SessionMiddleware).Name = "post-delete-crop-type"

	// Health endpoints.
	healthGrp := s.Echo.Group("/health")
	metricsHandler := fmetrics.NewMetricsHandlers(ctx)
	healthGrp.GET("/alive", metricsHandler.GetHealthAliveHandler).Name = "get-health-alive"

	// Debug endpoints.
	debugGrp := s.Echo.Group("/debug")
	debugGrp.GET("/vars", echo.WrapHandler(http.DefaultServeMux)).Name = "get-debug-vars"
}
