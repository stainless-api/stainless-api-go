// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"net/http"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// GenerateService contains methods and other services that help with interacting
// with the stainless API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGenerateService] method instead.
type GenerateService struct {
	Options []option.RequestOption
}

// NewGenerateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGenerateService(opts ...option.RequestOption) (r GenerateService) {
	r = GenerateService{}
	r.Options = opts
	return
}

func (r *GenerateService) NewSpec(ctx context.Context, body GenerateNewSpecParams, opts ...option.RequestOption) (res *GenerateNewSpecResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.Project, precfg.Project)
	path := "v0/generate/spec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GenerateNewSpecResponse struct {
	Spec any `json:"spec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerateNewSpecResponse) RawJSON() string { return r.JSON.raw }
func (r *GenerateNewSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerateNewSpecParams struct {
	Project param.Opt[string]                `json:"project,omitzero,required"`
	Source  GenerateNewSpecParamsSourceUnion `json:"source,omitzero,required"`
	paramObj
}

func (r GenerateNewSpecParams) MarshalJSON() (data []byte, err error) {
	type shadow GenerateNewSpecParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerateNewSpecParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GenerateNewSpecParamsSourceUnion struct {
	OfGenerateNewSpecsSourceObject *GenerateNewSpecParamsSourceObject `json:",omitzero,inline"`
	OfVariant2                     *GenerateNewSpecParamsSourceObject `json:",omitzero,inline"`
	paramUnion
}

func (u GenerateNewSpecParamsSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[GenerateNewSpecParamsSourceUnion](u.OfGenerateNewSpecsSourceObject, u.OfVariant2)
}
func (u *GenerateNewSpecParamsSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GenerateNewSpecParamsSourceUnion) asAny() any {
	if !param.IsOmitted(u.OfGenerateNewSpecsSourceObject) {
		return u.OfGenerateNewSpecsSourceObject
	} else if !param.IsOmitted(u.OfVariant2) {
		return u.OfVariant2
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GenerateNewSpecParamsSourceUnion) GetRevision() *string {
	if vt := u.OfGenerateNewSpecsSourceObject; vt != nil {
		return &vt.Revision
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GenerateNewSpecParamsSourceUnion) GetOpenAPISpec() *string {
	if vt := u.OfVariant2; vt != nil {
		return &vt.OpenAPISpec
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GenerateNewSpecParamsSourceUnion) GetStainlessConfig() *string {
	if vt := u.OfVariant2; vt != nil {
		return &vt.StainlessConfig
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GenerateNewSpecParamsSourceUnion) GetType() *string {
	if vt := u.OfGenerateNewSpecsSourceObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant2; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties Revision, Type are required.
type GenerateNewSpecParamsSourceObject struct {
	Revision string `json:"revision,required"`
	// Any of "git".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r GenerateNewSpecParamsSourceObject) MarshalJSON() (data []byte, err error) {
	type shadow GenerateNewSpecParamsSourceObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerateNewSpecParamsSourceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GenerateNewSpecParamsSourceObject](
		"type", "git",
	)
}
