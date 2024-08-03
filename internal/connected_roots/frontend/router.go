package frontend

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/activity"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/crop_types"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/orchard"
	"github.com/Kortivex/connected_roots/internal/connected_roots/frontend/sensor"
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

	// USERS endpoints.
	usersHandler := user.NewUsersHandlers(ctx)
	usersGrp := s.Echo.Group("/users")
	usersGrp.GET("/profile", usersHandler.GetUserProfileHandler, s.SessionMiddleware).Name = "get-user-profile"
	usersGrp.GET("/profile/edit", usersHandler.GetEditUserProfileHandler, s.SessionMiddleware).Name = "get-edit-user-profile"
	usersGrp.POST("/profile/edit", usersHandler.PostEditUserProfileHandler, s.SessionMiddleware).Name = "post-edit-user-profile"
	// |
	// |
	// + --- Activity endpoints:
	activitiesHandler := activity.NewActivitiesHandlers(ctx)
	userActivitiesGrp := usersGrp.Group("/activities")
	userActivitiesGrp.GET("/list", activitiesHandler.GetActivitiesListHandler, s.SessionMiddleware).Name = "get-list-activities"
	userActivitiesGrp.GET("/new", activitiesHandler.GetActivityCreateHandler, s.SessionMiddleware).Name = "get-new-activity"
	userActivitiesGrp.POST("/new", activitiesHandler.PostActivityCreateHandler, s.SessionMiddleware).Name = "post-new-activity"
	userActivitiesGrp.GET("/edit/:activity_id", activitiesHandler.GetActivityUpdateHandler, s.SessionMiddleware).Name = "get-edit-activity"
	userActivitiesGrp.POST("/edit/:activity_id", activitiesHandler.PostActivityUpdateHandler, s.SessionMiddleware).Name = "post-edit-activity"
	userActivitiesGrp.GET("/view/:activity_id", activitiesHandler.GetActivityViewHandler, s.SessionMiddleware).Name = "get-view-activity"
	userActivitiesGrp.GET("/delete/:activity_id", activitiesHandler.GetActivityDeleteHandler, s.SessionMiddleware).Name = "get-delete-activity"
	userActivitiesGrp.POST("/delete/:activity_id", activitiesHandler.PostActivityDeleteHandler, s.SessionMiddleware).Name = "post-delete-activity"
	// |
	// |
	// + --- Orchard endpoints:
	userOrchardsHandler := orchard.NewOrchardsHandlers(ctx)
	userOrchardsGrp := usersGrp.Group("/orchards")
	userOrchardsGrp.GET("/list", userOrchardsHandler.GetUserOrchardsListHandler, s.SessionMiddleware).Name = "get-list-user-orchards"
	userOrchardsGrp.GET("/view/:orchard_id", userOrchardsHandler.GetUserOrchardViewHandler, s.SessionMiddleware).Name = "get-view-user-orchard"
	userOrchardsGrp.GET("/report/:orchard_id", userOrchardsHandler.GetUserOrchardReportHandler, s.SessionMiddleware).Name = "get-report-orchard"
	// |
	// |
	// + --- Sensor endpoints:
	userSensorsHandler := sensor.NewSensorsHandlers(ctx)
	userSensorsGrp := usersGrp.Group("/sensors")
	userSensorsGrp.GET("/list", userSensorsHandler.GetUserSensorsListHandler, s.SessionMiddleware).Name = "get-list-user-sensors"
	userSensorsGrp.GET("/view/:sensor_id/data", userSensorsHandler.GetUserSensorDataViewHandler, s.SessionMiddleware).Name = "get-view-user-sensor-data"

	// ADMIN endpoints:
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
	adminOrchardsGrp.GET("/report/:orchard_id", orchardsHandler.GetOrchardReportHandler, s.SessionMiddleware).Name = "get-report-orchard"
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
	// |
	// |
	// + --- Sensors endpoints:
	sensorsHandler := sensor.NewSensorsHandlers(ctx)
	adminSensorsGrp := adminGrp.Group("/sensors")
	adminSensorsGrp.GET("/list", sensorsHandler.GetSensorsListHandler, s.SessionMiddleware).Name = "get-list-sensors"
	adminSensorsGrp.GET("/new", sensorsHandler.GetSensorCreateHandler, s.SessionMiddleware).Name = "get-new-sensor"
	adminSensorsGrp.POST("/new", sensorsHandler.PostSensorCreateHandler, s.SessionMiddleware).Name = "post-new-sensor"
	adminSensorsGrp.GET("/edit/:sensor_id", sensorsHandler.GetSensorUpdateHandler, s.SessionMiddleware).Name = "get-edit-sensor"
	adminSensorsGrp.POST("/edit/:sensor_id", sensorsHandler.PostSensorUpdateHandler, s.SessionMiddleware).Name = "post-edit-sensor"
	adminSensorsGrp.GET("/view/:sensor_id", sensorsHandler.GetSensorViewHandler, s.SessionMiddleware).Name = "get-view-sensor"
	adminSensorsGrp.GET("/delete/:sensor_id", sensorsHandler.GetSensorDeleteHandler, s.SessionMiddleware).Name = "get-delete-sensor"
	adminSensorsGrp.POST("/delete/:sensor_id", sensorsHandler.PostSensorDeleteHandler, s.SessionMiddleware).Name = "post-delete-sensor"
	adminSensorsGrp.GET("/view/:sensor_id/data", sensorsHandler.GetSensorDataViewHandler, s.SessionMiddleware).Name = "get-view-sensor-data"

	// Health endpoints.
	healthGrp := s.Echo.Group("/health")
	metricsHandler := fmetrics.NewMetricsHandlers(ctx)
	healthGrp.GET("/alive", metricsHandler.GetHealthAliveHandler).Name = "get-health-alive"

	// Debug endpoints.
	debugGrp := s.Echo.Group("/debug")
	debugGrp.GET("/vars", echo.WrapHandler(http.DefaultServeMux)).Name = "get-debug-vars"
}
