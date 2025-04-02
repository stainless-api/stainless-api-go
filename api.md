# Projects

## Config

### Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Commit">Commit</a>

Methods:

- <code title="post /v0/projects/{projectName}/config/commits">client.Projects.Config.Commits.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigCommitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigCommitNewParams">ProjectConfigCommitNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Commit">Commit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Branches

Methods:

- <code title="post /v0/projects/{projectName}/config/branches">client.Projects.Config.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigBranchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigBranchNewParams">ProjectConfigBranchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Commit">Commit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/projects/{projectName}/config/merge">client.Projects.Config.Branches.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigBranchService.Merge">Merge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigBranchMergeParams">ProjectConfigBranchMergeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Commit">Commit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Builds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Build">Build</a>
- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTarget">BuildTarget</a>

Methods:

- <code title="post /v0/builds">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildNewParams">BuildNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Build">Build</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/builds/{buildId}">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#Build">Build</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Targets

## Artifacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#TargetArtifactGetResponse">TargetArtifactGetResponse</a>

Methods:

- <code title="get /v0/builds/{buildId}/targets/{targetName}/artifacts/source">client.Targets.Artifacts.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#TargetArtifactService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>, targetName <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#TargetArtifactGetParamsTargetName">TargetArtifactGetParamsTargetName</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#TargetArtifactGetResponse">TargetArtifactGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
