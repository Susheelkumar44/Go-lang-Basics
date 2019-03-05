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

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO public."User"`).WithArgs("susheel1","susheel@gmail.com","45678").WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	if err = Signup("susheel1","susheel@gmail.com","45678");err!=nil{
		t.Errorf("error was not expected while inserting the stats: %s",err)
	}

	/*if err := mock.ExpectationsWereMet(); err!=nil {
		t.Errorf("there were unfilled expectations: %s",err)
	}*/
}
