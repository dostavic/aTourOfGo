package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

//type rot13ReaderError bool
//
//func (e rot13ReaderError) Error() string {
//	return fmt.Sprintf("Error")
//}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		if (b[i] >= 'a' && b[i] <= 'z') || (b[i] >= 'A' && b[i] <= 'Z') {
			b[i] = b[i] + 13
			if b[i] > 'z' {
				b[i] = b[i] - 26
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
