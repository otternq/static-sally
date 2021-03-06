package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
)

// Config represents the structure of the yaml file
type Config struct {
	URL      string             `yaml:"url"`
	Packages map[string]Package `yaml:"packages"`
}

// Package details the options available for each repo
type Package struct {
	Repo string `yaml:"repo"`
}

// ensureAlphabetical checks that the packages are listed alphabetically in the configuration.
func ensureAlphabetical(data []byte) bool {
	// A yaml.MapSlice perservers ordering of keys: https://godoc.org/gopkg.in/yaml.v2#MapSlice
	var c struct {
		Packages yaml.MapSlice `yaml:"packages"`
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		return false
	}

	packageNames := make([]string, 0, len(c.Packages))
	for _, v := range c.Packages {
		name, ok := v.Key.(string)
		if !ok {
			return false
		}
		packageNames = append(packageNames, name)
	}

	return sort.StringsAreSorted(packageNames)
}

// Parse takes a path to a yaml file and produces a parsed Config
func Parse(path string) (*Config, error) {
	var (
		err  error
		data []byte
		c    Config
	)

	if data, err = ioutil.ReadFile(path); err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	if !ensureAlphabetical(data) {
		return nil, fmt.Errorf("packages in %s must be alphabetically ordered", path)
	}

	return &c, err
}

// convertToPackageInfo converts a Config object into a slice of PackageInfo objects
func convertToPackageInfo(config *Config) []PackageInfo {
	var packages = []PackageInfo{}

	for packageName, packageDetails := range config.Packages {
		packages = append(
			packages,
			PackageInfo{
				Name:         packageName,
				Repo:         packageDetails.Repo,
				CanonicalURL: config.URL + "/" + packageName,
				GodocURL:     "https://godoc.org/" + config.URL + "/" + packageName,
			},
		)
	}

	return packages
}
