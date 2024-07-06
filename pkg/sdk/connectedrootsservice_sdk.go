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
	tracingConnectedRootsServiceSaveUser            = "connected-roots.save-user"
	tracingConnectedRootsServiceUpdateUser          = "connected-roots.update-user"
	tracingConnectedRootsServiceUpdatePartiallyUser = "connected-roots.update-partially-user"
	tracingConnectedRootsServiceObtainUser          = "connected-roots.obtain-user"
	tracingConnectedRootsServiceObtainUsers         = "connected-roots.obtain-users"
	tracingConnectedRootsServiceDeleteUser          = "connected-roots.delete-user"
	tracingConnectedRootsServiceAuthenticateUser    = "connected-roots.authenticate-user"

	tracingConnectedRootsServiceSaveRole    = "connected-roots.save-role"
	tracingConnectedRootsServiceUpdateRole  = "connected-roots.update-role"
	tracingConnectedRootsServiceObtainRole  = "connected-roots.obtain-role"
	tracingConnectedRootsServiceObtainRoles = "connected-roots.obtain-roles"
	tracingConnectedRootsServiceDeleteRole  = "connected-roots.delete-role"

	ErrMsgConnectedRootsServiceSaveRoleErr    = "saving role failure"
	ErrMsgConnectedRootsServiceUpdateRoleErr  = "updating role failure"
	ErrMsgConnectedRootsServiceObtainRoleErr  = "obtain role failure"
	ErrMsgConnectedRootsServiceObtainRolesErr = "obtain roles failure"

	ErrMsgConnectedRootsServiceSaveUserErr            = "saving user failure"
	ErrMsgConnectedRootsServiceUpdateUserErr          = "updating user failure"
	ErrMsgConnectedRootsServiceUpdatePartiallyUserErr = "updating partially user failure"
	ErrMsgConnectedRootsServiceObtainUserErr          = "obtain user failure"
	ErrMsgConnectedRootsServiceObtainUsersErr         = "obtain users failure"
	ErrMsgConnectedRootsServiceAuthenticateUserErr    = "authentication user failure"
)

type ConnectedRootsServiceSDK struct {
	api ConnectedRootsServiceAPI
}

type IConnectedRootsServiceSDK interface {
	////////////// USERS //////////////

	SaveUser(ctx context.Context, user *sdk_models.UsersBody) (*sdk_models.UsersResponse, error)
	UpdateUser(ctx context.Context, user *sdk_models.UsersBody) (*sdk_models.UsersResponse, error)
	UpdatePartiallyUser(ctx context.Context, user *sdk_models.UsersBody) (*sdk_models.UsersResponse, error)
	ObtainUser(ctx context.Context, userID string) (*sdk_models.UsersResponse, error)
	ObtainUsers(ctx context.Context, limit, nexCursor, prevCursor string, names, surnames, emails []string) ([]*sdk_models.UsersResponse, *pagination.Paging, error)
	DeleteUser(ctx context.Context, id string) error
	AuthenticateUser(ctx context.Context, userID string, authn *sdk_models.UsersAuthenticationBody) (*sdk_models.UsersAuthenticationResponse, error)

	////////////// ROLES //////////////

	SaveRole(ctx context.Context, role *sdk_models.RolesBody) (*sdk_models.RolesResponse, error)
	UpdateRole(ctx context.Context, role *sdk_models.RolesBody) (*sdk_models.RolesResponse, error)
	ObtainRole(ctx context.Context, id string) (*sdk_models.RolesResponse, error)
	ObtainRoles(ctx context.Context, limit, nexCursor, prevCursor string, names []string) ([]*sdk_models.RolesResponse, *pagination.Paging, error)
	DeleteRole(ctx context.Context, id string) error
}

////////////// USERS //////////////

func (c *ConnectedRootsServiceSDK) SaveUser(ctx context.Context, user *sdk_models.UsersBody) (*sdk_models.UsersResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceSaveUser)
	defer sp.End()

	resp, err := c.api.POSTUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceSaveUser, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceSaveUser, resp.Error().(*APIError))
	}

	respUser, ok := resp.Result().(*sdk_models.UsersResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceSaveUser, errors.New(ErrMsgConnectedRootsServiceSaveUserErr))
	}

	return respUser, nil
}

func (c *ConnectedRootsServiceSDK) UpdateUser(ctx context.Context, user *sdk_models.UsersBody) (*sdk_models.UsersResponse, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceUpdateUser)
	defer sp.End()

	resp, err := c.api.PUTUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdateUser, err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdateUser, resp.Error().(*APIError))
	}

	respUser, ok := resp.Result().(*sdk_models.UsersResponse)
	if !ok {
		return nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceUpdateUser, errors.New(ErrMsgConnectedRootsServiceUpdateUserErr))
	}

	return respUser, nil
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

func (c *ConnectedRootsServiceSDK) ObtainUsers(ctx context.Context, limit, nexCursor, prevCursor string, names, surnames, emails []string) ([]*sdk_models.UsersResponse, *pagination.Paging, error) {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceObtainUsers)
	defer sp.End()

	resp, err := c.api.GETUsers(ctx, limit, nexCursor, prevCursor, names, surnames, emails)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUsers, err)
	}
	if resp.IsError() {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUsers, resp.Error().(*APIError))
	}

	respUsers, ok := resp.Result().(*pagination.Pagination)
	if !ok {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUsers, errors.New(ErrMsgConnectedRootsServiceObtainUsersErr))
	}

	users := []*sdk_models.UsersResponse{}
	usersByte, err := json.Marshal(respUsers.Data)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUsers, err)
	}
	if err = json.Unmarshal(usersByte, &users); err != nil {
		return nil, nil, fmt.Errorf("%s: %w", tracingConnectedRootsServiceObtainUsers, err)
	}

	return users, &respUsers.Paging, nil
}

func (c *ConnectedRootsServiceSDK) DeleteUser(ctx context.Context, id string) error {
	ctx, sp := otel.Tracer("connected_roots").Start(ctx, tracingConnectedRootsServiceDeleteUser)
	defer sp.End()

	resp, err := c.api.DELETEUser(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", tracingConnectedRootsServiceDeleteUser, err)
	}
	if resp.IsError() {
		return fmt.Errorf("%s: %w", tracingConnectedRootsServiceDeleteUser, resp.Error().(*APIError))
	}

	return nil
}

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
		return fmt.Errorf("%s: %w", tracingConnectedRootsServiceDeleteRole, resp.Error().(*APIError))
	}

	return nil
}
