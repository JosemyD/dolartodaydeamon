package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var PATH_INDICADORES = "indicadores/"

type FakeDolarToday DolarToday

/*func (b DolarToday) AgregarMetadata()([]byte, error){
	return json.Marshal(struct {
		FakeDolarToday
		METADA Metadata
	}{
		FakeDolarToday: FakeDolarToday(b),
		METADA:
	})
}*/

type Indicadores struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	METADATA   Metadata
	DOLARTODAY DolarToday
}

type Metadata struct {
	Collection  *mgo.Collection
	Secuencia int `bson:"n"`
	URL_prev  string
	URL       string
	URL_sig   string
}

func (data *Metadata) AgregarMetadata(secuencia int) error{
	var indicadores Metadata
	if(secuencia == 0) {
		data.Secuencia = 0
		data.URL_prev = nil
		data.URL = PATH_INDICADORES + string(secuencia)
		data.URL_sig = nil
	}else {
		data.Secuencia = secuencia
		data.URL_prev = PATH_INDICADORES + string(secuencia - 1)
		data.URL = PATH_INDICADORES + string(secuencia)
		data.URL_sig = nil

		err := data.Collection.Find(nil).Sort("-secuencia").One(&indicadores)
		if err != nil {
			return err
		}
		indicadores.URL_sig = data.URL

	}
}

