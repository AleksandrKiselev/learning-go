// Выводит ответ на запрос по заданному url (легковесный curl)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) && !strings.HasPrefix(url, httpsPrefix) {
			url = httpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status: %s\n\n", resp.Status)
		written, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\n\nWritten: %d\n", written)
	}
}
