// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainless_test

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
	t.Skip("Prism tests are disabled")
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
	_, err := client.Builds.New(context.TODO(), stainless.BuildNewParams{
		Project: stainless.String("project"),
		Revision: stainless.BuildNewParamsRevisionUnion{
			OfString: stainless.String("string"),
		},
		AllowEmpty:    stainless.Bool(true),
		Branch:        stainless.String("branch"),
		CommitMessage: stainless.String("commit_message"),
		Targets:       []stainless.Target{stainless.TargetNode},
	})
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBuildGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Builds.Get(context.TODO(), "buildId")
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBuildListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Builds.List(context.TODO(), stainless.BuildListParams{
		Project: stainless.String("project"),
		Branch:  stainless.String("branch"),
		Cursor:  stainless.String("cursor"),
		Limit:   stainless.Float(1),
		Revision: stainless.BuildListParamsRevisionUnion{
			OfString: stainless.String("string"),
		},
	})
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBuildCompareWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Builds.Compare(context.TODO(), stainless.BuildCompareParams{
		Base: stainless.BuildCompareParamsBase{
			Revision: stainless.BuildCompareParamsBaseRevisionUnion{
				OfString: stainless.String("string"),
			},
			Branch:        stainless.String("branch"),
			CommitMessage: stainless.String("commit_message"),
		},
		Head: stainless.BuildCompareParamsHead{
			Revision: stainless.BuildCompareParamsHeadRevisionUnion{
				OfString: stainless.String("string"),
			},
			Branch:        stainless.String("branch"),
			CommitMessage: stainless.String("commit_message"),
		},
		Project: stainless.String("project"),
		Targets: []stainless.Target{stainless.TargetNode},
	})
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
