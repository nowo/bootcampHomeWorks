package main

// Call http://localhost:8080/greet/
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("server running on 8080")
	r := mux.NewRouter()
	r.HandleFunc("/greet/{name}", greet).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	var name = vars["name"]
	w.Write([]byte("Hello " + name + "!"))
}
