package main

import (
	"fmt"
	"github.com/clearblade/cbjson"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Printf("%s... ", arg)
		_, _, err := cbjson.GetJSONFile(arg)
		if err != nil {
			fmt.Printf("Not Good: %s\n", err.Error())
		} else {
			fmt.Printf("Good\n")
		}
	}
}
