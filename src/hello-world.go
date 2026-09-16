package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	message := "MESSAGE - cicd-main - 02"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})

	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("slow: bat dau xu ly luc", start.Format(time.RFC3339))
		time.Sleep(3 * time.Minute)
		end := time.Now()
		fmt.Println("slow: xu ly xong luc", end.Format(time.RFC3339))
		fmt.Fprintf(w, "OK - da xu ly xong sau %v (bat dau: %s, ket thuc: %s)\n",
			end.Sub(start), start.Format(time.RFC3339), end.Format(time.RFC3339))
	})

	fmt.Println("Starting server on port 8080.")
	http.ListenAndServe(":8080", nil)
}
