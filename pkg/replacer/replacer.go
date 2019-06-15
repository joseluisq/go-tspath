package replacer

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/joseluisq/go-tspath/pkg/tsconfig"

	"github.com/joseluisq/redel"
)

// Replace replaces every TS path occurence per file
func Replace(filePathAbs string, filePathRel string, outDir string, replacements []tsconfig.TSPathReplacement) {
	r, err := os.OpenFile(filePathAbs, os.O_RDONLY, 0)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer r.Close()

	w, err := os.OpenFile(filePathAbs, os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
			os.Exit(1)
		}
	}

	pathRel := filepath.Dir(filePathRel)

	filterFunc := func(matchValue []byte) []byte {
		for _, vtspath := range replacements {
			if len(vtspath.Replacement) <= 0 {
				continue
			}

			// TODO: Verify that every replaced path is valid (if path exists)
			if bytes.HasPrefix(matchValue, vtspath.Pattern) {
				repl := bytes.Replace(matchValue, vtspath.Pattern, []byte(outDir), 1)

				replacement, err := filepath.Rel(pathRel, string(repl))

				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				return []byte("./" + replacement)
			}
		}

		return matchValue
	}

	rep.ReplaceFilterWith(replaceFunc, filterFunc, true)

	writer.Flush()
}
