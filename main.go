package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luzcn/sudoku-solver-go/solver"
	"log"
	"net/http"
	"os"
	"time"
)

func health(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, r.Host)
}

// curl -X POST \
//  http://localhost:5000/solve \
//  -H 'Content-Type: application/json' \
//  -d '["53..7....","6..195...", ".98....6.","8...6...3", "4..8.3..1","7...2...6", ".6....28.","...419..5", "....8..79"]'
func solveHandler(res http.ResponseWriter, req *http.Request) {
	var body []string
	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		_, _ = fmt.Fprintf(res, fmt.Sprintf("%s", err))
		panic(err)
	}

	board := make([][]byte, 0)
	for _, row := range body {
		board = append(board, []byte(row))
	}

	if solver.Solver(board) {
		_, _ = fmt.Fprintf(res, fmt.Sprintf("%s", board))
	} else {
		_, _ = fmt.Fprintf(res, fmt.Sprintf("%s", "cannot solve it"))
	}
}
func main() {
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
