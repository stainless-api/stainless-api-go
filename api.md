# Projects

## Config

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ConfigCommit">ConfigCommit</a>

Methods:

- <code title="post /v0/projects/{projectName}/config/branches">client.Projects.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectConfigService.NewBranch">NewBranch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectConfigNewBranchParams">ProjectConfigNewBranchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ConfigCommit">ConfigCommit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/projects/{projectName}/config/commits">client.Projects.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectConfigService.NewCommit">NewCommit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectConfigNewCommitParams">ProjectConfigNewCommitParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ConfigCommit">ConfigCommit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/projects/{projectName}/config/merge">client.Projects.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectConfigService.Merge">Merge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ProjectConfigMergeParams">ProjectConfigMergeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#ConfigCommit">ConfigCommit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Builds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildGetResponse">BuildGetResponse</a>

Methods:

- <code title="get /v0/builds/{buildId}">client.Builds.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildGetResponse">BuildGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Target

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTarget">BuildTarget</a>

Methods:

- <code title="get /v0/builds/{buildId}/target/{targetName}">client.Builds.Target.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>, targetName <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetGetParamsTargetName">BuildTargetGetParamsTargetName</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTarget">BuildTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Artifacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetArtifactGetSourceResponse">BuildTargetArtifactGetSourceResponse</a>

Methods:

- <code title="get /v0/builds/{buildId}/target/{targetName}/artifacts/source">client.Builds.Target.Artifacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetArtifactService.GetSource">GetSource</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, buildID <a href="https://pkg.go.dev/builtin#string">string</a>, targetName <a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetArtifactGetSourceParamsTargetName">BuildTargetArtifactGetSourceParamsTargetName</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go">stainlessv0</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stainless-v0-go#BuildTargetArtifactGetSourceResponse">BuildTargetArtifactGetSourceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
