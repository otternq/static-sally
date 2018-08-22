// Command static-sally generates a set of static webpages supporting custom Golang import paths
package main

import (
	"flag"
	"log"
)

func main() {
	var (
		err error

		yamlConfigPath  string
		outputDirectory string
		importConfig    *Config

		packageInfoSlice []PackageInfo
	)

	flag.StringVar(&outputDirectory, "output", "source", "where to generate the static website")
	flag.StringVar(&yamlConfigPath, "yml", "sally.yaml", "yaml file to read config from")

	flag.Parse()

	if importConfig, err = Parse(yamlConfigPath); err != nil {
		log.Fatalf("unable to parse config: %s", err.Error())
	}

	packageInfoSlice = convertToPackageInfo(importConfig)

	if err = generateSite(packageInfoSlice, outputDirectory, packageTemplate); err != nil {
		log.Fatalf("unable to generate site: %s", err.Error())
	}
}
