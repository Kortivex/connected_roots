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

	tracingConnectedRootsPostRoleAPI   = "connected-roots-service.http-client: post /roles"
	tracingConnectedRootsPutRoleAPI    = "connected-roots-service.http-client: put /roles/:role_id"
	tracingConnectedRootsGetRoleAPI    = "connected-roots-service.http-client: get /roles/:role_id"
	tracingConnectedRootsGetRolesAPI   = "connected-roots-service.http-client: get /roles"
	tracingConnectedRootsDeleteRoleAPI = "connected-roots-service.http-client: delete /roles/:role_id"
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
	////////////// ROLES //////////////

	POSTRole(ctx context.Context, role *sdk_models.RolesBody) (*resty.Response, error)
	PUTRole(ctx context.Context, role *sdk_models.RolesBody) (*resty.Response, error)
	GETRole(ctx context.Context, id string) (*resty.Response, error)
	GETRoles(ctx context.Context, limit, nexCursor, prevCursor string, names []string) (*resty.Response, error)
	DELETERole(ctx context.Context, id string) (*resty.Response, error)
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
		SetHeader(HeaderContentType, tracingConnectedRootsServicePostUserAPI).
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
