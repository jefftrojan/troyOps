package ci

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// SetupCICDCmd configures CI/CD pipelines
func SetupCICDCmd() *cobra.Command {
	var platform string
	var repoPath string
	var appName string

	cmd := &cobra.Command{
		Use:   "cicd",
		Short: "Setup CI/CD pipeline (GitHub Actions/GitLab CI)",
		Long:  `Configure CI/CD pipelines using GitHub Actions or GitLab CI.`,
		Run: func(cmd *cobra.Command, args []string) {
			setupCICD(platform, repoPath, appName)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&platform, "platform", "p", "github", "CI/CD platform to use (github, gitlab)")
	cmd.Flags().StringVarP(&repoPath, "repo-path", "r", ".", "Path to the Git repository")
	cmd.Flags().StringVarP(&appName, "app-name", "a", "app", "Name of the application")

	return cmd
}

// setupCICD configures the CI/CD pipeline based on the platform
func setupCICD(platform, repoPath, appName string) {
	fmt.Printf("Setting up CI/CD pipeline for %s on %s platform...\n", appName, platform)

	switch platform {
	case "github":
		setupGitHubActions(repoPath, appName)
	case "gitlab":
		setupGitLabCI(repoPath, appName)
	default:
		fmt.Printf("Unsupported CI/CD platform: %s\n", platform)
	}
}

// setupGitHubActions configures GitHub Actions workflows
func setupGitHubActions(repoPath, appName string) {
	fmt.Println("Creating GitHub Actions workflow...")

	// Create .github/workflows directory if it doesn't exist
	workflowsDir := filepath.Join(repoPath, ".github", "workflows")
	if err := os.MkdirAll(workflowsDir, 0755); err != nil {
		fmt.Println("Error creating workflows directory:", err)
		return
	}

	// Create main workflow file
	workflowFile := filepath.Join(workflowsDir, fmt.Sprintf("%s-ci.yml", appName))
	workflowContent := fmt.Sprintf(`name: %s CI/CD

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 0

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Login to Docker Hub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_HUB_USERNAME }}
          password: ${{ secrets.DOCKER_HUB_TOKEN }}

      - name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          push: ${{ github.event_name != 'pull_request' }}
          tags: ${{ secrets.DOCKER_HUB_USERNAME }}/%s:latest

  deploy:
    needs: build
    if: github.event_name != 'pull_request'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 0

      - name: Setup Flux
        uses: fluxcd/flux2/action@main

      - name: Update Kubernetes manifests
        run: |
          cd ./kustomize/overlays/dev
          kustomize edit set image ${{ secrets.DOCKER_HUB_USERNAME }}/%s:latest
          git config --global user.name "Flux CD"
          git config --global user.email "flux@example.com"
          git add .
          git commit -m "Update %s image tag"
          git push
`, appName, appName, appName, appName)

	if err := os.WriteFile(workflowFile, []byte(workflowContent), 0644); err != nil {
		fmt.Println("Error creating workflow file:", err)
		return
	}

	fmt.Printf("GitHub Actions workflow created at %s\n", workflowFile)
	fmt.Println("Note: You need to set DOCKER_HUB_USERNAME and DOCKER_HUB_TOKEN secrets in your GitHub repository.")
}

// setupGitLabCI configures GitLab CI pipeline
func setupGitLabCI(repoPath, appName string) {
	fmt.Println("Creating GitLab CI pipeline...")

	// Create .gitlab-ci.yml file
	ciFile := filepath.Join(repoPath, ".gitlab-ci.yml")
	ciContent := fmt.Sprintf(`stages:
  - build
  - deploy

variables:
  DOCKER_DRIVER: overlay2
  DOCKER_TLS_CERTDIR: ""

build:
  stage: build
  image: docker:20.10.16
  services:
    - docker:20.10.16-dind
  before_script:
    - docker login -u $DOCKER_HUB_USERNAME -p $DOCKER_HUB_TOKEN
  script:
    - docker build -t $DOCKER_HUB_USERNAME/%s:latest .
    - docker push $DOCKER_HUB_USERNAME/%s:latest
  only:
    - main

deploy:
  stage: deploy
  image: 
    name: fluxcd/flux:latest
    entrypoint: [""]
  before_script:
    - apt-get update && apt-get install -y git curl
    - curl -s https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh | bash
    - mv kustomize /usr/local/bin/
  script:
    - cd ./kustomize/overlays/dev
    - kustomize edit set image $DOCKER_HUB_USERNAME/%s:latest
    - git config --global user.name "Flux CD"
    - git config --global user.email "flux@example.com"
    - git add .
    - git commit -m "Update %s image tag"
    - git push
  only:
    - main
`, appName, appName, appName, appName)

	if err := os.WriteFile(ciFile, []byte(ciContent), 0644); err != nil {
		fmt.Println("Error creating GitLab CI file:", err)
		return
	}

	fmt.Printf("GitLab CI pipeline created at %s\n", ciFile)
	fmt.Println("Note: You need to set DOCKER_HUB_USERNAME and DOCKER_HUB_TOKEN variables in your GitLab CI/CD settings.")
}
