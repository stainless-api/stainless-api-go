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
	"github.com/stainless-api/stainless-api-go/packages/resp"
)

// TargetArtifactService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTargetArtifactService] method instead.
type TargetArtifactService struct {
	Options []option.RequestOption
}

// NewTargetArtifactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTargetArtifactService(opts ...option.RequestOption) (r TargetArtifactService) {
	r = TargetArtifactService{}
	r.Options = opts
	return
}

// TODO
func (r *TargetArtifactService) Get(ctx context.Context, buildID string, targetName TargetArtifactGetParamsTargetName, opts ...option.RequestOption) (res *TargetArtifactGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if buildID == "" {
		err = errors.New("missing required buildId parameter")
		return
	}
	path := fmt.Sprintf("v0/builds/%s/targets/%v/artifacts/source", buildID, targetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TargetArtifactGetResponse struct {
	URL string `json:"url,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		URL         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TargetArtifactGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TargetArtifactGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TargetArtifactGetParamsTargetName string

const (
	TargetArtifactGetParamsTargetNameNode       TargetArtifactGetParamsTargetName = "node"
	TargetArtifactGetParamsTargetNameTypescript TargetArtifactGetParamsTargetName = "typescript"
	TargetArtifactGetParamsTargetNamePython     TargetArtifactGetParamsTargetName = "python"
	TargetArtifactGetParamsTargetNameGo         TargetArtifactGetParamsTargetName = "go"
	TargetArtifactGetParamsTargetNameJava       TargetArtifactGetParamsTargetName = "java"
	TargetArtifactGetParamsTargetNameKotlin     TargetArtifactGetParamsTargetName = "kotlin"
	TargetArtifactGetParamsTargetNameRuby       TargetArtifactGetParamsTargetName = "ruby"
	TargetArtifactGetParamsTargetNameTerraform  TargetArtifactGetParamsTargetName = "terraform"
	TargetArtifactGetParamsTargetNameCli        TargetArtifactGetParamsTargetName = "cli"
)
