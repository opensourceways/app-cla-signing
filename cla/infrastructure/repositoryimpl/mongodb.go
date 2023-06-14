package repositoryimpl

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type dao interface {
	IsDocNotExists(error) bool
	IsDocExists(error) bool

	InsertDoc(doc bson.M) (string, error)
	DeleteAll(filter bson.M) error
	FindOneAndDelete(filter bson.M, result interface{}) error
}

func genDoc(doc interface{}) (m bson.M, err error) {
	v, err := json.Marshal(doc)
	if err != nil {
		return
	}

	err = json.Unmarshal(v, &m)

	return
}
