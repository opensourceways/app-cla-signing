package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	errDocExists    = errors.New("doc exists")
	errDocNotExists = errors.New("doc doesn't exist")
)

func isErrOfNoDocuments(err error) bool {
	return err.Error() == mongo.ErrNoDocuments.Error()
}

type daoImpl struct {
	col     *mongo.Collection
	timeout time.Duration
}

func (impl *daoImpl) withContext(f func(context.Context) error) error {
	return withContext(f, impl.timeout)
}

func (impl *daoImpl) IsDocNotExists(err error) bool {
	return errors.Is(err, errDocNotExists)
}

func (impl *daoImpl) IsDocExists(err error) bool {
	return errors.Is(err, errDocExists)
}

func (impl *daoImpl) InsertDoc(doc bson.M) (string, error) {
	docId := ""

	err := impl.withContext(func(ctx context.Context) error {
		v, err := impl.col.InsertOne(ctx, doc)
		if err != nil {
			docId = toDocId(v.InsertedID)
		}

		return err
	})

	return docId, err
}

func (impl *daoImpl) DeleteAll(filter bson.M) error {
	return impl.withContext(func(ctx context.Context) error {
		_, err := impl.col.DeleteMany(ctx, filter)

		return err
	})
}

func (impl *daoImpl) FindOneAndDelete(filter bson.M, result interface{}) error {
	return impl.withContext(func(ctx context.Context) error {
		err := impl.col.FindOneAndDelete(ctx, filter).Decode(result)
		if err == nil {
			return nil
		}

		if isErrOfNoDocuments(err) {
			return errDocNotExists
		}

		return err
	})
}
