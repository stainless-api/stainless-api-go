// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	"encoding/json"
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
type InProgress string // Always "in_progress"
type NotStarted string // Always "not_started"
type Queued string     // Always "queued"

func (c Completed) Default() Completed   { return "completed" }
func (c InProgress) Default() InProgress { return "in_progress" }
func (c NotStarted) Default() NotStarted { return "not_started" }
func (c Queued) Default() Queued         { return "queued" }

func (c Completed) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c InProgress) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c NotStarted) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Queued) MarshalJSON() ([]byte, error)     { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return json.Marshal(string(v))
}
