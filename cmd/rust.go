package cmd

import (
	"errors"
	"os"

	rustext "github.com/bishopfox/sliver-sdk/sdk/rust/extension"
	"github.com/spf13/cobra"
)

var (
	rustExtCmd = &cobra.Command{
		Use:   "rust EXTENSION_NAME",
		Short: "Create a new Rust extension",
		Long: `Create a new Rust extension package.
EXTENSION_NAME is the name of the extension.

EXTENSION_NAME can only contain alphanumeric characters and underscores.`,
		Example: "sliver-sdk new rust my-sliver-ext",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("invalid number of arguments")
			}
			extName := args[0]
			zipData, err := rustext.RenderRustTemplate(extName)
			if err != nil {
				return err
			}
			err = os.WriteFile(extName+".zip", zipData, 0644)
			if err != nil {
				return err
			}
			cmd.Printf("[*] Your extension package ready: %s.zip\n", extName)
			return nil
		},
	}
)

func NewRustExtensionCmd() *cobra.Command {
	return rustExtCmd
}
