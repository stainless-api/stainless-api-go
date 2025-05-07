// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/apiquery"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// BuildTargetOutputService contains methods and other services that help with
// interacting with the stainless-v0 API.
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

// Download the output of a build target
func (r *BuildTargetOutputService) Get(ctx context.Context, query BuildTargetOutputGetParams, opts ...option.RequestOption) (res *BuildTargetOutputGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/build_target_outputs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// BuildTargetOutputGetResponseUnion contains all possible properties and values
// from [BuildTargetOutputGetResponseObject], [BuildTargetOutputGetResponseObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetOutputGetResponseUnion struct {
	// This field is from variant [BuildTargetOutputGetResponseObject].
	Output string `json:"output"`
	// This field is from variant [BuildTargetOutputGetResponseObject].
	URL string `json:"url"`
	// This field is from variant [BuildTargetOutputGetResponseObject].
	Token string `json:"token"`
	// This field is from variant [BuildTargetOutputGetResponseObject].
	Ref  string `json:"ref"`
	JSON struct {
		Output respjson.Field
		URL    respjson.Field
		Token  respjson.Field
		Ref    respjson.Field
		raw    string
	} `json:"-"`
}

func (u BuildTargetOutputGetResponseUnion) AsBuildTargetOutputGetResponseObject() (v BuildTargetOutputGetResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetOutputGetResponseUnion) AsVariant2() (v BuildTargetOutputGetResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetOutputGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetOutputGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetOutputGetResponseObject struct {
	// Any of "url".
	Output string `json:"output,required"`
	// URL for direct download
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetOutputGetResponseObject) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetOutputGetResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetOutputGetParams struct {
	// Build ID
	BuildID string `query:"build_id,required" json:"-"`
	// SDK language target name
	//
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli".
	Target BuildTargetOutputGetParamsTarget `query:"target,omitzero,required" json:"-"`
	// Type of output to download: source code
	//
	// Any of "source".
	Type BuildTargetOutputGetParamsType `query:"type,omitzero,required" json:"-"`
	// Output format: url (download URL) or git (temporary access token)
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
)

// Type of output to download: source code
type BuildTargetOutputGetParamsType string

const (
	BuildTargetOutputGetParamsTypeSource BuildTargetOutputGetParamsType = "source"
)

// Output format: url (download URL) or git (temporary access token)
type BuildTargetOutputGetParamsOutput string

const (
	BuildTargetOutputGetParamsOutputURL BuildTargetOutputGetParamsOutput = "url"
	BuildTargetOutputGetParamsOutputGit BuildTargetOutputGetParamsOutput = "git"
)
