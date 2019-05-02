package main
 
import (
    "fmt"
    "net/http"
	
    "log"
	//ms "./email"
	_ "github.com/lib/pq"
    "os"
   	"encoding/json"
    "strconv"
	"strings"
	"github.com/labstack/echo"
    consulapi "github.com/hashicorp/consul/api"
)


type event struct{
    Email string `json:"email"`
    Subject string `json:"subject"`
    Description string `json:"description"`
    Location string `json:"location"`
    StartDateTime string `json:"StartDateTime"`
    EndDateTime string `json:"EndDateTime"`
}


type JwtToken struct {
	Token string `json:"token"`
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "email-service"
	registration.Name = "email-service"
	address := hostname()
	registration.Address = address
	port, err := strconv.Atoi(port()[1:len(port())]) 
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = port
	consul.Agent().ServiceRegister(registration)
}

func port() string {
	p := os.Getenv("EMAIL_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8083"
	}
	return fmt.Sprintf(":%s", p)
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}

func main() {
	e:= echo.New()
    registerServiceWithConsul()
	fmt.Printf("user service is up on port: %s", port())
	e.POST("/GetEvent",GetEvent)
	e.Start(":8083")
    
}



//Sending an ICS as email

func lookupServiceWithConsul(serviceName string) (string, error) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := consul.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services["event-service"]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

func GetEvent(c echo.Context) error {
	e:= event{}
	url, err := lookupServiceWithConsul("email-service")
	fmt.Println("URL: ", url)
	if err != nil {
		return c.String(http.StatusInternalServerError,"Something went wrong1!")
	}
	client := &http.Client{}
	resp, err := client.Get(url + "/ListEvent")
	if err != nil {
		return c.String(http.StatusInternalServerError,"Something went wrong2!")
	}
	defer c.Request().Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&e) ; err != nil {
		return c.String(http.StatusInternalServerError,"Something went wrong3!")
	}

	//fmt.Printf("Check")

   
    email:=e.Email
    name := c.FormValue("name")
    subject := c.FormValue("subject")
    message := c.FormValue("message")
    //ms.SendEmail(name,email,subject,message)
	fmt.Println(email,name,subject,message)
    var file, err1 = os.Create(`calendar1.ics`)
    defer file.Close()
    fmt.Fprintf(file,"BEGIN:VCALENDAR\nMETHOD:PUBLISH\nVERSION:2.0\nPRODID:-//Company Name//Product//Language\nBEGIN:VEVENT")
    fmt.Fprintf(file,"\nSUMMARY:")
    fmt.Fprintf(file,e.Subject)
    fmt.Fprintf(file,"\nDTSTART:")
    fmt.Fprintf(file,e.StartDateTime)
    fmt.Fprintf(file,"\nDTEND:")
    fmt.Fprintf(file,e.EndDateTime)
    fmt.Fprintf(file,"\nDESCRIPTION:")
    fmt.Fprintf(file,e.Description)
    fmt.Fprintf(file,"\nLOCATION:")
    fmt.Fprintf(file,e.Location)
    fmt.Fprintf(file,"\nEND:VEVENT\nEND:VCALENDAR")   
    if err1 != nil {
        fmt.Println(err1)
	}
	return c.String(http.StatusOK,fmt.Sprintf("Your event details: %s", e))

}