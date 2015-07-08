package controllers

import "github.com/revel/revel"

type Ven struct {
	*revel.Controller
}


func (c Ven) Index(myName string) revel.Result {
	return c.Render(myName)
}