// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/stainless-api/stainless-api-go/internal/apijson"
	"github.com/stainless-api/stainless-api-go/internal/requestconfig"
	"github.com/stainless-api/stainless-api-go/option"
	"github.com/stainless-api/stainless-api-go/packages/param"
	"github.com/stainless-api/stainless-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type List[T any] struct {
	Data       []T    `json:"data"`
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r List[T]) RawJSON() string { return r.JSON.raw }
func (r *List[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *List[T]) GetNextPage() (res *List[T], err error) {
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *List[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &List[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ListAutoPager[T any] struct {
	page *List[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewListAutoPager[T any](page *List[T], err error) *ListAutoPager[T] {
	return &ListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ListAutoPager[T]) Current() T {
	return r.cur
}

func (r *ListAutoPager[T]) Err() error {
	return r.err
}

func (r *ListAutoPager[T]) Index() int {
	return r.run
}
