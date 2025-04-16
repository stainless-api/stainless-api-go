# Stainless V0 Go API Library

<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go"><img src="https://pkg.go.dev/badge/github.com/stainless-api/stainless-api-go.svg" alt="Go Reference"></a>

The Stainless V0 Go library provides convenient access to the [Stainless V0 REST API](https://docs.stainless-v0.com)
from applications written in Go.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

<!-- x-release-please-start-version -->

```go
import (
	"github.com/stainless-api/stainless-api-go" // imported as stainlessv0
)
```

<!-- x-release-please-end -->

Or to pin the version:

<!-- x-release-please-start-version -->

```sh
go get -u 'github.com/stainless-api/stainless-api-go@v0.1.0-alpha.3'
```

<!-- x-release-please-end -->

## Requirements

This library requires Go 1.18+.

## Usage

The full API of this library can be found in [api.md](api.md).

```go
package main

import (
	"context"
	"fmt"

	"github.com/stainless-api/stainless-api-go"
	"github.com/stainless-api/stainless-api-go/option"
)

func main() {
	client := stainlessv0.NewClient(
		option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("STAINLESS_V0_API_KEY")
	)
	commit, err := client.Projects.Config.Commits.New(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectConfigCommitNewParams{
			Branch:        "branch",
			CommitMessage: "commit_message",
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", commit.ID)
}

```

### Request fields

The stainlessv0 library uses the [`omitzero`](https://tip.golang.org/doc/go1.24#encodingjsonpkgencodingjson)
semantics from the Go 1.24+ `encoding/json` release for request fields.

Required primitive fields (`int64`, `string`, etc.) feature the tag <code>\`json:...,required\`</code>. These
fields are always serialized, even their zero values.

Optional primitive types are wrapped in a `param.Opt[T]`. Use the provided constructors set `param.Opt[T]` fields such as `stainlessv0.String(string)`, `stainlessv0.Int(int64)`, etc.

Optional primitives, maps, slices and structs and string enums (represented as `string`) always feature the
tag <code>\`json:"...,omitzero"\`</code>. Their zero values are considered omitted.

Any non-nil slice of length zero will serialize as an empty JSON array, `"[]"`. Similarly, any non-nil map with length zero with serialize as an empty JSON object, `"{}"`.

To send `null` instead of an `param.Opt[T]`, use `param.NullOpt[T]()`.
To send `null` instead of a struct, use `param.NullObj[T]()`, where `T` is a struct.
To send a custom value instead of a struct, use `param.OverrideObj[T](value)`.

To override request structs contain a `.WithExtraFields(map[string]any)` method which can be used to
send non-conforming fields in the request body. Extra fields overwrite any struct fields with a matching
key, so only use with trusted data.

```go
params := stainlessv0.ExampleParams{
	ID:          "id_xxx",                  // required property
	Name:        stainlessv0.String("..."), // optional property
	Description: param.NullOpt[string](),   // explicit null property

	Point: stainlessv0.Point{
		X: 0,                  // required field will serialize as 0
		Y: stainlessv0.Int(1), // optional field will serialize as 1
		// ... omitted non-required fields will not be serialized
	},

	Origin: stainlessv0.Origin{}, // the zero value of [Origin] is considered omitted
}

// In cases where the API specifies a given type,
// but you want to send something else, use [WithExtraFields]:
params.WithExtraFields(map[string]any{
	"x": 0.01, // send "x" as a float instead of int
})

// Send a number instead of an object
custom := param.OverrideObj[stainlessv0.FooParams](12)
```

When available, use the `.IsPresent()` method to check if an optional parameter is not omitted or `null`.
Otherwise, the `param.IsOmitted(any)` function can confirm the presence of any `omitzero` field.

### Request unions

Unions are represented as a struct with fields prefixed by "Of" for each of it's variants,
only one field can be non-zero. The non-zero field will be serialized.

Sub-properties of the union can be accessed via methods on the union struct.
These methods return a mutable pointer to the underlying data, if present.

```go
// Only one field can be non-zero, use param.IsOmitted() to check if a field is set
type AnimalUnionParam struct {
	OfCat *Cat `json:",omitzero,inline`
	OfDog *Dog `json:",omitzero,inline`
}

animal := AnimalUnionParam{
	OfCat: &Cat{
		Name: "Whiskers",
		Owner: PersonParam{
			Address: AddressParam{Street: "3333 Coyote Hill Rd", Zip: 0},
		},
	},
}

// Mutating a field
if address := animal.GetOwner().GetAddress(); address != nil {
	address.ZipCode = 94304
}
```

### Response objects

All fields in response structs are value types (not pointers or wrappers).

If a given field is `null`, not present, or invalid, the corresponding field
will simply be its zero value. To handle optional fields, see the `IsPresent()` method
below.

All response structs also include a special `JSON` field, containing more detailed
information about each property, which you can use like so:

```go
type Animal struct {
	Name   string `json:"name,nullable"`
	Owners int    `json:"owners"`
	Age    int    `json:"age"`
	JSON   struct {
		Name  resp.Field
		Owner resp.Field
		Age   resp.Field
	} `json:"-"`
}

var res Animal
json.Unmarshal([]byte(`{"name": null, "owners": 0}`), &res)

// Use the IsPresent() method to handle optional fields
res.Owners                  // 0
res.JSON.Owners.IsPresent() // true
res.JSON.Owners.Raw()       // "0"

res.Age                  // 0
res.JSON.Age.IsPresent() // false
res.JSON.Age.Raw()       // ""

// Use the IsExplicitNull() method to differentiate null and omitted
res.Name                       // ""
res.JSON.Name.IsPresent()      // false
res.JSON.Name.Raw()            // "null"
res.JSON.Name.IsExplicitNull() // true
```

These `.JSON` structs also include an `ExtraFields` map containing
any properties in the json response that were not specified
in the struct. This can be useful for API features not yet
present in the SDK.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### Response Unions

In responses, unions are represented by a flattened struct containing all possible fields from each of the
object variants.
To convert it to a variant use the `.AsFooVariant()` method or the `.AsAny()` method if present.

If a response value union contains primitive values, primitive fields will be alongside
the properties but prefixed with `Of` and feature the tag `json:"...,inline"`.

```go
type AnimalUnion struct {
	// From variants [Dog], [Cat]
	Owner Person `json:"owner"`
	// From variant [Dog]
	DogBreed string `json:"dog_breed"`
	// From variant [Cat]
	CatBreed string `json:"cat_breed"`
	// ...
	JSON struct {
		Owner resp.Field
		// ...
	} `json:"-"`
}

// If animal variant
if animal.Owner.Address.ZipCode == "" {
	panic("missing zip code")
}

// Switch on the variant
switch variant := animal.AsAny().(type) {
case Dog:
case Cat:
default:
	panic("unexpected type")
}
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := stainlessv0.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.Projects.Config.Commits.New(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

See the [full list of request options](https://pkg.go.dev/github.com/stainless-api/stainless-api-go/option).

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all pages:

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

### Errors

When the API returns a non-success status code, we return an error with type
`*stainlessv0.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.Projects.Config.Commits.New(
	context.TODO(),
	"projectName",
	stainlessv0.ProjectConfigCommitNewParams{
		Branch:        "branch",
		CommitMessage: "commit_message",
	},
)
if err != nil {
	var apierr *stainlessv0.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response
	}
	panic(err.Error()) // GET "/v0/projects/{projectName}/config/commits": 400 Bad Request { ... }
}
```

When other errors occur, they are returned unwrapped; for example,
if HTTP transport fails, you might receive `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a timeout for a request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not start over.
To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.Projects.Config.Commits.New(
	ctx,
	"projectName",
	stainlessv0.ProjectConfigCommitNewParams{
		Branch:        "branch",
		CommitMessage: "commit_message",
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

### File uploads

Request parameters that correspond to file uploads in multipart requests are typed as
`io.Reader`. The contents of the `io.Reader` will by default be sent as a multipart form
part with the file name of "anonymous_file" and content-type of "application/octet-stream".

The file name and content-type can be customized by implementing `Name() string` or `ContentType()
string` on the run-time type of `io.Reader`. Note that `os.File` implements `Name() string`, so a
file returned by `os.Open` will be sent with the file name on disk.

We also provide a helper `stainlessv0.File(reader io.Reader, filename string, contentType string)`
which can be used to wrap any `io.Reader` with the appropriate file name and content type.

### Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
We retry by default all connection errors, 408 Request Timeout, 409 Conflict, 429 Rate Limit,
and >=500 Internal errors.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := stainlessv0.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.Projects.Config.Commits.New(
	context.TODO(),
	"projectName",
	stainlessv0.ProjectConfigCommitNewParams{
		Branch:        "branch",
		CommitMessage: "commit_message",
	},
	option.WithMaxRetries(5),
)
```

### Accessing raw response data (e.g. response headers)

You can access the raw HTTP response data by using the `option.WithResponseInto()` request option. This is useful when
you need to examine response headers, status codes, or other details.

```go
// Create a variable to store the HTTP response
var response *http.Response
commit, err := client.Projects.Config.Commits.New(
	context.TODO(),
	"projectName",
	stainlessv0.ProjectConfigCommitNewParams{
		Branch:        "branch",
		CommitMessage: "commit_message",
	},
	option.WithResponseInto(&response),
)
if err != nil {
	// handle error
}
fmt.Printf("%+v\n", commit)

fmt.Printf("Status Code: %d\n", response.StatusCode)
fmt.Printf("Headers: %+#v\n", response.Header)
```

### Making custom/undocumented requests

This library is typed for convenient access to the documented API. If you need to access undocumented
endpoints, params, or response properties, the library can still be used.

#### Undocumented endpoints

To make requests to undocumented endpoints, you can use `client.Get`, `client.Post`, and other HTTP verbs.
`RequestOptions` on the client, such as retries, will be respected when making these requests.

```go
var (
    // params can be an io.Reader, a []byte, an encoding/json serializable object,
    // or a "…Params" struct defined in this library.
    params map[string]any

    // result can be an []byte, *http.Response, a encoding/json deserializable object,
    // or a model defined in this library.
    result *http.Response
)
err := client.Post(context.Background(), "/unspecified", params, &result)
if err != nil {
    …
}
```

#### Undocumented request params

To make requests using undocumented parameters, you may use either the `option.WithQuerySet()`
or the `option.WithJSONSet()` methods.

```go
params := FooNewParams{
    ID:   "id_xxxx",
    Data: FooNewParamsData{
        FirstName: stainlessv0.String("John"),
    },
}
client.Foo.New(context.Background(), params, option.WithJSONSet("data.last_name", "Doe"))
```

#### Undocumented response properties

To access undocumented response properties, you may either access the raw JSON of the response as a string
with `result.JSON.RawJSON()`, or get the raw JSON of a particular field on the result with
`result.JSON.Foo.Raw()`.

Any fields that are not present on the response struct will be saved and can be accessed by `result.JSON.ExtraFields()` which returns the extra fields as a `map[string]Field`.

### Middleware

We provide `option.WithMiddleware` which applies the given
middleware to requests.

```go
func Logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	// Before the request
	start := time.Now()
	LogReq(req)

	// Forward the request to the next handler
	res, err = next(req)

	// Handle stuff after the request
	end := time.Now()
	LogRes(res, err, start - end)

    return res, err
}

client := stainlessv0.NewClient(
	option.WithMiddleware(Logger),
)
```

When multiple middlewares are provided as variadic arguments, the middlewares
are applied left to right. If `option.WithMiddleware` is given
multiple times, for example first in the client then the method, the
middleware in the client will run first and the middleware given in the method
will run next.

You may also replace the default `http.Client` with
`option.WithHTTPClient(client)`. Only one http client is
accepted (this overwrites any previous client) and receives requests after any
middleware has been applied.

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/stainless-api/stainless-api-go/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
