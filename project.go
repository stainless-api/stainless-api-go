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
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// ProjectService contains methods and other services that help with interacting
// with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options  []option.RequestOption
	Branches ProjectBranchService
	Configs  ProjectConfigService
	Snippets ProjectSnippetService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	r.Branches = NewProjectBranchService(opts...)
	r.Configs = NewProjectConfigService(opts...)
	r.Snippets = NewProjectSnippetService(opts...)
	return
}

// Retrieve a project by name
func (r *ProjectService) Get(ctx context.Context, query ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Project, precfg.Project)
	if query.Project.Value == "" {
		err = errors.New("missing required project parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s", query.Project.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a project's properties
func (r *ProjectService) Update(ctx context.Context, params ProjectUpdateParams, opts ...option.RequestOption) (res *ProjectUpdateResponse, err error) {
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
	path := fmt.Sprintf("v0/projects/%s", params.Project.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List projects in an organization
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *ProjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProjectGetResponse struct {
	ConfigRepo  string `json:"config_repo,required"`
	DisplayName string `json:"display_name,required"`
	// Any of "project".
	Object ProjectGetResponseObject `json:"object,required"`
	Org    string                   `json:"org,required"`
	Slug   string                   `json:"slug,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigRepo  respjson.Field
		DisplayName respjson.Field
		Object      respjson.Field
		Org         respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetResponseObject string

const (
	ProjectGetResponseObjectProject ProjectGetResponseObject = "project"
)

type ProjectUpdateResponse struct {
	ConfigRepo  string `json:"config_repo,required"`
	DisplayName string `json:"display_name,required"`
	// Any of "project".
	Object ProjectUpdateResponseObject `json:"object,required"`
	Org    string                      `json:"org,required"`
	Slug   string                      `json:"slug,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigRepo  respjson.Field
		DisplayName respjson.Field
		Object      respjson.Field
		Org         respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateResponseObject string

const (
	ProjectUpdateResponseObjectProject ProjectUpdateResponseObject = "project"
)

type ProjectListResponse struct {
	Data       []ProjectListResponseData `json:"data,required"`
	HasMore    bool                      `json:"has_more,required"`
	NextCursor string                    `json:"next_cursor"`
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
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseData struct {
	ConfigRepo  string `json:"config_repo,required"`
	DisplayName string `json:"display_name,required"`
	// Any of "project".
	Object string `json:"object,required"`
	Org    string `json:"org,required"`
	Slug   string `json:"slug,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigRepo  respjson.Field
		DisplayName respjson.Field
		Object      respjson.Field
		Org         respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetParams struct {
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `path:"project,omitzero,required" json:"-"`
	paramObj
}

type ProjectUpdateParams struct {
	// Use [option.WithProject] on the client to set a global default for this field.
	Project     param.Opt[string] `path:"project,omitzero,required" json:"-"`
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListParams struct {
	Org string `query:"org,required" json:"-"`
	// Pagination cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of projects to return, defaults to 10 (maximum: 100)
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
