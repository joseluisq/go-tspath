package replacer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/joseluisq/go-tspath/pkg/tsconfig"
	"github.com/joseluisq/redel"
)

// Replace replaces every TS path occurence per file
func Replace(filePath string, replacements []tsconfig.TSPathReplacement) {
	r, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Close()

	// TODO: Write file content properly
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
		for _, vtspath := range replacements {
			if len(vtspath.Replacement) <= 0 {
				continue
			}

			// (!) LIMITATION: Take first Replacement slice only
			replacement := vtspath.Replacement[0]

			// TODO: Consider `baseUrl` value
			// TODO: Verify that every replaced path is valid (if path exists)
			if bytes.HasPrefix(matchValue, vtspath.Pattern) {
				return bytes.Replace(matchValue, vtspath.Pattern, replacement, 1)
			}
		}

		return matchValue
	}

	rep.ReplaceFilterWith(replaceFunc, filterFunc, true)

	writer.Flush()
}
