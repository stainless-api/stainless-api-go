// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/apiquery"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/resp"
)

// ProjectConfigService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectConfigService] method instead.
type ProjectConfigService struct {
	Options []option.RequestOption
}

// NewProjectConfigService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectConfigService(opts ...option.RequestOption) (r ProjectConfigService) {
	r = ProjectConfigService{}
	r.Options = opts
	return
}

// Retrieve configuration files for a project
func (r *ProjectConfigService) Get(ctx context.Context, project string, query ProjectConfigGetParams, opts ...option.RequestOption) (res *ProjectConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if project == "" {
		err = errors.New("missing required project parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/configs", project)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Generate configuration suggestions based on an OpenAPI spec
func (r *ProjectConfigService) Guess(ctx context.Context, project string, body ProjectConfigGuessParams, opts ...option.RequestOption) (res *ProjectConfigGuessResponse, err error) {
	opts = append(r.Options[:], opts...)
	if project == "" {
		err = errors.New("missing required project parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/configs/guess", project)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectConfigGetResponse map[string]ProjectConfigGetResponseItem

type ProjectConfigGetResponseItem struct {
	// The file content
	Content string `json:"content,required"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Content     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectConfigGetResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ProjectConfigGetResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectConfigGuessResponse map[string]ProjectConfigGuessResponseItem

type ProjectConfigGuessResponseItem struct {
	// The file content
	Content string `json:"content,required"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
	JSON struct {
		Content     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectConfigGuessResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ProjectConfigGuessResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectConfigGetParams struct {
	// Branch name, defaults to "main"
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectConfigGetParams]'s query parameters as `url.Values`.
func (r ProjectConfigGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectConfigGuessParams struct {
	// OpenAPI spec
	Spec string `json:"spec,required"`
	// Branch name
	Branch param.Opt[string] `json:"branch,omitzero"`
	paramObj
}

func (r ProjectConfigGuessParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectConfigGuessParams
	return param.MarshalObject(r, (*shadow)(&r))
}
