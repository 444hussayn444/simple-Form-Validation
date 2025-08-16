package controllers

import "net/http"
import "encoding/json"
import "fmt"
import "os"
import "strings"
import "log"


// load file and return it as http handler
func load_html() http.Handler {
	file_served := http.FileServer(http.Dir("../frontend"))
	return file_served
}

// show all users stored in the users.txt file 

// return respose with json data to the user that logged in
func handel_text_base( res http.ResponseWriter ,req *http.Request ,data map[string]string){
	if req.Method == http.MethodPost	{
		filename := "users.txt"
		created_file , err := os.OpenFile(filename,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
		defer created_file.Close()

		if err != nil{
			fmt.Print("Can not open file")
		}
		for key , val := range data {
			_, err := fmt.Fprintf(created_file,"%s : %s \n",key,val)
			if err != nil {
				fmt.Println("Error in writing.")
			}	
		}

		res.Header().Set("Content-Type","application/json")
		json.NewEncoder(res).Encode(&data)
		log.Printf("data: %s", data)
	}
}


func Check_errors(res http.ResponseWriter,req *http.Request,Name string , Password string , Email string) bool {
		if req.Method == http.MethodPost{
	     	if len(Name) < 1 {
				err_msg:= map[string]string{
					"message":"Please Enter A Valid Name",
				}
				res.Header().Set("Content-Type","application/json")
				res.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(res).Encode(err_msg)
				return false
			}
			if len(Password) < 8 {
				err_msg := map[string]string{
					"message":"Please enter a strong password",
				}
				res.Header().Set("Content-Type","application/json");
				res.WriteHeader(http.StatusBadRequest);
				json.NewEncoder(res).Encode(err_msg);
				return false
			}
			if !strings.Contains(Email,"@"){
				err_msg := map[string]string{
					"message":"Please enter a valid email",
				}
				res.Header().Set("Content-Type","application/json");
				res.WriteHeader(http.StatusBadRequest);
				json.NewEncoder(res).Encode(err_msg);
				return false
			}
	}
	return true
}

// login  handler and executes all of the last functions 
func Login(res http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodPost{
		var user Users
		// convert json to map
		err := json.NewDecoder(req.Body).Decode(&user)

		if err != nil {
			http.Error(res,"Invalid Json Format",http.StatusBadRequest)
			return 
		}

		// handling user data from here => (password, email, username)
	     response := map[string]string{
			"name":user.Name,
			"email":user.Email,
			"password":user.Password,
		}
		if !Check_errors(res,req,user.Name,user.Password,user.Email){
			log.Println("Error In Credinitals")
			return
		}else {

		
	    handel_text_base(res,req,response)
		}
	}else {
		http.Error(res,"Invalid Endpoint",http.StatusBadRequest)
	}
}

// this is the handler that preview the html file from the server
func Handler(res http.ResponseWriter,req *http.Request) {
	if req.Method == http.MethodGet{
		file := load_html()
		file.ServeHTTP(res,req)
	}else{
		http.Error(res,"Invalid Endpoint",http.StatusBadRequest)
	}
}





