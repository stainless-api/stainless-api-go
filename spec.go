// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
)

// SpecService contains methods and other services that help with interacting with
// the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpecService] method instead.
type SpecService struct {
	Options []option.RequestOption
}

// NewSpecService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpecService(opts ...option.RequestOption) (r SpecService) {
	r = SpecService{}
	r.Options = opts
	return
}

// Retrieve the decorated spec for a given application and project.
func (r *SpecService) GetDecoratedSpec(ctx context.Context, projectName string, query SpecGetDecoratedSpecParams, opts ...option.RequestOption) (res *SpecGetDecoratedSpecResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ClientID == "" {
		err = errors.New("missing required clientId parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required projectName parameter")
		return
	}
	path := fmt.Sprintf("v0/spec/application/%s/%s", url.PathEscape(query.ClientID), url.PathEscape(projectName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SpecGetDecoratedSpecResponse = any

type SpecGetDecoratedSpecParams struct {
	ClientID string `path:"clientId,required" json:"-"`
	paramObj
}
