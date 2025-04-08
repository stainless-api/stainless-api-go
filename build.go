// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/param"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/tidwall/gjson"
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
func NewBuildService(opts ...option.RequestOption) (r *BuildService) {
	r = &BuildService{}
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
	ID           string       `json:"id,required"`
	ConfigCommit string       `json:"config_commit,required"`
	Object       BuildObject  `json:"object,required"`
	Targets      BuildTargets `json:"targets,required"`
	JSON         buildJSON    `json:"-"`
}

// buildJSON contains the JSON metadata for the struct [Build]
type buildJSON struct {
	ID           apijson.Field
	ConfigCommit apijson.Field
	Object       apijson.Field
	Targets      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Build) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildJSON) RawJSON() string {
	return r.raw
}

type BuildObject string

const (
	BuildObjectBuild BuildObject = "build"
)

func (r BuildObject) IsKnown() bool {
	switch r {
	case BuildObjectBuild:
		return true
	}
	return false
}

type BuildTargets struct {
	Cli        BuildTarget      `json:"cli"`
	Go         BuildTarget      `json:"go"`
	Java       BuildTarget      `json:"java"`
	Kotlin     BuildTarget      `json:"kotlin"`
	Node       BuildTarget      `json:"node"`
	Python     BuildTarget      `json:"python"`
	Ruby       BuildTarget      `json:"ruby"`
	Terraform  BuildTarget      `json:"terraform"`
	Typescript BuildTarget      `json:"typescript"`
	JSON       buildTargetsJSON `json:"-"`
}

// buildTargetsJSON contains the JSON metadata for the struct [BuildTargets]
type buildTargetsJSON struct {
	Cli         apijson.Field
	Go          apijson.Field
	Java        apijson.Field
	Kotlin      apijson.Field
	Node        apijson.Field
	Python      apijson.Field
	Ruby        apijson.Field
	Terraform   apijson.Field
	Typescript  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetsJSON) RawJSON() string {
	return r.raw
}

type BuildTarget struct {
	Commit BuildTargetCommit `json:"commit,required"`
	Lint   BuildTargetLint   `json:"lint,required"`
	Object BuildTargetObject `json:"object,required"`
	Status BuildTargetStatus `json:"status,required"`
	Test   BuildTargetTest   `json:"test,required"`
	JSON   buildTargetJSON   `json:"-"`
}

// buildTargetJSON contains the JSON metadata for the struct [BuildTarget]
type buildTargetJSON struct {
	Commit      apijson.Field
	Lint        apijson.Field
	Object      apijson.Field
	Status      apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetJSON) RawJSON() string {
	return r.raw
}

type BuildTargetCommit struct {
	Status BuildTargetCommitStatus `json:"status,required"`
	// This field can have the runtime type of [BuildTargetCommitCompletedCompleted].
	Completed interface{}           `json:"completed"`
	JSON      buildTargetCommitJSON `json:"-"`
	union     BuildTargetCommitUnion
}

// buildTargetCommitJSON contains the JSON metadata for the struct
// [BuildTargetCommit]
type buildTargetCommitJSON struct {
	Status      apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r buildTargetCommitJSON) RawJSON() string {
	return r.raw
}

func (r *BuildTargetCommit) UnmarshalJSON(data []byte) (err error) {
	*r = BuildTargetCommit{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BuildTargetCommitUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BuildTargetCommitNotStarted],
// [BuildTargetCommitQueued], [BuildTargetCommitInProgress],
// [BuildTargetCommitCompleted].
func (r BuildTargetCommit) AsUnion() BuildTargetCommitUnion {
	return r.union
}

// Union satisfied by [BuildTargetCommitNotStarted], [BuildTargetCommitQueued],
// [BuildTargetCommitInProgress] or [BuildTargetCommitCompleted].
type BuildTargetCommitUnion interface {
	implementsBuildTargetCommit()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetCommitUnion)(nil)).Elem(),
		"status",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetCommitNotStarted{}),
			DiscriminatorValue: "not_started",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetCommitQueued{}),
			DiscriminatorValue: "queued",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetCommitInProgress{}),
			DiscriminatorValue: "in_progress",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetCommitCompleted{}),
			DiscriminatorValue: "completed",
		},
	)
}

