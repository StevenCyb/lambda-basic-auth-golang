# Lambda Basic Auth with Golang
After a long searching for a basic auth example for Lambda with Golang, I came across an [interesting repository](https://github.com/serverless/examples/tree/master/aws-golang-auth-examples).
Since it did not quite suit my requirements, I also derived a solution for my needs.

## What does this project contain
This example contains a basic auth handler and a greeting page. 
The username and password are set in the environment variables of `serverless.yaml`.
**This should not be done in production!**

Both Hendlers have simple tests that use the `ginkgo` framework.
These are not perfected, but show how they can be used.

In addition, there is a `Makefile` for basic operations:
- clean -> Delete old (unneeded) data
- test -> Run tests on all handlers
- build -> Build all handlers to deploy
- deploy-dev -> Deploy the APIs on `dev` stage
- deploy-live -> Deploy the APIs on `live` stage
- undeploy -> Undeploy the API on `dev` stage

I have also used the CodeCommit, CodeBuild and CodePipeline.
For this I used the `buildspec.yaml`.