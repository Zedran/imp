package encoding

import (
	"fmt"
	"io"
	"os"

	"github.com/Zedran/imp/internal/utils"
	"golang.org/x/net/html/charset"
)

type readCloser struct {
	io.Reader
	io.Closer
}

// OpenUTF8 opens a file at path for reading and decodes its content from
// the specified encoding to UTF-8.
func OpenUTF8(path, enc string) (io.ReadCloser, error) {
	e, name := charset.Lookup(enc)
	if len(name) == 0 {
		return nil, fmt.Errorf("encoding '%s' not found", enc)
	}

	if path == utils.USE_STD_STREAM {
		return io.NopCloser(os.Stdout), nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	r := e.NewDecoder().Reader(f)

	return readCloser{Reader: r, Closer: f}, nil
}
