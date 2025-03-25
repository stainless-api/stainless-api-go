// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/stainless-v0-go/internal/apijson"
	"github.com/stainless-sdks/stainless-v0-go/internal/requestconfig"
	"github.com/stainless-sdks/stainless-v0-go/option"
)

// BuildService contains methods and other services that help with interacting with
// the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuildService] method instead.
type BuildService struct {
	Options []option.RequestOption
	Target  *BuildTargetService
}

// NewBuildService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBuildService(opts ...option.RequestOption) (r *BuildService) {
	r = &BuildService{}
	r.Options = opts
	r.Target = NewBuildTargetService(opts...)
	return
}

// TODO
func (r *BuildService) Get(ctx context.Context, buildID string, opts ...option.RequestOption) (res *BuildGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s", buildID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BuildGetResponse struct {
	ID           string                  `json:"id,required"`
	ConfigCommit string                  `json:"config_commit,required"`
	Object       BuildGetResponseObject  `json:"object,required"`
	Targets      BuildGetResponseTargets `json:"targets,required"`
	JSON         buildGetResponseJSON    `json:"-"`
}

// buildGetResponseJSON contains the JSON metadata for the struct
// [BuildGetResponse]
type buildGetResponseJSON struct {
	ID           apijson.Field
	ConfigCommit apijson.Field
	Object       apijson.Field
	Targets      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BuildGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildGetResponseJSON) RawJSON() string {
	return r.raw
}

type BuildGetResponseObject string

const (
	BuildGetResponseObjectBuild BuildGetResponseObject = "build"
)

func (r BuildGetResponseObject) IsKnown() bool {
	switch r {
	case BuildGetResponseObjectBuild:
		return true
	}
	return false
}

type BuildGetResponseTargets struct {
	Cli        BuildTarget                 `json:"cli"`
	Go         BuildTarget                 `json:"go"`
	Java       BuildTarget                 `json:"java"`
	Kotlin     BuildTarget                 `json:"kotlin"`
	Node       BuildTarget                 `json:"node"`
	Python     BuildTarget                 `json:"python"`
	Ruby       BuildTarget                 `json:"ruby"`
	Terraform  BuildTarget                 `json:"terraform"`
	Typescript BuildTarget                 `json:"typescript"`
	JSON       buildGetResponseTargetsJSON `json:"-"`
}

// buildGetResponseTargetsJSON contains the JSON metadata for the struct
// [BuildGetResponseTargets]
type buildGetResponseTargetsJSON struct {
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

func (r *BuildGetResponseTargets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildGetResponseTargetsJSON) RawJSON() string {
	return r.raw
}
