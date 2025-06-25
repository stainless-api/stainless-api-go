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

func TestProjectNew(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Projects.New(context.TODO(), stainless.ProjectNewParams{
		DisplayName: "display_name",
		Org:         "org",
		Revision: map[string]stainless.ProjectNewParamsRevisionUnion{
			"foo": {
				OfProjectNewsRevisionContent: &stainless.ProjectNewParamsRevisionContent{
					Content: "content",
				},
			},
		},
		Slug:    "slug",
		Targets: []string{"node"},
	})
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Projects.Get(context.TODO(), stainless.ProjectGetParams{
		Project: stainless.String("project"),
	})
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Projects.Update(context.TODO(), stainless.ProjectUpdateParams{
		Project:     stainless.String("project"),
		DisplayName: stainless.String("display_name"),
	})
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Projects.List(context.TODO(), stainless.ProjectListParams{
		Cursor: stainless.String("cursor"),
		Limit:  stainless.Float(1),
		Org:    stainless.String("org"),
	})
	if err != nil {
		var apierr *stainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
