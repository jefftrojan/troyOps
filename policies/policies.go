package policies

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// SetupPoliciesCmd defines the command for setting up policy enforcement
func SetupPoliciesCmd() *cobra.Command {
	var policyEngine string
	var policyDir string

	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Setup policy enforcement (Kyverno/OPA)",
		Long:  `Configure and apply policy enforcement using Kyverno or OPA/Gatekeeper.`,
		Run: func(cmd *cobra.Command, args []string) {
			setupPolicyEnforcement(policyEngine, policyDir)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&policyEngine, "engine", "e", "kyverno", "Policy engine to use (kyverno, opa)")
	cmd.Flags().StringVarP(&policyDir, "directory", "d", "", "Directory containing policy definitions (defaults to policies/{engine})")

	return cmd
}

// setupPolicyEnforcement configures and applies policy enforcement
func setupPolicyEnforcement(engine, policyDir string) {
	fmt.Printf("Setting up %s policy enforcement...\n", engine)

	// If no policy directory is specified, use the default
	if policyDir == "" {
		policyDir = filepath.Join("policies", engine)
	}

	// Check if the policy directory exists
	if _, err := os.Stat(policyDir); os.IsNotExist(err) {
		fmt.Printf("Error: Policy directory does not exist: %s\n", policyDir)
		return
	}

	switch engine {
	case "kyverno":
		setupKyverno(policyDir)
	case "opa":
		setupOPA(policyDir)
	default:
		fmt.Printf("Unsupported policy engine: %s\n", engine)
	}
}

// setupKyverno installs and configures Kyverno
func setupKyverno(policyDir string) {
	// Install Kyverno using Helm
	fmt.Println("Installing Kyverno...")
	installCmd := exec.Command("helm", "repo", "add", "kyverno", "https://kyverno.github.io/kyverno/")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		fmt.Println("Error adding Kyverno Helm repo:", err)
		return
	}

	updateCmd := exec.Command("helm", "repo", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		fmt.Println("Error updating Helm repos:", err)
		return
	}

	helmCmd := exec.Command("helm", "install", "kyverno", "kyverno/kyverno", "--namespace", "kyverno", "--create-namespace")
	helmCmd.Stdout = os.Stdout
	helmCmd.Stderr = os.Stderr
	if err := helmCmd.Run(); err != nil {
		fmt.Println("Error installing Kyverno:", err)
		return
	}

	// Apply policies from the policy directory
	fmt.Printf("Applying Kyverno policies from %s...\n", policyDir)
	applyCmd := exec.Command("kubectl", "apply", "-f", policyDir)
	applyCmd.Stdout = os.Stdout
	applyCmd.Stderr = os.Stderr
	if err := applyCmd.Run(); err != nil {
		fmt.Println("Error applying Kyverno policies:", err)
		return
	}

	fmt.Println("Kyverno setup completed successfully!")
}

// setupOPA installs and configures OPA/Gatekeeper
func setupOPA(policyDir string) {
	// Install OPA/Gatekeeper using Helm
	fmt.Println("Installing OPA/Gatekeeper...")
	installCmd := exec.Command("helm", "repo", "add", "gatekeeper", "https://open-policy-agent.github.io/gatekeeper/charts")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		fmt.Println("Error adding Gatekeeper Helm repo:", err)
		return
	}

	updateCmd := exec.Command("helm", "repo", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		fmt.Println("Error updating Helm repos:", err)
		return
	}

	helmCmd := exec.Command("helm", "install", "gatekeeper", "gatekeeper/gatekeeper", "--namespace", "gatekeeper-system", "--create-namespace")
	helmCmd.Stdout = os.Stdout
	helmCmd.Stderr = os.Stderr
	if err := helmCmd.Run(); err != nil {
		fmt.Println("Error installing OPA/Gatekeeper:", err)
		return
	}

	// Apply policies from the policy directory
	fmt.Printf("Applying OPA/Gatekeeper policies from %s...\n", policyDir)
	applyCmd := exec.Command("kubectl", "apply", "-f", policyDir)
	applyCmd.Stdout = os.Stdout
	applyCmd.Stderr = os.Stderr
	if err := applyCmd.Run(); err != nil {
		fmt.Println("Error applying OPA/Gatekeeper policies:", err)
		return
	}

	fmt.Println("OPA/Gatekeeper setup completed successfully!")
}
