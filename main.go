package main

import (
	"dolartodaydeamon/controller"
	"dolartodaydeamon/model"
	"flag"
	"github.com/vharitonsky/iniflags"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"time"
	"log"

	"github.com/mitchellh/hashstructure"
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
	colIndicadores := db.C(*dbColIndicadores)

	if err != nil {
		panic(err)
	}

	//idg := controller.NewIDGenerator(db)

	for 1 == 1 {
		dolarToday := model.Indicadores{}
		controller.GetJson(*dolarTodayUrl, &dolarToday.DOLARTODAY)

		var results model.Indicadores
		err = colIndicadores.Find(nil).Sort("-metadata.n").One(&results)

		if err != nil {
			panic(err)
		}

		/*max := 1
		var result model.Indicadores
		for _, v := range results {
			if v.METADATA.Secuencia >= max{
				result = v
				max = v.METADATA.Secuencia
				fmt.Println(max)
			}
		}*/

		log.Printf("BD Secuencia: %d", results.METADATA.Secuencia)
		log.Printf("BD: %+v", results.DOLARTODAY)

		hashBD, err := hashstructure.Hash(results.DOLARTODAY, nil)
		if err != nil {
			panic(err)
		}

		log.Printf("Hash BD: %d", hashBD)


		log.Printf("DT: %+v", dolarToday.DOLARTODAY)

		hashDT, err := hashstructure.Hash(dolarToday.DOLARTODAY, nil)
		if err != nil {
			panic(err)
		}

		log.Printf("Hash DT: %d", hashDT)

		/*n, err := idg.Next("indicadores")
		if err != nil {
			panic(err)
		}
		withMetadata := model.Metadata{}
		withMetadata.Metadata(n)
		dolarToday.AgregarMetadata(withMetadata)
		err = colIndicadores.Insert(&dolarToday)
		if err != nil {
			panic(err)
		}*/
		time.Sleep(7 * time.Second)
	}
}
