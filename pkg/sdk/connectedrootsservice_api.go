package sdk

import (
	"context"
	"fmt"

	"github.com/Kortivex/connected_roots/pkg/sdk/sdk_models"

	"github.com/Kortivex/connected_roots/pkg/logger"
	"go.opentelemetry.io/otel"

	"github.com/go-resty/resty/v2"
)

const (
	tracingConnectedRootsServiceGetUserAPI = "connected-roots-service.http-client: get /users/:user_id"
	traciConnectedRootsPostUserAuthAPI     = "connected-roots-service.http-client: post /users/:user_id/auth"
)

type ConnectedRootsServiceAPI struct {
	Rest   Client
	logger *logger.Logger
}

type IConnectedRootsServiceAPI interface {
	GETUser(ctx context.Context, userID string) (*resty.Response, error)
	POSTUserAuthentication(ctx context.Context, userID string, authn *sdk_models.UsersAuthenticationBody) (*resty.Response, error)
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
