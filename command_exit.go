package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
