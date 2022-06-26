package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/bad", getBad).Methods(http.MethodGet)
	router.HandleFunc("/name/{param}", getName).Methods(http.MethodGet)
	router.HandleFunc("/data", postData).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func postData(w http.ResponseWriter, r *http.Request) {
	if data, err := io.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Fprintf(w, "I got message:\n"+string(data))
	}
}

func postHeaders(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		w.WriteHeader(http.StatusOK)
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusOK)
	}
	sum := strconv.Itoa(a + b)
	w.Header().Set("a+b", sum)
}

func getName(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["param"]
	fmt.Fprintf(w, "Hello, "+param+"!")
}

func getBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