type BuildTargetCommitNotStarted struct {
	Status BuildTargetCommitNotStartedStatus `json:"status,required"`
	JSON   buildTargetCommitNotStartedJSON   `json:"-"`
}

// buildTargetCommitNotStartedJSON contains the JSON metadata for the struct
// [BuildTargetCommitNotStarted]
type buildTargetCommitNotStartedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitNotStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitNotStartedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitNotStarted) implementsBuildTargetCommit() {}

type BuildTargetCommitNotStartedStatus string

const (
	BuildTargetCommitNotStartedStatusNotStarted BuildTargetCommitNotStartedStatus = "not_started"
)

func (r BuildTargetCommitNotStartedStatus) IsKnown() bool {
	switch r {
	case BuildTargetCommitNotStartedStatusNotStarted:
		return true
	}
	return false
}

type BuildTargetCommitQueued struct {
	Status BuildTargetCommitQueuedStatus `json:"status,required"`
	JSON   buildTargetCommitQueuedJSON   `json:"-"`
}

// buildTargetCommitQueuedJSON contains the JSON metadata for the struct
// [BuildTargetCommitQueued]
type buildTargetCommitQueuedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitQueued) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitQueuedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitQueued) implementsBuildTargetCommit() {}

type BuildTargetCommitQueuedStatus string

const (
	BuildTargetCommitQueuedStatusQueued BuildTargetCommitQueuedStatus = "queued"
)

func (r BuildTargetCommitQueuedStatus) IsKnown() bool {
	switch r {
	case BuildTargetCommitQueuedStatusQueued:
		return true
	}
	return false
}

type BuildTargetCommitInProgress struct {
	Status BuildTargetCommitInProgressStatus `json:"status,required"`
	JSON   buildTargetCommitInProgressJSON   `json:"-"`
}

// buildTargetCommitInProgressJSON contains the JSON metadata for the struct
// [BuildTargetCommitInProgress]
type buildTargetCommitInProgressJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitInProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitInProgressJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitInProgress) implementsBuildTargetCommit() {}

type BuildTargetCommitInProgressStatus string

const (
	BuildTargetCommitInProgressStatusInProgress BuildTargetCommitInProgressStatus = "in_progress"
)

func (r BuildTargetCommitInProgressStatus) IsKnown() bool {
	switch r {
	case BuildTargetCommitInProgressStatusInProgress:
		return true
	}
	return false
}

type BuildTargetCommitCompleted struct {
	Completed BuildTargetCommitCompletedCompleted `json:"completed,required"`
	Status    BuildTargetCommitCompletedStatus    `json:"status,required"`
	JSON      buildTargetCommitCompletedJSON      `json:"-"`
}

// buildTargetCommitCompletedJSON contains the JSON metadata for the struct
// [BuildTargetCommitCompleted]
type buildTargetCommitCompletedJSON struct {
	Completed   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitCompletedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitCompleted) implementsBuildTargetCommit() {}

type BuildTargetCommitCompletedCompleted struct {
	Conclusion BuildTargetCommitCompletedCompletedConclusion `json:"conclusion,required"`
	// This field can have the runtime type of
	// [BuildTargetCommitCompletedCompletedObjectCommit].
	Commit interface{} `json:"commit"`
	// This field can have the runtime type of
	// [BuildTargetCommitCompletedCompletedObjectMergeConflictPr].
	MergeConflictPr interface{}                             `json:"merge_conflict_pr"`
	JSON            buildTargetCommitCompletedCompletedJSON `json:"-"`
	union           BuildTargetCommitCompletedCompletedUnion
}

// buildTargetCommitCompletedCompletedJSON contains the JSON metadata for the
// struct [BuildTargetCommitCompletedCompleted]
type buildTargetCommitCompletedCompletedJSON struct {
	Conclusion      apijson.Field
	Commit          apijson.Field
	MergeConflictPr apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r buildTargetCommitCompletedCompletedJSON) RawJSON() string {
	return r.raw
}

