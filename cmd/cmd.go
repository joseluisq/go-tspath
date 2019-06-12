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
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	sourcePath := flag.String("source", "./dist/**/*.js", "Specifies path of Javascript files emitted by tsc.")
	configPath := flag.String("config", "./tsconfig.json", "Specifies the Typescript configuration file.")

	flag.Parse()

	tsconf := tsconfig.New(*configPath, *sourcePath)
	config := tsconf.Read()

	absSourcePath, err := filepath.Abs(*sourcePath)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	files, err := zglob.Glob(absSourcePath)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Create the replacement string array (pattern-replacement)
	var replacements []tsconfig.TSPathReplacement

	for kPathStr, vPathStr := range config.CompilerOptions.Paths {
		kParts := strings.Split(kPathStr, "/")

		// 0. Prevent no valid paths (key-value)
		if len(kParts) <= 0 || len(vPathStr) <= 0 {
			continue
		}

		// 1. Pattern placeholder: Take key parts skipping the last one
		patternStr := strings.TrimSpace(strings.Join(kParts[:len(kParts)-1], "/"))

		if len(patternStr) <= 0 {
			continue
		}

		// Store pattern string as byte
		pattern := []byte(patternStr)

		// 2. Replacement placeholder: Take value parts skipping the last one
		var replacementBytes [][]byte

		for _, vpathstr := range vPathStr {
			vparts := strings.Split(vpathstr, "/")

			// Prevent no valid replacement paths
			if len(vparts) <= 0 {
				continue
			}

			value := strings.TrimSpace(strings.Join(vparts[:len(vparts)-1], "/"))

			// Prevent empty replacement paths
			if len(value) <= 0 {
				continue
			}

			replacementBytes = append(replacementBytes, [][]byte{
				[]byte(value),
			}...)
		}

		replacements = append(replacements, tsconfig.TSPathReplacement{
			Pattern:     pattern,
			Replacement: replacementBytes,
		})
	}

	// Replace all occurrences per file
	for _, file := range files {
		replacer.Replace(file, replacements)
	}

	// TODO: Provide useful output information about the realized job
}
