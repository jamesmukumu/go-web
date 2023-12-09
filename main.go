package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)


func main(){
r:= mux.NewRouter()

r.HandleFunc("/test",Testcontroller).Methods("GET")





// http.Handle("/",r)



//listen to requests
fmt.Println("Server listening at 7000")

log.Fatal(http.ListenAndServe(":7000",r))

}




// create a controller to handle a request


func Testcontroller(w http.ResponseWriter, r*http.Request){
w.Write([]byte("hello world put this up"))


}