package controllers

import (
	"github.com/revel/revel"
	models "myapplication/app/models"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Account() revel.Result {
	return c.Render()
}

func (c App) Create() revel.Result {
	return c.RenderTemplate("App/CreateApp.html")
}

func (c App) CreatePost() revel.Result {
	var account models.Account
	c.Params.BindJSON(&account)
	fmt.Printf("Account info: %v\n", account)
	return c.RenderTemplate("App/CreateApp.html")
}
