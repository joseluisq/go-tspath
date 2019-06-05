package replacer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"redel"
)

// Replace replaces every TS path occurence per file
// TODO:
func Replace() {
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
		{Start: []byte("require(\""), End: []byte("\")")},
		{Start: []byte("from \""), End: []byte("\";")},
		{Start: []byte("from '"), End: []byte("';")},
	})

	replaceFunc := func(data []byte, atEOF bool) {
		_, err := writer.Write(data)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	prefixSearch := []byte("~/")
	replaceValue := []byte("+++")

	filterFunc := func(matchValue []byte) []byte {
		if bytes.HasPrefix(matchValue, prefixSearch) {
			return bytes.Replace(matchValue, prefixSearch, replaceValue, 1)
		}

		return matchValue
	}

	rep.ReplaceFilterWith(replaceFunc, filterFunc, true)

	writer.Flush()
}