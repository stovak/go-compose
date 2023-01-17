package cmd

import (
	"github.com/spf13/cobra"
)

var UpgradeCommand = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade composer.json",
	Long:  "Upgrade a composer.json file",
	Run:   runUpgrade,
}

func init() {
	RootCmd.AddCommand(UpgradeCommand)
}

func runUpgrade(cmd *cobra.Command, args []string) {

}
