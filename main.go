
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: Alert-Manager <file-or-dir>")
		os.Exit(1)
	}
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("name=%s size=%d modified=%s\n", fi.Name(), fi.Size(), fi.ModTime().Format(time.RFC3339))
}
