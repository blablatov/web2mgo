package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	Name string
	Data string
}

func ODatareq(Name, Data string, done chan bool) {
	session, err := mgo.Dial("mongodb://localhost:27017/testdb")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// is check name in dBase
	c := session.DB("testdb").C("name")
	chk := Mongo{}
	err = c.Find(bson.M{"name": Name, "data": Data}).One(&chk)
	if err == nil {
		log.Print("Name already is to DB", err)
		done <- false
	}
	if err != nil {
		log.Print("Data for write to DB:", err)
		err = c.Insert(&Mongo{Name, Data})
		if err != nil {
			log.Fatal(err)
			done <- false
		}
		fmt.Println("Name was written:", Name, "\nData was written:", Data)
		done <- true
	}
}
