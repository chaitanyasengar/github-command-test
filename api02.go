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
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Req: %s %s", r.URL.Host, r.URL.Path)
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

/*package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	txt, _ := reader.ReadString('\n')
	fmt.Println(reflect.TypeOf(txt)) // Get object type
	//url := fmt.Sprintf("http://%s",txt)
	url := "http://google.com"
	fmt.Println(reflect.TypeOf(url)) // Get object type
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", resp.Status)
	}
}*/
