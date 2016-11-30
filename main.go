package main

import (
	"dolartodaydeamon/controller"
	"dolartodaydeamon/model"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	db := session.DB("dolartoday")
	c := db.C("indicadores")

	defer session.Close()

	dolarToday := model.Indicadores{}
	controller.GetJson("https://s3.amazonaws.com/dolartoday/data.json", &dolarToday.DOLARTODAY)

	idg := controller.NewIDGenerator(db)

	for i := 0; i < 100; i++ {
		n, err := idg.Next("my-document")
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
		withMetadata := model.Metadata{}
		withMetadata.Metadata(n)
		dolarToday.AgregarMetadata(withMetadata)
		err = c.Insert(&dolarToday)
		if err != nil {
			panic(err)
		}
	}
}
