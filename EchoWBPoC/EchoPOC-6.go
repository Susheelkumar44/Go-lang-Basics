//Adding  Custom middleware with pre-defined one
package main

import(
	"net/http"
	"github.com/labstack/echo" //echo web framework package
	"github.com/labstack/echo/middleware" //middleware package
)

func mainAdmin(c echo.Context) error{
	return c.String(http.StatusOK,"Your into admin page")
}

// Creating a Custom middleware
/*The middleware created below just adds the server name to the incoming request*/
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc{ //Custom middleware takes handlerfunc as argument and returns the same
	return func(c echo.Context) error{ //passing echo context as argument
		c.Response().Header().Set(echo.HeaderServer,"My Custom name Server/1.0") //Echo built server headers
		c.Response().Header().Set("notReallyHeader","meansNothing") //Any thing you can add to your server header

		return next(c) //returning the handler func
	}
}



func main(){
	e:= echo.New() 

	e.Use(ServerHeader) //a
	g:=e.Group("/admin")

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${remote_ip} ${host}${path} ${uri} ${error} ${latency_human} `+"\n",
	}))

		g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username == "susheel" && password == "12345" {
				return true, nil
			}
			return false, nil
		}))
	g.GET("/main",mainAdmin)

	
	e.Start(":8000") 
}