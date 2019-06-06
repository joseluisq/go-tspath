package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	tsconfig := tsconfig.New(*configPath, *sourcePath)
	config := tsconfig.Read()

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

	// TODO: create a replacement string array
	var replacements tsconfig.TSPathReplacement

	// TODO: replace all ocurrences per file
	for _, file := range files {
		replacer.Replace(file, replacements)
	}

	fmt.Println(config)
	fmt.Println("SOURCE_PATH:", *sourcePath)
	fmt.Println("ABS_SOURCE_PATH:", absSourcePath)
	fmt.Println("FILES:", files)
}
