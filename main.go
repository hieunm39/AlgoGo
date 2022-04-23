package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

}

func main() {

	

	http.HandleFunc("/hello", hello);

	http.ListenAndServe(":8090", nil);
	
}	

