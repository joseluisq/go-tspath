package replacer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"redel"
)

type (
	// TSPathReplacement defines a single Typescript path with its replacement
	TSPathReplacement struct {
		Pattern     []byte
		Replacement []byte
	}
)

// Replace replaces every TS path occurence per file
func Replace(filePath string, replacements []TSPathReplacement) {
	r, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	w, err := os.Create(filePath + ".mod.js")

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

	filterFunc := func(matchValue []byte) []byte {
		for _, tspath := range replacements {
			if bytes.HasPrefix(matchValue, tspath.Pattern) {
				return bytes.Replace(matchValue, tspath.Pattern, tspath.Replacement, 1)
			}
		}

		return matchValue
	}

	rep.ReplaceFilterWith(replaceFunc, filterFunc, true)

	writer.Flush()
}
