package main

import (
	"expense-tracker/cmd"
	"fmt"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}
