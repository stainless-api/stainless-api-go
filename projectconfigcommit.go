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
func NewProjectConfigCommitService(opts ...option.RequestOption) (r ProjectConfigCommitService) {
	r = ProjectConfigCommitService{}
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
	ID     string  `json:"id,required"`
	Builds []Build `json:"builds,required"`
	// Any of "config_commit".
	Object CommitObject `json:"object,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Builds      resp.Field
		Object      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Commit) RawJSON() string { return r.JSON.raw }
func (r *Commit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitObject string

const (
	CommitObjectConfigCommit CommitObject = "config_commit"
)

type ProjectConfigCommitNewParams struct {
	Branch          string            `json:"branch,required"`
	CommitMessage   string            `json:"commit_message,required"`
	AllowEmpty      param.Opt[bool]   `json:"allow_empty,omitzero"`
	OpenAPISpec     param.Opt[string] `json:"openapi_spec,omitzero"`
	StainlessConfig param.Opt[string] `json:"stainless_config,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectConfigCommitNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectConfigCommitNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectConfigCommitNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
