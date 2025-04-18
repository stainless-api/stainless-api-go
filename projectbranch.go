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
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/resp"
)

// ProjectBranchService contains methods and other services that help with
// interacting with the stainless-v0 API.
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

// TODO
func (r *ProjectBranchService) New(ctx context.Context, project string, body ProjectBranchNewParams, opts ...option.RequestOption) (res *ProjectBranch, err error) {
	opts = append(r.Options[:], opts...)
	if project == "" {
		err = errors.New("missing required project parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/branches", project)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// TODO
func (r *ProjectBranchService) Get(ctx context.Context, project string, branch string, opts ...option.RequestOption) (res *ProjectBranch, err error) {
	opts = append(r.Options[:], opts...)
	if project == "" {
		err = errors.New("missing required project parameter")
		return
	}
	if branch == "" {
		err = errors.New("missing required branch parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/branches/%s", project, branch)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Branch       resp.Field
		ConfigCommit resp.Field
		LatestBuild  resp.Field
		Object       resp.Field
		Org          resp.Field
		Project      resp.Field
		ExtraFields  map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Repo        resp.Field
		Sha         resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Branch      resp.Field
		Name        resp.Field
		Owner       resp.Field
		ExtraFields map[string]resp.Field
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

type ProjectBranchNewParams struct {
	Branch     string `json:"branch,required"`
	BranchFrom string `json:"branch_from,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectBranchNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectBranchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectBranchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
