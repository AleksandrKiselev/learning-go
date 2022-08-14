// Выполняет парралельную выборку url и сообщает время выполнения и размер ответа
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}

	r := regexp.MustCompile(`(:\/\/|\.[a-zA-Z]+$)`)
	filename := r.Split(url, -1)[1]
	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("can not create file: %s, error: %v", filename, err)
	}

	nbites, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%10s, %.2fs, %7db, %s", resp.Status, secs, nbites, url)
}
