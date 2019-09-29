// Package cmd process command line arguments
package cmd

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joseluisq/go-tspath/pkg/replacer"
	"github.com/joseluisq/go-tspath/pkg/tsconfig"

	zglob "github.com/mattn/go-zglob"
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	configPath := flag.String("config", "./tsconfig.json", "Specifies the Typescript configuration file.")

	flag.Parse()

	config := tsconfig.New(*configPath).Read()

	outDir := config.CompilerOptions.OutDir
	basePath := outDir

	// Check if outDir is a relative dir path
	if !filepath.IsAbs(outDir) {
		basePath = filepath.Dir(*configPath)
	}

	outFilesPath := filepath.Join(basePath, outDir, "**/*.js")

	files, err := zglob.Glob(outFilesPath)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Create the replacement string array (pattern-replacement)
	replacements := make([]tsconfig.PathReplacement, 0, len(config.CompilerOptions.Paths))

	for keyPathStr, valuePathStr := range config.CompilerOptions.Paths {
		keyParts := strings.Split(keyPathStr, "/")

		// 0. Prevent no valid paths (key-value)
		if len(keyParts) == 0 || len(valuePathStr) == 0 {
			continue
		}

		// 1. Pattern placeholder: Take key parts skipping the last one
		patternStr := strings.TrimSpace(strings.Join(keyParts[:len(keyParts)-1], "/"))

		if len(patternStr) == 0 {
			continue
		}

		// Store pattern string as byte
		pattern := []byte(patternStr)

		// 2. Replacement placeholder: Take value parts skipping the last one
		replacementBytes := make([][]byte, 0, len(valuePathStr))

		for _, vpathstr := range valuePathStr {
			vparts := strings.Split(vpathstr, "/")

			// Prevent no valid replacement paths
			if len(vparts) == 0 {
				continue
			}

			value := strings.TrimSpace(strings.Join(vparts[:len(vparts)-1], "/"))

			// Prevent empty replacement paths
			if len(value) == 0 {
				continue
			}

			replacementBytes = append(replacementBytes, [][]byte{
				[]byte(value),
			}...)
		}

		replacements = append(replacements, tsconfig.PathReplacement{
			Pattern:     pattern,
			Replacement: replacementBytes,
		})
	}

	// Replace all occurrences per file
	for _, file := range files {
		relFilePath, err := filepath.Rel(basePath, file)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		replacer.Replace(file, relFilePath, outDir, replacements)
	}

	// TODO: Provide useful output information about the realized job
}
