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
	tracingConnectedRootsServiceGetUserAPI            = "connected-roots-service.http-client: get   /users/:user_id"
	tracingConnectedRootsServicePatchUserPartiallyAPI = "connected-roots-service.http-client: patch /users/:user_id"
	traciConnectedRootsPostUserAuthAPI                = "connected-roots-service.http-client: post  /users/:user_id/auth"

	traciConnectedRootsPostRoleAPI   = "connected-roots-service.http-client: post /roles"
	traciConnectedRootsPutRoleAPI    = "connected-roots-service.http-client: put /roles/:role_id"
	traciConnectedRootsGetRoleAPI    = "connected-roots-service.http-client: get /roles/:role_id"
	traciConnectedRootsGetRolesAPI   = "connected-roots-service.http-client: get /roles"
	traciConnectedRootsDeleteRoleAPI = "connected-roots-service.http-client: delete /roles/:role_id"
)

type ConnectedRootsServiceAPI struct {
	Rest   Client
	logger *logger.Logger
}

type IConnectedRootsServiceAPI interface {
	////////////// USERS //////////////

	GETUser(ctx context.Context, userID string) (*resty.Response, error)
	PATCHUserPartially(ctx context.Context, user *sdk_models.UsersBody) (*resty.Response, error)
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
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, traciConnectedRootsPostRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(traciConnectedRootsPostRoleAPI)

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
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, traciConnectedRootsPutRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(traciConnectedRootsPutRoleAPI)

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
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, traciConnectedRootsGetRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(traciConnectedRootsGetRoleAPI)

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
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, traciConnectedRootsGetRolesAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(traciConnectedRootsGetRolesAPI)

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
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, traciConnectedRootsDeleteRoleAPI)
	defer sp.End()

	loggerEmpty := c.logger.New()
	log := loggerEmpty.WithTag(traciConnectedRootsDeleteRoleAPI)

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
