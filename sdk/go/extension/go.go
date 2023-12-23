package goext

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
		return nil, sdk.ErrInvalidExtName
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
	walkErr := fs.WalkDir(templates.GoTemplates, goFolderName, func(path string, d fs.DirEntry, err error) error {
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
			fTemp, parseErr := template.ParseFS(templates.GoTemplates, path)
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
