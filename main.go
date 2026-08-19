package main

import (
	"fmt"
	"os"
)

// alert_manager - Intelligent alerting system
func alert_manager(path string) {
	fmt.Println("========================================")
	fmt.Println("  Alert-Manager")
	fmt.Println("  Intelligent alerting system")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	alert_manager(path)
}
