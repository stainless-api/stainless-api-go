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
	"time"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/apiquery"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/pagination"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
	"github.com/stainless-api/stainless-api-go/shared"
	"github.com/stainless-api/stainless-api-go/shared/constant"
)

// BuildService contains methods and other services that help with interacting with
// the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuildService] method instead.
type BuildService struct {
	Options       []option.RequestOption
	Diagnostics   BuildDiagnosticService
	TargetOutputs BuildTargetOutputService
}

// NewBuildService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBuildService(opts ...option.RequestOption) (r BuildService) {
	r = BuildService{}
	r.Options = opts
	r.Diagnostics = NewBuildDiagnosticService(opts...)
	r.TargetOutputs = NewBuildTargetOutputService(opts...)
	return
}

// Create a build, on top of a project branch, against a given input revision.
//
// The project branch will be modified so that its latest set of config files
// points to the one specified by the input revision.
func (r *BuildService) New(ctx context.Context, body BuildNewParams, opts ...option.RequestOption) (res *Build, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.Project, precfg.Project)
	path := "v0/builds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a build by its ID.
func (r *BuildService) Get(ctx context.Context, buildID string, opts ...option.RequestOption) (res *Build, err error) {
	opts = slices.Concat(r.Options, opts)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s", url.PathEscape(buildID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List user-triggered builds for a given project.
//
// An optional revision can be specified to filter by config commit SHA, or hashes
// of file contents.
func (r *BuildService) List(ctx context.Context, query BuildListParams, opts ...option.RequestOption) (res *pagination.Page[Build], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.Project, precfg.Project)
	path := "v0/builds"
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

// List user-triggered builds for a given project.
//
// An optional revision can be specified to filter by config commit SHA, or hashes
// of file contents.
func (r *BuildService) ListAutoPaging(ctx context.Context, query BuildListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Build] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create two builds whose outputs can be directly compared with each other.
//
// Created builds _modify_ their project branches so that their latest sets of
// config files point to the ones specified by the input revision.
//
// This endpoint is useful because a build has more inputs than the set of config
// files it uses, so comparing two builds directly may return spurious differences.
// Builds made via this endpoint are guaranteed to have differences arising from
// the set of config files, and any custom code.
func (r *BuildService) Compare(ctx context.Context, body BuildCompareParams, opts ...option.RequestOption) (res *BuildCompareResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.Project, precfg.Project)
	path := "v0/builds/compare"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Build struct {
	// Build ID
	ID             string                   `json:"id,required"`
	ConfigCommit   string                   `json:"config_commit,required"`
	CreatedAt      time.Time                `json:"created_at,required" format:"date-time"`
	DocumentedSpec BuildDocumentedSpecUnion `json:"documented_spec,required"`
	// Any of "build".
	Object    BuildObject  `json:"object,required"`
	Org       string       `json:"org,required"`
	Project   string       `json:"project,required"`
	Targets   BuildTargets `json:"targets,required"`
	UpdatedAt time.Time    `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ConfigCommit   respjson.Field
		CreatedAt      respjson.Field
		DocumentedSpec respjson.Field
		Object         respjson.Field
		Org            respjson.Field
		Project        respjson.Field
		Targets        respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Build) RawJSON() string { return r.JSON.raw }
func (r *Build) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BuildDocumentedSpecUnion contains all possible properties and values from
// [BuildDocumentedSpecObject], [BuildDocumentedSpecObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildDocumentedSpecUnion struct {
	// This field is from variant [BuildDocumentedSpecObject].
	Content string `json:"content"`
	// This field is from variant [BuildDocumentedSpecObject].
	Type string `json:"type"`
	// This field is from variant [BuildDocumentedSpecObject].
	Expires time.Time `json:"expires"`
	// This field is from variant [BuildDocumentedSpecObject].
	URL  string `json:"url"`
	JSON struct {
		Content respjson.Field
		Type    respjson.Field
		Expires respjson.Field
		URL     respjson.Field
		raw     string
	} `json:"-"`
}

func (u BuildDocumentedSpecUnion) AsBuildDocumentedSpecObject() (v BuildDocumentedSpecObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildDocumentedSpecUnion) AsVariant2() (v BuildDocumentedSpecObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildDocumentedSpecUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildDocumentedSpecUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildDocumentedSpecObject struct {
	Content string `json:"content,required"`
	// Any of "content".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildDocumentedSpecObject) RawJSON() string { return r.JSON.raw }
func (r *BuildDocumentedSpecObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildObject string

const (
	BuildObjectBuild BuildObject = "build"
)

type BuildTargets struct {
	Cli        BuildTarget `json:"cli"`
	Csharp     BuildTarget `json:"csharp"`
	Go         BuildTarget `json:"go"`
	Java       BuildTarget `json:"java"`
	Kotlin     BuildTarget `json:"kotlin"`
	Node       BuildTarget `json:"node"`
	OpenAPI    BuildTarget `json:"openapi"`
	Php        BuildTarget `json:"php"`
	Python     BuildTarget `json:"python"`
	Ruby       BuildTarget `json:"ruby"`
	Sql        BuildTarget `json:"sql"`
	Terraform  BuildTarget `json:"terraform"`
	Typescript BuildTarget `json:"typescript"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cli         respjson.Field
		Csharp      respjson.Field
		Go          respjson.Field
		Java        respjson.Field
		Kotlin      respjson.Field
		Node        respjson.Field
		OpenAPI     respjson.Field
		Php         respjson.Field
		Python      respjson.Field
		Ruby        respjson.Field
		Sql         respjson.Field
		Terraform   respjson.Field
		Typescript  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargets) RawJSON() string { return r.JSON.raw }
func (r *BuildTargets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTarget struct {
	Commit     BuildTargetCommitUnion `json:"commit,required"`
	InstallURL string                 `json:"install_url,required"`
	Lint       CheckStepUnion         `json:"lint,required"`
	// Any of "build_target".
	Object BuildTargetObject `json:"object,required"`
	// Any of "not_started", "codegen", "postgen", "completed".
	Status BuildTargetStatus `json:"status,required"`
	Test   CheckStepUnion    `json:"test,required"`
	Build  CheckStepUnion    `json:"build"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit      respjson.Field
		InstallURL  respjson.Field
		Lint        respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		Test        respjson.Field
		Build       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTarget) RawJSON() string { return r.JSON.raw }
func (r *BuildTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BuildTargetCommitUnion contains all possible properties and values from
// [BuildTargetCommitNotStarted], [BuildTargetCommitQueued],
// [BuildTargetCommitInProgress], [BuildTargetCommitCompleted].
//
// Use the [BuildTargetCommitUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetCommitUnion struct {
	// Any of "not_started", "queued", "in_progress", "completed".
	Status string `json:"status"`
	// This field is from variant [BuildTargetCommitCompleted].
	Commit shared.Commit `json:"commit"`
	// This field is from variant [BuildTargetCommitCompleted].
	Completed BuildTargetCommitCompletedCompleted `json:"completed"`
	// This field is from variant [BuildTargetCommitCompleted].
	Conclusion string `json:"conclusion"`
	// This field is from variant [BuildTargetCommitCompleted].
	MergeConflictPr BuildTargetCommitCompletedMergeConflictPr `json:"merge_conflict_pr"`
	JSON            struct {
		Status          respjson.Field
		Commit          respjson.Field
		Completed       respjson.Field
		Conclusion      respjson.Field
		MergeConflictPr respjson.Field
		raw             string
	} `json:"-"`
}

// anyBuildTargetCommit is implemented by each variant of [BuildTargetCommitUnion]
// to add type safety for the return type of [BuildTargetCommitUnion.AsAny]
type anyBuildTargetCommit interface {
	implBuildTargetCommitUnion()
}

func (BuildTargetCommitNotStarted) implBuildTargetCommitUnion() {}
func (BuildTargetCommitQueued) implBuildTargetCommitUnion()     {}
func (BuildTargetCommitInProgress) implBuildTargetCommitUnion() {}
func (BuildTargetCommitCompleted) implBuildTargetCommitUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BuildTargetCommitUnion.AsAny().(type) {
//	case stainless.BuildTargetCommitNotStarted:
//	case stainless.BuildTargetCommitQueued:
//	case stainless.BuildTargetCommitInProgress:
//	case stainless.BuildTargetCommitCompleted:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BuildTargetCommitUnion) AsAny() anyBuildTargetCommit {
	switch u.Status {
	case "not_started":
		return u.AsNotStarted()
	case "queued":
		return u.AsQueued()
	case "in_progress":
		return u.AsInProgress()
	case "completed":
		return u.AsCompleted()
	}
	return nil
}

func (u BuildTargetCommitUnion) AsNotStarted() (v BuildTargetCommitNotStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitUnion) AsQueued() (v BuildTargetCommitQueued) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitUnion) AsInProgress() (v BuildTargetCommitInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitUnion) AsCompleted() (v BuildTargetCommitCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetCommitUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetCommitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitNotStarted struct {
	Status constant.NotStarted `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitNotStarted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitNotStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitQueued struct {
	Status constant.Queued `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitQueued) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitQueued) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitInProgress struct {
	Status constant.InProgress `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitInProgress) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompleted struct {
	Commit shared.Commit `json:"commit,required"`
	// deprecated
	Completed BuildTargetCommitCompletedCompleted `json:"completed,required"`
	// Any of "error", "warning", "note", "success", "merge_conflict",
	// "upstream_merge_conflict", "fatal", "payment_required", "cancelled",
	// "timed_out", "noop", "version_bump".
	Conclusion      string                                    `json:"conclusion,required"`
	MergeConflictPr BuildTargetCommitCompletedMergeConflictPr `json:"merge_conflict_pr,required"`
	Status          constant.Completed                        `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit          respjson.Field
		Completed       respjson.Field
		Conclusion      respjson.Field
		MergeConflictPr respjson.Field
		Status          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// deprecated
type BuildTargetCommitCompletedCompleted struct {
	Commit shared.Commit `json:"commit,required"`
	// Any of "error", "warning", "note", "success", "merge_conflict",
	// "upstream_merge_conflict", "fatal", "payment_required", "cancelled",
	// "timed_out", "noop", "version_bump".
	Conclusion      string                                             `json:"conclusion,required"`
	MergeConflictPr BuildTargetCommitCompletedCompletedMergeConflictPr `json:"merge_conflict_pr,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit          respjson.Field
		Conclusion      respjson.Field
		MergeConflictPr respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedMergeConflictPr struct {
	Number float64                                                `json:"number,required"`
	Repo   BuildTargetCommitCompletedCompletedMergeConflictPrRepo `json:"repo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Repo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompletedMergeConflictPr) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedMergeConflictPr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedMergeConflictPrRepo struct {
	Name  string `json:"name,required"`
	Owner string `json:"owner,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Owner       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompletedMergeConflictPrRepo) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedMergeConflictPrRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedMergeConflictPr struct {
	Number float64                                       `json:"number,required"`
	Repo   BuildTargetCommitCompletedMergeConflictPrRepo `json:"repo,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Repo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedMergeConflictPr) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedMergeConflictPr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedMergeConflictPrRepo struct {
	Name  string `json:"name,required"`
	Owner string `json:"owner,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Owner       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedMergeConflictPrRepo) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedMergeConflictPrRepo) UnmarshalJSON(data []byte) error {
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

// CheckStepUnion contains all possible properties and values from
// [CheckStepNotStarted], [CheckStepQueued], [CheckStepInProgress],
// [CheckStepCompleted].
//
// Use the [CheckStepUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CheckStepUnion struct {
	// Any of "not_started", "queued", "in_progress", "completed".
	Status string `json:"status"`
	// This field is from variant [CheckStepCompleted].
	Completed CheckStepCompletedCompleted `json:"completed"`
	// This field is from variant [CheckStepCompleted].
	Conclusion string `json:"conclusion"`
	// This field is from variant [CheckStepCompleted].
	URL  string `json:"url"`
	JSON struct {
		Status     respjson.Field
		Completed  respjson.Field
		Conclusion respjson.Field
		URL        respjson.Field
		raw        string
	} `json:"-"`
}

// anyCheckStep is implemented by each variant of [CheckStepUnion] to add type
// safety for the return type of [CheckStepUnion.AsAny]
type anyCheckStep interface {
	implCheckStepUnion()
}

func (CheckStepNotStarted) implCheckStepUnion() {}
func (CheckStepQueued) implCheckStepUnion()     {}
func (CheckStepInProgress) implCheckStepUnion() {}
func (CheckStepCompleted) implCheckStepUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CheckStepUnion.AsAny().(type) {
//	case stainless.CheckStepNotStarted:
//	case stainless.CheckStepQueued:
//	case stainless.CheckStepInProgress:
//	case stainless.CheckStepCompleted:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CheckStepUnion) AsAny() anyCheckStep {
	switch u.Status {
	case "not_started":
		return u.AsNotStarted()
	case "queued":
		return u.AsQueued()
	case "in_progress":
		return u.AsInProgress()
	case "completed":
		return u.AsCompleted()
	}
	return nil
}

func (u CheckStepUnion) AsNotStarted() (v CheckStepNotStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckStepUnion) AsQueued() (v CheckStepQueued) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckStepUnion) AsInProgress() (v CheckStepInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckStepUnion) AsCompleted() (v CheckStepCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CheckStepUnion) RawJSON() string { return u.JSON.raw }

func (r *CheckStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckStepNotStarted struct {
	Status constant.NotStarted `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckStepNotStarted) RawJSON() string { return r.JSON.raw }
func (r *CheckStepNotStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckStepQueued struct {
	Status constant.Queued `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckStepQueued) RawJSON() string { return r.JSON.raw }
func (r *CheckStepQueued) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckStepInProgress struct {
	Status constant.InProgress `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckStepInProgress) RawJSON() string { return r.JSON.raw }
func (r *CheckStepInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckStepCompleted struct {
	// deprecated
	Completed CheckStepCompletedCompleted `json:"completed,required"`
	// Any of "success", "failure", "skipped", "cancelled", "action_required",
	// "neutral", "timed_out".
	Conclusion string             `json:"conclusion,required"`
	Status     constant.Completed `json:"status,required"`
	URL        string             `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Conclusion  respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckStepCompleted) RawJSON() string { return r.JSON.raw }
func (r *CheckStepCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// deprecated
type CheckStepCompletedCompleted struct {
	// Any of "success", "failure", "skipped", "cancelled", "action_required",
	// "neutral", "timed_out".
	Conclusion string `json:"conclusion,required"`
	URL        string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conclusion  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckStepCompletedCompleted) RawJSON() string { return r.JSON.raw }
func (r *CheckStepCompletedCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildCompareResponse struct {
	Base Build `json:"base,required"`
	Head Build `json:"head,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base        respjson.Field
		Head        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildCompareResponse) RawJSON() string { return r.JSON.raw }
func (r *BuildCompareResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildNewParams struct {
	// Project name
	Project param.Opt[string] `json:"project,omitzero,required"`
	// Specifies what to build: a branch name, commit SHA, merge command
	// ("base..head"), or file contents.
	Revision BuildNewParamsRevisionUnion `json:"revision,omitzero,required"`
	// Whether to allow empty commits (no changes). Defaults to false.
	AllowEmpty param.Opt[bool] `json:"allow_empty,omitzero"`
	// The project branch to use for the build. If not specified, the branch is
	// inferred from the `revision`, and will 400 when that is not possible.
	Branch param.Opt[string] `json:"branch,omitzero"`
	// Optional commit message to use when creating a new commit.
	CommitMessage param.Opt[string] `json:"commit_message,omitzero"`
	// Whether to generate AI-powered commit messages for the build. Cannot be combined
	// with `commit_message` or `target_commit_messages`.
	EnableAICommitMessage param.Opt[bool] `json:"enable_ai_commit_message,omitzero"`
	// Optional commit messages to use for each SDK when making a new commit. SDKs not
	// represented in this object will fallback to the optional `commit_message`
	// parameter, or will fallback further to the default commit message.
	TargetCommitMessages BuildNewParamsTargetCommitMessages `json:"target_commit_messages,omitzero"`
	// Optional list of SDK targets to build. If not specified, all configured targets
	// will be built.
	Targets []shared.Target `json:"targets,omitzero"`
	paramObj
}

func (r BuildNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BuildNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BuildNewParamsRevisionUnion struct {
	OfString       param.Opt[string]                     `json:",omitzero,inline"`
	OfFileInputMap map[string]shared.FileInputUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u BuildNewParamsRevisionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFileInputMap)
}
func (u *BuildNewParamsRevisionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BuildNewParamsRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFileInputMap) {
		return &u.OfFileInputMap
	}
	return nil
}

// Optional commit messages to use for each SDK when making a new commit. SDKs not
// represented in this object will fallback to the optional `commit_message`
// parameter, or will fallback further to the default commit message.
type BuildNewParamsTargetCommitMessages struct {
	Cli        param.Opt[string] `json:"cli,omitzero"`
	Csharp     param.Opt[string] `json:"csharp,omitzero"`
	Go         param.Opt[string] `json:"go,omitzero"`
	Java       param.Opt[string] `json:"java,omitzero"`
	Kotlin     param.Opt[string] `json:"kotlin,omitzero"`
	Node       param.Opt[string] `json:"node,omitzero"`
	OpenAPI    param.Opt[string] `json:"openapi,omitzero"`
	Php        param.Opt[string] `json:"php,omitzero"`
	Python     param.Opt[string] `json:"python,omitzero"`
	Ruby       param.Opt[string] `json:"ruby,omitzero"`
	Sql        param.Opt[string] `json:"sql,omitzero"`
	Terraform  param.Opt[string] `json:"terraform,omitzero"`
	Typescript param.Opt[string] `json:"typescript,omitzero"`
	paramObj
}

func (r BuildNewParamsTargetCommitMessages) MarshalJSON() (data []byte, err error) {
	type shadow BuildNewParamsTargetCommitMessages
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildNewParamsTargetCommitMessages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildListParams struct {
	// Project name
	//
	// Use [option.WithProject] on the client to set a global default for this field.
	Project param.Opt[string] `query:"project,omitzero,required" json:"-"`
	// Branch name
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Pagination cursor from a previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of builds to return, defaults to 10 (maximum: 100).
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// A config commit SHA used for the build
	Revision BuildListParamsRevisionUnion `query:"revision,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BuildListParams]'s query parameters as `url.Values`.
func (r BuildListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BuildListParamsRevisionUnion struct {
	OfString                   param.Opt[string]                         `query:",omitzero,inline"`
	OfBuildListsRevisionMapMap map[string]BuildListParamsRevisionMapItem `query:",omitzero,inline"`
	paramUnion
}

func (u *BuildListParamsRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBuildListsRevisionMapMap) {
		return &u.OfBuildListsRevisionMapMap
	}
	return nil
}

// The property Hash is required.
type BuildListParamsRevisionMapItem struct {
	// File content hash
	Hash string `query:"hash,required" json:"-"`
	paramObj
}

// URLQuery serializes [BuildListParamsRevisionMapItem]'s query parameters as
// `url.Values`.
func (r BuildListParamsRevisionMapItem) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BuildCompareParams struct {
	// Project name
	Project param.Opt[string] `json:"project,omitzero,required"`
	// Parameters for the base build
	Base BuildCompareParamsBase `json:"base,omitzero,required"`
	// Parameters for the head build
	Head BuildCompareParamsHead `json:"head,omitzero,required"`
	// Optional list of SDK targets to build. If not specified, all configured targets
	// will be built.
	Targets []shared.Target `json:"targets,omitzero"`
	paramObj
}

func (r BuildCompareParams) MarshalJSON() (data []byte, err error) {
	type shadow BuildCompareParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildCompareParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the base build
//
// The properties Branch, Revision are required.
type BuildCompareParamsBase struct {
	// Branch to use. When using a branch name as revision, this must match or be
	// omitted.
	Branch string `json:"branch,required"`
	// Specifies what to build: a branch name, a commit SHA, or file contents.
	Revision BuildCompareParamsBaseRevisionUnion `json:"revision,omitzero,required"`
	// Optional commit message to use when creating a new commit.
	CommitMessage param.Opt[string] `json:"commit_message,omitzero"`
	paramObj
}

func (r BuildCompareParamsBase) MarshalJSON() (data []byte, err error) {
	type shadow BuildCompareParamsBase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildCompareParamsBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BuildCompareParamsBaseRevisionUnion struct {
	OfString       param.Opt[string]                     `json:",omitzero,inline"`
	OfFileInputMap map[string]shared.FileInputUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u BuildCompareParamsBaseRevisionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFileInputMap)
}
func (u *BuildCompareParamsBaseRevisionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BuildCompareParamsBaseRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFileInputMap) {
		return &u.OfFileInputMap
	}
	return nil
}

// Parameters for the head build
//
// The properties Branch, Revision are required.
type BuildCompareParamsHead struct {
	// Branch to use. When using a branch name as revision, this must match or be
	// omitted.
	Branch string `json:"branch,required"`
	// Specifies what to build: a branch name, a commit SHA, or file contents.
	Revision BuildCompareParamsHeadRevisionUnion `json:"revision,omitzero,required"`
	// Optional commit message to use when creating a new commit.
	CommitMessage param.Opt[string] `json:"commit_message,omitzero"`
	paramObj
}

func (r BuildCompareParamsHead) MarshalJSON() (data []byte, err error) {
	type shadow BuildCompareParamsHead
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildCompareParamsHead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BuildCompareParamsHeadRevisionUnion struct {
	OfString       param.Opt[string]                     `json:",omitzero,inline"`
	OfFileInputMap map[string]shared.FileInputUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u BuildCompareParamsHeadRevisionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFileInputMap)
}
func (u *BuildCompareParamsHeadRevisionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BuildCompareParamsHeadRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFileInputMap) {
		return &u.OfFileInputMap
	}
	return nil
}
