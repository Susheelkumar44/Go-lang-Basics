package mydb
import "log"
import "database/sql"
import "fmt"

type User struct {
	email     string 
	pwd  string  	 
	username string  
	confirm  string  
} 

func Login(email, password string) (*User, error) {
	result := &User{}
	row := db.QueryRow(`
		SELECT USERNAME,EMAIL,PWD,CONFIRM
		FROM public."User"
		WHERE EMAIL = $1 
		  AND PWD = $2`, email, password)
	err := row.Scan(&result.username, &result.email, &result.pwd, &result.confirm)
	log.Printf("Error:%v",err)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	log.Printf("%v has logged in",result.username)

	return result, nil
}

func Signup(username, email, password, confirm string) {
	_,err := db.Exec(`
		INSERT INTO public."User" (USERNAME,EMAIL,PWD,CONFIRM)
		VALUES ($1,$2,$3,$4)`,username, email, password, confirm)
	
	if err != nil {
		log.Printf("Insertion Error : %v",err)
	}else{
		log.Printf("Registered successfully")
	}
}

