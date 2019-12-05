package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luzcn/sudoku-solver-go/solver"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, _ = fmt.Fprintf(w, r.Host)
	}
}

func readJson(data string) (res []string) {
	_ = json.Unmarshal([]byte(data), &res)
	return
}

func solveHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		_, _ = fmt.Fprintf(res, "Not Supported Method")
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		_, _ = fmt.Fprintf(res, fmt.Sprintf("%s", err))
		return
	}

	data := readJson(string(body))
	board := make([][]byte, 0)
	for _, row := range data {
		board = append(board, []byte(row))
	}

	if solver.Solver(board) {
		_, _ = fmt.Fprintf(res, fmt.Sprintf("%s", board))
	} else {
		_, _ = fmt.Fprintf(res, fmt.Sprintf("%s", "cannot solve it"))
	}
}
func main() {
	//http.HandleFunc("/health", health)
	//http.HandleFunc("/solve", solveHandler)
	//
	//log.Println("Start http server")
	//_ = http.ListenAndServe(":5000", nil)
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "5000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/solve", solveHandler).Methods("POST")
	server := &http.Server{
		Addr:         ":" + PORT,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      r,
	}
	log.Printf("action=start-server msg=\"Listening on port %s\"", PORT)
	_ = server.ListenAndServe()
}
