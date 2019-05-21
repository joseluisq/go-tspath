package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joseluisq/redel"
)

func main() {
	r, err := os.Open("case/src/index.js")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	w, err := os.Create("case/src/index.mod.js")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer w.Close()

	var writer = bufio.NewWriter(w)

	rep := redel.NewRedel(r, "require('", "')", "+++++")
	rep.FilterReplace(func(matchValue []byte, atEOF bool) bool {
		val := string(matchValue)
		fmt.Println("Match value:", val)
		fmt.Println("------")

		// if val == "~/222C" {
		// 	return true
		// }

		return true
	}, func(data []byte, atEOF bool) {
		_, err := writer.Write(data)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})

	writer.Flush()
}
