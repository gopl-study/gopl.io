package bzip

import "C"
import (
	"io"
	"os/exec"
)

type writer struct {
	w io.WriteCloser

}

func NewWriter(out io.Writer) io.WriteCloser {
	cmd := exec.Command("bzip2")
	cmd.Stdout = out
	wc, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	w := &writer{wc}
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	return w.w.Write(data)
}

func (w *writer) Close() error {
	return w.w.Close()
}
