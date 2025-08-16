package main

import "net/http"
import "server/controllers"
import "log"
import  "fmt"


// this is the server that handel all the functions 
func main(){
	http.HandleFunc("/", controllers.Handler)
	http.HandleFunc("/authintication/login" , controllers.Login)
	fmt.Println("Server Is Running On Port 3000")

	err := http.ListenAndServe(":3000",nil)
	if err != nil {
		log.Fatal("Error")
		return
	}

}
