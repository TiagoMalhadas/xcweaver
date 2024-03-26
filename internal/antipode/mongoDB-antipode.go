package antipode

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	clientOptions *options.ClientOptions
	database      string
	collection    string
}

type Document struct {
	key   string  `bson:"key"`
	value AntiObj `bson:"value"`
}

func CreateMongoDB(host string, port string, database string, collection string) MongoDB {
	return MongoDB{options.Client().ApplyURI("mongodb://" + host + ":" + port),
		database,
		collection,
	}
}

// devo verificar primeiro se já existe essa key? E caso exista fazer update?
// falta fechar a conecção
func (m MongoDB) write(ctx context.Context, key string, obj AntiObj) error {

	client, err := mongo.Connect(ctx, m.clientOptions)
	if err != nil {
		return err
	}

	mongoObj := Document{key, obj}

	fmt.Println("key: ", key)

	_, err = client.Database(m.database).Collection(m.collection).InsertOne(ctx, mongoObj)

	return err
}

// posso assumir que não há mais do que um objeto com a mesma key?
func (m MongoDB) read(ctx context.Context, key string) (AntiObj, error) {

	client, err := mongo.Connect(ctx, m.clientOptions)
	if err != nil {
		return AntiObj{}, err
	}

	fmt.Println("key read: ", key)

	filter := bson.D{{"key", key}}

	var result Document
	err = client.Database(m.database).Collection(m.collection).FindOne(context.Background(), filter).Decode(&result)

	return result.value, err
}

func (m MongoDB) barrier(ctx context.Context, lineage []WriteIdentifier, datastoreID string) error {

	client, err := mongo.Connect(ctx, m.clientOptions)
	if err != nil {
		return err
	}

	for _, writeIdentifier := range lineage {
		if writeIdentifier.Dtstid == datastoreID {
			for {
				filter := bson.D{{"key", writeIdentifier.Key}}

				var result Document
				err = client.Database(m.database).Collection(m.collection).FindOne(context.Background(), filter).Decode(&result)

				if !errors.Is(err, mongo.ErrNoDocuments) && err != nil {
					return err
				} else if errors.Is(err, mongo.ErrNoDocuments) { //the version replication process is not yet completed
					continue
				} else {
					if result.value.Version == writeIdentifier.Version { //the version replication process is already completed
						break
					} else { //the version replication process is not yet completed
						continue
					}
				}
			}
		}
	}

	return nil
}
