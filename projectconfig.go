// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"github.com/stainless-api/stainless-api-go/option"
)

// ProjectConfigService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectConfigService] method instead.
type ProjectConfigService struct {
	Options  []option.RequestOption
	Commits  ProjectConfigCommitService
	Branches ProjectConfigBranchService
}

// NewProjectConfigService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectConfigService(opts ...option.RequestOption) (r ProjectConfigService) {
	r = ProjectConfigService{}
	r.Options = opts
	r.Commits = NewProjectConfigCommitService(opts...)
	r.Branches = NewProjectConfigBranchService(opts...)
	return
}
