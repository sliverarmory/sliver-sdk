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

	goext "github.com/bishopfox/sliver-sdk/sdk/go/extension"
	"github.com/spf13/cobra"
)

var (
	goVersionFlag string
	goExtCmd      = &cobra.Command{
		Use:   "go PKG_NAME EXTENSION_NAME",
		Short: "Create a new Go extension",
		Long: `Create a new Go extension package.
PKG_NAME is the name of the Go package and EXTENSION_NAME is the name of the extension.`,
		Example: "sliver-sdk new go github.com/bishopfox/my-sliver-ext my-sliver-ext",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("invalid number of arguments")
			}
			pkgName := args[0]
			extName := args[1]
			zipData, err := goext.RenderGoTemplate(pkgName, extName, goVersionFlag)
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

func NewGoExtensionCmd() *cobra.Command {
	goExtCmd.Flags().StringVarP(&goVersionFlag, "go-version", "g", "1.21", "Go version to use")
	return goExtCmd
}
