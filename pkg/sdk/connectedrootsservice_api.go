package sdk

import (
	"context"
	"fmt"
	"github.com/Kortivex/connected_roots/pkg/logger"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"github.com/Kortivex/connected_roots/pkg/sdk/sdk_models"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel"
)

const (
	tracingConnectedRootsServicePostUserAPI           = "connected-roots-service.http-client: post /users"
	tracingConnectedRootsServicePutUserAPI            = "connected-roots-service.http-client: put /users/:user_id"
	tracingConnectedRootsServicePatchUserPartiallyAPI = "connected-roots-service.http-client: patch /users/:user_id"
	tracingConnectedRootsServiceGetUserAPI            = "connected-roots-service.http-client: get /users/:user_id"
	tracingConnectedRootsGetUsersAPI                  = "connected-roots-service.http-client: get /users"
	tracingConnectedRootsDeleteUserAPI                = "connected-roots-service.http-client: delete /users/:user_id"
	traciConnectedRootsPostUserAuthAPI                = "connected-roots-service.http-client: post /users/:user_id/auth"
	tracingConnectedRootsServiceGetUserCountAPI       = "connected-roots-service.http-client: get /users/count"

	tracingConnectedRootsPostRoleAPI     = "connected-roots-service.http-client: post /roles"
	tracingConnectedRootsPutRoleAPI      = "connected-roots-service.http-client: put /roles/:role_id"
	tracingConnectedRootsGetRoleAPI      = "connected-roots-service.http-client: get /roles/:role_id"
	tracingConnectedRootsGetRolesAPI     = "connected-roots-service.http-client: get /roles"
	tracingConnectedRootsDeleteRoleAPI   = "connected-roots-service.http-client: delete /roles/:role_id"
	tracingConnectedRootsGetRoleCountAPI = "connected-roots-service.http-client: get /roles/count"

	tracingConnectedRootsPostOrchardAPI         = "connected-roots-service.http-client: post /orchards"
	tracingConnectedRootsPutOrchardAPI          = "connected-roots-service.http-client: put /orchards/:orchard_id"
	tracingConnectedRootsGetOrchardAPI          = "connected-roots-service.http-client: get /orchards/:orchard_id"
	tracingConnectedRootsGetOrchardsAPI         = "connected-roots-service.http-client: get /orchards"
	tracingConnectedRootsDeleteOrchardAPI       = "connected-roots-service.http-client: delete /orchards/:orchard_id"
	tracingConnectedRootsGetOrchardCountAPI     = "connected-roots-service.http-client: get /orchards/count"
	tracingConnectedRootsGetUserOrchardCountAPI = "connected-roots-service.http-client: get /users/:user_id/orchards/count"

	tracingConnectedRootsGetUserOrchardAPI  = "connected-roots-service.http-client: get /users/:user_id/orchards/:orchard_id"
	tracingConnectedRootsGetUserOrchardsAPI = "connected-roots-service.http-client: get /users/:user_id/orchards"

	tracingConnectedRootsPostCropTypeAPI     = "connected-roots-service.http-client: post /crop-types"
	tracingConnectedRootsPutCropTypeAPI      = "connected-roots-service.http-client: put /crop-types/:crop_type_id"
	tracingConnectedRootsGetCropTypeAPI      = "connected-roots-service.http-client: get /crop-types/:crop_type_id"
	tracingConnectedRootsGetCropTypesAPI     = "connected-roots-service.http-client: get /crop-types"
	tracingConnectedRootsDeleteCropTypeAPI   = "connected-roots-service.http-client: delete /crop-types/:crop_type_id"
	tracingConnectedRootsGetCropTypeCountAPI = "connected-roots-service.http-client: get /crop-types/count"

	tracingConnectedRootsPostSensorAPI               = "connected-roots-service.http-client: post /sensors"
	tracingConnectedRootsPutSensorAPI                = "connected-roots-service.http-client: put /sensors/:sensor_id"
	tracingConnectedRootsGetSensorAPI                = "connected-roots-service.http-client: get /sensors/:sensor_id"
	tracingConnectedRootsGetSensorLastDataAPI        = "connected-roots-service.http-client: get /sensors/:sensor_id/last-data"
	tracingConnectedRootsGetSensorWeekDataAverageAPI = "connected-roots-service.http-client: get /orchards/:orchard_id/sensors/average"
	tracingConnectedRootsGetSensorsAPI               = "connected-roots-service.http-client: get /sensors"
	tracingConnectedRootsDeleteSensorAPI             = "connected-roots-service.http-client: delete /sensors/:sensor_id"
	tracingConnectedRootsGetUserSensorsAPI           = "connected-roots-service.http-client: get /users/:user_id/sensors"
	tracingConnectedRootsGetSensorCountAPI           = "connected-roots-service.http-client: get /sensors/count"
	tracingConnectedRootsGetUserSensorCountAPI       = "connected-roots-service.http-client: get /users/:user_id/sensors/count"

	tracingConnectedRootsPostActivityAPI         = "connected-roots-service.http-client: post /users/:user_id/activities"
	tracingConnectedRootsPutActivityAPI          = "connected-roots-service.http-client: put /users/:user_id/activities/:activity_id"
	tracingConnectedRootsGetActivityAPI          = "connected-roots-service.http-client: get /users/:user_id/activities/:activity_id"
	tracingConnectedRootsGetActivitiesAPI        = "connected-roots-service.http-client: get /users/:user_id/activities"
	tracingConnectedRootsDeleteActivityAPI       = "connected-roots-service.http-client: delete /users/:user_id/activities/:activity_id"
	tracingConnectedRootsGetActivityCountAPI     = "connected-roots-service.http-client: get /activities/count"
	tracingConnectedRootsGetUserActivityCountAPI = "connected-roots-service.http-client: get /users/:user_id/activities/count"
)

