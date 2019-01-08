package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(myHandleFunc))

	http.ListenAndServe(":8080", nil)
}

func myHandleFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Println("This is inside myHandleFunc")
}
