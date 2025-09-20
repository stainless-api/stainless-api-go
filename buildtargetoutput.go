// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/apiquery"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
	"github.com/stainless-api/stainless-api-go/shared"
	"github.com/stainless-api/stainless-api-go/shared/constant"
)

// BuildTargetOutputService contains methods and other services that help with
// interacting with the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuildTargetOutputService] method instead.
type BuildTargetOutputService struct {
	Options []option.RequestOption
}

// NewBuildTargetOutputService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBuildTargetOutputService(opts ...option.RequestOption) (r BuildTargetOutputService) {
	r = BuildTargetOutputService{}
	r.Options = opts
	return
}

// Retrieve a method to download an output for a given build target.
//
// If the requested type of output is `source`, and the requested output method is
// `url`, a download link to a tarball of the source files is returned. If the
// requested output method is `git`, a Git remote, ref, and access token (if
// necessary) is returned.
//
// Otherwise, the possible types of outputs are specific to the requested target,
// and the output method _must_ be `url`. See the documentation for `type` for more
// information.
func (r *BuildTargetOutputService) Get(ctx context.Context, query BuildTargetOutputGetParams, opts ...option.RequestOption) (res *BuildTargetOutputGetResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/build_target_outputs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// BuildTargetOutputGetResponseUnion contains all possible properties and values
// from [BuildTargetOutputGetResponseURL], [BuildTargetOutputGetResponseGit].
//
// Use the [BuildTargetOutputGetResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetOutputGetResponseUnion struct {
	// Any of "url", "git".
	Output string `json:"output"`
	// This field is from variant [BuildTargetOutputGetResponseURL].
	Target shared.Target `json:"target"`
	Type   string        `json:"type"`
	URL    string        `json:"url"`
	// This field is from variant [BuildTargetOutputGetResponseGit].
	Token string `json:"token"`
	// This field is from variant [BuildTargetOutputGetResponseGit].
	Ref  string `json:"ref"`
	JSON struct {
		Output respjson.Field
		Target respjson.Field
		Type   respjson.Field
		URL    respjson.Field
		Token  respjson.Field
		Ref    respjson.Field
		raw    string
	} `json:"-"`
}

// anyBuildTargetOutputGetResponse is implemented by each variant of
// [BuildTargetOutputGetResponseUnion] to add type safety for the return type of
// [BuildTargetOutputGetResponseUnion.AsAny]
type anyBuildTargetOutputGetResponse interface {
	implBuildTargetOutputGetResponseUnion()
}

func (BuildTargetOutputGetResponseURL) implBuildTargetOutputGetResponseUnion() {}
func (BuildTargetOutputGetResponseGit) implBuildTargetOutputGetResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BuildTargetOutputGetResponseUnion.AsAny().(type) {
//	case stainless.BuildTargetOutputGetResponseURL:
//	case stainless.BuildTargetOutputGetResponseGit:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BuildTargetOutputGetResponseUnion) AsAny() anyBuildTargetOutputGetResponse {
	switch u.Output {
	case "url":
		return u.AsURL()
	case "git":
		return u.AsGit()
	}
	return nil
}

