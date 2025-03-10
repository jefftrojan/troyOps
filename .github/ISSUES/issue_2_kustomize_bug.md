# [BUG] Kustomize deployment fails with large manifests

## Describe the bug
When deploying large Kubernetes manifests using the Kustomize functionality, the deployment fails with a timeout error. This happens consistently with manifests larger than approximately 1MB in size.

## To Reproduce
Steps to reproduce the behavior:
1. Create a large Kubernetes manifest (e.g., a Deployment with many replicas and environment variables)
2. Run 'troyops deploy -e dev'
3. Observe the timeout error after approximately 30 seconds

## Expected behavior
The deployment should complete successfully regardless of the manifest size, with appropriate timeout handling for larger manifests.

## Screenshots or Logs


## Environment
 - OS: Ubuntu 22.04
 - Kubernetes Version: 1.26.0
 - TroyOps Version: 0.1.0
 - Go Version: 1.24

## Additional context
This issue only occurs with manifests larger than approximately 1MB. Smaller manifests deploy successfully.

## Expected behavior
The deployment should complete successfully regardless of the manifest size, with appropriate timeout handling for larger manifests.

## Screenshots or Logs
```
Error: timeout waiting for deployment: context deadline exceeded
```

## Environment
 - OS: Ubuntu 22.04
 - Kubernetes Version: 1.26.0
 - TroyOps Version: 0.1.0
 - Go Version: 1.24

## Additional context
This issue only occurs with manifests larger than approximately 1MB. Smaller manifests deploy successfully.
