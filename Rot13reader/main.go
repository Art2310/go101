package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (l int, err error) {
	l, err = rr.r.Read(b)
	for i := 0; i < l; i++ {
		switch {
		case b[i] < 78:
			b[i] = b[i] + 13
		case b[i] <= 90:
			b[i] = b[i] - 13
		case b[i] < 110:
			b[i] = b[i] + 13
		default:
			b[i] = b[i] - 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
