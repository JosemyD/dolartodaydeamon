package main

import (
	"dolartodaydeamon/controller"
	"dolartodaydeamon/model"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	c := session.DB("dolartoday").C("indicadores")

	defer session.Close()

	indicador := model.Indicadores{}
	controller.GetJson("https://s3.amazonaws.com/dolartoday/data.json", &indicador)

	withMetadata =
	err = c.Insert(&indicador)

	if err != nil {
		log.Fatal(err)
	}


/*	foo2 := model.DolarToday{}
	controller.GetJson("https://s3.amazonaws.com/dolartoday/data.json", &foo2)
	println(foo2.BCV.Reservas)*/
}