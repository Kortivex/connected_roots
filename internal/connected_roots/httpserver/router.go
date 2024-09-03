package httpserver

import (
	"net/http"

	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/activity"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/crop_type"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/orchard"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/role"
	"github.com/Kortivex/connected_roots/internal/connected_roots/httpserver/sensor"

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
	usersGrp.POST("", usersHandler.PostUserHandler).Name = "post-user"
	usersGrp.PUT("/:user_id", usersHandler.PutUserHandler).Name = "put-user"
	usersGrp.PATCH("/:user_id", usersHandler.PatchUserPartiallyHandler).Name = "patch-user"
	usersGrp.GET("/:user_id", usersHandler.GetUserHandler).Name = "get-user"
	usersGrp.GET("", usersHandler.ListUsersHandler).Name = "list-users"
	usersGrp.DELETE("/:user_id", usersHandler.DeleteUserHandler).Name = "delete-user"
	usersGrp.POST("/:user_id/auth", usersHandler.PostUserAuthHandler).Name = "post-user-auth"
	usersGrp.GET("/count", usersHandler.GetCountUsersHandler).Name = "count-users"

	// Roles endpoints.
	rolesHandler := role.NewRolesHandlers(ctx)
	rolesGrp := s.Echo.Group("/roles")
	rolesGrp.POST("", rolesHandler.PostRolesHandler).Name = "post-role"
	rolesGrp.PUT("/:role_id", rolesHandler.PutRolesHandler).Name = "put-role"
	rolesGrp.GET("/:role_id", rolesHandler.GetRolesHandler).Name = "get-role"
	rolesGrp.GET("", rolesHandler.ListRolesHandler).Name = "list-roles"
	rolesGrp.DELETE("/:role_id", rolesHandler.DeleteRolesHandler).Name = "delete-role"
	rolesGrp.GET("/count", rolesHandler.GetCountRolesHandler).Name = "count-roles"

	// Orchard endpoints.
	orchardsHandler := orchard.NewOrchardsHandlers(ctx)
	orchardsGrp := s.Echo.Group("/orchards")
	orchardsGrp.POST("", orchardsHandler.PostOrchardHandler).Name = "post-orchard"
	orchardsGrp.PUT("/:orchard_id", orchardsHandler.PutOrchardHandler).Name = "put-orchard"
	orchardsGrp.GET("/:orchard_id", orchardsHandler.GetOrchardHandler).Name = "get-orchard"
	orchardsGrp.GET("", orchardsHandler.ListOrchardsHandler).Name = "list-orchards"
	orchardsGrp.DELETE("/:orchard_id", orchardsHandler.DeleteOrchardHandler).Name = "delete-orchard"
	orchardsGrp.GET("/count", orchardsHandler.GetCountOrchardsHandler).Name = "count-orchards"
	// User-Orchard endpoints.
	usersGrp.GET("/:user_id/orchards/:orchard_id", orchardsHandler.GetUserOrchardHandler).Name = "get-user-orchard"
	usersGrp.GET("/:user_id/orchards", orchardsHandler.ListUserOrchardsHandler).Name = "get-user-orchards"
	usersGrp.GET("/:user_id/orchards/count", orchardsHandler.GetCountUserOrchardsHandler).Name = "get-user-count-orchards"

	// Crop Types endpoints.
	cropTypesHandler := crop_type.NewCropTypesHandlers(ctx)
	cropTypesGrp := s.Echo.Group("/crop-types")
	cropTypesGrp.POST("", cropTypesHandler.PostCropTypeHandler).Name = "post-crop-type"
	cropTypesGrp.PUT("/:crop_type_id", cropTypesHandler.PutCropTypeHandler).Name = "put-crop-type"
	cropTypesGrp.GET("/:crop_type_id", cropTypesHandler.GetCropTypeHandler).Name = "get-crop-type"
	cropTypesGrp.GET("", cropTypesHandler.ListCropTypesHandler).Name = "list-crop-types"
	cropTypesGrp.DELETE("/:crop_type_id", cropTypesHandler.DeleteCropTypeHandler).Name = "delete-crop-type"
	cropTypesGrp.GET("/count", cropTypesHandler.GetCountCropTypesHandler).Name = "count-crop-types"

	// Sensors endpoints.
	sensorsHandler := sensor.NewSensorsHandlers(ctx)
	sensorsGrp := s.Echo.Group("/sensors")
	sensorsGrp.POST("", sensorsHandler.PostSensorHandler).Name = "post-sensor"
	sensorsGrp.PUT("/:sensor_id", sensorsHandler.PutSensorHandler).Name = "put-sensor"
	sensorsGrp.GET("/:sensor_id", sensorsHandler.GetSensorHandler).Name = "get-sensor"
	sensorsGrp.GET("", sensorsHandler.ListSensorsHandler).Name = "list-sensors"
	sensorsGrp.DELETE("/:sensor_id", sensorsHandler.DeleteSensorHandler).Name = "delete-sensor"
	sensorsGrp.GET("/count", sensorsHandler.GetCountSensorsHandler).Name = "count-sensors"
	// Sensors Data endpoints.
	sensorsGrp.POST("/:sensor_id/data", sensorsHandler.PostSensorDataHandler).Name = "post-sensor-data"
	sensorsGrp.GET("/:sensor_id/data", sensorsHandler.ListSensorsDataHandler).Name = "list-sensors-data"
	sensorsGrp.GET("/:sensor_id/last-data", sensorsHandler.GetSensorLastDataHandler).Name = "get-sensors-latest-data"
	// User-Sensors endpoints.
	usersGrp.GET("/:user_id/sensors", sensorsHandler.ListUserSensorsHandler).Name = "get-user-sensors"
	usersGrp.GET("/:user_id/sensors/count", sensorsHandler.GetCountUserSensorsHandler).Name = "get-user-count-sensors"
	orchardsGrp.GET("/:orchard_id/sensors/average", sensorsHandler.GetSensorWeekdayAverageDataHandler).Name = "get-user-sensor-weekday-average-data"

	// Activities endpoints.
	activitiesHandler := activity.NewActivitiesHandlers(ctx)
	activitiesGrp := usersGrp.Group("/:user_id/activities")
	activitiesGrp.POST("", activitiesHandler.PostActivityHandler).Name = "post-activity"
	activitiesGrp.PUT("/:activity_id", activitiesHandler.PutActivityHandler).Name = "put-activity"
	activitiesGrp.GET("/:activity_id", activitiesHandler.GetActivityHandler).Name = "get-activity"
	activitiesGrp.GET("", activitiesHandler.ListActivitiesHandler).Name = "list-activities"
	activitiesGrp.DELETE("/:activity_id", activitiesHandler.DeleteActivityHandler).Name = "delete-activity"
	s.Echo.GET("/activities/count", activitiesHandler.GetCountActivitiesHandler).Name = "count-activities"
	activitiesGrp.GET("/count", activitiesHandler.GetCountUserActivitiesHandler).Name = "count-user-activities"

	// Health endpoints.
	healthGrp := s.Echo.Group("/health")
	metricsHandler := metrics.NewMetricsHandlers(ctx)
	healthGrp.GET("/alive", metricsHandler.GetHealthAliveHandler).Name = "get-health-alive"

	// Debug endpoints.
	debugGrp := s.Echo.Group("/debug")
	debugGrp.GET("/vars", echo.WrapHandler(http.DefaultServeMux)).Name = "get-debug-vars"
}
