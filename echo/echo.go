// Печатает аргументы, с которыми была вызвана
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, ": ", arg)
	}
}
