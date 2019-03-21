package mydb

import (
	"log"
	
)


func CreateTable() (error) {
	flag := 0
	_,err := db.Exec(`
	CREATE TABLE public."events"(
    "EMAIL" text COLLATE pg_catalog."default" NOT NULL,
    "SUBJECT" text COLLATE pg_catalog."default" NOT NULL,
    "STARTDATETIME" timestamp with time zone,
    "ENDDATETIME" timestamp with time zone,
    "DESCRIPTION" text COLLATE pg_catalog."default",
    "LOCATION" text COLLATE pg_catalog."default"
)`)
	if err != nil {
		log.Printf("Creation Error : %v",err)
		flag++
	}else{
		log.Printf("Table Created successfully") 
	}
	return err
}




//Insert event details into database
func AddEvent(email,name,date,time,description,location string) (int){
	_,err := db.Exec(`
		INSERT INTO public."events" ("EMAIL","SUBJECT","STARTDATETIME","ENDDATETIME","DESCRIPTION","LOCATION")
		VALUES ($1,$2,$3,$4,$5,$6)`,email,name,date,time,description,location)
	
	if err != nil {
		log.Printf("Insertion Error : %v",err)
		return 1
	}else{
		log.Printf("Added successfully")
		return 0
}
}


