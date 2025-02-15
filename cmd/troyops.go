package main

import (
	"fmt"
	"os"

	"github.com/jefftrojan/troyops/ci"
	"github.com/jefftrojan/troyops/flux"
	"github.com/jefftrojan/troyops/kustomize"
	"github.com/jefftrojan/troyops/policies"
	"github.com/jefftrojan/troyops/secrets"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "troyops",
		Short: "TroyOps - GitOps-driven Kubernetes deployment tool",
		Long:  `An open-source GitOps-driven Kubernetes deployment tool for automating your Kubernetes lifecycle.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to TroyOps! Use 'troyops --help' for options.")
		},
	}

	// Add subcommands for different functionalities
	rootCmd.AddCommand(flux.SetupFluxCmd())
	rootCmd.AddCommand(ci.SetupCICDCmd())
	rootCmd.AddCommand(kustomize.DeployManifestsCmd())
	rootCmd.AddCommand(secrets.ConfigureSecretsCmd())
	rootCmd.AddCommand(policies.SetupPoliciesCmd())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing TroyOps:", err)
		os.Exit(1)
	}
}
