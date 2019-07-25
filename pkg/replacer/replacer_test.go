package replacer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joseluisq/go-tspath/pkg/tsconfig"
)

func TestReplace(t *testing.T) {
	type args struct {
		filePathAbs  string
		filePathRel  string
		outDir       string
		replacements []tsconfig.PathReplacement
	}

	filePathRel := "./../../sample/dist/index.js"
	outDir := "./../../sample/dist"
	filePathAbs, err := filepath.Abs(filePathRel)

	if err != nil {
		os.Exit(1)
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "replace files at sample/dist properly",
			args: args{
				filePathAbs: filePathAbs,
				filePathRel: filePathRel,
				outDir:      outDir,
				replacements: []tsconfig.PathReplacement{
					{
						Pattern: []byte("~/"),
						Replacement: [][]byte{
							[]byte("./"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Replace(tt.args.filePathAbs, tt.args.filePathRel, tt.args.outDir, tt.args.replacements)
		})
	}
}
