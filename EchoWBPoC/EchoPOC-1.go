//Introduction
package main

import(
	//"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func sayHello(c echo.Context) error{
	return c.String(http.StatusOK, "This is my first echo based go program")
}

func main(){
	e:= echo.New()

	e.GET("/",sayHello)
	e.Start(":8000")
}