package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!\n")
	/* var input1 string = "Hello"*/
	var input1 string
	/* var input2 string = "First"*/
	var input2 string
	fmt.Fprintf(w, "\nEntered 1st value is:")
	fmt.Scanln(&input1)
	fmt.Fprintf(w, input1)
	fmt.Fprintf(w, "\nEntered 2nd value is:")
	fmt.Scanln(&input2)
	fmt.Fprintf(w, input2)
	if input1 == input2 {
		fmt.Fprintf(w, "Equal\n")
	} else {
		fmt.Fprintf(w, "\nUnEqual")
	}
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
