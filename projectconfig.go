// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/stainless-v0-go/internal/apijson"
	"github.com/stainless-sdks/stainless-v0-go/internal/param"
	"github.com/stainless-sdks/stainless-v0-go/internal/requestconfig"
	"github.com/stainless-sdks/stainless-v0-go/option"
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
func NewProjectConfigService(opts ...option.RequestOption) (r *ProjectConfigService) {
	r = &ProjectConfigService{}
	r.Options = opts
	return
}

// TODO
func (r *ProjectConfigService) NewBranch(ctx context.Context, projectName string, body ProjectConfigNewBranchParams, opts ...option.RequestOption) (res *ConfigCommit, err error) {
	opts = append(r.Options[:], opts...)
	if projectName == "" {
		err = errors.New("missing required projectName parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/config/branches", projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// TODO
func (r *ProjectConfigService) NewCommit(ctx context.Context, projectName string, body ProjectConfigNewCommitParams, opts ...option.RequestOption) (res *ConfigCommit, err error) {
	opts = append(r.Options[:], opts...)
	if projectName == "" {
		err = errors.New("missing required projectName parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/config/commits", projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// TODO
func (r *ProjectConfigService) Merge(ctx context.Context, projectName string, body ProjectConfigMergeParams, opts ...option.RequestOption) (res *ConfigCommit, err error) {
	opts = append(r.Options[:], opts...)
	if projectName == "" {
		err = errors.New("missing required projectName parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/config/merge", projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ConfigCommit struct {
	ID     string              `json:"id,required"`
	Builds []ConfigCommitBuild `json:"builds,required"`
	Object ConfigCommitObject  `json:"object,required"`
	JSON   configCommitJSON    `json:"-"`
}

// configCommitJSON contains the JSON metadata for the struct [ConfigCommit]
type configCommitJSON struct {
	ID          apijson.Field
	Builds      apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configCommitJSON) RawJSON() string {
	return r.raw
}

type ConfigCommitBuild struct {
	ID     string                   `json:"id,required"`
	Object ConfigCommitBuildsObject `json:"object,required"`
	JSON   configCommitBuildJSON    `json:"-"`
}

// configCommitBuildJSON contains the JSON metadata for the struct
// [ConfigCommitBuild]
type configCommitBuildJSON struct {
	ID          apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigCommitBuild) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configCommitBuildJSON) RawJSON() string {
	return r.raw
}

type ConfigCommitBuildsObject string

const (
	ConfigCommitBuildsObjectBuild ConfigCommitBuildsObject = "build"
)

func (r ConfigCommitBuildsObject) IsKnown() bool {
	switch r {
	case ConfigCommitBuildsObjectBuild:
		return true
	}
	return false
}

type ConfigCommitObject string

const (
	ConfigCommitObjectConfigCommit ConfigCommitObject = "config_commit"
)

func (r ConfigCommitObject) IsKnown() bool {
	switch r {
	case ConfigCommitObjectConfigCommit:
		return true
	}
	return false
}

type ProjectConfigNewBranchParams struct {
	Branch     param.Field[string] `json:"branch,required"`
	BranchFrom param.Field[string] `json:"branch_from,required"`
}

func (r ProjectConfigNewBranchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectConfigNewCommitParams struct {
	Branch          param.Field[string] `json:"branch,required"`
	CommitMessage   param.Field[string] `json:"commit_message,required"`
	OpenAPISpec     param.Field[string] `json:"openapi_spec,required"`
	StainlessConfig param.Field[string] `json:"stainless_config,required"`
	AllowEmpty      param.Field[bool]   `json:"allow_empty"`
}

func (r ProjectConfigNewCommitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectConfigMergeParams struct {
	From param.Field[string] `json:"from,required"`
	Into param.Field[string] `json:"into,required"`
}

func (r ProjectConfigMergeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
