package main

import (
	"fmt"
	"log"
	"net/http"
	////"github.com/prometheus/common/log"
)

func main() {
	fmt.Println("--main-start--")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Host + r.RequestURI))
	})
	addr := ":8080"
	log.Fatal(http.ListenAndServe(addr, nil))
}
