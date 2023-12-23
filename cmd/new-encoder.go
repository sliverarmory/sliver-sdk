package cmd

import (
	"github.com/spf13/cobra"
)

// newExtensionCmd represents the newExtension command
var newEncoderCmd = &cobra.Command{
	Use:   "new-encoder",
	Short: "Create a new traffic encoder",
	Run: func(cmd *cobra.Command, args []string) {
	},
	ValidArgs: []string{"rust"},
	Args:      cobra.OnlyValidArgs,
}
