package cmd

/*
	Sliver Implant Framework
	Copyright (C) 2023  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"errors"
	"os"

	rustenc "github.com/sliverarmory/sliver-sdk/sdk/rust/encoder"
	rustext "github.com/sliverarmory/sliver-sdk/sdk/rust/extension"
	"github.com/spf13/cobra"
)

var (
	rustExtCmd = &cobra.Command{
		Use:   "rust EXTENSION_NAME",
		Short: "Create a new Rust extension",
		Long: `Create a new Rust extension package.
EXTENSION_NAME is the name of the extension.

EXTENSION_NAME can only contain alphanumeric characters and underscores.`,
		Example: "sliver-sdk new-extension rust my-sliver-ext",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("invalid number of arguments")
			}
			extName := args[0]
			zipData, err := rustext.RenderRustTemplate(extName)
			if err != nil {
				return err
			}
			err = os.WriteFile(extName+"_extension.zip", zipData, 0644)
			if err != nil {
				return err
			}
			cmd.Printf("[*] Your extension package ready: %s_extension.zip\n", extName)
			return nil
		},
	}
)

func NewRustExtensionCmd() *cobra.Command {
	return rustExtCmd
}

var (
	rustEncCmd = &cobra.Command{
		Use:   "rust ENCODER_NAME",
		Short: "Create a new Rust traffic encoder",
		Long: `Create a new Rust traffic encoder package.
ENCODER_NAME is the name of the traffic encoder.

ENCODER_NAME can only contain alphanumeric characters and underscores.`,
		Example: "sliver-sdk new-encoder rust my-encoder",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("invalid number of arguments")
			}
			encName := args[0]
			zipData, err := rustenc.RenderRustTemplate(encName)
			if err != nil {
				return err
			}
			err = os.WriteFile(encName+"_encoder.zip", zipData, 0644)
			if err != nil {
				return err
			}
			cmd.Printf("[*] Your encoder package ready: %s_encoder.zip\n", encName)
			return nil
		},
	}
)

func NewRustEncoderCmd() *cobra.Command {
	return rustEncCmd
}
