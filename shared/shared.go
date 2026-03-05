// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Commit struct {
	Repo    CommitRepo  `json:"repo" api:"required"`
	Sha     string      `json:"sha" api:"required"`
	Stats   CommitStats `json:"stats" api:"required"`
	TreeOid string      `json:"tree_oid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Repo        respjson.Field
		Sha         respjson.Field
		Stats       respjson.Field
		TreeOid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Commit) RawJSON() string { return r.JSON.raw }
func (r *Commit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitRepo struct {
	Branch string `json:"branch" api:"required"`
	Host   string `json:"host" api:"required"`
	Name   string `json:"name" api:"required"`
	Owner  string `json:"owner" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Branch      respjson.Field
		Host        respjson.Field
		Name        respjson.Field
		Owner       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitRepo) RawJSON() string { return r.JSON.raw }
func (r *CommitRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitStats struct {
	Additions int64 `json:"additions" api:"required"`
	Deletions int64 `json:"deletions" api:"required"`
	Total     int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Additions   respjson.Field
		Deletions   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitStats) RawJSON() string { return r.JSON.raw }
func (r *CommitStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func FileInputParamOfFileInputContent(content string) FileInputUnionParam {
	var variant FileInputContentParam
	variant.Content = content
	return FileInputUnionParam{OfFileInputContent: &variant}
}

func FileInputParamOfFileInputURL(url string) FileInputUnionParam {
	var variant FileInputURLParam
	variant.URL = url
	return FileInputUnionParam{OfFileInputURL: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileInputUnionParam struct {
	OfFileInputContent *FileInputContentParam `json:",omitzero,inline"`
	OfFileInputURL     *FileInputURLParam     `json:",omitzero,inline"`
	paramUnion
}

func (u FileInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileInputContent, u.OfFileInputURL)
}
func (u *FileInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFileInputContent) {
		return u.OfFileInputContent
	} else if !param.IsOmitted(u.OfFileInputURL) {
		return u.OfFileInputURL
	}
	return nil
}

// The property Content is required.
type FileInputContentParam struct {
	// File content
	Content string `json:"content" api:"required"`
	paramObj
}

func (r FileInputContentParam) MarshalJSON() (data []byte, err error) {
	type shadow FileInputContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileInputContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type FileInputURLParam struct {
	// URL to fetch file content from
	URL string `json:"url" api:"required"`
	paramObj
}

func (r FileInputURLParam) MarshalJSON() (data []byte, err error) {
	type shadow FileInputURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileInputURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Target string

const (
	TargetNode       Target = "node"
	TargetTypescript Target = "typescript"
	TargetPython     Target = "python"
	TargetGo         Target = "go"
	TargetJava       Target = "java"
	TargetKotlin     Target = "kotlin"
	TargetRuby       Target = "ruby"
	TargetTerraform  Target = "terraform"
	TargetCli        Target = "cli"
	TargetPhp        Target = "php"
	TargetCsharp     Target = "csharp"
	TargetSql        Target = "sql"
	TargetOpenAPI    Target = "openapi"
)
