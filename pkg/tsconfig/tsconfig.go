// Package tsconfig reads a tsconfig.json content
package tsconfig

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	// Config defines tsconfig instance
	Config struct {
		ConfigPath string
		Data       *ConfigData
	}

	// CompilerOptions defines a small Typescript compiler options
	CompilerOptions struct {
		BaseURL string              `json:"baseUrl"`
		Paths   map[string][]string `json:"paths"`
		OutDir  string              `json:"outDir"`
	}

	// ConfigData defines tsconfig json file properties
	ConfigData struct {
		CompilerOptions CompilerOptions `json:"compilerOptions"`
	}

	// PathReplacement defines a single Typescript path with its replacement
	PathReplacement struct {
		Pattern     []byte
		Replacement [][]byte
	}
)

// New creates a new Config instance.
func New(configPath string) *Config {
	if len(configPath) <= 0 {
		configPath = "./tsconfig.json"
	}

	return &Config{
		ConfigPath: configPath,
		Data: &ConfigData{
			CompilerOptions: CompilerOptions{
				BaseURL: "./",
			},
		},
	}
}

// Read reads a tsconfig.json file
func (tsconf *Config) Read() ConfigData {
	r, err := os.Open(tsconf.ConfigPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	dec := json.NewDecoder(r)

	var data ConfigData

	for {
		if err := dec.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(data.CompilerOptions.BaseURL); os.IsNotExist(err) {
		data.CompilerOptions.BaseURL = tsconf.Data.CompilerOptions.BaseURL
	}

	return data
}
