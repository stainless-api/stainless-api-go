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
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// ProjectSnippetService contains methods and other services that help with
// interacting with the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectSnippetService] method instead.
type ProjectSnippetService struct {
	Options []option.RequestOption
}

// NewProjectSnippetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectSnippetService(opts ...option.RequestOption) (r ProjectSnippetService) {
	r = ProjectSnippetService{}
	r.Options = opts
	return
}

func (r *ProjectSnippetService) NewRequest(ctx context.Context, projectName string, body ProjectSnippetNewRequestParams, opts ...option.RequestOption) (res *ProjectSnippetNewRequestResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectName == "" {
		err = errors.New("missing required projectName parameter")
		return
	}
	path := fmt.Sprintf("v0/projects/%s/snippets/request", projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProjectSnippetNewRequestResponse struct {
	Snippet string `json:"snippet,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Snippet     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectSnippetNewRequestResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectSnippetNewRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectSnippetNewRequestParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfObject *ProjectSnippetNewRequestParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfProjectSnippetNewRequestsBodyObject *ProjectSnippetNewRequestParamsBodyObject `json:",inline"`

	paramObj
}

func (u ProjectSnippetNewRequestParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ProjectSnippetNewRequestParams](u.OfObject, u.OfProjectSnippetNewRequestsBodyObject)
}
func (r *ProjectSnippetNewRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Request is required.
type ProjectSnippetNewRequestParamsBodyObject struct {
	Request ProjectSnippetNewRequestParamsBodyObjectRequest `json:"request,omitzero,required"`
	Har     any                                             `json:"har"`
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli", "php", "csharp".
	Language string `json:"language,omitzero"`
	// Any of "next", "latest_released".
	Version string `json:"version,omitzero"`
	paramObj
}

func (r ProjectSnippetNewRequestParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProjectSnippetNewRequestParamsBodyObject](
		"language", "node", "typescript", "python", "go", "java", "kotlin", "ruby", "terraform", "cli", "php", "csharp",
	)
	apijson.RegisterFieldValidator[ProjectSnippetNewRequestParamsBodyObject](
		"version", "next", "latest_released",
	)
}

// The properties Method, Parameters, Path are required.
type ProjectSnippetNewRequestParamsBodyObjectRequest struct {
	Method     string                                                     `json:"method,required"`
	Parameters []ProjectSnippetNewRequestParamsBodyObjectRequestParameter `json:"parameters,omitzero,required"`
	Path       string                                                     `json:"path,required"`
	Body       ProjectSnippetNewRequestParamsBodyObjectRequestBody        `json:"body,omitzero"`
	paramObj
}

func (r ProjectSnippetNewRequestParamsBodyObjectRequest) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsBodyObjectRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParamsBodyObjectRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties In, Name are required.
type ProjectSnippetNewRequestParamsBodyObjectRequestParameter struct {
	// Any of "path", "query", "header", "cookie".
	In    string `json:"in,omitzero,required"`
	Name  string `json:"name,required"`
	Value any    `json:"value,omitzero"`
	paramObj
}

func (r ProjectSnippetNewRequestParamsBodyObjectRequestParameter) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsBodyObjectRequestParameter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParamsBodyObjectRequestParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProjectSnippetNewRequestParamsBodyObjectRequestParameter](
		"in", "path", "query", "header", "cookie",
	)
}

type ProjectSnippetNewRequestParamsBodyObjectRequestBody struct {
	FileParam param.Opt[string] `json:"fileParam,omitzero"`
	FilePath  param.Opt[string] `json:"filePath,omitzero"`
	paramObj
}

func (r ProjectSnippetNewRequestParamsBodyObjectRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsBodyObjectRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParamsBodyObjectRequestBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
