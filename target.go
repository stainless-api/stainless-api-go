// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"github.com/stainless-api/stainless-api-go/option"
)

// TargetService contains methods and other services that help with interacting
// with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTargetService] method instead.
type TargetService struct {
	Options   []option.RequestOption
	Artifacts TargetArtifactService
}

// NewTargetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTargetService(opts ...option.RequestOption) (r TargetService) {
	r = TargetService{}
	r.Options = opts
	r.Artifacts = NewTargetArtifactService(opts...)
	return
}