func (r *BuildTargetCommitCompletedCompleted) UnmarshalJSON(data []byte) (err error) {
	*r = BuildTargetCommitCompletedCompleted{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BuildTargetCommitCompletedCompletedUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BuildTargetCommitCompletedCompletedObject],
// [BuildTargetCommitCompletedCompletedObject],
// [BuildTargetCommitCompletedCompletedConclusion].
func (r BuildTargetCommitCompletedCompleted) AsUnion() BuildTargetCommitCompletedCompletedUnion {
	return r.union
}

// Union satisfied by [BuildTargetCommitCompletedCompletedObject],
// [BuildTargetCommitCompletedCompletedObject] or
// [BuildTargetCommitCompletedCompletedConclusion].
type BuildTargetCommitCompletedCompletedUnion interface {
	implementsBuildTargetCommitCompletedCompleted()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetCommitCompletedCompletedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitCompletedCompletedObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitCompletedCompletedObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitCompletedCompletedConclusion{}),
		},
	)
}

type BuildTargetCommitCompletedCompletedObject struct {
	Commit     BuildTargetCommitCompletedCompletedObjectCommit     `json:"commit,required"`
	Conclusion BuildTargetCommitCompletedCompletedObjectConclusion `json:"conclusion,required"`
	JSON       buildTargetCommitCompletedCompletedObjectJSON       `json:"-"`
}

// buildTargetCommitCompletedCompletedObjectJSON contains the JSON metadata for the
// struct [BuildTargetCommitCompletedCompletedObject]
type buildTargetCommitCompletedCompletedObjectJSON struct {
	Commit      apijson.Field
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitCompletedCompletedObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitCompletedCompletedObjectJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitCompletedCompletedObject) implementsBuildTargetCommitCompletedCompleted() {}

type BuildTargetCommitCompletedCompletedObjectCommit struct {
	Repo BuildTargetCommitCompletedCompletedObjectCommitRepo `json:"repo,required"`
	Sha  string                                              `json:"sha,required"`
	JSON buildTargetCommitCompletedCompletedObjectCommitJSON `json:"-"`
}

// buildTargetCommitCompletedCompletedObjectCommitJSON contains the JSON metadata
// for the struct [BuildTargetCommitCompletedCompletedObjectCommit]
type buildTargetCommitCompletedCompletedObjectCommitJSON struct {
	Repo        apijson.Field
	Sha         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitCompletedCompletedObjectCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitCompletedCompletedObjectCommitJSON) RawJSON() string {
	return r.raw
}

type BuildTargetCommitCompletedCompletedObjectCommitRepo struct {
	Branch string                                                  `json:"branch,required"`
	Name   string                                                  `json:"name,required"`
	Owner  string                                                  `json:"owner,required"`
	JSON   buildTargetCommitCompletedCompletedObjectCommitRepoJSON `json:"-"`
}

// buildTargetCommitCompletedCompletedObjectCommitRepoJSON contains the JSON
// metadata for the struct [BuildTargetCommitCompletedCompletedObjectCommitRepo]
type buildTargetCommitCompletedCompletedObjectCommitRepoJSON struct {
	Branch      apijson.Field
	Name        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitCompletedCompletedObjectCommitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitCompletedCompletedObjectCommitRepoJSON) RawJSON() string {
	return r.raw
}

type BuildTargetCommitCompletedCompletedObjectConclusion string

const (
	BuildTargetCommitCompletedCompletedObjectConclusionError   BuildTargetCommitCompletedCompletedObjectConclusion = "error"
	BuildTargetCommitCompletedCompletedObjectConclusionWarning BuildTargetCommitCompletedCompletedObjectConclusion = "warning"
	BuildTargetCommitCompletedCompletedObjectConclusionNote    BuildTargetCommitCompletedCompletedObjectConclusion = "note"
	BuildTargetCommitCompletedCompletedObjectConclusionSuccess BuildTargetCommitCompletedCompletedObjectConclusion = "success"
)

func (r BuildTargetCommitCompletedCompletedObjectConclusion) IsKnown() bool {
	switch r {
	case BuildTargetCommitCompletedCompletedObjectConclusionError, BuildTargetCommitCompletedCompletedObjectConclusionWarning, BuildTargetCommitCompletedCompletedObjectConclusionNote, BuildTargetCommitCompletedCompletedObjectConclusionSuccess:
		return true
	}
	return false
}

