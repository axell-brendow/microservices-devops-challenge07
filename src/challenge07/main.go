package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
)

func greeting(str string) string {
	x := 0.0001

	for i := 0; i < 100000; i++ {
		x = math.Sqrt(x)
	}

	return fmt.Sprintf("<b>%s</b>", str)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func main() {
	portPtr := flag.Int("port", 8000, "Server port")
	flag.Parse()
	http.HandleFunc("/", greet)
	fmt.Println(fmt.Sprintf("Starting server on port %d", *portPtr))
	http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil)
}
