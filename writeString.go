package tint

import (
	"bytes"
	"fmt"
	"io"
)

func writeString(buf *buffer, s string) {
	iw := &indentWriter{
		Writer:      buf,
		indentBytes: []byte("  "), // two spaces
		isMultiline: false,
	}
	io.WriteString(iw, s)
	if iw.isMultiline {
		fmt.Println()
	}
}

type indentWriter struct {
	io.Writer
	indentBytes []byte
	isMultiline bool
}

func (w *indentWriter) Write(p []byte) (n int, err error) {
	w.isMultiline = false
	// In case of a multiline string, we need to insert an <indent> after each newline.
	// The most performant way to do this is to find the index of the newline,
	// write the slice up to and including the newline, write the <indent> and
	// continue with the rest of the slice.
	offset := 0
	for {
		// find the index of the newline
		index := bytes.IndexByte(p[offset:], '\n')
		if index == -1 {
			// no newline found, write the rest of the slice
			written, err := w.Writer.Write(p[offset:])
			return n + written, err
		}
		// write the slice up to and including the newline
		written, err := w.Writer.Write(p[offset : offset+index+1])
		n += written
		if err != nil {
			return n, err
		}
		// write the indent
		written, err = w.Writer.Write(w.indentBytes)
		n += written
		if err != nil {
			return n, err
		}
		w.isMultiline = true
		// continue with the rest of the slice
		offset += index + 1
	}
}
