// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/stainless-v0-go/internal/requestconfig"
	"github.com/stainless-sdks/stainless-v0-go/option"
)

// OpenAPIService contains methods and other services that help with interacting
// with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAPIService] method instead.
type OpenAPIService struct {
	Options []option.RequestOption
}

// NewOpenAPIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOpenAPIService(opts ...option.RequestOption) (r OpenAPIService) {
	r = OpenAPIService{}
	r.Options = opts
	return
}

// TODO
func (r *OpenAPIService) Get(ctx context.Context, opts ...option.RequestOption) (res *OpenAPIGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/openapi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OpenAPIGetResponse = any
