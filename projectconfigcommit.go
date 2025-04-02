// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/param"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
)

// ProjectConfigCommitService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectConfigCommitService] method instead.
type ProjectConfigCommitService struct {
	Options []option.RequestOption
}

// NewProjectConfigCommitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectConfigCommitService(opts ...option.RequestOption) (r *ProjectConfigCommitService) {
	r = &ProjectConfigCommitService{}
	r.Options = opts
	return
}

// TODO
func (r *ProjectConfigCommitService) New(ctx context.Context, projectName string, body ProjectConfigCommitNewParams, opts ...option.RequestOption) (res *Commit, err error) {
	opts = append(r.Options[:], opts...)
	if projectName == "" {
		err = errors.New("missing required projectName parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/config/commits", projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Commit struct {
	ID     string       `json:"id,required"`
	Builds []Build      `json:"builds,required"`
	Object CommitObject `json:"object,required"`
	JSON   commitJSON   `json:"-"`
}

// commitJSON contains the JSON metadata for the struct [Commit]
type commitJSON struct {
	ID          apijson.Field
	Builds      apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Commit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitJSON) RawJSON() string {
	return r.raw
}

type CommitObject string

const (
	CommitObjectConfigCommit CommitObject = "config_commit"
)

func (r CommitObject) IsKnown() bool {
	switch r {
	case CommitObjectConfigCommit:
		return true
	}
	return false
}

type ProjectConfigCommitNewParams struct {
	Branch          param.Field[string] `json:"branch,required"`
	CommitMessage   param.Field[string] `json:"commit_message,required"`
	AllowEmpty      param.Field[bool]   `json:"allow_empty"`
	OpenAPISpec     param.Field[string] `json:"openapi_spec"`
	StainlessConfig param.Field[string] `json:"stainless_config"`
}

func (r ProjectConfigCommitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
