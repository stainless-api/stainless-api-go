// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
)

// BuildTargetArtifactService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBuildTargetArtifactService] method instead.
type BuildTargetArtifactService struct {
	Options []option.RequestOption
}

// NewBuildTargetArtifactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBuildTargetArtifactService(opts ...option.RequestOption) (r *BuildTargetArtifactService) {
	r = &BuildTargetArtifactService{}
	r.Options = opts
	return
}

// TODO
func (r *BuildTargetArtifactService) GetSource(ctx context.Context, buildID string, targetName BuildTargetArtifactGetSourceParamsTargetName, opts ...option.RequestOption) (res *BuildTargetArtifactGetSourceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s/target/%v/artifacts/source", buildID, targetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BuildTargetArtifactGetSourceResponse struct {
	URL  string                                   `json:"url,required"`
	JSON buildTargetArtifactGetSourceResponseJSON `json:"-"`
}

// buildTargetArtifactGetSourceResponseJSON contains the JSON metadata for the
// struct [BuildTargetArtifactGetSourceResponse]
type buildTargetArtifactGetSourceResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BuildTargetArtifactGetSourceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r buildTargetArtifactGetSourceResponseJSON) RawJSON() string {
	return r.raw
}

type BuildTargetArtifactGetSourceParamsTargetName string

const (
	BuildTargetArtifactGetSourceParamsTargetNameNode       BuildTargetArtifactGetSourceParamsTargetName = "node"
	BuildTargetArtifactGetSourceParamsTargetNameTypescript BuildTargetArtifactGetSourceParamsTargetName = "typescript"
	BuildTargetArtifactGetSourceParamsTargetNamePython     BuildTargetArtifactGetSourceParamsTargetName = "python"
	BuildTargetArtifactGetSourceParamsTargetNameGo         BuildTargetArtifactGetSourceParamsTargetName = "go"
	BuildTargetArtifactGetSourceParamsTargetNameJava       BuildTargetArtifactGetSourceParamsTargetName = "java"
	BuildTargetArtifactGetSourceParamsTargetNameKotlin     BuildTargetArtifactGetSourceParamsTargetName = "kotlin"
	BuildTargetArtifactGetSourceParamsTargetNameRuby       BuildTargetArtifactGetSourceParamsTargetName = "ruby"
	BuildTargetArtifactGetSourceParamsTargetNameTerraform  BuildTargetArtifactGetSourceParamsTargetName = "terraform"
	BuildTargetArtifactGetSourceParamsTargetNameCli        BuildTargetArtifactGetSourceParamsTargetName = "cli"
)

func (r BuildTargetArtifactGetSourceParamsTargetName) IsKnown() bool {
	switch r {
	case BuildTargetArtifactGetSourceParamsTargetNameNode, BuildTargetArtifactGetSourceParamsTargetNameTypescript, BuildTargetArtifactGetSourceParamsTargetNamePython, BuildTargetArtifactGetSourceParamsTargetNameGo, BuildTargetArtifactGetSourceParamsTargetNameJava, BuildTargetArtifactGetSourceParamsTargetNameKotlin, BuildTargetArtifactGetSourceParamsTargetNameRuby, BuildTargetArtifactGetSourceParamsTargetNameTerraform, BuildTargetArtifactGetSourceParamsTargetNameCli:
		return true
	}
	return false
}