type BuildTargetCommitCompletedCompletedConclusion struct {
	Conclusion BuildTargetCommitCompletedCompletedConclusionConclusion `json:"conclusion,required"`
	JSON       buildTargetCommitCompletedCompletedConclusionJSON       `json:"-"`
}

// buildTargetCommitCompletedCompletedConclusionJSON contains the JSON metadata for
// the struct [BuildTargetCommitCompletedCompletedConclusion]
type buildTargetCommitCompletedCompletedConclusionJSON struct {
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitCompletedCompletedConclusion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitCompletedCompletedConclusionJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitCompletedCompletedConclusion) implementsBuildTargetCommitCompletedCompleted() {
}

type BuildTargetCommitCompletedCompletedConclusionConclusion string

const (
	BuildTargetCommitCompletedCompletedConclusionConclusionFatal           BuildTargetCommitCompletedCompletedConclusionConclusion = "fatal"
	BuildTargetCommitCompletedCompletedConclusionConclusionPaymentRequired BuildTargetCommitCompletedCompletedConclusionConclusion = "payment_required"
	BuildTargetCommitCompletedCompletedConclusionConclusionCancelled       BuildTargetCommitCompletedCompletedConclusionConclusion = "cancelled"
	BuildTargetCommitCompletedCompletedConclusionConclusionTimedOut        BuildTargetCommitCompletedCompletedConclusionConclusion = "timed_out"
	BuildTargetCommitCompletedCompletedConclusionConclusionNoop            BuildTargetCommitCompletedCompletedConclusionConclusion = "noop"
	BuildTargetCommitCompletedCompletedConclusionConclusionVersionBump     BuildTargetCommitCompletedCompletedConclusionConclusion = "version_bump"
)

func (r BuildTargetCommitCompletedCompletedConclusionConclusion) IsKnown() bool {
	switch r {
	case BuildTargetCommitCompletedCompletedConclusionConclusionFatal, BuildTargetCommitCompletedCompletedConclusionConclusionPaymentRequired, BuildTargetCommitCompletedCompletedConclusionConclusionCancelled, BuildTargetCommitCompletedCompletedConclusionConclusionTimedOut, BuildTargetCommitCompletedCompletedConclusionConclusionNoop, BuildTargetCommitCompletedCompletedConclusionConclusionVersionBump:
		return true
	}
	return false
}

type BuildTargetCommitCompletedStatus string

const (
	BuildTargetCommitCompletedStatusCompleted BuildTargetCommitCompletedStatus = "completed"
)

func (r BuildTargetCommitCompletedStatus) IsKnown() bool {
	switch r {
	case BuildTargetCommitCompletedStatusCompleted:
		return true
	}
	return false
}

type BuildTargetCommitStatus string

const (
	BuildTargetCommitStatusNotStarted BuildTargetCommitStatus = "not_started"
	BuildTargetCommitStatusQueued     BuildTargetCommitStatus = "queued"
	BuildTargetCommitStatusInProgress BuildTargetCommitStatus = "in_progress"
	BuildTargetCommitStatusCompleted  BuildTargetCommitStatus = "completed"
)

func (r BuildTargetCommitStatus) IsKnown() bool {
	switch r {
	case BuildTargetCommitStatusNotStarted, BuildTargetCommitStatusQueued, BuildTargetCommitStatusInProgress, BuildTargetCommitStatusCompleted:
		return true
	}
	return false
}

type BuildTargetLint struct {
	Status BuildTargetLintStatus `json:"status,required"`
	// This field can have the runtime type of [BuildTargetLintCompletedCompleted].
	Completed interface{}         `json:"completed"`
	JSON      buildTargetLintJSON `json:"-"`
	union     BuildTargetLintUnion
}

// buildTargetLintJSON contains the JSON metadata for the struct [BuildTargetLint]
type buildTargetLintJSON struct {
	Status      apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r buildTargetLintJSON) RawJSON() string {
	return r.raw
}

