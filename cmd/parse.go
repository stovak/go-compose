package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stovak/go-compose/pkg/composer"
)

var ParseCommand = &cobra.Command{
	Use:   "parse",
	Short: "Parse composer.json",
	Long:  "Parse a composer.json file",
	Run:   runParse,
}

func init() {
	RootCmd.AddCommand(ParseCommand)
}

func runParse(cmd *cobra.Command, args []string) {
	file := args[0]
	parser := composer.New(file)
	fmt.Printf("Parsed: %+v", parser.Require)
}
