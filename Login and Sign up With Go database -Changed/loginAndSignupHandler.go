package main
 
import (
    "fmt"
	"net/http"
	"log"

	mydb "./mydb"
)

type Users struct {
	username    string `json:"username"`
	email string `json:"email"`
	pwd string `json:"pwd"`
	confirm string `json:"confirm"`
}

func signuphandler(w http.ResponseWriter, r *http.Request) {
	user := &Users{}
	err := r.ParseForm()
	if err!=nil{
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.username=r.Form.Get("username")
	user.email=r.Form.Get("email")
	user.pwd=r.Form.Get("password")
	user.confirm=r.Form.Get("confirm")

	mydb.Signup(&user)
	
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uName, email, pwd, pwdConfirm := "", "", "", ""

	email = r.FormValue("email")  // Data from the form
	pwd = r.FormValue("password") // Data from the form

	// Empty data checking
	

	if user, err:= mydb.Login(email, pwd); err == nil {
		log.Printf("%v\n", user)
	
	  
		return
	} else {
		log.Printf("Failed to log user in with email: %v %v, error was: %v\n", email,pwd, err)
	}
}