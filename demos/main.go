package main

import (
	"fmt"
	"os"
)

func main() {
	demos := map[string]func(){
		"simpleintersection": SimpleIntersection,
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: demo [demoname]")
		return
	}

	demoName := os.Args[1]

	demo, exists := demos[demoName]
	if !exists {
		err := fmt.Errorf("no demo exists with name '%s'", demoName)
		panic(err)
	}

	fmt.Printf("Running demo '%s'...\n", demoName)
	demo()
}
