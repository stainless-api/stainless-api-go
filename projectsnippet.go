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
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli", "php", "csharp".
	Language ProjectSnippetNewRequestParamsLanguage     `json:"language,omitzero,required"`
	Request  ProjectSnippetNewRequestParamsRequestUnion `json:"request,omitzero,required"`
	// Any of "next", "latest_released".
	Version ProjectSnippetNewRequestParamsVersion `json:"version,omitzero,required"`
	paramObj
}

func (r ProjectSnippetNewRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ProjectSnippetNewRequestParamsLanguagePhp        ProjectSnippetNewRequestParamsLanguage = "php"
	ProjectSnippetNewRequestParamsLanguageCsharp     ProjectSnippetNewRequestParamsLanguage = "csharp"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProjectSnippetNewRequestParamsRequestUnion struct {
	OfProjectSnippetNewRequestsRequestObject *ProjectSnippetNewRequestParamsRequestObject `json:",omitzero,inline"`
	OfVariant2                               *ProjectSnippetNewRequestParamsRequestObject `json:",omitzero,inline"`
	paramUnion
}

func (u ProjectSnippetNewRequestParamsRequestUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ProjectSnippetNewRequestParamsRequestUnion](u.OfProjectSnippetNewRequestsRequestObject, u.OfVariant2)
}
func (u *ProjectSnippetNewRequestParamsRequestUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ProjectSnippetNewRequestParamsRequestUnion) asAny() any {
	if !param.IsOmitted(u.OfProjectSnippetNewRequestsRequestObject) {
		return u.OfProjectSnippetNewRequestsRequestObject
	} else if !param.IsOmitted(u.OfVariant2) {
		return u.OfVariant2
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ProjectSnippetNewRequestParamsRequestUnion) GetParameters() []ProjectSnippetNewRequestParamsRequestObjectParameter {
	if vt := u.OfProjectSnippetNewRequestsRequestObject; vt != nil {
		return vt.Parameters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ProjectSnippetNewRequestParamsRequestUnion) GetPath() *string {
	if vt := u.OfProjectSnippetNewRequestsRequestObject; vt != nil {
		return &vt.Path
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ProjectSnippetNewRequestParamsRequestUnion) GetBody() *ProjectSnippetNewRequestParamsRequestObjectBody {
	if vt := u.OfProjectSnippetNewRequestsRequestObject; vt != nil {
		return &vt.Body
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ProjectSnippetNewRequestParamsRequestUnion) GetQueryString() []ProjectSnippetNewRequestParamsRequestObjectQueryString {
	if vt := u.OfVariant2; vt != nil {
		return vt.QueryString
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ProjectSnippetNewRequestParamsRequestUnion) GetURL() *string {
	if vt := u.OfVariant2; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ProjectSnippetNewRequestParamsRequestUnion) GetPostData() *ProjectSnippetNewRequestParamsRequestObjectPostData {
	if vt := u.OfVariant2; vt != nil {
		return &vt.PostData
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ProjectSnippetNewRequestParamsRequestUnion) GetMethod() *string {
	if vt := u.OfProjectSnippetNewRequestsRequestObject; vt != nil {
		return (*string)(&vt.Method)
	} else if vt := u.OfVariant2; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// The properties Method, Parameters, Path are required.
type ProjectSnippetNewRequestParamsRequestObject struct {
	Method     string                                                 `json:"method,required"`
	Parameters []ProjectSnippetNewRequestParamsRequestObjectParameter `json:"parameters,omitzero,required"`
	Path       string                                                 `json:"path,required"`
	Body       ProjectSnippetNewRequestParamsRequestObjectBody        `json:"body,omitzero"`
	paramObj
}

func (r ProjectSnippetNewRequestParamsRequestObject) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsRequestObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParamsRequestObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties In, Name are required.
type ProjectSnippetNewRequestParamsRequestObjectParameter struct {
	// Any of "path", "query", "header", "cookie".
	In    string `json:"in,omitzero,required"`
	Name  string `json:"name,required"`
	Value any    `json:"value,omitzero"`
	paramObj
}

func (r ProjectSnippetNewRequestParamsRequestObjectParameter) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsRequestObjectParameter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParamsRequestObjectParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProjectSnippetNewRequestParamsRequestObjectParameter](
		"in", "path", "query", "header", "cookie",
	)
}

type ProjectSnippetNewRequestParamsRequestObjectBody struct {
	FileParam param.Opt[string] `json:"fileParam,omitzero"`
	FilePath  param.Opt[string] `json:"filePath,omitzero"`
	paramObj
}

func (r ProjectSnippetNewRequestParamsRequestObjectBody) MarshalJSON() (data []byte, err error) {
	type shadow ProjectSnippetNewRequestParamsRequestObjectBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectSnippetNewRequestParamsRequestObjectBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectSnippetNewRequestParamsVersion string

const (
	ProjectSnippetNewRequestParamsVersionNext           ProjectSnippetNewRequestParamsVersion = "next"
	ProjectSnippetNewRequestParamsVersionLatestReleased ProjectSnippetNewRequestParamsVersion = "latest_released"
)
