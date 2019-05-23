package main

import (
	"bufio"
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

	rep := redel.NewRedel(r, "require('", "')")
	// rep := redel.NewRedel(r, "START", "END", "+++++")

	replaceFunc := func(data []byte, atEOF bool) {
		v := data

		// if len(values) > 0 {
		// 	d := append([]byte(nil), data...)
		// 	v = append(d, values...)
		// }

		_, err := writer.Write(v)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	filterFunc := func(matchValue []byte) bool {
		value := string(matchValue)

		// fmt.Println("MATCH VALUE::", value)
		// fmt.Println("------")

		// if value == "~/222C" || value == "~/111B" {
		// 	return false
		// }

		// if value == "~/222C" || value == "~/111B" {
		// 	return false
		// }

		// if value == "~/222C" || value == "~/111B" {
		// 	a := bytes.Replace(matchValue, []byte("~/"), []byte("___"), 1)

		// 	return a
		// }

		if value == " slice, " {
			// TODO: Fix Redel to support extra chars
			// return append(matchValue, []byte("====")...)
			// return []byte("...........................")
		}

		if value == "~/222C" {
			return false
		}

		// if value == " slice, " {
		// 	return false
		// }

		return true

		// return matchValue
	}

	rep.FilterReplace([]byte("1234567"), replaceFunc, filterFunc, true)

	writer.Flush()
}
