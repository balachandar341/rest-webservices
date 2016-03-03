package main
import (

	"rest-service/content"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	
	router.HandleFunc("/",content.Hello).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
