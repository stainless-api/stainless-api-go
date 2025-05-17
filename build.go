// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"encoding/json"
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
	"github.com/stainless-api/stainless-api-go/shared/constant"
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

// Create a new build
func (r *BuildService) New(ctx context.Context, body BuildNewParams, opts ...option.RequestOption) (res *BuildObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/builds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a build by ID
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

// List builds for a project
func (r *BuildService) List(ctx context.Context, query BuildListParams, opts ...option.RequestOption) (res *BuildListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/builds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Creates two builds whose outputs can be compared directly
func (r *BuildService) Compare(ctx context.Context, body BuildCompareParams, opts ...option.RequestOption) (res *BuildCompareResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/builds/compare"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BuildObject struct {
	ID           string `json:"id,required"`
	ConfigCommit string `json:"config_commit,required"`
	// Any of "build".
	Object  BuildObjectObject  `json:"object,required"`
	Org     string             `json:"org,required"`
	Project string             `json:"project,required"`
	Targets BuildObjectTargets `json:"targets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ConfigCommit respjson.Field
		Object       respjson.Field
		Org          respjson.Field
		Project      respjson.Field
		Targets      respjson.Field
		ExtraFields  map[string]respjson.Field
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
	Csharp     BuildTarget `json:"csharp"`
	Go         BuildTarget `json:"go"`
	Java       BuildTarget `json:"java"`
	Kotlin     BuildTarget `json:"kotlin"`
	Node       BuildTarget `json:"node"`
	Php        BuildTarget `json:"php"`
	Python     BuildTarget `json:"python"`
	Ruby       BuildTarget `json:"ruby"`
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
		Php         respjson.Field
		Python      respjson.Field
		Ruby        respjson.Field
		Terraform   respjson.Field
		Typescript  respjson.Field
		ExtraFields map[string]respjson.Field
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit      respjson.Field
		Lint        respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		Test        respjson.Field
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
	Completed BuildTargetCommitCompletedCompleted `json:"completed"`
	JSON      struct {
		Status    respjson.Field
		Completed respjson.Field
		raw       string
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
//	case stainlessv0.BuildTargetCommitNotStarted:
//	case stainlessv0.BuildTargetCommitQueued:
//	case stainlessv0.BuildTargetCommitInProgress:
//	case stainlessv0.BuildTargetCommitCompleted:
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
	Completed BuildTargetCommitCompletedCompleted `json:"completed,required"`
	Status    constant.Completed                  `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompleted struct {
	Commit BuildTargetCommitCompletedCompletedCommit `json:"commit,required"`
	// Any of "success", "failure", "skipped", "cancelled", "action_required",
	// "neutral", "timed_out", "error", "warning", "note", "merge_conflict",
	// "upstream_merge_conflict", "fatal", "payment_required", "noop", "version_bump".
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

type BuildTargetCommitCompletedCompletedCommit struct {
	Repo BuildTargetCommitCompletedCompletedCommitRepo `json:"repo,required"`
	Sha  string                                        `json:"sha,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Repo        respjson.Field
		Sha         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompletedCommit) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedCommitRepo struct {
	Branch string `json:"branch,required"`
	Name   string `json:"name,required"`
	Owner  string `json:"owner,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Branch      respjson.Field
		Name        respjson.Field
		Owner       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompletedCommitRepo) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedCommitRepo) UnmarshalJSON(data []byte) error {
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

// BuildTargetLintUnion contains all possible properties and values from
// [BuildTargetLintNotStarted], [BuildTargetLintQueued],
// [BuildTargetLintInProgress], [BuildTargetLintCompleted].
//
// Use the [BuildTargetLintUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetLintUnion struct {
	// Any of "not_started", "queued", "in_progress", "completed".
	Status string `json:"status"`
	// This field is from variant [BuildTargetLintCompleted].
	Completed BuildTargetLintCompletedCompleted `json:"completed"`
	JSON      struct {
		Status    respjson.Field
		Completed respjson.Field
		raw       string
	} `json:"-"`
}

// anyBuildTargetLint is implemented by each variant of [BuildTargetLintUnion] to
// add type safety for the return type of [BuildTargetLintUnion.AsAny]
type anyBuildTargetLint interface {
	implBuildTargetLintUnion()
}

func (BuildTargetLintNotStarted) implBuildTargetLintUnion() {}
func (BuildTargetLintQueued) implBuildTargetLintUnion()     {}
func (BuildTargetLintInProgress) implBuildTargetLintUnion() {}
func (BuildTargetLintCompleted) implBuildTargetLintUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BuildTargetLintUnion.AsAny().(type) {
//	case stainlessv0.BuildTargetLintNotStarted:
//	case stainlessv0.BuildTargetLintQueued:
//	case stainlessv0.BuildTargetLintInProgress:
//	case stainlessv0.BuildTargetLintCompleted:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BuildTargetLintUnion) AsAny() anyBuildTargetLint {
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

func (u BuildTargetLintUnion) AsNotStarted() (v BuildTargetLintNotStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetLintUnion) AsQueued() (v BuildTargetLintQueued) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetLintUnion) AsInProgress() (v BuildTargetLintInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetLintUnion) AsCompleted() (v BuildTargetLintCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetLintUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetLintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintNotStarted struct {
	Status constant.NotStarted `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintNotStarted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintNotStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintQueued struct {
	Status constant.Queued `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintQueued) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintQueued) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintInProgress struct {
	Status constant.InProgress `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintInProgress) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintCompleted struct {
	Completed BuildTargetLintCompletedCompleted `json:"completed,required"`
	Status    constant.Completed                `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintCompletedCompleted struct {
	// Any of "success", "failure", "skipped", "cancelled", "action_required",
	// "neutral", "timed_out", "error", "warning", "note", "merge_conflict",
	// "upstream_merge_conflict", "fatal", "payment_required", "noop", "version_bump".
	Conclusion string `json:"conclusion,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conclusion  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetLintCompletedCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintCompletedCompleted) UnmarshalJSON(data []byte) error {
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
// [BuildTargetTestNotStarted], [BuildTargetTestQueued],
// [BuildTargetTestInProgress], [BuildTargetTestCompleted].
//
// Use the [BuildTargetTestUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetTestUnion struct {
	// Any of "not_started", "queued", "in_progress", "completed".
	Status string `json:"status"`
	// This field is from variant [BuildTargetTestCompleted].
	Completed BuildTargetTestCompletedCompleted `json:"completed"`
	JSON      struct {
		Status    respjson.Field
		Completed respjson.Field
		raw       string
	} `json:"-"`
}

// anyBuildTargetTest is implemented by each variant of [BuildTargetTestUnion] to
// add type safety for the return type of [BuildTargetTestUnion.AsAny]
type anyBuildTargetTest interface {
	implBuildTargetTestUnion()
}

func (BuildTargetTestNotStarted) implBuildTargetTestUnion() {}
func (BuildTargetTestQueued) implBuildTargetTestUnion()     {}
func (BuildTargetTestInProgress) implBuildTargetTestUnion() {}
func (BuildTargetTestCompleted) implBuildTargetTestUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BuildTargetTestUnion.AsAny().(type) {
//	case stainlessv0.BuildTargetTestNotStarted:
//	case stainlessv0.BuildTargetTestQueued:
//	case stainlessv0.BuildTargetTestInProgress:
//	case stainlessv0.BuildTargetTestCompleted:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BuildTargetTestUnion) AsAny() anyBuildTargetTest {
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

func (u BuildTargetTestUnion) AsNotStarted() (v BuildTargetTestNotStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetTestUnion) AsQueued() (v BuildTargetTestQueued) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetTestUnion) AsInProgress() (v BuildTargetTestInProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetTestUnion) AsCompleted() (v BuildTargetTestCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetTestUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetTestUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestNotStarted struct {
	Status constant.NotStarted `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestNotStarted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestNotStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestQueued struct {
	Status constant.Queued `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestQueued) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestQueued) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestInProgress struct {
	Status constant.InProgress `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestInProgress) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestCompleted struct {
	Completed BuildTargetTestCompletedCompleted `json:"completed,required"`
	Status    constant.Completed                `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestCompletedCompleted struct {
	// Any of "success", "failure", "skipped", "cancelled", "action_required",
	// "neutral", "timed_out", "error", "warning", "note", "merge_conflict",
	// "upstream_merge_conflict", "fatal", "payment_required", "noop", "version_bump".
	Conclusion string `json:"conclusion,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conclusion  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetTestCompletedCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestCompletedCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildListResponse struct {
	Data       []BuildObject `json:"data,required"`
	HasMore    bool          `json:"has_more,required"`
	NextCursor string        `json:"next_cursor"`
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
func (r BuildListResponse) RawJSON() string { return r.JSON.raw }
func (r *BuildListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildCompareResponse struct {
	Base BuildObject `json:"base,required"`
	Head BuildObject `json:"head,required"`
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
	// "terraform", "cli", "php", "csharp".
	Targets []string `json:"targets,omitzero"`
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
	OfString                  param.Opt[string]                        `json:",omitzero,inline"`
	OfBuildNewsRevisionMapMap map[string]BuildNewParamsRevisionMapItem `json:",omitzero,inline"`
	paramUnion
}

func (u BuildNewParamsRevisionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BuildNewParamsRevisionUnion](u.OfString, u.OfBuildNewsRevisionMapMap)
}
func (u *BuildNewParamsRevisionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BuildNewParamsRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBuildNewsRevisionMapMap) {
		return &u.OfBuildNewsRevisionMapMap
	}
	return nil
}

// The property Content is required.
type BuildNewParamsRevisionMapItem struct {
	// The file content
	Content string `json:"content,required"`
	paramObj
}

func (r BuildNewParamsRevisionMapItem) MarshalJSON() (data []byte, err error) {
	type shadow BuildNewParamsRevisionMapItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildNewParamsRevisionMapItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildListParams struct {
	// Project name
	Project string `query:"project,required" json:"-"`
	// Branch name
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Pagination cursor from a previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of builds to return, defaults to 10 (maximum: 100)
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
	// Parameters for the base build
	Base BuildCompareParamsBase `json:"base,omitzero,required"`
	// Parameters for the head build
	Head BuildCompareParamsHead `json:"head,omitzero,required"`
	// Project name
	Project string `json:"project,required"`
	// Optional list of SDK targets to build. If not specified, all configured targets
	// will be built.
	//
	// Any of "node", "typescript", "python", "go", "java", "kotlin", "ruby",
	// "terraform", "cli", "php", "csharp".
	Targets []string `json:"targets,omitzero"`
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
// The property Revision is required.
type BuildCompareParamsBase struct {
	// Specifies what to build: a branch name, a commit SHA, or file contents
	Revision BuildCompareParamsBaseRevisionUnion `json:"revision,omitzero,required"`
	// Optional branch to use. If not specified, defaults to "main". When using a
	// branch name as revision, this must match or be omitted.
	Branch param.Opt[string] `json:"branch,omitzero"`
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
	OfString                          param.Opt[string]                                `json:",omitzero,inline"`
	OfBuildComparesBaseRevisionMapMap map[string]BuildCompareParamsBaseRevisionMapItem `json:",omitzero,inline"`
	paramUnion
}

func (u BuildCompareParamsBaseRevisionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BuildCompareParamsBaseRevisionUnion](u.OfString, u.OfBuildComparesBaseRevisionMapMap)
}
func (u *BuildCompareParamsBaseRevisionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BuildCompareParamsBaseRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBuildComparesBaseRevisionMapMap) {
		return &u.OfBuildComparesBaseRevisionMapMap
	}
	return nil
}

// The property Content is required.
type BuildCompareParamsBaseRevisionMapItem struct {
	// The file content
	Content string `json:"content,required"`
	paramObj
}

func (r BuildCompareParamsBaseRevisionMapItem) MarshalJSON() (data []byte, err error) {
	type shadow BuildCompareParamsBaseRevisionMapItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildCompareParamsBaseRevisionMapItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for the head build
//
// The property Revision is required.
type BuildCompareParamsHead struct {
	// Specifies what to build: a branch name, a commit SHA, or file contents
	Revision BuildCompareParamsHeadRevisionUnion `json:"revision,omitzero,required"`
	// Optional branch to use. If not specified, defaults to "main". When using a
	// branch name as revision, this must match or be omitted.
	Branch param.Opt[string] `json:"branch,omitzero"`
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
	OfString                          param.Opt[string]                                `json:",omitzero,inline"`
	OfBuildComparesHeadRevisionMapMap map[string]BuildCompareParamsHeadRevisionMapItem `json:",omitzero,inline"`
	paramUnion
}

func (u BuildCompareParamsHeadRevisionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BuildCompareParamsHeadRevisionUnion](u.OfString, u.OfBuildComparesHeadRevisionMapMap)
}
func (u *BuildCompareParamsHeadRevisionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BuildCompareParamsHeadRevisionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBuildComparesHeadRevisionMapMap) {
		return &u.OfBuildComparesHeadRevisionMapMap
	}
	return nil
}

// The property Content is required.
type BuildCompareParamsHeadRevisionMapItem struct {
	// The file content
	Content string `json:"content,required"`
	paramObj
}

func (r BuildCompareParamsHeadRevisionMapItem) MarshalJSON() (data []byte, err error) {
	type shadow BuildCompareParamsHeadRevisionMapItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BuildCompareParamsHeadRevisionMapItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
