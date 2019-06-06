package tsconfig

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	// TSConfig defines tsconfig instance
	TSConfig struct {
		ConfigPath string
		SourcePath string
		Data       *TSConfigData
	}

	// TSCompilerOptions defines a small Typescript compiler options
	TSCompilerOptions struct {
		BaseURL string              `json:"baseUrl"`
		Paths   map[string][]string `json:"paths"`
	}

	// TSConfigData defines tsconfig json file properties
	TSConfigData struct {
		CompilerOptions TSCompilerOptions `json:"compilerOptions"`
	}

	// TSPathReplacement defines a single Typescript path with its replacement
	TSPathReplacement struct {
		Pattern     []byte
		Replacement []byte
	}
)

// New creates a new TSConfig instance.
func New(configPath string, sourcePath string) *TSConfig {
	return &TSConfig{
		ConfigPath: configPath,
		SourcePath: sourcePath,
		Data: &TSConfigData{
			CompilerOptions: TSCompilerOptions{
				BaseURL: "./",
			},
		},
	}
}

// Read reads a tsconfig.json file
func (tsconf *TSConfig) Read() TSConfigData {
	r, err := os.Open(tsconf.ConfigPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	dec := json.NewDecoder(r)

	var data TSConfigData

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
