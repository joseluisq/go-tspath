package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/joseluisq/redel"
)

func main() {
	r, err := os.Open("case/src/test.js")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	w, err := os.Create("case/src/test.mod.js")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer w.Close()

	var writer = bufio.NewWriter(w)

	rep := redel.New(r, []redel.Delimiter{
		{Start: []byte("require('"), End: []byte("')")},
		{Start: []byte("log('"), End: []byte("')")},
	})

	replaceFunc := func(data []byte, atEOF bool) {
		_, err := writer.Write(data)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// filterFunc := func(matchValue []byte) bool {
	// 	value := string(matchValue)

	// 	if value == "~/111B" {
	// 		return false
	// 	}

	// 	return true
	// }

	filterFunc := func(matchValue []byte) []byte {
		value := string(matchValue)

		if value == "~/111B" || value == "~/333D" {
			b := bytes.Replace(matchValue, []byte("~/"), []byte("./"), 1)

			return b
		}

		return matchValue
	}

	// rep.Replace([]byte("1234567"), replaceFunc)
	// rep.ReplaceFilter([]byte("1234567"), replaceFunc, filterFunc, true)
	rep.ReplaceFilterWith(replaceFunc, filterFunc, true)

	writer.Flush()
}
