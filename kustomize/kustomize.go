package kustomize

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// DeployManifestsCmd defines the command for deploying Kubernetes manifests using Kustomize
func DeployManifestsCmd() *cobra.Command {
	var environment string
	var namespace string

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy Kubernetes manifests using Kustomize",
		Long:  `Deploy Kubernetes manifests to a cluster using Kustomize overlays for different environments.`,
		Run: func(cmd *cobra.Command, args []string) {
			deployWithKustomize(environment, namespace)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&environment, "environment", "e", "dev", "Environment to deploy (dev, staging, prod)")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Kubernetes namespace to deploy to")

	return cmd
}

// deployWithKustomize applies Kustomize overlays for the specified environment
func deployWithKustomize(environment, namespace string) {
	fmt.Printf("Deploying to %s environment in namespace %s...\n", environment, namespace)

	// Determine the overlay path
	overlayPath := filepath.Join("kustomize", "overlays", environment)

	// Check if the overlay exists
	if _, err := os.Stat(overlayPath); os.IsNotExist(err) {
		fmt.Printf("Error: Overlay for environment '%s' does not exist at path: %s\n", environment, overlayPath)
		return
	}

	// Run kubectl apply with kustomize
	cmd := exec.Command("kubectl", "apply", "-k", overlayPath, "-n", namespace)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Executing:", cmd.String())
	if err := cmd.Run(); err != nil {
		fmt.Println("Error deploying manifests:", err)
		return
	}

	fmt.Println("Deployment completed successfully!")
}
