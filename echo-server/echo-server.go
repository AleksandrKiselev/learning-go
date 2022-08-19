// Минимальный echo сервер, со счетчиком запросов
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	http.HandleFunc("/count", func(w http.ResponseWriter, _ *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count = %d\n", count)
		mu.Unlock()
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
