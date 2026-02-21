// imp -- tool for rewriting CSV files and normalizing encoding.
// Copyright (C) 2026  Wojciech Głąb (github.com/Zedran)
//
// This file is part of imp.
//
// imp is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// imp is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with imp.  If not, see <https://www.gnu.org/licenses/>.

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