func (r *BuildTargetLint) UnmarshalJSON(data []byte) (err error) {
	*r = BuildTargetLint{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BuildTargetLintUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BuildTargetLintNotStarted],
// [BuildTargetLintQueued], [BuildTargetLintInProgress],
// [BuildTargetLintCompleted].
func (r BuildTargetLint) AsUnion() BuildTargetLintUnion {
	return r.union
}

// Union satisfied by [BuildTargetLintNotStarted], [BuildTargetLintQueued],
// [BuildTargetLintInProgress] or [BuildTargetLintCompleted].
type BuildTargetLintUnion interface {
	implementsBuildTargetLint()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetLintUnion)(nil)).Elem(),
		"status",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetLintNotStarted{}),
			DiscriminatorValue: "not_started",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetLintQueued{}),
			DiscriminatorValue: "queued",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetLintInProgress{}),
			DiscriminatorValue: "in_progress",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetLintCompleted{}),
			DiscriminatorValue: "completed",
		},
	)
}

type BuildTargetLintNotStarted struct {
	Status BuildTargetLintNotStartedStatus `json:"status,required"`
	JSON   buildTargetLintNotStartedJSON   `json:"-"`
}

// buildTargetLintNotStartedJSON contains the JSON metadata for the struct
// [BuildTargetLintNotStarted]
type buildTargetLintNotStartedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintNotStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintNotStartedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetLintNotStarted) implementsBuildTargetLint() {}

type BuildTargetLintNotStartedStatus string

const (
	BuildTargetLintNotStartedStatusNotStarted BuildTargetLintNotStartedStatus = "not_started"
)

func (r BuildTargetLintNotStartedStatus) IsKnown() bool {
	switch r {
	case BuildTargetLintNotStartedStatusNotStarted:
		return true
	}
	return false
}

type BuildTargetLintQueued struct {
	Status BuildTargetLintQueuedStatus `json:"status,required"`
	JSON   buildTargetLintQueuedJSON   `json:"-"`
}

// buildTargetLintQueuedJSON contains the JSON metadata for the struct
// [BuildTargetLintQueued]
type buildTargetLintQueuedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintQueued) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintQueuedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetLintQueued) implementsBuildTargetLint() {}

type BuildTargetLintQueuedStatus string

const (
	BuildTargetLintQueuedStatusQueued BuildTargetLintQueuedStatus = "queued"
)

func (r BuildTargetLintQueuedStatus) IsKnown() bool {
	switch r {
	case BuildTargetLintQueuedStatusQueued:
		return true
	}
	return false
}

type BuildTargetLintInProgress struct {
	Status BuildTargetLintInProgressStatus `json:"status,required"`
	JSON   buildTargetLintInProgressJSON   `json:"-"`
}

// buildTargetLintInProgressJSON contains the JSON metadata for the struct
// [BuildTargetLintInProgress]
type buildTargetLintInProgressJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintInProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintInProgressJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetLintInProgress) implementsBuildTargetLint() {}

type BuildTargetLintInProgressStatus string

const (
	BuildTargetLintInProgressStatusInProgress BuildTargetLintInProgressStatus = "in_progress"
)

func (r BuildTargetLintInProgressStatus) IsKnown() bool {
	switch r {
	case BuildTargetLintInProgressStatusInProgress:
		return true
	}
	return false
}

type BuildTargetLintCompleted struct {
	Completed BuildTargetLintCompletedCompleted `json:"completed,required"`
	Status    BuildTargetLintCompletedStatus    `json:"status,required"`
	JSON      buildTargetLintCompletedJSON      `json:"-"`
}

// buildTargetLintCompletedJSON contains the JSON metadata for the struct
// [BuildTargetLintCompleted]
type buildTargetLintCompletedJSON struct {
	Completed   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintCompletedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetLintCompleted) implementsBuildTargetLint() {}

type BuildTargetLintCompletedCompleted struct {
	Conclusion BuildTargetLintCompletedCompletedConclusion `json:"conclusion,required"`
	JSON       buildTargetLintCompletedCompletedJSON       `json:"-"`
}

// buildTargetLintCompletedCompletedJSON contains the JSON metadata for the struct
// [BuildTargetLintCompletedCompleted]
type buildTargetLintCompletedCompletedJSON struct {
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintCompletedCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintCompletedCompletedJSON) RawJSON() string {
	return r.raw
}

