package controllers

import (
	"github.com/revel/revel"
)

type Orders struct {
	*revel.Controller
}

func (c Orders) Create() revel.Result {
	return c.Render()
}

func (c Orders) GetPaymentId(orderId int) revel.Result {
	println("Order Id is: ", orderId)
	return c.RenderTemplate("orders/payment.html")
}

func (c Orders) testforcatchall() revel.Result{
	return c.Render()
}