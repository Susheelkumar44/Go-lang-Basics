//Adding middlewares introduction
package main

import(
	"net/http"
	"github.com/labstack/echo" //echo web framework package
	"github.com/labstack/echo/middleware" //middleware package
)

func mainAdmin(c echo.Context) error{
	return c.String(http.StatusOK,"Your into admin page")
}
func main(){
	e:= echo.New() //initation of echo object

	/*g:= e.Group("/admin",middleware.Logger())
	g.GET("/main",mainAdmin)*/ //first way of adding middleware

	/*g:=e.Group("/admin")
	g.Use(middleware.Logger())
	g.GET("/main",mainAdmin) //second way of adding middleware and best way to use*/

	//refined way of using Logger as per second way
	g:=e.Group("/admin")
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${remote_ip} ${host}${path} ${uri} ${error} ${latency_human} `+"\n",
	}))
	g.GET("/main",mainAdmin)

	// e.GET("/main",mainAdmin,middleware.Logger()) //third way of adding middleware
	e.Start(":8000") //start the server on 8000 port
}