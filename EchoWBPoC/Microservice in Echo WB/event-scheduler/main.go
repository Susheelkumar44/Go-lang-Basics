package main
 
import (
    "fmt"
    "net/http"
	"database/sql"
    "log"
	config "EchoWBPoC/event-scheduler/config"
	mydb "EchoWBPoC/event-scheduler/mydb"
	//helper "Microservice-service-discovery/event-scheduler//helpers"
	_ "github.com/lib/pq"
    "os"
    "strconv"
    "strings"
    "encoding/json"
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

var mail, sub, desc, loc,sDate, eDate string

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

	registration.ID = "event-service"
	registration.Name = "event-service"
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
	p := os.Getenv("EVENT_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8082"
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
    db := connectToDatabase()
    e.GET("/CreateTable", CreateTable)
    e.POST("/AddEvent",AddEvent)
    e.GET("/ListEvent",ListEvent)
	fmt.Printf("user service is up on port: %s", port())
	e.Start(":8082")
    defer db.Close()
}
//Database connection
func connectToDatabase() *sql.DB {
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    config.HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.PORT)
    db, err := sql.Open("postgres", dbinfo)
    if err != nil {
        fmt.Println(err)
    }
    log.Printf("Postgres started at %s PORT", config.PORT)
    mydb.SetDatabase(db)		
    return db
}

//Inserting events to database
func AddEvent(c echo.Context) error {
    email:= c.FormValue("email")     // Data from the form
    subject := c.FormValue("subject")   // Data from the form
    description := c.FormValue("description")
    location := c.FormValue("location")
    StartDateTime := c.FormValue("StartDateTime")   // Data from the form
    EndDateTime := c.FormValue("EndDateTime") // Data from the form

    /*emailCheck := helper.IsEmpty(email)  //Check if the data is empty to prevent inserting them
    subjectCheck := helper.IsEmpty(subject)
    StartDateTimeCheck := helper.IsEmpty(StartDateTime)
    EndDateTimeCheck := helper.IsEmpty(EndDateTime)
    descriptionCheck := helper.IsEmpty(description)
    locationCheck := helper.IsEmpty(location)

    if emailCheck || subjectCheck || StartDateTimeCheck || EndDateTimeCheck || descriptionCheck || locationCheck{
        fmt.Fprintf(w, "There is empty data.")
        return
    }*/

    e := event{
        Email: email,
        Subject: subject,
        Description: description,
        Location: location,
        StartDateTime: StartDateTime,
        EndDateTime: EndDateTime,
    }
    //fmt.Println(e)
    status:=mydb.AddEvent(email,subject,StartDateTime,EndDateTime,description,location)
    if status==0{
        
        c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
        c.Response().WriteHeader(http.StatusOK)
        mail=email
        sub=subject
        desc=description
        loc=location
        sDate=StartDateTime
        eDate=EndDateTime   
        return json.NewEncoder(c.Response()).Encode(&e)
    } 
    return c.String(http.StatusInternalServerError,"Something went wrong")     
}

func ListEvent(c echo.Context) error{
    list:=event{
        Email: mail,
        Subject: sub,
        Description: desc,
        Location: loc,
        StartDateTime: sDate,
        EndDateTime: eDate,
    }
    c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
    c.Response().WriteHeader(http.StatusOK)
    return json.NewEncoder(c.Response()).Encode(&list)
    
}

func CreateTable(c echo.Context) error{
    mydb.CreateTable()
    return c.String(http.StatusOK,"Table Created")
}
