# [FEATURE] Improve Flux CD integration with multi-cluster support

## Is your feature request related to a problem? Please describe.
Currently, TroyOps only supports deploying to a single Kubernetes cluster with Flux CD. In enterprise environments, it's common to have multiple clusters for different purposes (e.g., development, staging, production) or different regions.

## Describe the solution you'd like
Enhance the Flux CD integration to support managing multiple Kubernetes clusters from a single Git repository. This should include:

1. The ability to specify different clusters in the Flux configuration
2. Support for cluster-specific configurations and overlays
3. A way to target specific clusters when running commands
4. Documentation on multi-cluster GitOps best practices

## Describe alternatives you've considered
- Using separate Git repositories for each cluster
- Using Git branches for different clusters
- Using Argo CD instead of Flux CD

## Additional context
This feature would make TroyOps more suitable for enterprise environments where multi-cluster deployments are common.
