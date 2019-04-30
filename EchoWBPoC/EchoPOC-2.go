//basic get and post 
package main

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo" //echo web framework package
)

/*func sayHello(c echo.Context ) error{ 
	return c.String(http.StatusOK, "This is my first echo based go program")
}*/

func getBookDetails(c echo.Context /* echo context parameter*/) error{
	bookName := c.QueryParam("name") //to extract query parameter from query string
	bookType := c.QueryParam("type")

	dataType := c.Param("data") //to extract path parameter from query string

	if dataType =="string"{
		return c.String(http.StatusOK,fmt.Sprintf("Your book name is: %s\nYour book type is: %s\n", bookName, bookType))
	}

	if dataType =="JSON" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": bookName,
			"type": bookType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "You have to select the data type as either string or JSON",
	})
}

func postBookDetails(c echo.Context) error{
	bookName := c.FormValue("name")
	bookType := c.FormValue("type")

	return c.String(http.StatusOK," bookname: "+bookName+" booktype: "+bookType)
}

func main(){
	e:= echo.New() //initation of echo object

	//e.GET("/",sayHello)
	e.GET("/bookDetails/:data",getBookDetails) //mapping route to specific controller method
	e.POST("/postbooks",postBookDetails)
	e.Start(":8000") //start the server on 8000 port
}