# OpenAPI

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#OpenAPIGetResponse">OpenAPIGetResponse</a>

Methods:

- <code title="get /v0/openapi">client.OpenAPI.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#OpenAPIService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#OpenAPIGetResponse">OpenAPIGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectGetResponse">ProjectGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectUpdateResponse">ProjectUpdateResponse</a>

Methods:

- <code title="get /projects/{projectName}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectGetResponse">ProjectGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v0/projects/{projectName}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectUpdateResponse">ProjectUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Branches

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectBranch">ProjectBranch</a>

Methods:

- <code title="post /v0/projects/{project}/branches">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectBranchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, project <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectBranchNewParams">ProjectBranchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/projects/{project}/branches/{branch}">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectBranchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, project <a href="https://pkg.go.dev/builtin#string">string</a>, branch <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Snippets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectSnippetNewRequestResponse">ProjectSnippetNewRequestResponse</a>

Methods:

- <code title="post /v0/projects/{projectName}/snippets/request">client.Projects.Snippets.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectSnippetService.NewRequest">NewRequest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectSnippetNewRequestParams">ProjectSnippetNewRequestParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectSnippetNewRequestResponse">ProjectSnippetNewRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Builds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildObject">BuildObject</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTarget">BuildTarget</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildListResponse">BuildListResponse</a>

Methods:

- <code title="post /v0/builds">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildNewParams">BuildNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildObject">BuildObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/builds/{buildId}">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildObject">BuildObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/builds">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildListParams">BuildListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildListResponse">BuildListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BuildTargetOutputs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetOutputListResponseUnion">BuildTargetOutputListResponseUnion</a>

Methods:

- <code title="get /v0/build_target_outputs">client.BuildTargetOutputs.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetOutputService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetOutputListParams">BuildTargetOutputListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetOutputListResponseUnion">BuildTargetOutputListResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

## Postman

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#WebhookPostmanNewNotificationResponse">WebhookPostmanNewNotificationResponse</a>

Methods:

- <code title="post /v0/webhooks/postman/notifications">client.Webhooks.Postman.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#WebhookPostmanService.NewNotification">NewNotification</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#WebhookPostmanNewNotificationParams">WebhookPostmanNewNotificationParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#WebhookPostmanNewNotificationResponse">WebhookPostmanNewNotificationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
