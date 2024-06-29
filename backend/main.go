package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Golang Web!")
	})

	// サーバーをポート8080で起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
