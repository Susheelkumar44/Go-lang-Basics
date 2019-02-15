package main
 
import (
    "fmt"
	"net/http"
    "log"
	_ "github.com/lib/pq"
	"database/sql"
    mydb "./mydb"
    helper "./helpers"
    
)
func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres","postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database %v",err))
	}
	mydb.SetDatabase(db)
	return db
}
func main() {
 
	uName, email, pwd, pwdConfirm := "", "", "", ""
	db := connectToDatabase()
	defer db.Close()
    mux := http.NewServeMux()
 
    // Signup
    mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
 
        /*
            // Available for testing.
            for key, value := range r.Form {
                fmt.Printf("%s = %s\n", key, value)
            }
        */
 
        uName = r.FormValue("username")     // Data from the form
        email = r.FormValue("email")        // Data from the form
        pwd = r.FormValue("password")       // Data from the form
        pwdConfirm = r.FormValue("confirm") // Data from the form
 
        // Empty data checking
        uNameCheck := helper.IsEmpty(uName)
        emailCheck := helper.IsEmpty(email)
        pwdCheck := helper.IsEmpty(pwd)
        pwdConfirmCheck := helper.IsEmpty(pwdConfirm)
 
        if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
            fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
            return
        }
 
        if pwd == pwdConfirm {
            // Save to database (username, email and password)
            mydb.Signup(uName,email,pwd,pwdConfirm)
        } else {
            fmt.Fprintln(w, "Password information must be the same.")
        }
    })
 
    // Login
    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
 
        email = r.FormValue("email")  // Data from the form
        pwd = r.FormValue("password") // Data from the form
 
        // Empty data checking
        emailCheck := helper.IsEmpty(email)
        pwdCheck := helper.IsEmpty(pwd)
 
        if emailCheck || pwdCheck {
            fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
            return
        }

        if user, err:= mydb.Login(email, pwd); err == nil {
            log.Printf("%v\n", user)
        
          //log.Printf("%T",user)
			//http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("Failed to log user in with email: %v %v, error was: %v\n", email,pwd, err)
		}
    })
    
    http.ListenAndServe(":8080", mux)
}