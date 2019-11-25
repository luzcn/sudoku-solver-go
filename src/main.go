package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, r.Host)
}

func readJson() {
	dataJson := `["1,2,3", "4,5,6,.,7"]`
	var arr [][]string
	_ = json.Unmarshal([]byte(dataJson), &arr)
	fmt.Println(arr)
}

func main() {
	//fmt.Println("hello")
	//http.HandleFunc("/health", health)

	//_ = http.ListenAndServe(":5000", nil)
}
