package flux

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// SetupFluxCmd defines the Flux setup command
func SetupFluxCmd() *cobra.Command {
	var gitRepo string
	var gitBranch string
	var namespace string
	var path string

	cmd := &cobra.Command{
		Use:   "flux",
		Short: "Setup Flux CD for GitOps",
		Long:  `Setup and configure Flux CD for GitOps-driven deployments.`,
		Run: func(cmd *cobra.Command, args []string) {
			setupFlux(gitRepo, gitBranch, namespace, path)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&gitRepo, "repo", "r", "", "Git repository URL (required)")
	cmd.Flags().StringVarP(&gitBranch, "branch", "b", "main", "Git branch to use")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "flux-system", "Kubernetes namespace for Flux")
	cmd.Flags().StringVarP(&path, "path", "p", "./flux", "Path to Flux manifests in the repository")
	cmd.MarkFlagRequired("repo")

	// Add subcommands
	cmd.AddCommand(syncCmd())
	cmd.AddCommand(checkCmd())

	return cmd
}

// setupFlux installs and configures Flux CD
func setupFlux(gitRepo, gitBranch, namespace, path string) {
	fmt.Println("Setting up Flux CD...")

	// Check if flux CLI is installed
	_, err := exec.LookPath("flux")
	if err != nil {
		fmt.Println("Error: Flux CLI is not installed. Please install it first.")
		fmt.Println("Installation instructions: https://fluxcd.io/docs/installation/")
		return
	}

	// Install Flux components
	fmt.Println("Installing Flux components...")
	installCmd := exec.Command("flux", "install", "--namespace", namespace)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		fmt.Println("Error installing Flux:", err)
		return
	}

	// Bootstrap Flux with the Git repository
	fmt.Printf("Bootstrapping Flux with repository %s (branch: %s)...\n", gitRepo, gitBranch)
	bootstrapCmd := exec.Command(
		"flux", "bootstrap", "git",
		"--url", gitRepo,
		"--branch", gitBranch,
		"--path", path,
		"--namespace", namespace,
	)
	bootstrapCmd.Stdout = os.Stdout
	bootstrapCmd.Stderr = os.Stderr
	if err := bootstrapCmd.Run(); err != nil {
		fmt.Println("Error bootstrapping Flux:", err)
		return
	}

	fmt.Println("Flux CD setup completed successfully!")
}

// syncCmd creates a command to manually trigger Flux synchronization
func syncCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Trigger Flux synchronization",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Triggering Flux synchronization...")
			syncCmd := exec.Command("flux", "reconcile", "source", "git", "--all")
			syncCmd.Stdout = os.Stdout
			syncCmd.Stderr = os.Stderr
			if err := syncCmd.Run(); err != nil {
				fmt.Println("Error triggering synchronization:", err)
				return
			}
			fmt.Println("Flux synchronization triggered successfully!")
		},
	}
}

// checkCmd creates a command to check Flux status
func checkCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "Check Flux status",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking Flux status...")
			checkCmd := exec.Command("flux", "check")
			checkCmd.Stdout = os.Stdout
			checkCmd.Stderr = os.Stderr
			if err := checkCmd.Run(); err != nil {
				fmt.Println("Error checking Flux status:", err)
				return
			}
		},
	}
}
