package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// computeE computes the approximation of e by running a fixed number of iterations.
func computeE(iterations int64) float64 {
	res := 2.0
	fact := 1.0

	for i := int64(2); i < iterations; i++ {
		fact *= float64(i)
		res += 1 / fact
	}
	return res
}

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there from %s!", r.Host)
}

func main() {
	http.HandleFunc("/e", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		iters := int64(100)
		keys, ok := r.URL.Query()["iters"]
		if ok && len(keys[0]) >= 1 {
			val, err := strconv.ParseInt(keys[0], 10, 64)
			if err != nil || val <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			iters = val
		}

		w.Write([]byte(fmt.Sprintf("e = %0.4f\n", computeE(iters))))
	})

	http.HandleFunc("/", handlerFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
