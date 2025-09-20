// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-api/stainless-api-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Completed string  // Always "completed"
type Git string        // Always "git"
type InProgress string // Always "in_progress"
type Markdown string   // Always "markdown"
type NotStarted string // Always "not_started"
type Queued string     // Always "queued"
type Raw string        // Always "raw"
type URL string        // Always "url"

func (c Completed) Default() Completed   { return "completed" }
func (c Git) Default() Git               { return "git" }
func (c InProgress) Default() InProgress { return "in_progress" }
func (c Markdown) Default() Markdown     { return "markdown" }
func (c NotStarted) Default() NotStarted { return "not_started" }
func (c Queued) Default() Queued         { return "queued" }
func (c Raw) Default() Raw               { return "raw" }
func (c URL) Default() URL               { return "url" }

func (c Completed) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Git) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c InProgress) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Markdown) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c NotStarted) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Queued) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Raw) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c URL) MarshalJSON() ([]byte, error)        { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
