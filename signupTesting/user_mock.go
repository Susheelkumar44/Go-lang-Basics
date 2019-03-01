package main

import (
	"github.com/stretchr/testify/mock"
	"fmt"
)


type MockStore struct {
	mock.Mock
}

func (m *MockStore) Signup(user *Users) {
	
	rets := m.Called(user)
	fmt.Printf("%v",rets)
	
}

func (m *MockStore) Login(user *Users) (*Users,error) {
	rets := m.Called()
	return rets.Get(0).(*Users), rets.Error(1)
}

func InitMockStore() *MockStore {
	/*
		Like the InitStore function we defined earlier, this function
		also initializes the store variable, but this time, it assigns
		a new MockStore instance to it, instead of an actual store
	*/
	us := new(MockStore)
	u = us
	return us
}
