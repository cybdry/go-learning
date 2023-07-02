// @Title
// @Description
// @Author
// @Update

package interfaces

import (
	"io"
	"os"
)

// PipesExample helps give some more examples of using io interfaces
func PipesExample() error {
	// the pipe reader and pipe write implement io.Reader and io.Writer

	r, w := io.Pipe()
	// this needs to be run in separate go routine as it will block waiting for
	// reader .Close at end for cleanup

	go func() {
		// for now we'll write something basic
		//this could be also used to encode json
		//base 64 encode etc.
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
