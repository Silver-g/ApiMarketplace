package main

import (
	"fmt"
	"net/http"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Нулевой ендпоинт")
}

func main() {
	var err error

	http.HandleFunc("/", ServerHandler)
	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске")
	}
}
