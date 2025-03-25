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

func TestProjectConfigNewBranch(t *testing.T) {
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
	_, err := client.Projects.Config.NewBranch(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectConfigNewBranchParams{
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

func TestProjectConfigNewCommitWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Config.NewCommit(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectConfigNewCommitParams{
			Branch:          stainlessv0.F("branch"),
			CommitMessage:   stainlessv0.F("commit_message"),
			OpenAPISpec:     stainlessv0.F("openapi_spec"),
			StainlessConfig: stainlessv0.F("stainless_config"),
			AllowEmpty:      stainlessv0.F(true),
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

func TestProjectConfigMerge(t *testing.T) {
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
	_, err := client.Projects.Config.Merge(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectConfigMergeParams{
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
