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

// ProjectSnippetService contains methods and other services that help with
// interacting with the stainless-v0 API.
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Snippet     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectSnippetNewRequestResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectSnippetNewRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectSnippetNewRequestParams struct {
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli".
	Language ProjectSnippetNewRequestParamsLanguage `json:"language,omitzero,required"`
	Request  ProjectSnippetNewRequestParamsRequest  `json:"request,omitzero,required"`
	// Any of "next", "latest_released".
	Version ProjectSnippetNewRequestParamsVersion `json:"version,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectSnippetNewRequestParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectSnippetNewRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ProjectSnippetNewRequestParamsLanguage string

const (
	ProjectSnippetNewRequestParamsLanguageNode       ProjectSnippetNewRequestParamsLanguage = "node"
	ProjectSnippetNewRequestParamsLanguageTypescript ProjectSnippetNewRequestParamsLanguage = "typescript"
	ProjectSnippetNewRequestParamsLanguagePython     ProjectSnippetNewRequestParamsLanguage = "python"
	ProjectSnippetNewRequestParamsLanguageGo         ProjectSnippetNewRequestParamsLanguage = "go"
	ProjectSnippetNewRequestParamsLanguageJava       ProjectSnippetNewRequestParamsLanguage = "java"
	ProjectSnippetNewRequestParamsLanguageKotlin     ProjectSnippetNewRequestParamsLanguage = "kotlin"
	ProjectSnippetNewRequestParamsLanguageRuby       ProjectSnippetNewRequestParamsLanguage = "ruby"
	ProjectSnippetNewRequestParamsLanguageTerraform  ProjectSnippetNewRequestParamsLanguage = "terraform"
	ProjectSnippetNewRequestParamsLanguageCli        ProjectSnippetNewRequestParamsLanguage = "cli"
)

// The properties Method, Parameters, Path are required.
type ProjectSnippetNewRequestParamsRequest struct {
	Method     string                                           `json:"method,required"`
	Parameters []ProjectSnippetNewRequestParamsRequestParameter `json:"parameters,omitzero,required"`
	Path       string                                           `json:"path,required"`
	Body       ProjectSnippetNewRequestParamsRequestBody        `json:"body,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectSnippetNewRequestParamsRequest) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ProjectSnippetNewRequestParamsRequest) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsRequest
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties In, Name are required.
type ProjectSnippetNewRequestParamsRequestParameter struct {
	// Any of "path", "query", "header", "cookie".
	In    string `json:"in,omitzero,required"`
	Name  string `json:"name,required"`
	Value any    `json:"value,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectSnippetNewRequestParamsRequestParameter) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ProjectSnippetNewRequestParamsRequestParameter) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsRequestParameter
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[ProjectSnippetNewRequestParamsRequestParameter](
		"In", false, "path", "query", "header", "cookie",
	)
}

type ProjectSnippetNewRequestParamsRequestBody struct {
	FileParam param.Opt[string] `json:"fileParam,omitzero"`
	FilePath  param.Opt[string] `json:"filePath,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectSnippetNewRequestParamsRequestBody) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ProjectSnippetNewRequestParamsRequestBody) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsRequestBody
	return param.MarshalObject(r, (*shadow)(&r))
}

type ProjectSnippetNewRequestParamsVersion string

const (
	ProjectSnippetNewRequestParamsVersionNext           ProjectSnippetNewRequestParamsVersion = "next"
	ProjectSnippetNewRequestParamsVersionLatestReleased ProjectSnippetNewRequestParamsVersion = "latest_released"
)
