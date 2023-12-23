package rust_encoder

import (
	"archive/zip"
	"bytes"
	"io/fs"
	"regexp"
	"strings"
	"text/template"

	"github.com/bishopfox/sliver-sdk/sdk"
	"github.com/bishopfox/sliver-sdk/templates"
)

const (
	encPlaceholder = "rename"
	rustFolderName = "rust"
)

type RustEncoder struct {
	EncoderName string
}

func validateEncName(extName string) bool {
	regex := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	return !regex.Match([]byte(extName))
}

func RenderRustTemplate(extName string) ([]byte, error) {
	if !validateEncName(extName) {
		return nil, sdk.ErrInvalidExtName
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	walkErr := fs.WalkDir(templates.RustTrafficEncoderTemplates, rustFolderName, func(path string, d fs.DirEntry, err error) error {
		zipPath := path

		zipPath = strings.Replace(zipPath, rustFolderName+"/", "", 1)
		if zipPath == rustFolderName {
			return nil
		}
		if d.IsDir() {
			zipPath += "/"
		}

		zipPath = strings.ReplaceAll(zipPath, encPlaceholder, extName)
		f, zipErr := zipWriter.Create(zipPath)
		if zipErr != nil {
			return zipErr
		}
		if !d.IsDir() {
			fTemp, parseErr := template.ParseFS(templates.RustExtensionTemplates, path)
			if err != nil {
				return parseErr
			}
			data := RustEncoder{
				EncoderName: extName,
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