type BuildTargetLintCompletedCompletedConclusion string

const (
	BuildTargetLintCompletedCompletedConclusionSuccess        BuildTargetLintCompletedCompletedConclusion = "success"
	BuildTargetLintCompletedCompletedConclusionFailure        BuildTargetLintCompletedCompletedConclusion = "failure"
	BuildTargetLintCompletedCompletedConclusionSkipped        BuildTargetLintCompletedCompletedConclusion = "skipped"
	BuildTargetLintCompletedCompletedConclusionCancelled      BuildTargetLintCompletedCompletedConclusion = "cancelled"
	BuildTargetLintCompletedCompletedConclusionActionRequired BuildTargetLintCompletedCompletedConclusion = "action_required"
	BuildTargetLintCompletedCompletedConclusionNeutral        BuildTargetLintCompletedCompletedConclusion = "neutral"
	BuildTargetLintCompletedCompletedConclusionTimedOut       BuildTargetLintCompletedCompletedConclusion = "timed_out"
)

func (r BuildTargetLintCompletedCompletedConclusion) IsKnown() bool {
	switch r {
	case BuildTargetLintCompletedCompletedConclusionSuccess, BuildTargetLintCompletedCompletedConclusionFailure, BuildTargetLintCompletedCompletedConclusionSkipped, BuildTargetLintCompletedCompletedConclusionCancelled, BuildTargetLintCompletedCompletedConclusionActionRequired, BuildTargetLintCompletedCompletedConclusionNeutral, BuildTargetLintCompletedCompletedConclusionTimedOut:
		return true
	}
	return false
}

type BuildTargetLintCompletedStatus string

const (
	BuildTargetLintCompletedStatusCompleted BuildTargetLintCompletedStatus = "completed"
)

func (r BuildTargetLintCompletedStatus) IsKnown() bool {
	switch r {
	case BuildTargetLintCompletedStatusCompleted:
		return true
	}
	return false
}

type BuildTargetLintStatus string

const (
	BuildTargetLintStatusNotStarted BuildTargetLintStatus = "not_started"
	BuildTargetLintStatusQueued     BuildTargetLintStatus = "queued"
	BuildTargetLintStatusInProgress BuildTargetLintStatus = "in_progress"
	BuildTargetLintStatusCompleted  BuildTargetLintStatus = "completed"
)

func (r BuildTargetLintStatus) IsKnown() bool {
	switch r {
	case BuildTargetLintStatusNotStarted, BuildTargetLintStatusQueued, BuildTargetLintStatusInProgress, BuildTargetLintStatusCompleted:
		return true
	}
	return false
}

type BuildTargetObject string

const (
	BuildTargetObjectBuildTarget BuildTargetObject = "build_target"
)

func (r BuildTargetObject) IsKnown() bool {
	switch r {
	case BuildTargetObjectBuildTarget:
		return true
	}
	return false
}

type BuildTargetStatus string

const (
	BuildTargetStatusNotStarted BuildTargetStatus = "not_started"
	BuildTargetStatusCodegen    BuildTargetStatus = "codegen"
	BuildTargetStatusPostgen    BuildTargetStatus = "postgen"
	BuildTargetStatusCompleted  BuildTargetStatus = "completed"
)

func (r BuildTargetStatus) IsKnown() bool {
	switch r {
	case BuildTargetStatusNotStarted, BuildTargetStatusCodegen, BuildTargetStatusPostgen, BuildTargetStatusCompleted:
		return true
	}
	return false
}

type BuildTargetTest struct {
	Status BuildTargetTestStatus `json:"status,required"`
	// This field can have the runtime type of [BuildTargetTestCompletedCompleted].
	Completed interface{}         `json:"completed"`
	JSON      buildTargetTestJSON `json:"-"`
	union     BuildTargetTestUnion
}

// buildTargetTestJSON contains the JSON metadata for the struct [BuildTargetTest]
type buildTargetTestJSON struct {
	Status      apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r buildTargetTestJSON) RawJSON() string {
	return r.raw
}

