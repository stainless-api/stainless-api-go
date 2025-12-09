// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/apiquery"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/pagination"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
	"github.com/stainless-api/stainless-api-go/shared/constant"
)

// BuildDiagnosticService contains methods and other services that help with
// interacting with the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuildDiagnosticService] method instead.
type BuildDiagnosticService struct {
	Options []option.RequestOption
}

// NewBuildDiagnosticService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBuildDiagnosticService(opts ...option.RequestOption) (r BuildDiagnosticService) {
	r = BuildDiagnosticService{}
	r.Options = opts
	return
}

// Get the list of diagnostics for a given build.
//
// If no language targets are specified, diagnostics for all languages are
// returned.
func (r *BuildDiagnosticService) List(ctx context.Context, buildID string, query BuildDiagnosticListParams, opts ...option.RequestOption) (res *pagination.Page[BuildDiagnostic], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s/diagnostics", url.PathEscape(buildID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get the list of diagnostics for a given build.
//
// If no language targets are specified, diagnostics for all languages are
// returned.
func (r *BuildDiagnosticService) ListAutoPaging(ctx context.Context, buildID string, query BuildDiagnosticListParams, opts ...option.RequestOption) *pagination.PageAutoPager[BuildDiagnostic] {
	return pagination.NewPageAutoPager(r.List(ctx, buildID, query, opts...))
}

type BuildDiagnostic struct {
	// The kind of diagnostic.
	Code string `json:"code,required"`
	// Whether the diagnostic is ignored in the Stainless config.
	Ignored bool `json:"ignored,required"`
	// The severity of the diagnostic.
	//
	// Any of "fatal", "error", "warning", "note".
	Level BuildDiagnosticLevel `json:"level,required"`
	// A description of the diagnostic.
	Message string                   `json:"message,required"`
	More    BuildDiagnosticMoreUnion `json:"more,required"`
	// A JSON pointer to a relevant field in the Stainless config.
	ConfigRef string `json:"config_ref"`
	// A JSON pointer to a relevant field in the OpenAPI spec.
	OasRef string `json:"oas_ref"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Ignored     respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		More        respjson.Field
		ConfigRef   respjson.Field
		OasRef      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildDiagnostic) RawJSON() string { return r.JSON.raw }
func (r *BuildDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The severity of the diagnostic.
type BuildDiagnosticLevel string

const (
	BuildDiagnosticLevelFatal   BuildDiagnosticLevel = "fatal"
	BuildDiagnosticLevelError   BuildDiagnosticLevel = "error"
	BuildDiagnosticLevelWarning BuildDiagnosticLevel = "warning"
	BuildDiagnosticLevelNote    BuildDiagnosticLevel = "note"
)

// BuildDiagnosticMoreUnion contains all possible properties and values from
// [BuildDiagnosticMoreMarkdown], [BuildDiagnosticMoreRaw].
//
// Use the [BuildDiagnosticMoreUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildDiagnosticMoreUnion struct {
	// This field is from variant [BuildDiagnosticMoreMarkdown].
	Markdown string `json:"markdown"`
	// Any of "markdown", "raw".
	Type string `json:"type"`
	// This field is from variant [BuildDiagnosticMoreRaw].
	Raw  string `json:"raw"`
	JSON struct {
		Markdown respjson.Field
		Type     respjson.Field
		Raw      respjson.Field
		raw      string
	} `json:"-"`
}

// anyBuildDiagnosticMore is implemented by each variant of
// [BuildDiagnosticMoreUnion] to add type safety for the return type of
// [BuildDiagnosticMoreUnion.AsAny]
type anyBuildDiagnosticMore interface {
	implBuildDiagnosticMoreUnion()
}

func (BuildDiagnosticMoreMarkdown) implBuildDiagnosticMoreUnion() {}
func (BuildDiagnosticMoreRaw) implBuildDiagnosticMoreUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BuildDiagnosticMoreUnion.AsAny().(type) {
//	case stainless.BuildDiagnosticMoreMarkdown:
//	case stainless.BuildDiagnosticMoreRaw:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BuildDiagnosticMoreUnion) AsAny() anyBuildDiagnosticMore {
	switch u.Type {
	case "markdown":
		return u.AsMarkdown()
	case "raw":
		return u.AsRaw()
	}
	return nil
}

func (u BuildDiagnosticMoreUnion) AsMarkdown() (v BuildDiagnosticMoreMarkdown) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildDiagnosticMoreUnion) AsRaw() (v BuildDiagnosticMoreRaw) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildDiagnosticMoreUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildDiagnosticMoreUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildDiagnosticMoreMarkdown struct {
	Markdown string            `json:"markdown,required"`
	Type     constant.Markdown `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildDiagnosticMoreMarkdown) RawJSON() string { return r.JSON.raw }
func (r *BuildDiagnosticMoreMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildDiagnosticMoreRaw struct {
	Raw  string       `json:"raw,required"`
	Type constant.Raw `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Raw         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildDiagnosticMoreRaw) RawJSON() string { return r.JSON.raw }
func (r *BuildDiagnosticMoreRaw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildDiagnosticListParams struct {
	// Pagination cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of diagnostics to return, defaults to 100 (maximum: 100)
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Optional comma-delimited list of language targets to filter diagnostics by
	Targets param.Opt[string] `query:"targets,omitzero" json:"-"`
	// Includes the given severity and above (fatal > error > warning > note).
	//
	// Any of "fatal", "error", "warning", "note".
	Severity BuildDiagnosticListParamsSeverity `query:"severity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BuildDiagnosticListParams]'s query parameters as
// `url.Values`.
func (r BuildDiagnosticListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Includes the given severity and above (fatal > error > warning > note).
type BuildDiagnosticListParamsSeverity string

const (
	BuildDiagnosticListParamsSeverityFatal   BuildDiagnosticListParamsSeverity = "fatal"
	BuildDiagnosticListParamsSeverityError   BuildDiagnosticListParamsSeverity = "error"
	BuildDiagnosticListParamsSeverityWarning BuildDiagnosticListParamsSeverity = "warning"
	BuildDiagnosticListParamsSeverityNote    BuildDiagnosticListParamsSeverity = "note"
)
