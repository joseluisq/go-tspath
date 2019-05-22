package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joseluisq/redel"
)

func main() {
	r, err := os.Open("case/src/test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	w, err := os.Create("case/src/test.mod.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer w.Close()

	var writer = bufio.NewWriter(w)

	rep := redel.NewRedel(r, "require('", "')", "+++++")
	// rep := redel.NewRedel(r, "START", "END", "+++++")

	replaceFunc := func(data []byte, atEOF bool) {
		_, err := writer.Write(data)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	filterFunc := func(matchValue []byte) bool {
		fmt.Println("Match value:", string(matchValue))
		fmt.Println("------")

		if string(matchValue) == "~/222C" {
			return false
		}

		return true
	}

	rep.FilterReplace(replaceFunc, filterFunc, true)

	writer.Flush()
}
