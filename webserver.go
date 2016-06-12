package main

import ("fmt"
"net/http"
)

func main(){

	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)

	//port to listen on
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Hello World\n")	
	fmt.Fprintf(w, "Hello World2\n")	

}


func handler2(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Hello Earth\n")	

}