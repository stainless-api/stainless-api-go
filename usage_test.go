// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/stainless-v0-go"
	"github.com/stainless-sdks/stainless-v0-go/internal/testutil"
	"github.com/stainless-sdks/stainless-v0-go/option"
)

func TestUsage(t *testing.T) {
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
	configCommit, err := client.Projects.Config.NewBranch(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectConfigNewBranchParams{
			Branch:     stainlessv0.F("branch"),
			BranchFrom: stainlessv0.F("branch_from"),
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", configCommit.ID)
}
