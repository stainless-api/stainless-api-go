// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

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
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// ProjectConfigService contains methods and other services that help with
// interacting with the stainless API.
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

// Retrieve the configuration files for a given project.
func (r *ProjectConfigService) Get(ctx context.Context, params ProjectConfigGetParams, opts ...option.RequestOption) (res *ProjectConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Project, precfg.Project)
	if params.Project.Value == "" {
		err = errors.New("missing required project parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/configs", url.PathEscape(params.Project.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Generate suggestions for changes to config files based on an OpenAPI spec.
func (r *ProjectConfigService) Guess(ctx context.Context, params ProjectConfigGuessParams, opts ...option.RequestOption) (res *ProjectConfigGuessResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Project, precfg.Project)
	if params.Project.Value == "" {
		err = errors.New("missing required project parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/configs/guess", url.PathEscape(params.Project.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ProjectConfigGetResponse map[string]ProjectConfigGetResponseItem

type ProjectConfigGetResponseItem struct {
	// The file content
	Content string `json:"content,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectConfigGuessResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ProjectConfigGuessResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectConfigGetParams struct {
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `path:"project,omitzero,required" json:"-"`
	// Branch name, defaults to "main".
	Branch  param.Opt[string] `query:"branch,omitzero" json:"-"`
	Include param.Opt[string] `query:"include,omitzero" json:"-"`
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
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `path:"project,omitzero,required" json:"-"`
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
func (r *ProjectConfigGuessParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
