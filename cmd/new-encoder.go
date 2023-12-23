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
