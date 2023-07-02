package main

import (
	"bytes"
	"fmt"

	"github.com/cybdry/go-learning/chap_1/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
}

// @Title
// @Description
// @Author
// @Update