func (u BuildTargetOutputGetResponseUnion) AsURL() (v BuildTargetOutputGetResponseURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetOutputGetResponseUnion) AsGit() (v BuildTargetOutputGetResponseGit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetOutputGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetOutputGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetOutputGetResponseURL struct {
	Output constant.URL `json:"output,required"`
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli", "php", "csharp".
	Target shared.Target `json:"target,required"`
	// Any of "source", "dist", "wheel".
	Type BuildTargetOutputGetResponseURLType `json:"type,required"`
	// URL for direct download
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		Target      respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetOutputGetResponseURL) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetOutputGetResponseURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetOutputGetResponseURLType string

const (
	BuildTargetOutputGetResponseURLTypeSource BuildTargetOutputGetResponseURLType = "source"
	BuildTargetOutputGetResponseURLTypeDist   BuildTargetOutputGetResponseURLType = "dist"
	BuildTargetOutputGetResponseURLTypeWheel  BuildTargetOutputGetResponseURLType = "wheel"
)

type BuildTargetOutputGetResponseGit struct {
	// Temporary GitHub access token
	Token  string       `json:"token,required"`
	Output constant.Git `json:"output,required"`
	// Git reference (commit SHA, branch, or tag)
	Ref string `json:"ref,required"`
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli", "php", "csharp".
	Target shared.Target `json:"target,required"`
	// Any of "source", "dist", "wheel".
	Type BuildTargetOutputGetResponseGitType `json:"type,required"`
	// URL to git remote
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Output      respjson.Field
		Ref         respjson.Field
		Target      respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetOutputGetResponseGit) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetOutputGetResponseGit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetOutputGetResponseGitType string

const (
	BuildTargetOutputGetResponseGitTypeSource BuildTargetOutputGetResponseGitType = "source"
	BuildTargetOutputGetResponseGitTypeDist   BuildTargetOutputGetResponseGitType = "dist"
	BuildTargetOutputGetResponseGitTypeWheel  BuildTargetOutputGetResponseGitType = "wheel"
)

type BuildTargetOutputGetResponseType string

const (
	BuildTargetOutputGetResponseTypeSource BuildTargetOutputGetResponseType = "source"
	BuildTargetOutputGetResponseTypeDist   BuildTargetOutputGetResponseType = "dist"
	BuildTargetOutputGetResponseTypeWheel  BuildTargetOutputGetResponseType = "wheel"
)

type BuildTargetOutputGetParams struct {
	// Build ID
	BuildID string `query:"build_id,required" json:"-"`
	// SDK language target name
	//
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli", "php", "csharp".
	Target BuildTargetOutputGetParamsTarget `query:"target,omitzero,required" json:"-"`
	// Any of "source", "dist", "wheel".
	Type BuildTargetOutputGetParamsType `query:"type,omitzero,required" json:"-"`
	// Output format: url (download URL) or git (temporary access token).
	//
	// Any of "url", "git".
	Output BuildTargetOutputGetParamsOutput `query:"output,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BuildTargetOutputGetParams]'s query parameters as
// `url.Values`.
func (r BuildTargetOutputGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// SDK language target name
type BuildTargetOutputGetParamsTarget string

const (
	BuildTargetOutputGetParamsTargetNode       BuildTargetOutputGetParamsTarget = "node"
	BuildTargetOutputGetParamsTargetTypescript BuildTargetOutputGetParamsTarget = "typescript"
	BuildTargetOutputGetParamsTargetPython     BuildTargetOutputGetParamsTarget = "python"
	BuildTargetOutputGetParamsTargetGo         BuildTargetOutputGetParamsTarget = "go"
	BuildTargetOutputGetParamsTargetJava       BuildTargetOutputGetParamsTarget = "java"
	BuildTargetOutputGetParamsTargetKotlin     BuildTargetOutputGetParamsTarget = "kotlin"
	BuildTargetOutputGetParamsTargetRuby       BuildTargetOutputGetParamsTarget = "ruby"
	BuildTargetOutputGetParamsTargetTerraform  BuildTargetOutputGetParamsTarget = "terraform"
	BuildTargetOutputGetParamsTargetCli        BuildTargetOutputGetParamsTarget = "cli"
	BuildTargetOutputGetParamsTargetPhp        BuildTargetOutputGetParamsTarget = "php"
	BuildTargetOutputGetParamsTargetCsharp     BuildTargetOutputGetParamsTarget = "csharp"
)

type BuildTargetOutputGetParamsType string

const (
	BuildTargetOutputGetParamsTypeSource BuildTargetOutputGetParamsType = "source"
	BuildTargetOutputGetParamsTypeDist   BuildTargetOutputGetParamsType = "dist"
	BuildTargetOutputGetParamsTypeWheel  BuildTargetOutputGetParamsType = "wheel"
)

// Output format: url (download URL) or git (temporary access token).
type BuildTargetOutputGetParamsOutput string

const (
	BuildTargetOutputGetParamsOutputURL BuildTargetOutputGetParamsOutput = "url"
	BuildTargetOutputGetParamsOutputGit BuildTargetOutputGetParamsOutput = "git"
)
