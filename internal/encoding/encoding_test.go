package encoding

import (
	"io"
	"os"
	"path"
	"slices"
	"testing"

	"github.com/Zedran/imp/internal/tests"
)

func TestOpenUTF8(t *testing.T) {
	reference, err := os.ReadFile(path.Join(tests.TEST_DATA_DIR, "utf-8.txt"))
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	refSum := tests.SHA256(reference)

	reader, err := OpenUTF8(path.Join(tests.TEST_DATA_DIR, "windows-1250.txt"), "windows-1250")
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
		t.Fatalf("normalized data is not equal to the UTF-8 reference")
	}
}
