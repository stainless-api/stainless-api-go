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

func TestProjectSnippetNewRequestWithOptionalParams(t *testing.T) {
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
		option.WithProject("example-project"),
	)
	_, err := client.Projects.Snippets.NewRequest(
		context.TODO(),
		"projectName",
		stainlessv0.ProjectSnippetNewRequestParams{
			Language: stainlessv0.ProjectSnippetNewRequestParamsLanguageNode,
			Request: stainlessv0.ProjectSnippetNewRequestParamsRequestUnion{
				OfProjectSnippetNewRequestsRequestObject: &stainlessv0.ProjectSnippetNewRequestParamsRequestObject{
					Method: "method",
					Parameters: []stainlessv0.ProjectSnippetNewRequestParamsRequestObjectParameter{{
						In:    "path",
						Name:  "name",
						Value: map[string]interface{}{},
					}},
					Path: "path",
					Body: stainlessv0.ProjectSnippetNewRequestParamsRequestObjectBody{
						FileParam: stainlessv0.String("fileParam"),
						FilePath:  stainlessv0.String("filePath"),
					},
				},
			},
			Version: stainlessv0.ProjectSnippetNewRequestParamsVersionNext,
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
