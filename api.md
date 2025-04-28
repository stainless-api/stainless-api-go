# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectGetResponse">ProjectGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectUpdateResponse">ProjectUpdateResponse</a>

Methods:

- <code title="get /v0/projects/{projectName}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectGetResponse">ProjectGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v0/projects/{projectName}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectUpdateResponse">ProjectUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Branches

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>

Methods:

- <code title="post /v0/projects/{project}/branches">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, project <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchNewParams">ProjectBranchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/projects/{project}/branches/{branch}">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, project <a href="https://pkg.go.dev/builtin#string">string</a>, branch <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Configs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGetResponse">ProjectConfigGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGuessResponse">ProjectConfigGuessResponse</a>

Methods:

- <code title="get /v0/projects/{project}/configs">client.Projects.Configs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, project <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGetParams">ProjectConfigGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGetResponse">ProjectConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/projects/{project}/configs/guess">client.Projects.Configs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigService.Guess">Guess</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, project <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGuessParams">ProjectConfigGuessParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGuessResponse">ProjectConfigGuessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Builds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildObject">BuildObject</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTarget">BuildTarget</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildListResponse">BuildListResponse</a>

Methods:

- <code title="post /v0/builds">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildNewParams">BuildNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildObject">BuildObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/builds/{buildId}">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildObject">BuildObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/builds">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildListParams">BuildListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildListResponse">BuildListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BuildTargetOutputs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputGetResponseUnion">BuildTargetOutputGetResponseUnion</a>

Methods:

- <code title="get /v0/build_target_outputs">client.BuildTargetOutputs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputGetParams">BuildTargetOutputGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputGetResponseUnion">BuildTargetOutputGetResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
