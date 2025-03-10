package secrets

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// ConfigureSecretsCmd defines the command for configuring secret management
func ConfigureSecretsCmd() *cobra.Command {
	var secretEngine string
	var secretsDir string

	cmd := &cobra.Command{
		Use:   "secrets",
		Short: "Configure secret management (SOPS/Sealed Secrets)",
		Long:  `Configure and apply secret management using SOPS or Sealed Secrets.`,
		Run: func(cmd *cobra.Command, args []string) {
			configureSecretManagement(secretEngine, secretsDir)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&secretEngine, "engine", "e", "sops", "Secret management engine to use (sops, sealed-secrets)")
	cmd.Flags().StringVarP(&secretsDir, "directory", "d", "", "Directory containing secret definitions (defaults to secrets/{engine})")

	return cmd
}

// configureSecretManagement sets up the specified secret management solution
func configureSecretManagement(engine, secretsDir string) {
	fmt.Printf("Setting up %s secret management...\n", engine)

	// If no secrets directory is specified, use the default
	if secretsDir == "" {
		secretsDir = filepath.Join("secrets", engine)
	}

	// Check if the secrets directory exists
	if _, err := os.Stat(secretsDir); os.IsNotExist(err) {
		fmt.Printf("Error: Secrets directory does not exist: %s\n", secretsDir)
		return
	}

	switch engine {
	case "sops":
		setupSOPS(secretsDir)
	case "sealed-secrets":
		setupSealedSecrets(secretsDir)
	default:
		fmt.Printf("Unsupported secret management engine: %s\n", engine)
	}
}

// setupSOPS configures SOPS for secret management
func setupSOPS(secretsDir string) {
	fmt.Println("Setting up SOPS for secret management...")

	// Check if SOPS is installed
	_, err := exec.LookPath("sops")
	if err != nil {
		fmt.Println("Error: SOPS is not installed. Please install SOPS first.")
		fmt.Println("Installation instructions: https://github.com/mozilla/sops#installation")
		return
	}

	// Create a .sops.yaml configuration file if it doesn't exist
	sopsConfigPath := ".sops.yaml"
	if _, err := os.Stat(sopsConfigPath); os.IsNotExist(err) {
		fmt.Println("Creating SOPS configuration file...")
		sopsConfig := `
creation_rules:
  - path_regex: secrets/.*\.yaml
    encrypted_regex: ^(data|stringData)$
    pgp: <YOUR_PGP_KEY_FINGERPRINT>
`
		err := os.WriteFile(sopsConfigPath, []byte(sopsConfig), 0644)
		if err != nil {
			fmt.Println("Error creating SOPS configuration:", err)
			return
		}
		fmt.Println("Created SOPS configuration file. Please update it with your encryption keys.")
	}

	// Apply encrypted secrets from the secrets directory
	fmt.Printf("Applying encrypted secrets from %s...\n", secretsDir)
	filepath.Walk(secretsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			fmt.Printf("Applying secret: %s\n", path)
			cmd := exec.Command("sops", "--decrypt", path, "|", "kubectl", "apply", "-f", "-")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error applying secret %s: %v\n", path, err)
			}
		}
		return nil
	})

	fmt.Println("SOPS setup completed successfully!")
}

// setupSealedSecrets installs and configures Sealed Secrets
func setupSealedSecrets(secretsDir string) {
	fmt.Println("Setting up Sealed Secrets for secret management...")

	// Install Sealed Secrets controller using Helm
	fmt.Println("Installing Sealed Secrets controller...")
	installCmd := exec.Command("helm", "repo", "add", "sealed-secrets", "https://bitnami-labs.github.io/sealed-secrets")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		fmt.Println("Error adding Sealed Secrets Helm repo:", err)
		return
	}

	updateCmd := exec.Command("helm", "repo", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		fmt.Println("Error updating Helm repos:", err)
		return
	}

	helmCmd := exec.Command("helm", "install", "sealed-secrets", "sealed-secrets/sealed-secrets", "--namespace", "kube-system")
	helmCmd.Stdout = os.Stdout
	helmCmd.Stderr = os.Stderr
	if err := helmCmd.Run(); err != nil {
		fmt.Println("Error installing Sealed Secrets controller:", err)
		return
	}

	// Apply sealed secrets from the secrets directory
	fmt.Printf("Applying sealed secrets from %s...\n", secretsDir)
	applyCmd := exec.Command("kubectl", "apply", "-f", secretsDir)
	applyCmd.Stdout = os.Stdout
	applyCmd.Stderr = os.Stderr
	if err := applyCmd.Run(); err != nil {
		fmt.Println("Error applying sealed secrets:", err)
		return
	}

	fmt.Println("Sealed Secrets setup completed successfully!")
}
