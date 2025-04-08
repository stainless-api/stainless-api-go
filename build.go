// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/resp"
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

// TODO
func (r *BuildService) New(ctx context.Context, body BuildNewParams, opts ...option.RequestOption) (res *Build, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/builds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// TODO
func (r *BuildService) Get(ctx context.Context, buildID string, opts ...option.RequestOption) (res *Build, err error) {
	opts = append(r.Options[:], opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s", buildID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Build struct {
	ID           string `json:"id,required"`
	ConfigCommit string `json:"config_commit,required"`
	// Any of "build".
	Object  BuildObject  `json:"object,required"`
	Targets BuildTargets `json:"targets,required"`
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
func (r Build) RawJSON() string { return r.JSON.raw }
func (r *Build) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildObject string

const (
	BuildObjectBuild BuildObject = "build"
)

type BuildTargets struct {
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
func (r BuildTargets) RawJSON() string { return r.JSON.raw }
func (r *BuildTargets) UnmarshalJSON(data []byte) error {
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
	Completed BuildTargetCommitCompletedCompletedUnion `json:"completed"`
	JSON      struct {
		Status    resp.Field
		Completed resp.Field
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
//	case BuildTargetCommitNotStarted:
//	case BuildTargetCommitQueued:
//	case BuildTargetCommitInProgress:
//	case BuildTargetCommitCompleted:
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitInProgress) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitInProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompleted struct {
	Completed BuildTargetCommitCompletedCompletedUnion `json:"completed,required"`
	Status    constant.Completed                       `json:"status,required"`
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
func (r BuildTargetCommitCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BuildTargetCommitCompletedCompletedUnion contains all possible properties and
// values from [BuildTargetCommitCompletedCompletedObject],
// [BuildTargetCommitCompletedCompletedMergeConflict],
// [BuildTargetCommitCompletedCompletedConclusion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuildTargetCommitCompletedCompletedUnion struct {
	// This field is from variant [BuildTargetCommitCompletedCompletedObject].
	Commit     BuildTargetCommitCompletedCompletedObjectCommit `json:"commit"`
	Conclusion string                                          `json:"conclusion"`
	// This field is from variant [BuildTargetCommitCompletedCompletedMergeConflict].
	MergeConflictPr BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPr `json:"merge_conflict_pr"`
	JSON            struct {
		Commit          resp.Field
		Conclusion      resp.Field
		MergeConflictPr resp.Field
		raw             string
	} `json:"-"`
}

func (u BuildTargetCommitCompletedCompletedUnion) AsBuildTargetCommitCompletedCompletedObject() (v BuildTargetCommitCompletedCompletedObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitCompletedCompletedUnion) AsMergeConflict() (v BuildTargetCommitCompletedCompletedMergeConflict) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuildTargetCommitCompletedCompletedUnion) AsBuildTargetCommitCompletedCompletedConclusion() (v BuildTargetCommitCompletedCompletedConclusion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuildTargetCommitCompletedCompletedUnion) RawJSON() string { return u.JSON.raw }

func (r *BuildTargetCommitCompletedCompletedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedObject struct {
	Commit BuildTargetCommitCompletedCompletedObjectCommit `json:"commit,required"`
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
func (r BuildTargetCommitCompletedCompletedObject) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedObjectCommit struct {
	Repo BuildTargetCommitCompletedCompletedObjectCommitRepo `json:"repo,required"`
	Sha  string                                              `json:"sha,required"`
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
func (r BuildTargetCommitCompletedCompletedObjectCommit) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedObjectCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedObjectCommitRepo struct {
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
func (r BuildTargetCommitCompletedCompletedObjectCommitRepo) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedObjectCommitRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedMergeConflict struct {
	// Any of "merge_conflict", "upstream_merge_conflict".
	Conclusion      string                                                          `json:"conclusion,required"`
	MergeConflictPr BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPr `json:"merge_conflict_pr,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Conclusion      resp.Field
		MergeConflictPr resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompletedMergeConflict) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedMergeConflict) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPr struct {
	Number float64                                                             `json:"number,required"`
	Repo   BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPrRepo `json:"repo,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Number      resp.Field
		Repo        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPr) RawJSON() string {
	return r.JSON.raw
}
func (r *BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPrRepo struct {
	Name  string `json:"name,required"`
	Owner string `json:"owner,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Owner       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPrRepo) RawJSON() string {
	return r.JSON.raw
}
func (r *BuildTargetCommitCompletedCompletedMergeConflictMergeConflictPrRepo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetCommitCompletedCompletedConclusion struct {
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
func (r BuildTargetCommitCompletedCompletedConclusion) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetCommitCompletedCompletedConclusion) UnmarshalJSON(data []byte) error {
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
		Status    resp.Field
		Completed resp.Field
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
//	case BuildTargetLintNotStarted:
//	case BuildTargetLintQueued:
//	case BuildTargetLintInProgress:
//	case BuildTargetLintCompleted:
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
func (r BuildTargetLintCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetLintCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetLintCompletedCompleted struct {
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
		Status    resp.Field
		Completed resp.Field
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
//	case BuildTargetTestNotStarted:
//	case BuildTargetTestQueued:
//	case BuildTargetTestInProgress:
//	case BuildTargetTestCompleted:
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
func (r BuildTargetTestCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildTargetTestCompletedCompleted struct {
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
func (r BuildTargetTestCompletedCompleted) RawJSON() string { return r.JSON.raw }
func (r *BuildTargetTestCompletedCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BuildNewParams struct {
	Branch       string `json:"branch,required"`
	ConfigCommit string `json:"config_commit,required"`
	Project      string `json:"project,required"`
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
