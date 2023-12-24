package rustenc

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
	"archive/zip"
	"bytes"
	"io/fs"
	"regexp"
	"strings"
	"text/template"

	"github.com/sliverarmory/sliver-sdk/sdk"
	"github.com/sliverarmory/sliver-sdk/templates"
)

const (
	encPlaceholder = "rename"
	rustFolderName = "encoders/rust"
)

type RustEncoder struct {
	EncoderName string
}

func validateEncName(encName string) bool {
	regex := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	return !regex.Match([]byte(encName))
}

func RenderRustTemplate(encName string) ([]byte, error) {
	if !validateEncName(encName) {
		return nil, sdk.ErrInvalidName
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Walk the templates directory and write each file to the zip archive
	walkErr := fs.WalkDir(templates.RustTrafficEncoderTemplates, rustFolderName, func(path string, d fs.DirEntry, err error) error {
		zipPath := path
		// remove the top level "rust" folder
		zipPath = strings.Replace(zipPath, rustFolderName+"/", "", 1)
		if zipPath == rustFolderName {
			return nil
		}
		if d.IsDir() {
			zipPath += "/"
		}

		zipPath = strings.ReplaceAll(zipPath, encPlaceholder, encName)
		f, zipErr := zipWriter.Create(zipPath)
		if zipErr != nil {
			return zipErr
		}
		if !d.IsDir() {
			fTemp, parseErr := template.ParseFS(templates.RustTrafficEncoderTemplates, path)
			if err != nil {
				return parseErr
			}
			data := RustEncoder{
				EncoderName: encName,
			}
			execErr := fTemp.Execute(f, data)
			if err != nil {
				return execErr
			}
		}
		return err
	})
	if walkErr != nil {
		return nil, walkErr
	}
	zipWriter.Close()
	return buf.Bytes(), nil
}
