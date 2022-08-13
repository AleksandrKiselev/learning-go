// Подсчитывает и выводит количество дублированных строк (из набора файлов)
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		counts := make(map[string]int)
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if len(line) > 0 {
				counts[line]++
			}
		}

		for _, n := range counts {
			if n > 1 {
				fmt.Printf("%s\n", filename)
				break
			}
		}
	}
}
