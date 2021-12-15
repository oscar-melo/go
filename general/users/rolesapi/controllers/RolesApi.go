package controllers

import (
	"github.com/gorilla/mux"
	"general/users/rolesapi/business"
    "fmt"
    "log"
    "net/http"

)



// using gorilla
func HandleRequests() {

    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", business.ReturnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", business.ReturnSingleArticle).Methods("GET")

	myRouter.HandleFunc("/articles", business.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", business.DeleteArticle).Methods("DELETE")

	myRouter.HandleFunc("/articles/{id}", business.UpdateArticle).Methods("PUT")

    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":8003", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}