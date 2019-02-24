package main
 
import (
    "fmt"
	"net/http"
    "log"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
	"database/sql"
    mydb "./mydb"
   
)
func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres","postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database %v",err))
	}
	mydb.SetDatabase(db)
	return db
}

func newRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/signup", signuphandler).Methods("POST")
    r.HandleFunc("/login", signuphandler).Methods("POST")
    return r
 
}

func main() {
 
	db := connectToDatabase()
	defer db.Close()
    mux := http.NewServeMux()
 
    r := newRouter()
    
    http.ListenAndServe(":8080", r)
}