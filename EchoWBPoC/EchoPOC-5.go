//Adding  Basic Auth middleware
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
	e:= echo.New() 

	g:=e.Group("/admin")

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${remote_ip} ${host}${path} ${uri} ${error} ${latency_human} `+"\n",
	}))

	/*BasicAuth middleware take a call back function with 
		credentials as arguments and returns bool based on validation*/
		g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username == "susheel" && password == "12345" {
				return true, nil
			}
			return false, nil
		}))
	g.GET("/main",mainAdmin)

	
	e.Start(":8000") 
}