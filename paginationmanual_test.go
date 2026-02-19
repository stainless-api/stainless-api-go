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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Builds.List(context.TODO(), stainless.BuildListParams{
		Project: stainless.String("stainless"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, build := range page.Data {
		t.Logf("%+v\n", build.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, build := range page.Data {
			t.Logf("%+v\n", build.ID)
		}
	}
}
