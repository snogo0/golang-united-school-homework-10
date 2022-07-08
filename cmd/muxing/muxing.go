package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"io/ioutil"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func hello(wr http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)["PARAM"]
	wr.WriteHeader(http.StatusOK)
	fmt.Fprintf(wr, "Hello, %s!", p)
}

func bad(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusInternalServerError)
}

func gotmess(wr http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
	wr.WriteHeader(http.StatusOK)
	fmt.Fprintf(wr, "I got message:\n%s", string(b))
}

func sum(wr http.ResponseWriter, req *http.Request) {
	a, err := strconv.Atoi(req.Header.Get("a"))
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := strconv.Atoi(req.Header.Get("b"))
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
	wr.Header().Add("a+b", strconv.Itoa(a+b))
	wr.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", hello).Methods(http.MethodGet)
	router.HandleFunc("/bad", bad).Methods(http.MethodGet)
	router.HandleFunc("/data", gotmess).Methods(http.MethodPost)
	router.HandleFunc("/headers", sum).Methods(http.MethodPost).HeadersRegexp("a", "[0-9].*", "b", "[0-9].*")

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
