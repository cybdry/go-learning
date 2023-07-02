package interfaces

import (
	"io"
	"os"
)

// Copy copies dat from in to out first directly,
// then using a buffer . It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	// we write to out ,but also stdout
	w := io.MultiWriter(out, os.Stdout)
	// a standart copy ,this can be dangerous if there's a
	// lot of data in in

	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	//buffered write using 64 bytes chuncks
	buffer := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buffer); err != nil {
		return err
	}
	return nil

}
