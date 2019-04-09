package controllers

import (
	"github.com/revel/revel"
)

type Events struct {
	*revel.Controller
}

func (e Events) Create() revel.Result {
	return e.Render()
}