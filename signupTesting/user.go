package main
import "log"
import "database/sql"
import "fmt"

type User interface{
	Login(user *Users) (*Users,error)
	Signup(user *Users)
}
type dbUser struct {
	db *sql.DB
}

func(u *dbUser) Login(user *Users) (*Users,error) {  
	result := &Users{}
	row := u.db.QueryRow(`
		SELECT USERNAME,EMAIL,PWD
		FROM public."User"
		WHERE EMAIL = $1 
		  AND PWD = $2`, user.email, user.pwd)

	err := row.Scan(&result.username, &result.email, &result.pwd)
	log.Printf("Error:%v",err)
	switch {
	case err == sql.ErrNoRows:
		return nil,fmt.Errorf("User not found")
	case err != nil:
		return  nil,err
	}
	log.Printf("%v has logged in",result.username)

	return result,nil
}

func (u *dbUser) Signup(user *Users) {
	_,err := u.db.Exec(`
		INSERT INTO public."User" (USERNAME,EMAIL,PWD)
		VALUES ($1,$2,$3)`,user.username, user.email, user.pwd) 
	
	if err != nil {
		log.Printf("Insertion Error : %v",err)
	}else{
		log.Printf("Registered successfully")
	}
	}
	var u User
	func InitStore(us User) {
		u=us
}

