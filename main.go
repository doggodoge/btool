package main

import (
	"flag"
	"fmt"
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
				fmt.Printf("decompressed %s\n", path)
			}(path)
		} else {
			go func(path string) {
				defer wg.Done()

				fmt.Printf("compressing %s\n", path)
				err := compress.File(path)
				if err != nil {
					panic(err)
				}
				fmt.Printf("finished %s\n", path)
			}(path)
		}
	}

	wg.Wait()
}
