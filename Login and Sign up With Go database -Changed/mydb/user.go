package mydb
import "log"
import "database/sql"
import "fmt"



func Login(user *Users) error {  
	result := &Users{}
	row := db.QueryRow(`
		SELECT USERNAME,EMAIL,PWD,CONFIRM
		FROM public."User"
		WHERE EMAIL = $1 
		  AND PWD = $2`, user.email, user.pwd)

	err := row.Scan(&result.username, &result.email, &result.pwd, &result.confirm)
	log.Printf("Error:%v",err)
	switch {
	case err == sql.ErrNoRows:
		return fmt.Errorf("User not found")
	case err != nil:
		return  err
	}
	log.Printf("%v has logged in",result.username)

	return result
}

func Signup(user *Users) {
	_,err := db.Exec(`
		INSERT INTO public."User" (USERNAME,EMAIL,PWD,CONFIRM)
		VALUES ($1,$2,$3,$4)`,user.username, user.email, user.pwd, user.confirm)
	
	if err != nil {
		log.Printf("Insertion Error : %v",err)
	}else{
		log.Printf("Registered successfully")
	}
}

