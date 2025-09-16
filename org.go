// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// OrgService contains methods and other services that help with interacting with
// the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgService] method instead.
type OrgService struct {
	Options []option.RequestOption
}

// NewOrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgService(opts ...option.RequestOption) (r OrgService) {
	r = OrgService{}
	r.Options = opts
	return
}

// Retrieve an organization by name.
func (r *OrgService) Get(ctx context.Context, org string, opts ...option.RequestOption) (res *Org, err error) {
	opts = append(r.Options[:], opts...)
	if org == "" {
		err = errors.New("missing required org parameter")
		return
	}
	path := fmt.Sprintf("v0/orgs/%s", url.PathEscape(org))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List organizations accessible to the current authentication method.
func (r *OrgService) List(ctx context.Context, opts ...option.RequestOption) (res *OrgListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Org struct {
	DisplayName string `json:"display_name,required"`
	// Any of "org".
	Object OrgObject `json:"object,required"`
	Slug   string    `json:"slug,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Object      respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Org) RawJSON() string { return r.JSON.raw }
func (r *Org) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrgObject string

const (
	OrgObjectOrg OrgObject = "org"
)

type OrgListResponse struct {
	Data       []Org  `json:"data,required"`
	HasMore    bool   `json:"has_more,required"`
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrgListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrgListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
