// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/stainless-v0-go/internal/apijson"
	"github.com/stainless-sdks/stainless-v0-go/internal/requestconfig"
	"github.com/stainless-sdks/stainless-v0-go/option"
	"github.com/tidwall/gjson"
)

// BuildTargetService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuildTargetService] method instead.
type BuildTargetService struct {
	Options   []option.RequestOption
	Artifacts *BuildTargetArtifactService
}

// NewBuildTargetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBuildTargetService(opts ...option.RequestOption) (r *BuildTargetService) {
	r = &BuildTargetService{}
	r.Options = opts
	r.Artifacts = NewBuildTargetArtifactService(opts...)
	return
}

// TODO
func (r *BuildTargetService) Get(ctx context.Context, buildID string, targetName BuildTargetGetParamsTargetName, opts ...option.RequestOption) (res *BuildTarget, err error) {
	opts = append(r.Options[:], opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s/target/%v", buildID, targetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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
	// This field can have the runtime type of [BuildTargetCommitObjectCompleted].
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
// Possible runtime types of the union are [BuildTargetCommitStatus],
// [BuildTargetCommitStatus], [BuildTargetCommitStatus], [BuildTargetCommitObject].
func (r BuildTargetCommit) AsUnion() BuildTargetCommitUnion {
	return r.union
}

// Union satisfied by [BuildTargetCommitStatus], [BuildTargetCommitStatus],
// [BuildTargetCommitStatus] or [BuildTargetCommitObject].
type BuildTargetCommitUnion interface {
	implementsBuildTargetCommit()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetCommitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitObject{}),
		},
	)
}

type BuildTargetCommitStatus struct {
	Status BuildTargetCommitStatusStatus `json:"status,required"`
	JSON   buildTargetCommitStatusJSON   `json:"-"`
}

// buildTargetCommitStatusJSON contains the JSON metadata for the struct
// [BuildTargetCommitStatus]
type buildTargetCommitStatusJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitStatusJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitStatus) implementsBuildTargetCommit() {}

type BuildTargetCommitStatusStatus string

const (
	BuildTargetCommitStatusStatusNotStarted BuildTargetCommitStatusStatus = "not_started"
)

func (r BuildTargetCommitStatusStatus) IsKnown() bool {
	switch r {
	case BuildTargetCommitStatusStatusNotStarted:
		return true
	}
	return false
}

type BuildTargetCommitObject struct {
	Completed BuildTargetCommitObjectCompleted `json:"completed,required"`
	Status    BuildTargetCommitObjectStatus    `json:"status,required"`
	JSON      buildTargetCommitObjectJSON      `json:"-"`
}

// buildTargetCommitObjectJSON contains the JSON metadata for the struct
// [BuildTargetCommitObject]
type buildTargetCommitObjectJSON struct {
	Completed   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitObjectJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitObject) implementsBuildTargetCommit() {}

type BuildTargetCommitObjectCompleted struct {
	Conclusion BuildTargetCommitObjectCompletedConclusion `json:"conclusion,required"`
	// This field can have the runtime type of
	// [BuildTargetCommitObjectCompletedObjectCommit].
	Commit interface{} `json:"commit"`
	// This field can have the runtime type of
	// [BuildTargetCommitObjectCompletedObjectMergeConflictPr].
	MergeConflictPr interface{}                          `json:"merge_conflict_pr"`
	JSON            buildTargetCommitObjectCompletedJSON `json:"-"`
	union           BuildTargetCommitObjectCompletedUnion
}

// buildTargetCommitObjectCompletedJSON contains the JSON metadata for the struct
// [BuildTargetCommitObjectCompleted]
type buildTargetCommitObjectCompletedJSON struct {
	Conclusion      apijson.Field
	Commit          apijson.Field
	MergeConflictPr apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r buildTargetCommitObjectCompletedJSON) RawJSON() string {
	return r.raw
}

