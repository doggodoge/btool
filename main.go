package main

import (
	"flag"
	"sync"

	"github.com/doggodoge/btool/compress"
	"github.com/doggodoge/btool/decompress"
)

func main() {
	decompressFlag := flag.Bool("d", false, "set flag to decompress a file or folder.")

	flag.Parse()

	paths := flag.Args()

	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)

		if *decompressFlag {
			go func(path string) {
				defer wg.Done()

				err := decompress.File(path)
				if err != nil {
					panic(err)
				}
			}(path)
		} else {
			go func(path string) {
				defer wg.Done()

				err := compress.File(path)
				if err != nil {
					panic(err)
				}
			}(path)
		}
	}

	wg.Wait()
}
