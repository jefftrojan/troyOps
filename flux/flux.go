package flux

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// SetupFluxCmd defines the Flux setup command
func SetupFluxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "flux",
		Short: "Setup Flux CD for GitOps",
		Run: func(cmd *cobra.Command, args []string) {
			// exsmple of Flux installation using `flux bootstrap`
			fmt.Println("Setting up Flux CD...")

			// command to install Flux
			installCmd := exec.Command("flux", "install")
			if err := installCmd.Run(); err != nil {
				fmt.Println("Error installing Flux:", err)
			} else {
				fmt.Println("Flux CD setup complete!")
			}
		},
	}
}