func (r *BuildTargetCommitObjectCompleted) UnmarshalJSON(data []byte) (err error) {
	*r = BuildTargetCommitObjectCompleted{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BuildTargetCommitObjectCompletedUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BuildTargetCommitObjectCompletedObject],
// [BuildTargetCommitObjectCompletedObject],
// [BuildTargetCommitObjectCompletedConclusion].
func (r BuildTargetCommitObjectCompleted) AsUnion() BuildTargetCommitObjectCompletedUnion {
	return r.union
}

// Union satisfied by [BuildTargetCommitObjectCompletedObject],
// [BuildTargetCommitObjectCompletedObject] or
// [BuildTargetCommitObjectCompletedConclusion].
type BuildTargetCommitObjectCompletedUnion interface {
	implementsBuildTargetCommitObjectCompleted()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetCommitObjectCompletedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitObjectCompletedObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitObjectCompletedObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetCommitObjectCompletedConclusion{}),
		},
	)
}

type BuildTargetCommitObjectCompletedObject struct {
	Commit     BuildTargetCommitObjectCompletedObjectCommit     `json:"commit,required"`
	Conclusion BuildTargetCommitObjectCompletedObjectConclusion `json:"conclusion,required"`
	JSON       buildTargetCommitObjectCompletedObjectJSON       `json:"-"`
}

// buildTargetCommitObjectCompletedObjectJSON contains the JSON metadata for the
// struct [BuildTargetCommitObjectCompletedObject]
type buildTargetCommitObjectCompletedObjectJSON struct {
	Commit      apijson.Field
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitObjectCompletedObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitObjectCompletedObjectJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitObjectCompletedObject) implementsBuildTargetCommitObjectCompleted() {}

type BuildTargetCommitObjectCompletedObjectCommit struct {
	Repo BuildTargetCommitObjectCompletedObjectCommitRepo `json:"repo,required"`
	Sha  string                                           `json:"sha,required"`
	JSON buildTargetCommitObjectCompletedObjectCommitJSON `json:"-"`
}

// buildTargetCommitObjectCompletedObjectCommitJSON contains the JSON metadata for
// the struct [BuildTargetCommitObjectCompletedObjectCommit]
type buildTargetCommitObjectCompletedObjectCommitJSON struct {
	Repo        apijson.Field
	Sha         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitObjectCompletedObjectCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitObjectCompletedObjectCommitJSON) RawJSON() string {
	return r.raw
}

type BuildTargetCommitObjectCompletedObjectCommitRepo struct {
	Branch string                                               `json:"branch,required"`
	Name   string                                               `json:"name,required"`
	Owner  string                                               `json:"owner,required"`
	JSON   buildTargetCommitObjectCompletedObjectCommitRepoJSON `json:"-"`
}

// buildTargetCommitObjectCompletedObjectCommitRepoJSON contains the JSON metadata
// for the struct [BuildTargetCommitObjectCompletedObjectCommitRepo]
type buildTargetCommitObjectCompletedObjectCommitRepoJSON struct {
	Branch      apijson.Field
	Name        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitObjectCompletedObjectCommitRepo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitObjectCompletedObjectCommitRepoJSON) RawJSON() string {
	return r.raw
}

type BuildTargetCommitObjectCompletedObjectConclusion string

const (
	BuildTargetCommitObjectCompletedObjectConclusionError   BuildTargetCommitObjectCompletedObjectConclusion = "error"
	BuildTargetCommitObjectCompletedObjectConclusionWarning BuildTargetCommitObjectCompletedObjectConclusion = "warning"
	BuildTargetCommitObjectCompletedObjectConclusionNote    BuildTargetCommitObjectCompletedObjectConclusion = "note"
	BuildTargetCommitObjectCompletedObjectConclusionSuccess BuildTargetCommitObjectCompletedObjectConclusion = "success"
)

func (r BuildTargetCommitObjectCompletedObjectConclusion) IsKnown() bool {
	switch r {
	case BuildTargetCommitObjectCompletedObjectConclusionError, BuildTargetCommitObjectCompletedObjectConclusionWarning, BuildTargetCommitObjectCompletedObjectConclusionNote, BuildTargetCommitObjectCompletedObjectConclusionSuccess:
		return true
	}
	return false
}

type BuildTargetCommitObjectCompletedConclusion struct {
	Conclusion BuildTargetCommitObjectCompletedConclusionConclusion `json:"conclusion,required"`
	JSON       buildTargetCommitObjectCompletedConclusionJSON       `json:"-"`
}

