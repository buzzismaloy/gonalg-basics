package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type LimitingReader struct {
	reader    io.Reader
	bytesLeft int
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitingReader{reader: r, bytesLeft: n}
}

func (r *LimitingReader) Read(p []byte) (int, error) {
	if r.bytesLeft == 0 {
		return 0, io.EOF
	}

	if r.bytesLeft < len(p) {
		p = p[0:r.bytesLeft]
	}

	n, err := r.reader.Read(p)
	r.bytesLeft -= n

	return n, err
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
