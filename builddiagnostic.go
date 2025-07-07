// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/apiquery"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
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

// Get diagnostics for a build
func (r *BuildDiagnosticService) List(ctx context.Context, buildID string, query BuildDiagnosticListParams, opts ...option.RequestOption) (res *BuildDiagnosticListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s/diagnostics", buildID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BuildDiagnosticListResponse struct {
	Data       []BuildDiagnosticListResponseData `json:"data,required"`
	HasMore    bool                              `json:"has_more,required"`
	NextCursor string                            `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildDiagnosticListResponse) RawJSON() string { return r.JSON.raw }
func (r *BuildDiagnosticListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildDiagnosticListResponseData struct {
	Code    string `json:"code,required"`
	Ignored bool   `json:"ignored,required"`
	// Any of "fatal", "error", "warning", "note".
	Level     string `json:"level,required"`
	Message   string `json:"message,required"`
	ConfigRef string `json:"config_ref"`
	OasRef    string `json:"oas_ref"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Ignored     respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		ConfigRef   respjson.Field
		OasRef      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildDiagnosticListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BuildDiagnosticListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildDiagnosticListParams struct {
	// Pagination cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of diagnostics to return, defaults to 100 (maximum: 100)
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Includes the given severity and above (fatal > error > warning > note).
	//
	// Any of "fatal", "error", "warning", "note".
	Severity BuildDiagnosticListParamsSeverity `query:"severity,omitzero" json:"-"`
	// Optional list of language targets to filter diagnostics by
	//
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli", "php", "csharp".
	Targets []string `query:"targets,omitzero" json:"-"`
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