type ConnectedRootsServiceAPI struct {
	Rest   Client
	logger *logger.Logger
}

type IConnectedRootsServiceAPI interface {
	////////////// USERS //////////////

	POSTUser(ctx context.Context, user *sdk_models.UsersBody) (*resty.Response, error)
	PUTUser(ctx context.Context, user *sdk_models.UsersBody) (*resty.Response, error)
	GETUser(ctx context.Context, userID string) (*resty.Response, error)
	GETUsers(ctx context.Context, limit, nexCursor, prevCursor string, names, surnames, emails []string) (*resty.Response, error)
	PATCHUserPartially(ctx context.Context, user *sdk_models.UsersBody) (*resty.Response, error)
	DELETEUser(ctx context.Context, id string) (*resty.Response, error)
	POSTUserAuthentication(ctx context.Context, userID string, authn *sdk_models.UsersAuthenticationBody) (*resty.Response, error)
	GETUserCount(ctx context.Context) (*resty.Response, error)

	////////////// ROLES //////////////

	POSTRole(ctx context.Context, role *sdk_models.RolesBody) (*resty.Response, error)
	PUTRole(ctx context.Context, role *sdk_models.RolesBody) (*resty.Response, error)
	GETRole(ctx context.Context, id string) (*resty.Response, error)
	GETRoles(ctx context.Context, limit, nexCursor, prevCursor string, names []string) (*resty.Response, error)
	DELETERole(ctx context.Context, id string) (*resty.Response, error)
	GETRoleCount(ctx context.Context) (*resty.Response, error)

	////////////// ORCHARDS //////////////

	POSTOrchard(ctx context.Context, orchard *sdk_models.OrchardsBody) (*resty.Response, error)
	PUTOrchard(ctx context.Context, orchard *sdk_models.OrchardsBody) (*resty.Response, error)
	GETOrchard(ctx context.Context, id string) (*resty.Response, error)
	DELETEOrchard(ctx context.Context, id string) (*resty.Response, error)
	GETOrchardCount(ctx context.Context) (*resty.Response, error)
	GETUserOrchardCount(ctx context.Context, userID string) (*resty.Response, error)

	////////////// USERS - ORCHARDS //////////////

	GETUserOrchard(ctx context.Context, userID, id string) (*resty.Response, error)
	GETUserOrchards(ctx context.Context, userID, limit, nexCursor, prevCursor string, names, locations []string) (*resty.Response, error)

	////////////// CROP TYPES //////////////

	POSTCropType(ctx context.Context, cropType *sdk_models.CropTypesBody) (*resty.Response, error)
	PUTCropType(ctx context.Context, cropType *sdk_models.CropTypesBody) (*resty.Response, error)
	GETCropType(ctx context.Context, id string) (*resty.Response, error)
	GETCropTypes(ctx context.Context, limit, nexCursor, prevCursor string, names, scientificNames, plantingSeasons, harvestSeasons []string) (*resty.Response, error)
	DELETECropType(ctx context.Context, id string) (*resty.Response, error)
	GETCropTypeCount(ctx context.Context) (*resty.Response, error)

	////////////// SENSORS //////////////

	POSTSensor(ctx context.Context, sensor *sdk_models.SensorsBody) (*resty.Response, error)
	PUTSensor(ctx context.Context, sensor *sdk_models.SensorsBody) (*resty.Response, error)
	GETSensor(ctx context.Context, id string) (*resty.Response, error)
	GETSensors(ctx context.Context, limit, nexCursor, prevCursor string, names, firmwareVersions, manufacturers, batteryLifes, statuses []string) (*resty.Response, error)
	DELETESensor(ctx context.Context, id string) (*resty.Response, error)
	GETSensorCount(ctx context.Context) (*resty.Response, error)
	GETUserSensorCount(ctx context.Context, userID string) (*resty.Response, error)

	////////////// USERS - SENSORS //////////////

	GETUserSensors(ctx context.Context, userID, limit, nexCursor, prevCursor string, names, firmwareVersions, manufacturers, batteryLifes, statuses []string) (*resty.Response, error)

	////////////// ACTIVITIES //////////////

	POSTActivity(ctx context.Context, userID string, activity *sdk_models.ActivitiesBody) (*resty.Response, error)
	PUTActivity(ctx context.Context, userID string, activity *sdk_models.ActivitiesBody) (*resty.Response, error)
	GETActivity(ctx context.Context, userID, id string) (*resty.Response, error)
	GETActivities(ctx context.Context, userID, limit, nexCursor, prevCursor string, names, orchardIDs []string) (*resty.Response, error)
	DELETEActivity(ctx context.Context, userID, id string) (*resty.Response, error)
	GETActivityCount(ctx context.Context) (*resty.Response, error)
	GETUserActivityCount(ctx context.Context, userID string) (*resty.Response, error)
}

