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
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/resp"
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

// TODO
func (r *BuildTargetOutputService) List(ctx context.Context, query BuildTargetOutputListParams, opts ...option.RequestOption) (res *BuildTargetOutputListResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/build_target_outputs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// BuildTargetOutputListResponseUnion contains all possible properties and values
// from [BuildTargetOutputListResponseObject],
// [BuildTargetOutputListResponseObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetOutputListResponseUnion struct {
	// This field is from variant [BuildTargetOutputListResponseObject].
	Output string `json:"output"`
	// This field is from variant [BuildTargetOutputListResponseObject].
	URL string `json:"url"`
	// This field is from variant [BuildTargetOutputListResponseObject].
	Token string `json:"token"`
	// This field is from variant [BuildTargetOutputListResponseObject].
	Ref string `json:"ref"`
	// This field is from variant [BuildTargetOutputListResponseObject].
	Repo string `json:"repo"`
	JSON struct {
		Output resp.Field
		URL    resp.Field
		Token  resp.Field
		Ref    resp.Field
		Repo   resp.Field
		raw    string
	} `json:"-"`
}

func (u BuildTargetOutputListResponseUnion) AsBuildTargetOutputListResponseObject() (v BuildTargetOutputListResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetOutputListResponseUnion) AsunionMember2() (v BuildTargetOutputListResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetOutputListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetOutputListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetOutputListResponseObject struct {
	// Any of "url".
	Output string `json:"output,required"`
	// URL for direct download
	URL string `json:"url,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Output      resp.Field
		URL         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetOutputListResponseObject) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetOutputListResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetOutputListParams struct {
	// Build ID
	BuildID string `query:"build_id,required" json:"-"`
	// SDK language target name
	//
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli".
	Target BuildTargetOutputListParamsTarget `query:"target,omitzero,required" json:"-"`
	// Type of output to download: source code
	//
	// Any of "source".
	Type BuildTargetOutputListParamsType `query:"type,omitzero,required" json:"-"`
	// Output format: url (download URL) or git (temporary access token)
	//
	// Any of "url", "git".
	Output BuildTargetOutputListParamsOutput `query:"output,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BuildTargetOutputListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [BuildTargetOutputListParams]'s query parameters as
// `url.Values`.
func (r BuildTargetOutputListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// SDK language target name
type BuildTargetOutputListParamsTarget string

const (
	BuildTargetOutputListParamsTargetNode       BuildTargetOutputListParamsTarget = "node"
	BuildTargetOutputListParamsTargetTypescript BuildTargetOutputListParamsTarget = "typescript"
	BuildTargetOutputListParamsTargetPython     BuildTargetOutputListParamsTarget = "python"
	BuildTargetOutputListParamsTargetGo         BuildTargetOutputListParamsTarget = "go"
	BuildTargetOutputListParamsTargetJava       BuildTargetOutputListParamsTarget = "java"
	BuildTargetOutputListParamsTargetKotlin     BuildTargetOutputListParamsTarget = "kotlin"
	BuildTargetOutputListParamsTargetRuby       BuildTargetOutputListParamsTarget = "ruby"
	BuildTargetOutputListParamsTargetTerraform  BuildTargetOutputListParamsTarget = "terraform"
	BuildTargetOutputListParamsTargetCli        BuildTargetOutputListParamsTarget = "cli"
)

// Type of output to download: source code
type BuildTargetOutputListParamsType string

const (
	BuildTargetOutputListParamsTypeSource BuildTargetOutputListParamsType = "source"
)

// Output format: url (download URL) or git (temporary access token)
type BuildTargetOutputListParamsOutput string

const (
	BuildTargetOutputListParamsOutputURL BuildTargetOutputListParamsOutput = "url"
	BuildTargetOutputListParamsOutputGit BuildTargetOutputListParamsOutput = "git"
)
