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

func TestBuildNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Builds.New(context.TODO(), stainlessv0.BuildNewParams{
		Project: "project",
		Revision: stainlessv0.BuildNewParamsRevisionUnion{
			OfString: stainlessv0.String("string"),
		},
		AllowEmpty:    stainlessv0.Bool(true),
		Branch:        stainlessv0.String("branch"),
		CommitMessage: stainlessv0.String("commit_message"),
		Targets:       []string{"node"},
	})
	if err != nil {
		var apierr *stainlessv0.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBuildGet(t *testing.T) {
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
	_, err := client.Builds.Get(context.TODO(), "buildId")
	if err != nil {
		var apierr *stainlessv0.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBuildListWithOptionalParams(t *testing.T) {
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
	_, err := client.Builds.List(context.TODO(), stainlessv0.BuildListParams{
		Project: "project",
		Branch:  stainlessv0.String("branch"),
		Cursor:  stainlessv0.String("cursor"),
		Limit:   stainlessv0.Float(1),
		Revision: stainlessv0.BuildListParamsRevisionUnion{
			OfString: stainlessv0.String("string"),
		},
	})
	if err != nil {
		var apierr *stainlessv0.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBuildCompareWithOptionalParams(t *testing.T) {
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
	_, err := client.Builds.Compare(context.TODO(), stainlessv0.BuildCompareParams{
		Base: stainlessv0.BuildCompareParamsBase{
			Revision: stainlessv0.BuildCompareParamsBaseRevisionUnion{
				OfString: stainlessv0.String("string"),
			},
			Branch:        stainlessv0.String("branch"),
			CommitMessage: stainlessv0.String("commit_message"),
		},
		Head: stainlessv0.BuildCompareParamsHead{
			Revision: stainlessv0.BuildCompareParamsHeadRevisionUnion{
				OfString: stainlessv0.String("string"),
			},
			Branch:        stainlessv0.String("branch"),
			CommitMessage: stainlessv0.String("commit_message"),
		},
		Project: "project",
		Targets: []string{"node"},
	})
	if err != nil {
		var apierr *stainlessv0.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
