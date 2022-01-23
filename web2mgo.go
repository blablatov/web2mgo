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

func ODatareq(Name, Data string) {
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
	err = c.Find(bson.M{"name": Name}).One(&chk)
	if err == nil {
		log.Print("Name already is to DB", err)
	}
	if err != nil {
		log.Print("Data for write to DB:", err)
		err = c.Insert(&Mongo{Name, Data})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Name was written:", Name, "\nData was written:", Data)
	}
}