// buildTargetCommitObjectCompletedConclusionJSON contains the JSON metadata for
// the struct [BuildTargetCommitObjectCompletedConclusion]
type buildTargetCommitObjectCompletedConclusionJSON struct {
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetCommitObjectCompletedConclusion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetCommitObjectCompletedConclusionJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetCommitObjectCompletedConclusion) implementsBuildTargetCommitObjectCompleted() {}

type BuildTargetCommitObjectCompletedConclusionConclusion string

const (
	BuildTargetCommitObjectCompletedConclusionConclusionFatal           BuildTargetCommitObjectCompletedConclusionConclusion = "fatal"
	BuildTargetCommitObjectCompletedConclusionConclusionPaymentRequired BuildTargetCommitObjectCompletedConclusionConclusion = "payment_required"
	BuildTargetCommitObjectCompletedConclusionConclusionCancelled       BuildTargetCommitObjectCompletedConclusionConclusion = "cancelled"
	BuildTargetCommitObjectCompletedConclusionConclusionTimedOut        BuildTargetCommitObjectCompletedConclusionConclusion = "timed_out"
	BuildTargetCommitObjectCompletedConclusionConclusionNoop            BuildTargetCommitObjectCompletedConclusionConclusion = "noop"
	BuildTargetCommitObjectCompletedConclusionConclusionVersionBump     BuildTargetCommitObjectCompletedConclusionConclusion = "version_bump"
)

func (r BuildTargetCommitObjectCompletedConclusionConclusion) IsKnown() bool {
	switch r {
	case BuildTargetCommitObjectCompletedConclusionConclusionFatal, BuildTargetCommitObjectCompletedConclusionConclusionPaymentRequired, BuildTargetCommitObjectCompletedConclusionConclusionCancelled, BuildTargetCommitObjectCompletedConclusionConclusionTimedOut, BuildTargetCommitObjectCompletedConclusionConclusionNoop, BuildTargetCommitObjectCompletedConclusionConclusionVersionBump:
		return true
	}
	return false
}

type BuildTargetCommitObjectStatus string

const (
	BuildTargetCommitObjectStatusCompleted BuildTargetCommitObjectStatus = "completed"
)

func (r BuildTargetCommitObjectStatus) IsKnown() bool {
	switch r {
	case BuildTargetCommitObjectStatusCompleted:
		return true
	}
	return false
}

type BuildTargetLint struct {
	Status BuildTargetLintStatus `json:"status,required"`
	// This field can have the runtime type of [BuildTargetLintObjectCompleted].
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
// Possible runtime types of the union are [BuildTargetLintStatus],
// [BuildTargetLintStatus], [BuildTargetLintStatus], [BuildTargetLintObject].
func (r BuildTargetLint) AsUnion() BuildTargetLintUnion {
	return r.union
}

// Union satisfied by [BuildTargetLintStatus], [BuildTargetLintStatus],
// [BuildTargetLintStatus] or [BuildTargetLintObject].
type BuildTargetLintUnion interface {
	implementsBuildTargetLint()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetLintUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetLintStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetLintStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetLintStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetLintObject{}),
		},
	)
}

type BuildTargetLintStatus struct {
	Status BuildTargetLintStatusStatus `json:"status,required"`
	JSON   buildTargetLintStatusJSON   `json:"-"`
}

// buildTargetLintStatusJSON contains the JSON metadata for the struct
// [BuildTargetLintStatus]
type buildTargetLintStatusJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintStatusJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetLintStatus) implementsBuildTargetLint() {}

type BuildTargetLintStatusStatus string

const (
	BuildTargetLintStatusStatusNotStarted BuildTargetLintStatusStatus = "not_started"
)

func (r BuildTargetLintStatusStatus) IsKnown() bool {
	switch r {
	case BuildTargetLintStatusStatusNotStarted:
		return true
	}
	return false
}

type BuildTargetLintObject struct {
	Completed BuildTargetLintObjectCompleted `json:"completed,required"`
	Status    BuildTargetLintObjectStatus    `json:"status,required"`
	JSON      buildTargetLintObjectJSON      `json:"-"`
}

// buildTargetLintObjectJSON contains the JSON metadata for the struct
// [BuildTargetLintObject]
type buildTargetLintObjectJSON struct {
	Completed   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintObjectJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetLintObject) implementsBuildTargetLint() {}

