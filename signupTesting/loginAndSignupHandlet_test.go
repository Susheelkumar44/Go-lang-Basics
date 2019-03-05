package main

import (
	"bytes"
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"fmt"
)

/*func TestGetLoginHandler(t *testing.T){
	mockStore := InitMockStore()
	mockStore.On("Login").Return(Users{"chetajai","chethan@gmail.com","12345"},nil).Once()
	req, err := http.NewRequest("POST", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(loginhandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := Users{"chetajai","chethan@gmail.com","12345"}
	u := []Users{}
	err = json.NewDecoder(recorder.Body).Decode(&u)

	if err != nil {
		t.Fatal(err)
	}

	actual := u[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// the expectations that we defined in the `On` method are asserted here
	mockStore.AssertExpectations(t)
}*/

func TestSignUphandler(t *testing.T) {

	mockStore := InitMockStore()
	
	mockStore.On("Signup", &Users{"chetajai", "chethan@gmail.com",""}).Return(nil)
	fmt.Println("111111")
	form := newSignUpForm() 
	req, err := http.NewRequest("POST", " ", bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
		fmt.Println("1111")
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(signuphandler)
	hf.ServeHTTP(recorder, req)
 
	if status := recorder.Code; status == http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func newSignUpForm() *url.Values{
	form := url.Values{}
	form.Set("username", "chetajai")
	form.Set("email", "chethan@gmail.com")
	form.Set("pwd","12345")
	
	return &form

}
