package main

import (
	"dolartodaydeamon/controller"
	"dolartodaydeamon/model"
	"flag"
	"github.com/kataras/iris"
	"github.com/mitchellh/hashstructure"
	"github.com/square/go-jose/json"
	"github.com/vharitonsky/iniflags"
	"gopkg.in/mgo.v2"
	"log"
	"time"
	"os"
)

var (
	dolarTodayUrl    = flag.String("dolarTodayUrl", "https://s3.amazonaws.com/dolartoday/data.json", "URL que provee los indicadores de DolarToday")
	dbHost           = flag.String("dbHost", "localhost:27017", "Host de la base de datos")
	dbName           = flag.String("dbName", "db", "Nombre de la base de datos")
	dbUser           = flag.String("dbUser", "", "Nombre del usuario de la base de datos")
	dbPassword       = flag.String("dbPassword", "", "Password del usuario de la base de datos")
	dbSource         = flag.String("dbSource", "", "Base de datos utilizada para establecer credenciales y privilegios")
	dbColIndicadores = flag.String("dbColIndicadores", "indicadores", "Colección para los indicadores")
	minCheckDT       = flag.Int("minCheckDT", 1, "Minutos para revisar DolarToday por cambios")
)

func main() {
	iniflags.Parse()

	mongoDBDialInfo := controller.MongoDBDialInfo(*dbHost, *dbName, *dbUser, *dbPassword, *dbSource)

	db, err := controller.Connect(mongoDBDialInfo)
	defer db.Session.Close()
	colIndicadores := db.C(*dbColIndicadores)

	if err != nil {
		panic(err)
	}

	idg := controller.NewIDGenerator(db)

	go func() {
		for 1 == 1 {
			dolarToday := model.Indicadores{}
			controller.GetJson(*dolarTodayUrl, &dolarToday.DOLARTODAY)

			var results model.Indicadores
			err = colIndicadores.Find(nil).Sort("-metadata.n").One(&results)
			if err == mgo.ErrNotFound || err == nil {

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

				if hashDT != hashBD {
					log.Printf("HASH DT != BD: %d != %d", hashDT, hashBD)
					n, err := idg.Next("indicadores")
					if err != nil {
						panic(err)
					}
					withMetadata := model.Metadata{}
					withMetadata.Metadata(n)
					dolarToday.AgregarMetadata(withMetadata)
					err = colIndicadores.Insert(&dolarToday)
					if err != nil {
						panic(err)
					}

				}
			} else {
				panic(err)
			}
			time.Sleep(time.Duration(*minCheckDT) * time.Minute)
		}
	}()

	iris.Get("/", func(c *iris.Context) {
		var results model.Indicadores
		err := colIndicadores.Find(nil).Sort("-metadata.n").One(&results)
		if err == mgo.ErrNotFound || err == nil {
			out, err := json.MarshalIndent(results, "", "  ")
			if err != nil {
				c.HTML(iris.StatusInternalServerError, err.Error())
			} else {
				c.HTML(iris.StatusOK, string(out))

			}
		} else {
			c.HTML(iris.StatusInternalServerError, err.Error())
		}
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	iris.Listen(":" + port)
}
