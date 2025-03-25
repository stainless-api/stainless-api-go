// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/stainless-v0-go"
	"github.com/stainless-sdks/stainless-v0-go/internal/testutil"
	"github.com/stainless-sdks/stainless-v0-go/option"
)

func TestBuildTargetGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stainlessv0.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Builds.Target.Get(
		context.TODO(),
		"buildId",
		stainlessv0.BuildTargetGetParamsTargetNameNode,
	)
	if err != nil {
		var apierr *stainlessv0.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
