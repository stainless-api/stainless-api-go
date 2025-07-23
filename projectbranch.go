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
	"github.com/stainless-api/stainless-api-go/packages/pagination"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// ProjectBranchService contains methods and other services that help with
// interacting with the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectBranchService] method instead.
type ProjectBranchService struct {
	Options []option.RequestOption
}

// NewProjectBranchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectBranchService(opts ...option.RequestOption) (r ProjectBranchService) {
	r = ProjectBranchService{}
	r.Options = opts
	return
}

// Create a new branch for a project
func (r *ProjectBranchService) New(ctx context.Context, params ProjectBranchNewParams, opts ...option.RequestOption) (res *ProjectBranch, err error) {
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
	path := fmt.Sprintf("v0/projects/%s/branches", params.Project.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a project branch
func (r *ProjectBranchService) Get(ctx context.Context, branch string, query ProjectBranchGetParams, opts ...option.RequestOption) (res *ProjectBranch, err error) {
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
	if branch == "" {
		err = errors.New("missing required branch parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/branches/%s", query.Project.Value, branch)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List project branches
func (r *ProjectBranchService) List(ctx context.Context, params ProjectBranchListParams, opts ...option.RequestOption) (res *pagination.Page[ProjectBranchListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.Project, precfg.Project)
	if params.Project.Value == "" {
		err = errors.New("missing required project parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/branches", params.Project.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List project branches
func (r *ProjectBranchService) ListAutoPaging(ctx context.Context, params ProjectBranchListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ProjectBranchListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a project branch
func (r *ProjectBranchService) Delete(ctx context.Context, branch string, body ProjectBranchDeleteParams, opts ...option.RequestOption) (res *ProjectBranchDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.Project, precfg.Project)
	if body.Project.Value == "" {
		err = errors.New("missing required project parameter")
		return
	}
	if branch == "" {
		err = errors.New("missing required branch parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/branches/%s", body.Project.Value, branch)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ProjectBranch struct {
	Branch       string                    `json:"branch,required"`
	ConfigCommit ProjectBranchConfigCommit `json:"config_commit,required"`
	LatestBuild  BuildObject               `json:"latest_build,required"`
	// Any of "project_branch".
	Object  ProjectBranchObject `json:"object,required"`
	Org     string              `json:"org,required"`
	Project string              `json:"project,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Branch       respjson.Field
		ConfigCommit respjson.Field
		LatestBuild  respjson.Field
		Object       respjson.Field
		Org          respjson.Field
		Project      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectBranch) RawJSON() string { return r.JSON.raw }
func (r *ProjectBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectBranchConfigCommit struct {
	Repo ProjectBranchConfigCommitRepo `json:"repo,required"`
	Sha  string                        `json:"sha,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Repo        respjson.Field
		Sha         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectBranchConfigCommit) RawJSON() string { return r.JSON.raw }
func (r *ProjectBranchConfigCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectBranchConfigCommitRepo struct {
	Branch string `json:"branch,required"`
	Name   string `json:"name,required"`
	Owner  string `json:"owner,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Branch      respjson.Field
		Name        respjson.Field
		Owner       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectBranchConfigCommitRepo) RawJSON() string { return r.JSON.raw }
func (r *ProjectBranchConfigCommitRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectBranchObject string

const (
	ProjectBranchObjectProjectBranch ProjectBranchObject = "project_branch"
)

type ProjectBranchListResponse struct {
	Branch        string                                `json:"branch,required"`
	ConfigCommit  ProjectBranchListResponseConfigCommit `json:"config_commit,required"`
	LatestBuildID string                                `json:"latest_build_id,required"`
	// Any of "project_branch".
	Object  ProjectBranchListResponseObject `json:"object,required"`
	Org     string                          `json:"org,required"`
	Project string                          `json:"project,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Branch        respjson.Field
		ConfigCommit  respjson.Field
		LatestBuildID respjson.Field
		Object        respjson.Field
		Org           respjson.Field
		Project       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectBranchListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectBranchListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectBranchListResponseConfigCommit struct {
	Repo ProjectBranchListResponseConfigCommitRepo `json:"repo,required"`
	Sha  string                                    `json:"sha,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Repo        respjson.Field
		Sha         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectBranchListResponseConfigCommit) RawJSON() string { return r.JSON.raw }
func (r *ProjectBranchListResponseConfigCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectBranchListResponseConfigCommitRepo struct {
	Branch string `json:"branch,required"`
	Name   string `json:"name,required"`
	Owner  string `json:"owner,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Branch      respjson.Field
		Name        respjson.Field
		Owner       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectBranchListResponseConfigCommitRepo) RawJSON() string { return r.JSON.raw }
func (r *ProjectBranchListResponseConfigCommitRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectBranchListResponseObject string

const (
	ProjectBranchListResponseObjectProjectBranch ProjectBranchListResponseObject = "project_branch"
)

type ProjectBranchDeleteResponse = any

type ProjectBranchNewParams struct {
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `path:"project,omitzero,required" json:"-"`
	// Name of the new project branch.
	Branch string `json:"branch,required"`
	// Branch or commit SHA to branch from.
	BranchFrom string `json:"branch_from,required"`
	// Whether to throw an error if the branch already exists. Defaults to false.
	Force param.Opt[bool] `json:"force,omitzero"`
	paramObj
}

func (r ProjectBranchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectBranchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectBranchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectBranchGetParams struct {
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `path:"project,omitzero,required" json:"-"`
	paramObj
}

type ProjectBranchListParams struct {
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `path:"project,omitzero,required" json:"-"`
	// Pagination cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return, defaults to 10 (maximum: 100)
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectBranchListParams]'s query parameters as
// `url.Values`.
func (r ProjectBranchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectBranchDeleteParams struct {
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `path:"project,omitzero,required" json:"-"`
	paramObj
}
