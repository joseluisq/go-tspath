package cmd

import (
	"flag"
	"fmt"

	tsconfig "go-tspath/pkg/tsconfig"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	sourcePath := flag.String("source", "./dist/**/*.js", "Specifies path of Javascript files emitted by tsc.")
	configPath := flag.String("config", "./tsconfig.json", "Specifies the Typescript configuration file.")

	flag.Parse()

	tsconfig := tsconfig.New(*configPath, *sourcePath)
	config := tsconfig.Read()

	fmt.Println(config.CompilerOptions)
}
