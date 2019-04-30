//Adding cookie
package main
/*
NOTE:  First run localhost:8000/cookie/main --> You will notice that there will be no cookie
	   then run localhost:8000/login?name=susheel&password=12345 --> cookie will be created and stored
	   Now you run localhost:8000/cookie/main --> You will see your cookie will be created
*/
import(
	"net/http"
	"time"
	"log"
	"github.com/labstack/echo" 
	"github.com/labstack/echo/middleware" 
	"strings"
)
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc{ 
	return func(c echo.Context) error{ 
		c.Response().Header().Set(echo.HeaderServer,"My Custom name Server/1.0") 
		c.Response().Header().Set("notReallyHeader","meansNothing") 

		return next(c) 
	}
}
func mainAdmin(c echo.Context) error{
	return c.String(http.StatusOK,"You are into admin page")
}
func mainCookie(c echo.Context) error{
	return c.String(http.StatusOK,"You are into secret cookie page")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	
	if username == "susheel" && password == "12345"{
		cookie := &http.Cookie{}
		/* OR
		cookie := new(http.Cookie) this is will also do the same as above
		*/
		cookie.Name = "sessionID"
		cookie.Value = "some_string" 
		
		cookie.Expires = time.Now().Add(48 * time.Hour) 
		c.SetCookie(cookie)
		return c.String(http.StatusOK,"You were logged in!!")
	}

	return c.String(http.StatusUnauthorized,"Your username or password doesn't match!!")
}


func checkCookie(next echo.HandlerFunc)echo.HandlerFunc{
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err!=nil{
			if strings.Contains(err.Error(), "named cookie not present"){
				return c.String(http.StatusUnauthorized,"You don't have the any cookie!!")
			}
			log.Println(err)
			return err
		}
		if cookie.Value == "some_string" { //here you need to pull seesionID from session store to get the user data
			return next(c)
		} 
		return c.String(http.StatusUnauthorized,"You don't have right cookie!!")
	}
}


func main(){
	e:= echo.New() 
	adminGroup:=e.Group("/admin")
	cookieGroup:=e.Group("/cookie")
	cookieGroup.Use(checkCookie)

	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${remote_ip} ${host}${path} ${uri} ${error} ${latency_human} `+"\n",
	}))

	
		adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username == "susheel" && password == "12345" {
				return true, nil
			}
			return false, nil
		}))
	adminGroup.GET("/main",mainAdmin)
	cookieGroup.GET("/main",mainCookie)
	
	e.GET("/login",login)

	e.Start(":8000") 
}