# Shared Params Types

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared#FileInputUnionParam">FileInputUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared#Target">Target</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared#Commit">Commit</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/shared#Target">Target</a>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Project">Project</a>

Methods:

- <code title="post /v0/projects">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectNewParams">ProjectNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/projects/{project}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectGetParams">ProjectGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v0/projects/{project}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectUpdateParams">ProjectUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/projects">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectListParams">ProjectListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Project">Project</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Branches

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchListResponse">ProjectBranchListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchDeleteResponse">ProjectBranchDeleteResponse</a>

Methods:

- <code title="post /v0/projects/{project}/branches">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchNewParams">ProjectBranchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/projects/{project}/branches/{branch}">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, branch <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchGetParams">ProjectBranchGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/projects/{project}/branches">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchListParams">ProjectBranchListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchListResponse">ProjectBranchListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v0/projects/{project}/branches/{branch}">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, branch <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchDeleteParams">ProjectBranchDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchDeleteResponse">ProjectBranchDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v0/projects/{project}/branches/{branch}/rebase">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.Rebase">Rebase</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, branch <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchRebaseParams">ProjectBranchRebaseParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v0/projects/{project}/branches/{branch}/reset">client.Projects.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchService.Reset">Reset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, branch <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranchResetParams">ProjectBranchResetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectBranch">ProjectBranch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Configs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGetResponse">ProjectConfigGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGuessResponse">ProjectConfigGuessResponse</a>

Methods:

- <code title="get /v0/projects/{project}/configs">client.Projects.Configs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGetParams">ProjectConfigGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGetResponse">ProjectConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/projects/{project}/configs/guess">client.Projects.Configs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigService.Guess">Guess</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGuessParams">ProjectConfigGuessParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigGuessResponse">ProjectConfigGuessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Builds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Build">Build</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTarget">BuildTarget</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#CheckStepUnion">CheckStepUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildCompareResponse">BuildCompareResponse</a>

Methods:

- <code title="post /v0/builds">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildNewParams">BuildNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Build">Build</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/builds/{buildId}">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Build">Build</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/builds">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildListParams">BuildListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Build">Build</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/builds/compare">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.Compare">Compare</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildCompareParams">BuildCompareParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildCompareResponse">BuildCompareResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Diagnostics

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildDiagnostic">BuildDiagnostic</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildDiagnosticMoreUnion">BuildDiagnosticMoreUnion</a>

Methods:

- <code title="get /v0/builds/{buildId}/diagnostics">client.Builds.Diagnostics.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildDiagnosticService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildDiagnosticListParams">BuildDiagnosticListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go/packages/pagination#Page">Page</a>[<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildDiagnostic">BuildDiagnostic</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TargetOutputs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputGetResponseUnion">BuildTargetOutputGetResponseUnion</a>

Methods:

- <code title="get /v0/build_target_outputs">client.Builds.TargetOutputs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputGetParams">BuildTargetOutputGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetOutputGetResponseUnion">BuildTargetOutputGetResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Org">Org</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#OrgListResponse">OrgListResponse</a>

Methods:

- <code title="get /v0/orgs/{org}">client.Orgs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, org <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Org">Org</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/orgs">client.Orgs.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#OrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainless</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#OrgListResponse">OrgListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
