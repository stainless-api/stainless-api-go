// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-api/stainless-api-go"
	"github.com/stainless-api/stainless-api-go/internal/testutil"
	"github.com/stainless-api/stainless-api-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stainless.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	buildObject, err := client.Builds.New(context.TODO(), stainless.BuildNewParams{
		Project: stainless.String("example"),
		Revision: stainless.BuildNewParamsRevisionUnion{
			OfString: stainless.String("string"),
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", buildObject.ID)
}
