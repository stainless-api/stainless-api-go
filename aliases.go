// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless

import (
	"github.com/stainless-api/stainless-api-go/internal/apierror"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Commit = shared.Commit

// This is an alias to an internal type.
type CommitRepo = shared.CommitRepo

// This is an alias to an internal type.
type FileInputUnionParam = shared.FileInputUnionParam

// This is an alias to an internal type.
type FileInputContentParam = shared.FileInputContentParam

// This is an alias to an internal type.
type FileInputURLParam = shared.FileInputURLParam
