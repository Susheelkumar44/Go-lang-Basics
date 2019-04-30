//Parsing JSON Data
package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/labstack/echo" //echo web framework package
)

type Book struct{
	Name string `json:"name"`
	Type string `json:"type"`
}

type Paper struct{
	Name string `json:"name"`
	Type string `json:"type"`
}

type CD struct{
	Name string `json:"name"`
	Type string `json:"type"`
}

/*first and second methods are default methods provided by plain go to parse JSON data
third method is echo provided method to parse the data*/

// one type of parsing JSON data
 func addBooks(c echo.Context) error{
	book := Book{} //Create an object of book structure
	
	defer c.Request().Body.Close() //Closing the body when the function scope exits 
	body, err := ioutil.ReadAll(c.Request().Body) //returns Request body and error

	if err!=nil{
		log.Printf("Failed in reading the request body in addBook method: %s", err) 
		return c.String(http.StatusInternalServerError,"")
	}

	err = json.Unmarshal(body, &book) //unmarshal- parses the json data read in the body and store it in the object passed for example: &book object
	if err!=nil{
		log.Printf("Failed in unmarshal the request body in addBooks method: %s", err)
		return c.String(http.StatusInternalServerError,"") //sends error string
	}
	log.Printf("this is your book added: %s",book) //here you can added it to database
	return c.String(http.StatusOK,fmt.Sprintf("We have added your book successfully\nYour added book details: %s", book))
 }

//second type of parsing JSON data
func addPapers(c echo.Context) error{
	paper := Paper{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&paper) //decodes the request body and stores it into object
	if err!=nil{
		log.Printf("Failed processing request body in addPapers method: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError) //sends error object with help of echo package
	}
	log.Printf("this is your paper added: %s",paper) //here you can added it to database
	return c.String(http.StatusOK,fmt.Sprintf("We have added your paper successfully\nYour added paper details: %s", paper))
}

//Third method using Echo package's bind method to parse JSON Data
func addCDs(c echo.Context) error{
	cd := CD{}

	err:= c.Bind(&cd)
	if err!=nil{
		log.Printf("Failed processing request body in addPapers method: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError) //sends error object with help of echo package
	}
	log.Printf("this is your cd added: %s",cd) //here you can added it to database
	return c.String(http.StatusOK,fmt.Sprintf("We have added your cd successfully\nYour added cd details: %s", cd))
}

/*
In terms efficiency: 1st and 2nd method are more efficient than 3rd method but for smaller data reading 3rd method is used.
*/

func main(){
	e:= echo.New() //initation of echo object

	e.POST("/addbook",addBooks)
	e.POST("/addpaper",addPapers)
	e.POST("/addcd",addCDs)
	e.Start(":8000") //start the server on 8000 port
}