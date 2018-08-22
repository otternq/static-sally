package main

import (
	"html/template"
	"os"
	"path/filepath"
)

// PackageInfo contains information required for the package template
type PackageInfo struct {
	Name         string
	Repo         string
	CanonicalURL string
	GodocURL     string
}

// generateSite creates a directory and index.html files per PackageInfo item
func generateSite(packageInfoSlice []PackageInfo, outputDirectory string, indexTemplate *template.Template) error {
	for _, packageInfo := range packageInfoSlice {
		var (
			err       error
			outputDir = filepath.Join(outputDirectory, packageInfo.Name)
			fileName  = filepath.Join(outputDir + "/index.html")
			file      *os.File
		)

		if err = os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return err
		}

		if file, err = os.Create(fileName); err != nil {
			return err
		}

		if err = indexTemplate.Execute(file, packageInfo); err != nil {
			return err
		}
	}

	return nil
}
