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

// ProjectConfigBranchService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectConfigBranchService] method instead.
type ProjectConfigBranchService struct {
	Options []option.RequestOption
}

// NewProjectConfigBranchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectConfigBranchService(opts ...option.RequestOption) (r *ProjectConfigBranchService) {
	r = &ProjectConfigBranchService{}
	r.Options = opts
	return
}

// TODO
func (r *ProjectConfigBranchService) New(ctx context.Context, projectName string, body ProjectConfigBranchNewParams, opts ...option.RequestOption) (res *Commit, err error) {
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
func (r *ProjectConfigBranchService) Merge(ctx context.Context, projectName string, body ProjectConfigBranchMergeParams, opts ...option.RequestOption) (res *Commit, err error) {
	opts = append(r.Options[:], opts...)
	if projectName == "" {
		err = errors.New("missing required projectName parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/config/merge", projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectConfigBranchNewParams struct {
	Branch     param.Field[string] `json:"branch,required"`
	BranchFrom param.Field[string] `json:"branch_from,required"`
}

func (r ProjectConfigBranchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectConfigBranchMergeParams struct {
	From param.Field[string] `json:"from,required"`
	Into param.Field[string] `json:"into,required"`
}

func (r ProjectConfigBranchMergeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
