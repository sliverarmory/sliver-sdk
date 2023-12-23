package goext

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
	"html/template"
	"io/fs"
	"regexp"
	"strings"

	"github.com/sliverarmory/sliver-sdk/sdk"
	"github.com/sliverarmory/sliver-sdk/templates"
)

const (
	extPlaceholder   = "myextension"
	goFolderName     = "extensions/go"
	defaultGoVersion = "1.21"
)

type GoExtension struct {
	ExtensionName string
	PackageName   string
	GoVersion     string
}

func validateExtname(extName string) bool {
	regex := regexp.MustCompile(`[^a-zA-Z0-9_\-]+`)
	return !regex.Match([]byte(extName))
}

func validatePkgName(pkgName string) bool {
	regex := regexp.MustCompile(`[^a-zA-Z0-9\/_\-\.]+`)
	return !regex.Match([]byte(pkgName))
}

func RenderGoTemplate(pkgName, extName, goVersion string) ([]byte, error) {

	if !validateExtname(extName) {
		return nil, sdk.ErrInvalidName
	}
	if !validatePkgName(pkgName) {
		return nil, sdk.ErrInvalidPackageName
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	if goVersion == "" {
		goVersion = defaultGoVersion
	}
	// Add go.mod
	goModBytes, err := genGoMod(GoExtension{
		ExtensionName: extName,
		PackageName:   pkgName,
		GoVersion:     goVersion,
	})
	if err != nil {
		return nil, err
	}
	goModFile, err := zipWriter.Create("go.mod")
	if err != nil {
		return nil, err
	}
	_, err = goModFile.Write(goModBytes)
	if err != nil {
		return nil, err
	}
	// Walk the templates directory and write each file to the zip archive
	walkErr := fs.WalkDir(templates.GoExtensionTemplates, goFolderName, func(path string, d fs.DirEntry, err error) error {
		zipPath := path
		// remove the top level "go" folder
		zipPath = strings.Replace(zipPath, goFolderName+"/", "", 1)
		if zipPath == goFolderName {
			return nil
		}
		if d.IsDir() {
			zipPath += "/"
		}
		// rename the pkg/myextension/myextension.go file and directory
		zipPath = strings.ReplaceAll(zipPath, extPlaceholder, extName)
		f, zipErr := zipWriter.Create(zipPath)
		if zipErr != nil {
			return zipErr
		}
		if !d.IsDir() {
			fTemp, parseErr := template.ParseFS(templates.GoExtensionTemplates, path)
			if err != nil {
				return parseErr
			}
			data := GoExtension{
				ExtensionName: extName,
				PackageName:   pkgName,
				GoVersion:     goVersion,
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

// genGoMod generates a go.mod file to add to the
// extension package because we can't include one
// in the template folder, otherwise embed won't work
func genGoMod(goExt GoExtension) ([]byte, error) {
	goModTmpl := `module {{.PackageName}}

go {{.GoVersion}}}

require (
	golang.org/x/text v0.14.0
)
	`
	tmpl, err := template.New("go.mod").Parse(goModTmpl)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, goExt)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
