// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
// Ref: https://go.dev/tour/methods/22
// Answer: https://stackoverflow.com/questions/27839140/tour-of-go-exercise-22-reader-what-does-the-question-mean

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read (b []byte) (int, error) {
    for x := range b {
        b[x] = 'A'
    }
    return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