func (r *BuildTargetTest) UnmarshalJSON(data []byte) (err error) {
	*r = BuildTargetTest{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BuildTargetTestUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BuildTargetTestNotStarted],
// [BuildTargetTestQueued], [BuildTargetTestInProgress],
// [BuildTargetTestCompleted].
func (r BuildTargetTest) AsUnion() BuildTargetTestUnion {
	return r.union
}

// Union satisfied by [BuildTargetTestNotStarted], [BuildTargetTestQueued],
// [BuildTargetTestInProgress] or [BuildTargetTestCompleted].
type BuildTargetTestUnion interface {
	implementsBuildTargetTest()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetTestUnion)(nil)).Elem(),
		"status",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetTestNotStarted{}),
			DiscriminatorValue: "not_started",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetTestQueued{}),
			DiscriminatorValue: "queued",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetTestInProgress{}),
			DiscriminatorValue: "in_progress",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BuildTargetTestCompleted{}),
			DiscriminatorValue: "completed",
		},
	)
}

type BuildTargetTestNotStarted struct {
	Status BuildTargetTestNotStartedStatus `json:"status,required"`
	JSON   buildTargetTestNotStartedJSON   `json:"-"`
}

// buildTargetTestNotStartedJSON contains the JSON metadata for the struct
// [BuildTargetTestNotStarted]
type buildTargetTestNotStartedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestNotStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestNotStartedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetTestNotStarted) implementsBuildTargetTest() {}

type BuildTargetTestNotStartedStatus string

const (
	BuildTargetTestNotStartedStatusNotStarted BuildTargetTestNotStartedStatus = "not_started"
)

func (r BuildTargetTestNotStartedStatus) IsKnown() bool {
	switch r {
	case BuildTargetTestNotStartedStatusNotStarted:
		return true
	}
	return false
}

type BuildTargetTestQueued struct {
	Status BuildTargetTestQueuedStatus `json:"status,required"`
	JSON   buildTargetTestQueuedJSON   `json:"-"`
}

// buildTargetTestQueuedJSON contains the JSON metadata for the struct
// [BuildTargetTestQueued]
type buildTargetTestQueuedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestQueued) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestQueuedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetTestQueued) implementsBuildTargetTest() {}

type BuildTargetTestQueuedStatus string

const (
	BuildTargetTestQueuedStatusQueued BuildTargetTestQueuedStatus = "queued"
)

func (r BuildTargetTestQueuedStatus) IsKnown() bool {
	switch r {
	case BuildTargetTestQueuedStatusQueued:
		return true
	}
	return false
}

type BuildTargetTestInProgress struct {
	Status BuildTargetTestInProgressStatus `json:"status,required"`
	JSON   buildTargetTestInProgressJSON   `json:"-"`
}

// buildTargetTestInProgressJSON contains the JSON metadata for the struct
// [BuildTargetTestInProgress]
type buildTargetTestInProgressJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestInProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestInProgressJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetTestInProgress) implementsBuildTargetTest() {}

type BuildTargetTestInProgressStatus string

const (
	BuildTargetTestInProgressStatusInProgress BuildTargetTestInProgressStatus = "in_progress"
)

func (r BuildTargetTestInProgressStatus) IsKnown() bool {
	switch r {
	case BuildTargetTestInProgressStatusInProgress:
		return true
	}
	return false
}

type BuildTargetTestCompleted struct {
	Completed BuildTargetTestCompletedCompleted `json:"completed,required"`
	Status    BuildTargetTestCompletedStatus    `json:"status,required"`
	JSON      buildTargetTestCompletedJSON      `json:"-"`
}

// buildTargetTestCompletedJSON contains the JSON metadata for the struct
// [BuildTargetTestCompleted]
type buildTargetTestCompletedJSON struct {
	Completed   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestCompletedJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetTestCompleted) implementsBuildTargetTest() {}

type BuildTargetTestCompletedCompleted struct {
	Conclusion BuildTargetTestCompletedCompletedConclusion `json:"conclusion,required"`
	JSON       buildTargetTestCompletedCompletedJSON       `json:"-"`
}

// buildTargetTestCompletedCompletedJSON contains the JSON metadata for the struct
// [BuildTargetTestCompletedCompleted]
type buildTargetTestCompletedCompletedJSON struct {
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestCompletedCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestCompletedCompletedJSON) RawJSON() string {
	return r.raw
}

