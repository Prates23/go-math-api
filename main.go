package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/hello", helloHandler)
	mux.HandleFunc("/api/v1/sum", makeMathHandler(sum))
	mux.HandleFunc("/api/v1/subtract", makeMathHandler(subtract))
	mux.HandleFunc("/api/v1/multiply", makeMathHandler(multiply))
	mux.HandleFunc("/api/v1/divide", makeMathHandler(divide))

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request on /api/v1/hello")
	fmt.Fprintln(w, "Hello from the Go application running inside the container!!!!!")
}

func makeMathHandler(op func(float64, float64) (float64, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aStr := r.URL.Query().Get("a")
		bStr := r.URL.Query().Get("b")

		a, err1 := strconv.ParseFloat(aStr, 64)
		b, err2 := strconv.ParseFloat(bStr, 64)

		if err1 != nil || err2 != nil {
			log.Printf("Invalid parameters: a=%s, b=%s", aStr, bStr)
			json.NewEncoder(w).Encode(response{Error: "Invalid parameters"})
			return
		}

		result, err := op(a, b)
		if err != nil {
			log.Printf("Error performing operation: %v", err)
			json.NewEncoder(w).Encode(response{Error: err.Error()})
			return
		}

		log.Printf("Operation result: %.2f", result)
		json.NewEncoder(w).Encode(response{Result: result})
	}
}

func sum(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
