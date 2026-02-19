// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Retrieve the currently authenticated user's information.
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserGetResponse struct {
	ID     string                `json:"id,required"`
	Email  string                `json:"email,required"`
	GitHub UserGetResponseGitHub `json:"github,required"`
	Name   string                `json:"name,required"`
	// Any of "user".
	Object UserGetResponseObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		GitHub      respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetResponseGitHub struct {
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponseGitHub) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseGitHub) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetResponseObject string

const (
	UserGetResponseObjectUser UserGetResponseObject = "user"
)
