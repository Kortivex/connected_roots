package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Kortivex/connected_roots/pkg/pagination"
	"github.com/Kortivex/connected_roots/pkg/sdk/sdk_models"
	"go.opentelemetry.io/otel"
)

const (
	tracingConnectedRootsServiceObtainUser          = "connected-roots.obtain-user"
	tracingConnectedRootsServiceUpdatePartiallyUser = "connected-roots.update-partially-user"
	tracingConnectedRootsServiceAuthenticateUser    = "connected-roots.authenticate-user"

	tracingConnectedRootsServiceSaveRole    = "connected-roots.save-role"
	tracingConnectedRootsServiceUpdateRole  = "connected-roots.update-role"
	tracingConnectedRootsServiceObtainRole  = "connected-roots.obtain-role"
	tracingConnectedRootsServiceObtainRoles = "connected-roots.obtain-roles"
	tracingConnectedRootsServiceDeleteRole  = "connected-roots.delete-role"

	ErrMsgConnectedRootsServiceAuthenticateUserErr    = "authentication user failure"
	ErrMsgConnectedRootsServiceObtainUserErr          = "obtain user failure"
	ErrMsgConnectedRootsServiceSaveRoleErr            = "saving role failure"
	ErrMsgConnectedRootsServiceUpdateRoleErr          = "updating role failure"
	ErrMsgConnectedRootsServiceObtainRoleErr          = "obtain role failure"
	ErrMsgConnectedRootsServiceObtainRolesErr         = "obtain roles failure"
	ErrMsgConnectedRootsServiceUpdatePartiallyUserErr = "updating partially user failure"
)

type ConnectedRootsServiceSDK struct {
	api ConnectedRootsServiceAPI
}

type IConnectedRootsServiceSDK interface {
	////////////// USERS //////////////

	ObtainUser(ctx context.Context, userID string) (*sdk_models.UsersResponse, error)
	UpdatePartiallyUser(ctx context.Context, user *sdk_models.UsersBody) (*sdk_models.UsersResponse, error)
	AuthenticateUser(ctx context.Context, userID string, authn *sdk_models.UsersAuthenticationBody) (*sdk_models.UsersAuthenticationResponse, error)

	////////////// ROLES //////////////

	SaveRole(ctx context.Context, role *sdk_models.RolesBody) (*sdk_models.RolesResponse, error)
	ObtainRole(ctx context.Context, id string) (*sdk_models.RolesResponse, error)
	ObtainRoles(ctx context.Context, limit, nexCursor, prevCursor string, names []string) ([]*sdk_models.RolesResponse, *pagination.Paging, error)
	DeleteRole(ctx context.Context, id string) error
}

////////////// USERS //////////////

func (c *ConnectedRootsServiceSDK) AuthenticateUser(ctx context.Context, userID string, authn *sdk_models.UsersAuthenticationBody) (*sdk_models.UsersAuthenticationResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceAuthenticateUser)
	defer sp.End()

	resp, err := c.api.POSTUserAuthentication(ctx, userID, authn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceAuthenticateUser, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceAuthenticateUser, resp.Error().(*APIError))
	}

	respAuthn, ok := resp.Result().(*sdk_models.UsersAuthenticationResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceAuthenticateUser, errors.New(ErrMsgConnectedRootsServiceAuthenticateUserErr))
	}

	return respAuthn, nil
}

func (c *ConnectedRootsServiceSDK) ObtainUser(ctx context.Context, userID string) (*sdk_models.UsersResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceObtainUser)
	defer sp.End()

	resp, err := c.api.GETUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUser, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUser, resp.Error().(*APIError))
	}

	respAuthn, ok := resp.Result().(*sdk_models.UsersResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUser, errors.New(ErrMsgConnectedRootsServiceObtainUserErr))
	}

	return respAuthn, nil
}

func (c *ConnectedRootsServiceSDK) UpdatePartiallyUser(ctx context.Context, user *sdk_models.UsersBody) (*sdk_models.UsersResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceUpdatePartiallyUser)
	defer sp.End()

	resp, err := c.api.PATCHUserPartially(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdatePartiallyUser, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdatePartiallyUser, resp.Error().(*APIError))
	}

	respUpdate, ok := resp.Result().(*sdk_models.UsersResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdatePartiallyUser, errors.New(ErrMsgConnectedRootsServiceUpdatePartiallyUserErr))
	}

	return respUpdate, nil
}

////////////// ROLES //////////////

func (c *ConnectedRootsServiceSDK) SaveRole(ctx context.Context, role *sdk_models.RolesBody) (*sdk_models.RolesResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceSaveRole)
	defer sp.End()

	resp, err := c.api.POSTRole(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceSaveRole, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceSaveRole, resp.Error().(*APIError))
	}

	respRole, ok := resp.Result().(*sdk_models.RolesResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceSaveRole, errors.New(ErrMsgConnectedRootsServiceSaveRoleErr))
	}

	return respRole, nil
}

func (c *ConnectedRootsServiceSDK) UpdateRole(ctx context.Context, role *sdk_models.RolesBody) (*sdk_models.RolesResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceUpdateRole)
	defer sp.End()

	resp, err := c.api.PUTRole(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdateRole, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdateRole, resp.Error().(*APIError))
	}

	respRole, ok := resp.Result().(*sdk_models.RolesResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdateRole, errors.New(ErrMsgConnectedRootsServiceUpdateRoleErr))
	}

	return respRole, nil
}

func (c *ConnectedRootsServiceSDK) ObtainRole(ctx context.Context, id string) (*sdk_models.RolesResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceObtainRole)
	defer sp.End()

	resp, err := c.api.GETRole(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRole, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRole, resp.Error().(*APIError))
	}

	respRole, ok := resp.Result().(*sdk_models.RolesResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRole, errors.New(ErrMsgConnectedRootsServiceObtainRoleErr))
	}

	return respRole, nil
}

func (c *ConnectedRootsServiceSDK) ObtainRoles(ctx context.Context, limit, nexCursor, prevCursor string, names []string) ([]*sdk_models.RolesResponse, *pagination.Paging, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceObtainRoles)
	defer sp.End()

	resp, err := c.api.GETRoles(ctx, limit, nexCursor, prevCursor, names)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRoles, err)
	}
	if resp.IsError() {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRoles, resp.Error().(*APIError))
	}

	respRoles, ok := resp.Result().(*pagination.Pagination)
	if !ok {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRoles, errors.New(ErrMsgConnectedRootsServiceObtainRolesErr))
	}

	roles := []*sdk_models.RolesResponse{}
	rolesByte, err := json.Marshal(respRoles.Data)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRoles, err)
	}
	if err = json.Unmarshal(rolesByte, &roles); err != nil {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRoles, err)
	}

	return roles, &respRoles.Paging, nil
}

func (c *ConnectedRootsServiceSDK) DeleteRole(ctx context.Context, id string) error {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceDeleteRole)
	defer sp.End()

	resp, err := c.api.DELETERole(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", tracingConnectedRootsServiceDeleteRole, err)
	}
	if resp.IsError() {
		return fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainRole, resp.Error().(*APIError))
	}

	return nil
}
