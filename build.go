// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/stainless-v0-go/internal/apijson"
	"github.com/stainless-sdks/stainless-v0-go/internal/apiquery"
	"github.com/stainless-sdks/stainless-v0-go/internal/requestconfig"
	"github.com/stainless-sdks/stainless-v0-go/option"
	"github.com/stainless-sdks/stainless-v0-go/packages/param"
	"github.com/stainless-sdks/stainless-v0-go/packages/resp"
)

// BuildService contains methods and other services that help with interacting with
// the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuildService] method instead.
type BuildService struct {
	Options []option.RequestOption
}

// NewBuildService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBuildService(opts ...option.RequestOption) (r BuildService) {
	r = BuildService{}
	r.Options = opts
	return
}

// TODO
func (r *BuildService) New(ctx context.Context, body BuildNewParams, opts ...option.RequestOption) (res *BuildObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/builds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// TODO
func (r *BuildService) Get(ctx context.Context, buildID string, opts ...option.RequestOption) (res *BuildObject, err error) {
	opts = append(r.Options[:], opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s", buildID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// TODO
func (r *BuildService) List(ctx context.Context, query BuildListParams, opts ...option.RequestOption) (res *BuildListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/builds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BuildObject struct {
	ID           string `json:"id,required"`
	ConfigCommit string `json:"config_commit,required"`
	// Any of "build".
	Object  BuildObjectObject  `json:"object,required"`
	Targets BuildObjectTargets `json:"targets,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		ConfigCommit resp.Field
		Object       resp.Field
		Targets      resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildObject) RawJSON() string { return r.JSON.raw }
func (r *BuildObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildObjectObject string

const (
	BuildObjectObjectBuild BuildObjectObject = "build"
)

type BuildObjectTargets struct {
	Cli        BuildTarget `json:"cli"`
	Go         BuildTarget `json:"go"`
	Java       BuildTarget `json:"java"`
	Kotlin     BuildTarget `json:"kotlin"`
	Node       BuildTarget `json:"node"`
	Python     BuildTarget `json:"python"`
	Ruby       BuildTarget `json:"ruby"`
	Terraform  BuildTarget `json:"terraform"`
	Typescript BuildTarget `json:"typescript"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Cli         resp.Field
		Go          resp.Field
		Java        resp.Field
		Kotlin      resp.Field
		Node        resp.Field
		Python      resp.Field
		Ruby        resp.Field
		Terraform   resp.Field
		Typescript  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildObjectTargets) RawJSON() string { return r.JSON.raw }
func (r *BuildObjectTargets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTarget struct {
	Commit BuildTargetCommitUnion `json:"commit,required"`
	Lint   BuildTargetLintUnion   `json:"lint,required"`
	// Any of "build_target".
	Object BuildTargetObject `json:"object,required"`
	// Any of "not_started", "codegen", "postgen", "completed".
	Status BuildTargetStatus    `json:"status,required"`
	Test   BuildTargetTestUnion `json:"test,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Commit      resp.Field
		Lint        resp.Field
		Object      resp.Field
		Status      resp.Field
		Test        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTarget) RawJSON() string { return r.JSON.raw }
func (r *BuildTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BuildTargetCommitUnion contains all possible properties and values from
// [BuildTargetCommitStatus], [BuildTargetCommitStatus], [BuildTargetCommitStatus],
// [BuildTargetCommitObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetCommitUnion struct {
	Status string `json:"status"`
	// This field is from variant [BuildTargetCommitObject].
	Completed BuildTargetCommitObjectCompletedUnion `json:"completed"`
	JSON      struct {
		Status    resp.Field
		Completed resp.Field
		raw       string
	} `json:"-"`
}

func (u BuildTargetCommitUnion) AsBuildTargetCommitStatus() (v BuildTargetCommitStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitUnion) AsunionMember2() (v BuildTargetCommitStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitUnion) AsunionMember3() (v BuildTargetCommitStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitUnion) AsBuildTargetCommitObject() (v BuildTargetCommitObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetCommitUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetCommitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitStatus struct {
	// Any of "not_started".
	Status string `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitStatus) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitObject struct {
	Completed BuildTargetCommitObjectCompletedUnion `json:"completed,required"`
	// Any of "completed".
	Status string `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Completed   resp.Field
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitObject) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BuildTargetCommitObjectCompletedUnion contains all possible properties and
// values from [BuildTargetCommitObjectCompletedObject],
// [BuildTargetCommitObjectCompletedObject],
// [BuildTargetCommitObjectCompletedConclusion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetCommitObjectCompletedUnion struct {
	// This field is from variant [BuildTargetCommitObjectCompletedObject].
	Commit     BuildTargetCommitObjectCompletedObjectCommit `json:"commit"`
	Conclusion string                                       `json:"conclusion"`
	// This field is from variant [BuildTargetCommitObjectCompletedObject].
	MergeConflictPr BuildTargetCommitObjectCompletedObjectMergeConflictPr `json:"merge_conflict_pr"`
	JSON            struct {
		Commit          resp.Field
		Conclusion      resp.Field
		MergeConflictPr resp.Field
		raw             string
	} `json:"-"`
}

func (u BuildTargetCommitObjectCompletedUnion) AsBuildTargetCommitObjectCompletedObject() (v BuildTargetCommitObjectCompletedObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitObjectCompletedUnion) AsunionMember2() (v BuildTargetCommitObjectCompletedObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitObjectCompletedUnion) AsBuildTargetCommitObjectCompletedConclusion() (v BuildTargetCommitObjectCompletedConclusion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetCommitObjectCompletedUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetCommitObjectCompletedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitObjectCompletedObject struct {
	Commit BuildTargetCommitObjectCompletedObjectCommit `json:"commit,required"`
	// Any of "error", "warning", "note", "success".
	Conclusion string `json:"conclusion,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Commit      resp.Field
		Conclusion  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitObjectCompletedObject) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitObjectCompletedObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitObjectCompletedObjectCommit struct {
	Repo BuildTargetCommitObjectCompletedObjectCommitRepo `json:"repo,required"`
	Sha  string                                           `json:"sha,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Repo        resp.Field
		Sha         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitObjectCompletedObjectCommit) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitObjectCompletedObjectCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitObjectCompletedObjectCommitRepo struct {
	Branch string `json:"branch,required"`
	Name   string `json:"name,required"`
	Owner  string `json:"owner,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Branch      resp.Field
		Name        resp.Field
		Owner       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitObjectCompletedObjectCommitRepo) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitObjectCompletedObjectCommitRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitObjectCompletedConclusion struct {
	// Any of "fatal", "payment_required", "cancelled", "timed_out", "noop",
	// "version_bump".
	Conclusion string `json:"conclusion,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Conclusion  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitObjectCompletedConclusion) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitObjectCompletedConclusion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BuildTargetLintUnion contains all possible properties and values from
// [BuildTargetLintStatus], [BuildTargetLintStatus], [BuildTargetLintStatus],
// [BuildTargetLintObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetLintUnion struct {
	Status string `json:"status"`
	// This field is from variant [BuildTargetLintObject].
	Completed BuildTargetLintObjectCompleted `json:"completed"`
	JSON      struct {
		Status    resp.Field
		Completed resp.Field
		raw       string
	} `json:"-"`
}

func (u BuildTargetLintUnion) AsBuildTargetLintStatus() (v BuildTargetLintStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetLintUnion) AsunionMember2() (v BuildTargetLintStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetLintUnion) AsunionMember3() (v BuildTargetLintStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetLintUnion) AsBuildTargetLintObject() (v BuildTargetLintObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetLintUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetLintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintStatus struct {
	// Any of "not_started".
	Status string `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintStatus) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintObject struct {
	Completed BuildTargetLintObjectCompleted `json:"completed,required"`
	// Any of "completed".
	Status string `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Completed   resp.Field
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintObject) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintObjectCompleted struct {
	// Any of "success", "failure", "skipped", "cancelled", "action_required",
	// "neutral", "timed_out".
	Conclusion string `json:"conclusion,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Conclusion  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintObjectCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintObjectCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetObject string

const (
	BuildTargetObjectBuildTarget BuildTargetObject = "build_target"
)

type BuildTargetStatus string

const (
	BuildTargetStatusNotStarted BuildTargetStatus = "not_started"
	BuildTargetStatusCodegen    BuildTargetStatus = "codegen"
	BuildTargetStatusPostgen    BuildTargetStatus = "postgen"
	BuildTargetStatusCompleted  BuildTargetStatus = "completed"
)

// BuildTargetTestUnion contains all possible properties and values from
// [BuildTargetTestStatus], [BuildTargetTestStatus], [BuildTargetTestStatus],
// [BuildTargetTestObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetTestUnion struct {
	Status string `json:"status"`
	// This field is from variant [BuildTargetTestObject].
	Completed BuildTargetTestObjectCompleted `json:"completed"`
	JSON      struct {
		Status    resp.Field
		Completed resp.Field
		raw       string
	} `json:"-"`
}

func (u BuildTargetTestUnion) AsBuildTargetTestStatus() (v BuildTargetTestStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetTestUnion) AsunionMember2() (v BuildTargetTestStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetTestUnion) AsunionMember3() (v BuildTargetTestStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetTestUnion) AsBuildTargetTestObject() (v BuildTargetTestObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetTestUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetTestUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestStatus struct {
	// Any of "not_started".
	Status string `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestStatus) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestObject struct {
	Completed BuildTargetTestObjectCompleted `json:"completed,required"`
	// Any of "completed".
	Status string `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Completed   resp.Field
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestObject) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestObjectCompleted struct {
	// Any of "success", "failure", "skipped", "cancelled", "action_required",
	// "neutral", "timed_out".
	Conclusion string `json:"conclusion,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Conclusion  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestObjectCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestObjectCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildListResponse struct {
	Data       []BuildObject `json:"data,required"`
	HasMore    bool          `json:"has_more,required"`
	NextCursor string        `json:"next_cursor"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data        resp.Field
		HasMore     resp.Field
		NextCursor  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildListResponse) RawJSON() string { return r.JSON.raw }
func (r *BuildListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildNewParams struct {
	// Project name
	Project string `json:"project,required"`
	// Specifies what to build: a branch name, commit SHA, merge command
	// ("base..head"), or file contents
	Revision BuildNewParamsRevisionUnion `json:"revision,omitzero,required"`
	// Whether to allow empty commits (no changes). Defaults to false.
	AllowEmpty param.Opt[bool] `json:"allow_empty,omitzero"`
	// Optional branch to use. If not specified, defaults to "main". When using a
	// branch name or merge command as revision, this must match or be omitted.
	Branch param.Opt[string] `json:"branch,omitzero"`
	// Optional commit message to use when creating a new commit.
	CommitMessage param.Opt[string] `json:"commit_message,omitzero"`
	// Optional list of SDK targets to build. If not specified, all configured targets
	// will be built.
	//
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli".
	Targets []string `json:"targets,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BuildNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r BuildNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BuildNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BuildNewParamsRevisionUnion struct {
	OfString               param.Opt[string]                        `json:",omitzero,inline"`
	OfBuildNewsRevisionMap map[string]BuildNewParamsRevisionMapItem `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u BuildNewParamsRevisionUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u BuildNewParamsRevisionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BuildNewParamsRevisionUnion](u.OfString, u.OfBuildNewsRevisionMap)
}

func (u *BuildNewParamsRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBuildNewsRevisionMap) {
		return &u.OfBuildNewsRevisionMap
	}
	return nil
}

// The property Content is required.
type BuildNewParamsRevisionMapItem struct {
	// The file content
	Content string `json:"content,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BuildNewParamsRevisionMapItem) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r BuildNewParamsRevisionMapItem) MarshalJSON() (data []byte, err error) {
	type shadow BuildNewParamsRevisionMapItem
	return param.MarshalObject(r, (*shadow)(&r))
}

type BuildListParams struct {
	// Project name
	Project string `query:"project,required" json:"-"`
	// Branch name, defaults to "main"
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Pagination cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of builds to return, defaults to 10
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BuildListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [BuildListParams]'s query parameters as `url.Values`.
func (r BuildListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
