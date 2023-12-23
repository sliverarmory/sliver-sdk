/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sliver-sdk",
	Short: "Small utility to bootstrap Sliver extensions writing",
}

func init() {
	newExtensionCmd.AddCommand(NewGoExtensionCmd())
	newExtensionCmd.AddCommand(NewRustExtensionCmd())
	rootCmd.AddCommand(newExtensionCmd)

	newEncoderCmd.AddCommand(NewRustEncoderCmd())
	rootCmd.AddCommand(newEncoderCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
