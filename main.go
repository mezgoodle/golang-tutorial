package main

import ("fmt" 
		"net/http")

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Go is a super lang!")
}

func main()  {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":5000", nil)
}
