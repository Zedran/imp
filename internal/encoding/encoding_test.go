package encoding

import (
	"io"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/Zedran/imp/internal/tests"
)

func TestOpenUTF8(t *testing.T) {
	reference, err := os.ReadFile(filepath.Join(tests.TEST_DATA_DIR, "utf-8.txt"))
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	refSum := tests.SHA256(reference)

	cases := []string{
		"windows-1250",
		"utf-8",
	}

	for _, c := range cases {
		reader, err := OpenUTF8(filepath.Join(tests.TEST_DATA_DIR, c+".txt"), c)
		if err != nil {
			t.Fatalf("failed to load test data: '%v'", err)
		}
		defer reader.Close()

		stream, err := io.ReadAll(reader)
		if err != nil {
			t.Fatalf("failed to read normalized data: '%v'", err)
		}

		outSum := tests.SHA256(stream)

		if !slices.Equal(refSum, outSum) {
			t.Fatalf("normalized '%s' file is not equal to the UTF-8 reference", c)
		}
	}

}
