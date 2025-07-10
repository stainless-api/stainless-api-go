// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/packages/param"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

func FileInputParamOfFileInputContent(content string) FileInputUnionParam {
	var variant FileInputContentParam
	variant.Content = content
	return FileInputUnionParam{OfFileInputContent: &variant}
}

func FileInputParamOfFileInputURL(url string) FileInputUnionParam {
	var variant FileInputURLParam
	variant.URL = url
	return FileInputUnionParam{OfFileInputURL: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileInputUnionParam struct {
	OfFileInputContent *FileInputContentParam `json:",omitzero,inline"`
	OfFileInputURL     *FileInputURLParam     `json:",omitzero,inline"`
	paramUnion
}

func (u FileInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFileInputContent, u.OfFileInputURL)
}
func (u *FileInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFileInputContent) {
		return u.OfFileInputContent
	} else if !param.IsOmitted(u.OfFileInputURL) {
		return u.OfFileInputURL
	}
	return nil
}

// The property Content is required.
type FileInputContentParam struct {
	// File content
	Content string `json:"content,required"`
	paramObj
}

func (r FileInputContentParam) MarshalJSON() (data []byte, err error) {
	type shadow FileInputContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileInputContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type FileInputURLParam struct {
	// URL to fetch file content from
	URL string `json:"url,required"`
	paramObj
}

func (r FileInputURLParam) MarshalJSON() (data []byte, err error) {
	type shadow FileInputURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileInputURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
