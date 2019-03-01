package main
 
import (
    "fmt"
	"net/http"
	"log"
   // "encoding/json"

	
)

type Users struct {
	username string 
	email string 
	pwd string 
}

func signuphandler(w http.ResponseWriter, r *http.Request) {
	user := Users{}
	err := r.ParseForm()
	if err!=nil{
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.username=r.Form.Get("username")
	user.email=r.Form.Get("email")
	user.pwd=r.Form.Get("password")
	

	u.Signup(&user) 
	
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	
	user := &Users{}
	err := r.ParseForm()
	if err!=nil{
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.email=r.Form.Get("email")
	user.pwd=r.Form.Get("password")
	if user, err:= u.Login(user); err==nil{
		log.Printf("%v\n", user)
	
	  
		return
	} else {
		log.Printf("Failed to log user in with email: %v %v, error was: %v\n", user.email,user.pwd, err)	
	}
}