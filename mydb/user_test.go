package mydb

import(
	
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSignUpFunction(t *testing.T){
 	db, mock, err := sqlmock.New()
	if err!=nil {
		t.Fatalf("an Error %s was not expected when opening a stub database connection",err)
	} 
	defer db.Close()
	username :="nikil"
	email :="nikil@gmail.com"
	password:="45678"

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO public."User"`).WithArgs(username,email,password).WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	if err=Signup(db,"nikil","nikil@gmail.com","45678");err!=nil{
		t.Errorf("error was not expected while inserting the stats: %s",err)
	}

}
