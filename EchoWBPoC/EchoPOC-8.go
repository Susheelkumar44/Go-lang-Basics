//Creating JWT Token
package main
/*
NOTE: first run localhost:8000/jwt/main --> You will get Unauthorized or malformed JWT
	Now in Authorization, select No Auth and in headers against Authorization key type in the value Bearer and paste your JWT Token,
	Now try to run localhost:8000/jwt/main --> You will get the access to this page
*/
import(
	"net/http"
	"time"
	"log"
	//"encoding/json"
	"github.com/labstack/echo" 
	"github.com/labstack/echo/middleware" 
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type JwtClaims struct{
	Name		string `json:"name"`
	jwt.StandardClaims				 //Default JWT fields can be accessed with the help of this
}


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

func mainJwt( c echo.Context) error{
	//here we are accessing the Claims i.e, data in the payload
	/* the way echo middleware works here is it will take The JWT Token decode using base64 the token seperated by '.' 
	and put it on to the user object*/
	user := c.Get("user") 
	/*this Get method when used with echo's context 'c' will go to the store use the key and get the interface which nothing but user 
	i.e claims in interface format of map[string]{interface}*/

	/*Now we need to convert into original structure as shown below*/
	token := user.(*jwt.Token) /* and this token is structure of following members
	like raw string, signing method, header, Claims, Signature and whether token is valid or not*/

	//Now we need to access those Claims as done below
	claim := token.Claims.(jwt.MapClaims)

	log.Println("Username: ", claim["name"], "UserID: ",claim["jti"])
	/*Note in above line claim[name] is an interface, if you need a string you need to typecast it as claim["name"].(string)
	So any datatype you need as per the claims data, you need to type cast it explicitly ex: integer etc.. */

	return c.String(http.StatusOK,"You are on the top secret jwt page!!")
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

		//Using JWT Token
		token, err := createJwtToken()
		if err!=nil{
			log.Println("Error Creating JWT Token ",err)
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}


		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!",
			"token": token,
		})
	}

	return c.String(http.StatusUnauthorized,"Your username or password doesn't match!!")
}

func createJwtToken()(string,error){
	/* In order create JWT Token, we need to create claims which are nothing but data in the payload (refer JWT Website), 
	for that you need to create a structure which holds claims, here in our code structure is JwtClaims*/
	claims := JwtClaims{
		"susheel",
		jwt.StandardClaims{ //default structure which has various fields, here we include some sample fields of it
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	//Creating a JWT token
	/*NewWithClaims method takes two arguments 
		1. Signing method
		2. Which data to Sign*/
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims) 

	//Hashing the token with the secret key.
	/*Note: Here we are entrying the key in the form of string ex: mySecret
		  	but never try doing that, instead load a secret key from where it is 
		  	stored secretly.	*/
	
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err!=nil{
		return "",err
	}
	return token, nil
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
		if cookie.Value == "some_string" { 
			return next(c)
		} 
		return c.String(http.StatusUnauthorized,"You don't have right cookie!!")
	}
}



func main(){
	e:= echo.New() 
	adminGroup:=e.Group("/admin")
	cookieGroup:=e.Group("/cookie")
	jwtGroup:=e.Group("/jwt") 


	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${remote_ip} ${host}${path} ${uri} ${error} ${latency_human} `+"\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "susheel" && password == "12345" {
			return true, nil
		}
		return false, nil
	}))

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		//This middleware will compare whether same signing method and key are used to generate JWT Token
		SigningMethod: "HS512", 
		SigningKey: []byte("mySecret"),
	}))

	cookieGroup.Use(checkCookie)

	adminGroup.GET("/main",mainAdmin)

	cookieGroup.GET("/main",mainCookie)

	jwtGroup.GET("/main",mainJwt)
	
	e.GET("/login",login)

	e.Start(":8000") 
}