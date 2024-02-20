package main

import "fmt"

type MyReader struct{}
type MyReaderError bool

func (m MyReaderError) Error() string {
	return fmt.Sprintf("Error")
}

func (m MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderError(true)
	}
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
