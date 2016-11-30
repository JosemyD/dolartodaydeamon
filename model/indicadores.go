package model

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

var PATH_INDICADORES = "indicadores/"

type Indicadores struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	METADATA   Metadata
	DOLARTODAY DolarToday
}

func (d *Indicadores) AgregarMetadata(metadata Metadata) {
	d.METADATA = metadata
}

type Metadata struct {
	Secuencia int `bson:"n"`
	URL_prev  string
	URL       string
	URL_sig   string
}

func (data *Metadata) Metadata(secuencia int) {
	if secuencia == 0 {
		secuencia = secuencia
		data.URL = PATH_INDICADORES
		data.URL_sig = PATH_INDICADORES + strconv.Itoa(secuencia+1)
	} else if secuencia == 1 {
		data.Secuencia = secuencia
		data.URL_prev = PATH_INDICADORES
		data.URL = PATH_INDICADORES + strconv.Itoa(secuencia)
		data.URL_sig = PATH_INDICADORES + strconv.Itoa(secuencia+1)
	} else {
		data.Secuencia = secuencia
		data.URL_prev = PATH_INDICADORES + strconv.Itoa(secuencia-1)
		data.URL = PATH_INDICADORES + strconv.Itoa(secuencia)
		data.URL_sig = PATH_INDICADORES + strconv.Itoa(secuencia+1)
	}
}
