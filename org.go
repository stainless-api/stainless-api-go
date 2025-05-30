// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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

// Retrieve an organization by name
func (r *OrgService) Get(ctx context.Context, orgName string, opts ...option.RequestOption) (res *OrgGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required orgName parameter")
		return
	}
	path := fmt.Sprintf("v0/orgs/%s", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List organizations the user has access to
func (r *OrgService) List(ctx context.Context, opts ...option.RequestOption) (res *OrgListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrgGetResponse struct {
	DisplayName string `json:"display_name,required"`
	// Any of "org".
	Object OrgGetResponseObject `json:"object,required"`
	Slug   string               `json:"slug,required"`
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
func (r OrgGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrgGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrgGetResponseObject string

const (
	OrgGetResponseObjectOrg OrgGetResponseObject = "org"
)

type OrgListResponse struct {
	Data       []OrgListResponseData `json:"data,required"`
	HasMore    bool                  `json:"has_more,required"`
	NextCursor string                `json:"next_cursor"`
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

type OrgListResponseData struct {
	DisplayName string `json:"display_name,required"`
	// Any of "org".
	Object string `json:"object,required"`
	Slug   string `json:"slug,required"`
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
func (r OrgListResponseData) RawJSON() string { return r.JSON.raw }
func (r *OrgListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
