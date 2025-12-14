package main

import (
	"fmt"
	"io"
	"os"
)

const bufSize = 32 * 1024 // 32 KB

func catReader(r io.Reader, buf []byte) error {
	_, err := io.CopyBuffer(os.Stdout, r, buf)
	return err
}

func main() {
	args := os.Args[1:]

	buf := make([]byte, bufSize)

	if len(args) == 0 {
		if err := catReader(os.Stdin, buf); err != nil {
			fmt.Fprintf(os.Stderr, "Error copiando stdin: %v\n", err)
			os.Exit(1)
		}
		return
	}

	for _, name := range args {
		if name == "-" {
			if err := catReader(os.Stdin, buf); err != nil {
				fmt.Fprintf(os.Stderr, "Error copiando stdin: %v\n", err)
			}
			continue
		}

		f, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error abriendo %s: %v\n", name, err)
			continue
		}

		if err := catReader(f, buf); err != nil {
			fmt.Fprintf(os.Stderr, "Error leyendo %s: %v\n", name, err)
		}

		f.Close()
	}
}
