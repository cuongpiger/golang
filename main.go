package main

import (
	"fmt"
	"net/http"
)

func readyToLearn(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("<h1>Ready to learn!</h1>"))
	fmt.Println("Server is running...")
}

func main() {
	http.HandleFunc("/", readyToLearn)
	_ = http.ListenAndServe(":8000", nil)
}
