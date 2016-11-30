package main

import (
	"dolartodaydeamon/controller"
	"dolartodaydeamon/model"
	"flag"
	"fmt"
	"github.com/vharitonsky/iniflags"
)

var (
	dolarTodayUrl    = flag.String("dolarTodayUrl", "https://s3.amazonaws.com/dolartoday/data.json", "URL que provee los indicadores de DolarToday")
	dbHost           = flag.String("dbHost", "localhost", "Host de la base de datos")
	dbPort           = flag.String("dbPort", "27017", "Puerto de la base de datos")
	dbName           = flag.String("dbName", "db", "Nombre de la base de datos")
	dbColIndicadores = flag.String("dbColIndicadores", "indicadores", "Colección para los indicadores")
)

func main() {
	iniflags.Parse()

	db, err := controller.Connect(dbHost, dbPort, dbName)
	defer db.Session.Close()

	if err != nil {
		panic(err)
	}

	dolarToday := model.Indicadores{}
	controller.GetJson(*dolarTodayUrl, &dolarToday.DOLARTODAY)

	idg := controller.NewIDGenerator(db)

	for i := 0; i < 100; i++ {
		n, err := idg.Next("indicadores")
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
		withMetadata := model.Metadata{}
		withMetadata.Metadata(n)
		dolarToday.AgregarMetadata(withMetadata)
		err = db.C(*dbColIndicadores).Insert(&dolarToday)
		if err != nil {
			panic(err)
		}
	}
}