func NewConnectedRootsClient(host string, client *resty.Client, logr *logger.Logger) *ConnectedRootsService {
	connectedRootsService := Client{
		Host:   host,
		Client: client.SetBaseURL(host),
		logger: logr,
	}

	return &ConnectedRootsService{
		API: ConnectedRootsServiceAPI{Rest: connectedRootsService, logger: connectedRootsService.logger},
		SDK: ConnectedRootsServiceSDK{api: ConnectedRootsServiceAPI{Rest: connectedRootsService, logger: connectedRootsService.logger}},
	}
}

////////////// USERS //////////////

func (c *ConnectedRootsServiceAPI) POSTUser(ctx context.Context, user *sdk_models.UsersBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServicePostUserAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsServicePostUserAPI)

	log.Debug("request [POST] /users")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(user).
		SetResult(&sdk_models.UsersResponse{}).
		SetError(&APIError{}).
		Post("/users")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) PUTUser(ctx context.Context, user *sdk_models.UsersBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServicePutUserAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsServicePutUserAPI)

	log.Debug("request [PUT] /users/:user_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(user).
		SetResult(&sdk_models.UsersResponse{}).
		SetError(&APIError{}).
		Put(fmt.Sprintf("/users/%s", user.ID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETUser(ctx context.Context, userID string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceGetUserAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsServiceGetUserAPI)

	log.Debug(fmt.Sprintf("request [GET] /users/%s", userID))

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.UsersResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETUsers(ctx context.Context, limit, nexCursor, prevCursor string, names, surnames, emails []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetUsersAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetUsersAPI)

	log.Debug("request [GET] /users")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	for _, surname := range surnames {
		request.SetQueryParam("surname[]", surname)
	}

	for _, email := range emails {
		request.SetQueryParam("email[]", email)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get("/users")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) PATCHUserPartially(ctx context.Context, user *sdk_models.UsersBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServicePatchUserPartiallyAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsServicePatchUserPartiallyAPI)

	log.Debug(fmt.Sprintf("request [POST] /users/%s with body: %v", user.Email, user))

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(user).
		SetResult(&sdk_models.UsersResponse{}).
		SetError(&APIError{}).
		Patch(fmt.Sprintf("/users/%s", user.Email))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) DELETEUser(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsDeleteUserAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsDeleteUserAPI)

	log.Debug("request [DELETE] /users/:user_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetError(&APIError{}).
		Delete(fmt.Sprintf("/users/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) POSTUserAuthentication(ctx context.Context, userID string, authn *sdk_models.UsersAuthenticationBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, traciConnectedRootsPostUserAuthAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(traciConnectedRootsPostUserAuthAPI)

	log.Debug(fmt.Sprintf("request [POST] /users/%s/auth with body: %v", userID, authn))

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(authn).
		SetResult(&sdk_models.UsersAuthenticationResponse{}).
		SetError(&APIError{}).
		Post(fmt.Sprintf("/users/%s/auth", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETUserCount(ctx context.Context) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceGetUserCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsServiceGetUserCountAPI)

	log.Debug("request [GET] /users/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalUsersResponse{}).
		SetError(&APIError{}).
		Get("/users/count")
	if err != nil {
		return nil, err
	}
	return response, nil
}

////////////// ROLES //////////////

func (c *ConnectedRootsServiceAPI) POSTRole(ctx context.Context, role *sdk_models.RolesBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPostRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPostRoleAPI)

	log.Debug("request [POST] /roles")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(role).
		SetResult(&sdk_models.RolesResponse{}).
		SetError(&APIError{}).
		Post("/roles")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) PUTRole(ctx context.Context, role *sdk_models.RolesBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPutRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPutRoleAPI)

	log.Debug("request [PUT] /roles/:role_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(role).
		SetResult(&sdk_models.RolesResponse{}).
		SetError(&APIError{}).
		Put(fmt.Sprintf("/roles/%s", role.ID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETRole(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetRoleAPI)

	log.Debug("request [GET] /roles/:role_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.RolesResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/roles/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETRoles(ctx context.Context, limit, nexCursor, prevCursor string, names []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetRolesAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetRolesAPI)

	log.Debug("request [GET] /roles")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get("/roles")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) DELETERole(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsDeleteRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsDeleteRoleAPI)

	log.Debug("request [DELETE] /roles/:role_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetError(&APIError{}).
		Delete(fmt.Sprintf("/roles/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETRoleCount(ctx context.Context) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetRoleCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetRoleCountAPI)

	log.Debug("request [GET] /roles/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalRolesResponse{}).
		SetError(&APIError{}).
		Get("/roles/count")
	if err != nil {
		return nil, err
	}
	return response, nil
}

////////////// ORCHARDS //////////////

func (c *ConnectedRootsServiceAPI) POSTOrchard(ctx context.Context, orchard *sdk_models.OrchardsBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPostOrchardAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPostOrchardAPI)

	log.Debug("request [POST] /orchards")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(orchard).
		SetResult(&sdk_models.OrchardsResponse{}).
		SetError(&APIError{}).
		Post("/orchards")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) PUTOrchard(ctx context.Context, orchard *sdk_models.OrchardsBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPutOrchardAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPutOrchardAPI)

	log.Debug("request [PUT] /orchards/:orchard_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(orchard).
		SetResult(&sdk_models.OrchardsResponse{}).
		SetError(&APIError{}).
		Put(fmt.Sprintf("/orchards/%s", orchard.ID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETOrchard(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetOrchardAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetOrchardAPI)

	log.Debug("request [GET] /orchards/:orchard_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.OrchardsResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/orchards/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETOrchards(ctx context.Context, limit, nexCursor, prevCursor string, names, locations, userIDs []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetOrchardsAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetOrchardsAPI)

	log.Debug("request [GET] /orchards")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	for _, location := range locations {
		request.SetQueryParam("location[]", location)
	}

	for _, userID := range userIDs {
		request.SetQueryParam("user_id[]", userID)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get("/orchards")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) DELETEOrchard(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsDeleteOrchardAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsDeleteOrchardAPI)

	log.Debug("request [DELETE] /orchards/:orchard_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetError(&APIError{}).
		Delete(fmt.Sprintf("/orchards/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETOrchardCount(ctx context.Context) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetOrchardCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetOrchardCountAPI)

	log.Debug("request [GET] /orchards/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalOrchardsResponse{}).
		SetError(&APIError{}).
		Get("/orchards/count")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETUserOrchardCount(ctx context.Context, userID string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetUserOrchardCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetUserOrchardCountAPI)

	log.Debug("request [GET] /users/:user_id/orchards/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalOrchardsResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/orchards/count", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

////////////// USERS - ORCHARDS //////////////

func (c *ConnectedRootsServiceAPI) GETUserOrchard(ctx context.Context, userID, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetUserOrchardAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetUserOrchardAPI)

	log.Debug("request [GET] /users/:user_id/orchards/:orchard_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.OrchardsResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/orchards/%s", userID, id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETUserOrchards(ctx context.Context, userID, limit, nexCursor, prevCursor string, names, locations []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetUserOrchardsAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetUserOrchardsAPI)

	log.Debug("request [GET] /users/:user_id/orchards")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	for _, location := range locations {
		request.SetQueryParam("location[]", location)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/orchards", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

////////////// CROP TYPES //////////////

func (c *ConnectedRootsServiceAPI) POSTCropType(ctx context.Context, cropType *sdk_models.CropTypesBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPostCropTypeAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPostCropTypeAPI)

	log.Debug("request [POST] /crop-types")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(cropType).
		SetResult(&sdk_models.CropTypesResponse{}).
		SetError(&APIError{}).
		Post("/crop-types")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) PUTCropType(ctx context.Context, cropType *sdk_models.CropTypesBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPutCropTypeAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPutCropTypeAPI)

	log.Debug("request [PUT] /crop-types/:crop_type_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(cropType).
		SetResult(&sdk_models.CropTypesResponse{}).
		SetError(&APIError{}).
		Put(fmt.Sprintf("/crop-types/%s", cropType.ID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETCropType(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetCropTypeAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetCropTypeAPI)

	log.Debug("request [GET] /crop-types/:crop_type_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.CropTypesResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/crop-types/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETCropTypes(ctx context.Context, limit, nexCursor, prevCursor string, names, scientificNames, plantingSeasons, harvestSeasons []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetCropTypesAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetCropTypesAPI)

	log.Debug("request [GET] /crop-types")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	for _, scientificName := range scientificNames {
		request.SetQueryParam("scientific_name[]", scientificName)
	}

	for _, plantingSeason := range plantingSeasons {
		request.SetQueryParam("planting_season[]", plantingSeason)
	}

	for _, harvestSeason := range harvestSeasons {
		request.SetQueryParam("harvest_season[]", harvestSeason)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get("/crop-types")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) DELETECropType(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsDeleteCropTypeAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsDeleteCropTypeAPI)

	log.Debug("request [DELETE] /crop-types/:crop_type_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetError(&APIError{}).
		Delete(fmt.Sprintf("/crop-types/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETCropTypeCount(ctx context.Context) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetCropTypeCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetCropTypeCountAPI)

	log.Debug("request [GET] /crop-types/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalCropTypesResponse{}).
		SetError(&APIError{}).
		Get("/crop-types/count")
	if err != nil {
		return nil, err
	}
	return response, nil
}

////////////// SENSORS //////////////

func (c *ConnectedRootsServiceAPI) POSTSensor(ctx context.Context, sensor *sdk_models.SensorsBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPostSensorAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPostSensorAPI)

	log.Debug("request [POST] /sensors")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(sensor).
		SetResult(&sdk_models.SensorsResponse{}).
		SetError(&APIError{}).
		Post("/sensors")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) PUTSensor(ctx context.Context, sensor *sdk_models.SensorsBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPutSensorAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPutSensorAPI)

	log.Debug("request [PUT] /sensors/:sensor_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(sensor).
		SetResult(&sdk_models.SensorsResponse{}).
		SetError(&APIError{}).
		Put(fmt.Sprintf("/sensors/%s", sensor.ID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETSensor(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetSensorAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetSensorAPI)

	log.Debug("request [GET] /sensors/:sensor_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.SensorsResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/sensors/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETSensors(ctx context.Context, limit, nexCursor, prevCursor string, names, firmwareVersions, manufacturers, batteryLifes, statuses []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetSensorsAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetSensorsAPI)

	log.Debug("request [GET] /sensors")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	for _, firmwareVersion := range firmwareVersions {
		request.SetQueryParam("firmware_version[]", firmwareVersion)
	}

	for _, manufacturer := range manufacturers {
		request.SetQueryParam("manufacturer[]", manufacturer)
	}

	for _, batteryLife := range batteryLifes {
		request.SetQueryParam("battery_life[]", batteryLife)
	}

	for _, status := range statuses {
		request.SetQueryParam("status[]", status)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get("/sensors")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) DELETESensor(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsDeleteSensorAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsDeleteSensorAPI)

	log.Debug("request [DELETE] /sensors/:sensor_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetError(&APIError{}).
		Delete(fmt.Sprintf("/sensors/%s", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETSensorCount(ctx context.Context) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetSensorCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetSensorCountAPI)

	log.Debug("request [GET] /sensors/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalSensorsResponse{}).
		SetError(&APIError{}).
		Get("/sensors/count")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETUserSensorCount(ctx context.Context, userID string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetUserSensorCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetUserSensorCountAPI)

	log.Debug("request [GET] /users/:user_id/sensors/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalSensorsResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/sensors/count", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETSensorLastData(ctx context.Context, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetSensorLastDataAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetSensorLastDataAPI)

	log.Debug("request [GET] /sensors/:sensor_id/last-data")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.SensorsDataResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/sensors/%s/last-data", id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETSensorWeekDataAverage(ctx context.Context, orchardID string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetSensorWeekDataAverageAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetSensorWeekDataAverageAPI)

	log.Debug("request [GET] /orchards/:orchard_id/sensors/average")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult([]*sdk_models.SensorsDataWeekdayAverageResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/orchards/%s/sensors/average", orchardID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

////////////// USERS - SENSORS //////////////

func (c *ConnectedRootsServiceAPI) GETUserSensors(ctx context.Context, userID, limit, nexCursor, prevCursor string, names, firmwareVersions, manufacturers, batteryLifes, statuses []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetUserSensorsAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetUserSensorsAPI)

	log.Debug("request [GET] /users/:user_id/sensors")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	for _, firmwareVersion := range firmwareVersions {
		request.SetQueryParam("firmware_version[]", firmwareVersion)
	}

	for _, manufacturer := range manufacturers {
		request.SetQueryParam("manufacturer[]", manufacturer)
	}

	for _, batteryLife := range batteryLifes {
		request.SetQueryParam("battery_life[]", batteryLife)
	}

	for _, status := range statuses {
		request.SetQueryParam("status[]", status)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/sensors", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

////////////// ACTIVITIES //////////////

func (c *ConnectedRootsServiceAPI) POSTActivity(ctx context.Context, userID string, activity *sdk_models.ActivitiesBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPostActivityAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPostActivityAPI)

	log.Debug("request [POST] /users/:user_id/activities")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(activity).
		SetResult(&sdk_models.ActivitiesResponse{}).
		SetError(&APIError{}).
		Post(fmt.Sprintf("/users/%s/activities", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) PUTActivity(ctx context.Context, userID string, activity *sdk_models.ActivitiesBody) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsPutActivityAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsPutActivityAPI)

	log.Debug("request [PUT] /users/:user_id/activities/:activity_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetHeader(HeaderContentType, ContentTypeApplicationJSON).
		SetBody(activity).
		SetResult(&sdk_models.ActivitiesResponse{}).
		SetError(&APIError{}).
		Put(fmt.Sprintf("/users/%s/activities/%s", userID, activity.ID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETActivity(ctx context.Context, userID, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetActivityAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetActivityAPI)

	log.Debug("request [GET] /users/:user_id/activities/:activity_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.ActivitiesResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/activities/%s", userID, id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETActivities(ctx context.Context, userID, limit, nexCursor, prevCursor string, names, orchardIDs []string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetActivitiesAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetActivitiesAPI)

	log.Debug("request [GET] /users/:user_id/activities")

	request := c.Rest.Client.R()

	if limit != "" {
		request.SetQueryParam("limit", limit)
	}

	if nexCursor != "" {
		request.SetQueryParam("next_cursor", nexCursor)
	}

	if prevCursor != "" {
		request.SetQueryParam("previous_cursor", prevCursor)
	}

	for _, name := range names {
		request.SetQueryParam("name[]", name)
	}

	for _, orchardID := range orchardIDs {
		request.SetQueryParam("orchard_id[]", orchardID)
	}

	response, err := request.
		SetContext(ctx).
		SetResult(&pagination.Pagination{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/activities", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) DELETEActivity(ctx context.Context, userID, id string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsDeleteActivityAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsDeleteActivityAPI)

	log.Debug("request [DELETE] /users/:user_id/activities/:activity_id")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetError(&APIError{}).
		Delete(fmt.Sprintf("/users/%s/activities/%s", userID, id))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETActivityCount(ctx context.Context) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetActivityCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetActivityCountAPI)

	log.Debug("request [GET] /activities/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalActivitiesResponse{}).
		SetError(&APIError{}).
		Get("/activities/count")
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *ConnectedRootsServiceAPI) GETUserActivityCount(ctx context.Context, userID string) (*resty.Response, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsGetUserActivityCountAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(tracingConnectedRootsGetUserActivityCountAPI)

	log.Debug("request [GET] /users/:user_id/activities/count")

	request := c.Rest.Client.R()
	response, err := request.
		SetContext(ctx).
		SetResult(&sdk_models.TotalActivitiesResponse{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/users/%s/activities/count", userID))
	if err != nil {
		return nil, err
	}
	return response, nil
}
