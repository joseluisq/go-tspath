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
	r, err := os.Open(filePathAbs)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer r.Close()

	filePathTemp := filePathAbs + ".js"

	w, err := os.Create(filePathTemp)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer w.Close()

	var writer = bufio.NewWriter(w)

	pathRel := filepath.Dir(filePathRel)

	replaceFunc := func(data []byte, atEOF bool) {
		_, err := writer.Write(data)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if atEOF {
			err := os.Remove(filePathAbs)

			if err != nil {
				log.Fatal(err)
			}

			err = os.Rename(filePathTemp, filePathAbs)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	filterFunc := func(matchValue []byte) []byte {
		for _, vtspath := range replacements {
			if len(vtspath.Replacement) <= 0 {
				continue
			}

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

	rep := redel.New(r, []redel.Delimiter{
		{Start: []byte("require(\""), End: []byte("\");")},
		{Start: []byte("from \""), End: []byte("\";")},
	})

	rep.ReplaceFilterWith(replaceFunc, filterFunc, true)

	writer.Flush()
}
