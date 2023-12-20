package rustext

import (
	"archive/zip"
	"bytes"
	"html/template"
	"io/fs"
	"regexp"
	"strings"

	"github.com/bishopfox/sliver-sdk/sdk"
	"github.com/bishopfox/sliver-sdk/templates"
)

const (
	extPlaceholder = "myextension"
	rustFolderName = "rust"
)

type RustExtension struct {
	ExtensionName string
}

func validateExtname(extName string) bool {
	regex := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	return !regex.Match([]byte(extName))
}

func RenderRustTemplate(extName string) ([]byte, error) {
	if !validateExtname(extName) {
		return nil, sdk.ErrInvalidExtName
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	// Add go.mod
	// Walk the templates directory and write each file to the zip archive
	walkErr := fs.WalkDir(templates.RustTemplates, rustFolderName, func(path string, d fs.DirEntry, err error) error {
		zipPath := path
		// remove the top level "go" folder
		zipPath = strings.Replace(zipPath, rustFolderName+"/", "", 1)
		if zipPath == rustFolderName {
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
			fTemp, parseErr := template.ParseFS(templates.RustTemplates, path)
			if err != nil {
				return parseErr
			}
			data := RustExtension{
				ExtensionName: extName,
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
