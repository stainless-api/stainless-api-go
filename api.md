# Projects

## Config

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ConfigCommit">ConfigCommit</a>

Methods:

- <code title="post /v0/projects/{projectName}/config/branches">client.Projects.Config.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigService.NewBranch">NewBranch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigNewBranchParams">ProjectConfigNewBranchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ConfigCommit">ConfigCommit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/projects/{projectName}/config/commits">client.Projects.Config.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigService.NewCommit">NewCommit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigNewCommitParams">ProjectConfigNewCommitParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ConfigCommit">ConfigCommit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/projects/{projectName}/config/merge">client.Projects.Config.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigService.Merge">Merge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ProjectConfigMergeParams">ProjectConfigMergeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#ConfigCommit">ConfigCommit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Builds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildGetResponse">BuildGetResponse</a>

Methods:

- <code title="get /v0/builds/{buildId}">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildGetResponse">BuildGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Target

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTarget">BuildTarget</a>

Methods:

- <code title="get /v0/builds/{buildId}/target/{targetName}">client.Builds.Target.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>, targetName <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetGetParamsTargetName">BuildTargetGetParamsTargetName</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTarget">BuildTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Artifacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetArtifactGetSourceResponse">BuildTargetArtifactGetSourceResponse</a>

Methods:

- <code title="get /v0/builds/{buildId}/target/{targetName}/artifacts/source">client.Builds.Target.Artifacts.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetArtifactService.GetSource">GetSource</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>, targetName <a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetArtifactGetSourceParamsTargetName">BuildTargetArtifactGetSourceParamsTargetName</a>) (<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-api/stainless-api-go#BuildTargetArtifactGetSourceResponse">BuildTargetArtifactGetSourceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