type BuildTargetTestCompletedCompletedConclusion string

const (
	BuildTargetTestCompletedCompletedConclusionSuccess        BuildTargetTestCompletedCompletedConclusion = "success"
	BuildTargetTestCompletedCompletedConclusionFailure        BuildTargetTestCompletedCompletedConclusion = "failure"
	BuildTargetTestCompletedCompletedConclusionSkipped        BuildTargetTestCompletedCompletedConclusion = "skipped"
	BuildTargetTestCompletedCompletedConclusionCancelled      BuildTargetTestCompletedCompletedConclusion = "cancelled"
	BuildTargetTestCompletedCompletedConclusionActionRequired BuildTargetTestCompletedCompletedConclusion = "action_required"
	BuildTargetTestCompletedCompletedConclusionNeutral        BuildTargetTestCompletedCompletedConclusion = "neutral"
	BuildTargetTestCompletedCompletedConclusionTimedOut       BuildTargetTestCompletedCompletedConclusion = "timed_out"
)

func (r BuildTargetTestCompletedCompletedConclusion) IsKnown() bool {
	switch r {
	case BuildTargetTestCompletedCompletedConclusionSuccess, BuildTargetTestCompletedCompletedConclusionFailure, BuildTargetTestCompletedCompletedConclusionSkipped, BuildTargetTestCompletedCompletedConclusionCancelled, BuildTargetTestCompletedCompletedConclusionActionRequired, BuildTargetTestCompletedCompletedConclusionNeutral, BuildTargetTestCompletedCompletedConclusionTimedOut:
		return true
	}
	return false
}

type BuildTargetTestCompletedStatus string

const (
	BuildTargetTestCompletedStatusCompleted BuildTargetTestCompletedStatus = "completed"
)

func (r BuildTargetTestCompletedStatus) IsKnown() bool {
	switch r {
	case BuildTargetTestCompletedStatusCompleted:
		return true
	}
	return false
}

type BuildTargetTestStatus string

const (
	BuildTargetTestStatusNotStarted BuildTargetTestStatus = "not_started"
	BuildTargetTestStatusQueued     BuildTargetTestStatus = "queued"
	BuildTargetTestStatusInProgress BuildTargetTestStatus = "in_progress"
	BuildTargetTestStatusCompleted  BuildTargetTestStatus = "completed"
)

func (r BuildTargetTestStatus) IsKnown() bool {
	switch r {
	case BuildTargetTestStatusNotStarted, BuildTargetTestStatusQueued, BuildTargetTestStatusInProgress, BuildTargetTestStatusCompleted:
		return true
	}
	return false
}

type BuildNewParams struct {
	Branch       param.Field[string]                 `json:"branch,required"`
	ConfigCommit param.Field[string]                 `json:"config_commit,required"`
	Project      param.Field[string]                 `json:"project,required"`
	Targets      param.Field[[]BuildNewParamsTarget] `json:"targets"`
}

func (r BuildNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BuildNewParamsTarget string

const (
	BuildNewParamsTargetNode       BuildNewParamsTarget = "node"
	BuildNewParamsTargetTypescript BuildNewParamsTarget = "typescript"
	BuildNewParamsTargetPython     BuildNewParamsTarget = "python"
	BuildNewParamsTargetGo         BuildNewParamsTarget = "go"
	BuildNewParamsTargetJava       BuildNewParamsTarget = "java"
	BuildNewParamsTargetKotlin     BuildNewParamsTarget = "kotlin"
	BuildNewParamsTargetRuby       BuildNewParamsTarget = "ruby"
	BuildNewParamsTargetTerraform  BuildNewParamsTarget = "terraform"
	BuildNewParamsTargetCli        BuildNewParamsTarget = "cli"
)

func (r BuildNewParamsTarget) IsKnown() bool {
	switch r {
	case BuildNewParamsTargetNode, BuildNewParamsTargetTypescript, BuildNewParamsTargetPython, BuildNewParamsTargetGo, BuildNewParamsTargetJava, BuildNewParamsTargetKotlin, BuildNewParamsTargetRuby, BuildNewParamsTargetTerraform, BuildNewParamsTargetCli:
		return true
	}
	return false
}
