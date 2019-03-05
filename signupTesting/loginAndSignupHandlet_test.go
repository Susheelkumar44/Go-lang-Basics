package main

import (
	"bytes"
	
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"fmt"
)



func TestSignUphandler(t *testing.T) {

	mockStore := InitMockStore()
	
	mockStore.On("Signup", &Users{"chetajai", "chethan@gmail.com",""}).Return(nil)
	
	form := newSignUpForm() 
	req, err := http.NewRequest("POST", " ", bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
		fmt.Println("Error")
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
