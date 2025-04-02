// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-api/stainless-api-go"
	"github.com/stainless-api/stainless-api-go/internal/testutil"
	"github.com/stainless-api/stainless-api-go/option"
)

func TestProjectConfigBranchNew(t *testing.T) {
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
	_, err := client.Projects.Config.Branches.New(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectConfigBranchNewParams{
			Branch:     stainlessv0.F("branch"),
			BranchFrom: stainlessv0.F("branch_from"),
		},
	)
	if err != nil {
		var apierr *stainlessv0.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectConfigBranchMerge(t *testing.T) {
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
	_, err := client.Projects.Config.Branches.Merge(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectConfigBranchMergeParams{
			From: stainlessv0.F("from"),
			Into: stainlessv0.F("into"),
		},
	)
	if err != nil {
		var apierr *stainlessv0.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
