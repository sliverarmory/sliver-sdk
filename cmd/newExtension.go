/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/bishopfox/sliver-sdk/cmd/goext"
	"github.com/spf13/cobra"
)

// newExtensionCmd represents the newExtension command
var newExtensionCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new extension",
	Run: func(cmd *cobra.Command, args []string) {
	},
	ValidArgs: []string{"go", "rust", "nim", "c"},
	Args:      cobra.OnlyValidArgs,
}

func init() {
	newExtensionCmd.AddCommand(goext.NewGoExtensionCmd())
	rootCmd.AddCommand(newExtensionCmd)
}
