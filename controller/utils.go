package controller

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

/*Autoincrement ID Mongo*/
type idGenerator struct {
	db  *mgo.Database
	N   int    `bson:"n"`
	Key string // Don't forget to add an unique index to this field.
}

func NewIDGenerator(db *mgo.Database) *idGenerator {
	return &idGenerator{
		db: db,
	}
}

// Generate a auto increment version ID for the given key
func (idg *idGenerator) Next(key string) (int, error) {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"n": 1}},
		ReturnNew: true,
	}
	r := &idGenerator{}
	_, err := idg.db.C("idgen").Find(bson.M{"key": key}).Apply(change, &r)

	if err == mgo.ErrNotFound {
		err := idg.db.C("idgen").Insert(bson.M{"key": key, "n": 1})
		if err != nil {
			return -1, err
		}
		return 1, nil
	} else if err != nil {
		return -1, err
	}

	return r.N, nil
}

/*Get Json from API*/
func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

/*Connect Mongodb*/
func Connect(dialInfo *mgo.DialInfo) (*mgo.Database, error) {

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	db := session.DB("")

	return db, err
}