type BuildTargetLintObjectCompleted struct {
	Conclusion BuildTargetLintObjectCompletedConclusion `json:"conclusion,required"`
	JSON       buildTargetLintObjectCompletedJSON       `json:"-"`
}

// buildTargetLintObjectCompletedJSON contains the JSON metadata for the struct
// [BuildTargetLintObjectCompleted]
type buildTargetLintObjectCompletedJSON struct {
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetLintObjectCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetLintObjectCompletedJSON) RawJSON() string {
	return r.raw
}

type BuildTargetLintObjectCompletedConclusion string

const (
	BuildTargetLintObjectCompletedConclusionSuccess        BuildTargetLintObjectCompletedConclusion = "success"
	BuildTargetLintObjectCompletedConclusionFailure        BuildTargetLintObjectCompletedConclusion = "failure"
	BuildTargetLintObjectCompletedConclusionSkipped        BuildTargetLintObjectCompletedConclusion = "skipped"
	BuildTargetLintObjectCompletedConclusionCancelled      BuildTargetLintObjectCompletedConclusion = "cancelled"
	BuildTargetLintObjectCompletedConclusionActionRequired BuildTargetLintObjectCompletedConclusion = "action_required"
	BuildTargetLintObjectCompletedConclusionNeutral        BuildTargetLintObjectCompletedConclusion = "neutral"
	BuildTargetLintObjectCompletedConclusionTimedOut       BuildTargetLintObjectCompletedConclusion = "timed_out"
)

func (r BuildTargetLintObjectCompletedConclusion) IsKnown() bool {
	switch r {
	case BuildTargetLintObjectCompletedConclusionSuccess, BuildTargetLintObjectCompletedConclusionFailure, BuildTargetLintObjectCompletedConclusionSkipped, BuildTargetLintObjectCompletedConclusionCancelled, BuildTargetLintObjectCompletedConclusionActionRequired, BuildTargetLintObjectCompletedConclusionNeutral, BuildTargetLintObjectCompletedConclusionTimedOut:
		return true
	}
	return false
}

type BuildTargetLintObjectStatus string

const (
	BuildTargetLintObjectStatusCompleted BuildTargetLintObjectStatus = "completed"
)

func (r BuildTargetLintObjectStatus) IsKnown() bool {
	switch r {
	case BuildTargetLintObjectStatusCompleted:
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
	// This field can have the runtime type of [BuildTargetTestObjectCompleted].
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
// Possible runtime types of the union are [BuildTargetTestStatus],
// [BuildTargetTestStatus], [BuildTargetTestStatus], [BuildTargetTestObject].
func (r BuildTargetTest) AsUnion() BuildTargetTestUnion {
	return r.union
}

// Union satisfied by [BuildTargetTestStatus], [BuildTargetTestStatus],
// [BuildTargetTestStatus] or [BuildTargetTestObject].
type BuildTargetTestUnion interface {
	implementsBuildTargetTest()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BuildTargetTestUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetTestStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetTestStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetTestStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BuildTargetTestObject{}),
		},
	)
}

type BuildTargetTestStatus struct {
	Status BuildTargetTestStatusStatus `json:"status,required"`
	JSON   buildTargetTestStatusJSON   `json:"-"`
}

// buildTargetTestStatusJSON contains the JSON metadata for the struct
// [BuildTargetTestStatus]
type buildTargetTestStatusJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestStatusJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetTestStatus) implementsBuildTargetTest() {}

type BuildTargetTestStatusStatus string

const (
	BuildTargetTestStatusStatusNotStarted BuildTargetTestStatusStatus = "not_started"
)

func (r BuildTargetTestStatusStatus) IsKnown() bool {
	switch r {
	case BuildTargetTestStatusStatusNotStarted:
		return true
	}
	return false
}

type BuildTargetTestObject struct {
	Completed BuildTargetTestObjectCompleted `json:"completed,required"`
	Status    BuildTargetTestObjectStatus    `json:"status,required"`
	JSON      buildTargetTestObjectJSON      `json:"-"`
}

