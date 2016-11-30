package main

import (
	"dolartodaydeamon/controller"
	"dolartodaydeamon/model"
	"flag"
	"fmt"
	"github.com/vharitonsky/iniflags"
	"gopkg.in/mgo.v2"
)

var (
	dolarTodayUrl    = flag.String("dolarTodayUrl", "https://s3.amazonaws.com/dolartoday/data.json", "URL que provee los indicadores de DolarToday")
	dbHost           = flag.String("dbHost", "localhost:27017", "Host de la base de datos")
	dbName           = flag.String("dbName", "db", "Nombre de la base de datos")
	dbUser           = flag.String("dbUser", "", "Nombre del usuario de la base de datos")
	dbPassword       = flag.String("dbPassword", "", "Password del usuario de la base de datos")
	dbSource         = flag.String("dbSource", "", "Base de datos utilizada para establecer credenciales y privilegios")
	dbColIndicadores = flag.String("dbColIndicadores", "indicadores", "Colección para los indicadores")
)

func main() {
	iniflags.Parse()

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{*dbHost},
		Database: *dbName,
		Username: *dbUser,
		Password: *dbPassword,
		Source:   *dbSource,
	}

	db, err := controller.Connect(mongoDBDialInfo)
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