// buildTargetTestObjectJSON contains the JSON metadata for the struct
// [BuildTargetTestObject]
type buildTargetTestObjectJSON struct {
	Completed   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestObjectJSON) RawJSON() string {
	return r.raw
}

func (r BuildTargetTestObject) implementsBuildTargetTest() {}

type BuildTargetTestObjectCompleted struct {
	Conclusion BuildTargetTestObjectCompletedConclusion `json:"conclusion,required"`
	JSON       buildTargetTestObjectCompletedJSON       `json:"-"`
}

// buildTargetTestObjectCompletedJSON contains the JSON metadata for the struct
// [BuildTargetTestObjectCompleted]
type buildTargetTestObjectCompletedJSON struct {
	Conclusion  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetTestObjectCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetTestObjectCompletedJSON) RawJSON() string {
	return r.raw
}

type BuildTargetTestObjectCompletedConclusion string

const (
	BuildTargetTestObjectCompletedConclusionSuccess        BuildTargetTestObjectCompletedConclusion = "success"
	BuildTargetTestObjectCompletedConclusionFailure        BuildTargetTestObjectCompletedConclusion = "failure"
	BuildTargetTestObjectCompletedConclusionSkipped        BuildTargetTestObjectCompletedConclusion = "skipped"
	BuildTargetTestObjectCompletedConclusionCancelled      BuildTargetTestObjectCompletedConclusion = "cancelled"
	BuildTargetTestObjectCompletedConclusionActionRequired BuildTargetTestObjectCompletedConclusion = "action_required"
	BuildTargetTestObjectCompletedConclusionNeutral        BuildTargetTestObjectCompletedConclusion = "neutral"
	BuildTargetTestObjectCompletedConclusionTimedOut       BuildTargetTestObjectCompletedConclusion = "timed_out"
)

func (r BuildTargetTestObjectCompletedConclusion) IsKnown() bool {
	switch r {
	case BuildTargetTestObjectCompletedConclusionSuccess, BuildTargetTestObjectCompletedConclusionFailure, BuildTargetTestObjectCompletedConclusionSkipped, BuildTargetTestObjectCompletedConclusionCancelled, BuildTargetTestObjectCompletedConclusionActionRequired, BuildTargetTestObjectCompletedConclusionNeutral, BuildTargetTestObjectCompletedConclusionTimedOut:
		return true
	}
	return false
}

type BuildTargetTestObjectStatus string

const (
	BuildTargetTestObjectStatusCompleted BuildTargetTestObjectStatus = "completed"
)

func (r BuildTargetTestObjectStatus) IsKnown() bool {
	switch r {
	case BuildTargetTestObjectStatusCompleted:
		return true
	}
	return false
}

type BuildTargetGetParamsTargetName string

const (
	BuildTargetGetParamsTargetNameNode       BuildTargetGetParamsTargetName = "node"
	BuildTargetGetParamsTargetNameTypescript BuildTargetGetParamsTargetName = "typescript"
	BuildTargetGetParamsTargetNamePython     BuildTargetGetParamsTargetName = "python"
	BuildTargetGetParamsTargetNameGo         BuildTargetGetParamsTargetName = "go"
	BuildTargetGetParamsTargetNameJava       BuildTargetGetParamsTargetName = "java"
	BuildTargetGetParamsTargetNameKotlin     BuildTargetGetParamsTargetName = "kotlin"
	BuildTargetGetParamsTargetNameRuby       BuildTargetGetParamsTargetName = "ruby"
	BuildTargetGetParamsTargetNameTerraform  BuildTargetGetParamsTargetName = "terraform"
	BuildTargetGetParamsTargetNameCli        BuildTargetGetParamsTargetName = "cli"
)

func (r BuildTargetGetParamsTargetName) IsKnown() bool {
	switch r {
	case BuildTargetGetParamsTargetNameNode, BuildTargetGetParamsTargetNameTypescript, BuildTargetGetParamsTargetNamePython, BuildTargetGetParamsTargetNameGo, BuildTargetGetParamsTargetNameJava, BuildTargetGetParamsTargetNameKotlin, BuildTargetGetParamsTargetNameRuby, BuildTargetGetParamsTargetNameTerraform, BuildTargetGetParamsTargetNameCli:
		return true
	}
	return false
}
