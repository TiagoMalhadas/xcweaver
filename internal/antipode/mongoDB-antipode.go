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
	mongoClient *mongo.Client
	database    string
}

type Document struct {
	Key   string  `bson:"key"`
	Value AntiObj `bson:"value"`
}

func CreateMongoDB(host string, port string, database string) MongoDB {
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port + "/?directConnection=true")
	client, err := mongo.Connect(_, clientOptions)
	if err != nil {
		panic(err)
	}
	return MongoDB{client,
		database,
	}
}

// devo verificar primeiro se já existe essa key? E caso exista fazer update?
// falta fechar a conexão
func (m MongoDB) write(ctx context.Context, collectionName string, key string, obj AntiObj) error {

	collection := m.mongoClient.Database(m.database).Collection(collectionName)

	filter := bson.D{{"key", key}}

	replacement := Document{
		Key:   key,
		Value: obj,
	}

	_, err := collection.ReplaceOne(context.Background(), filter, replacement, options.Replace().SetUpsert(true))

	/*mongoObj := Document{
		Key:   key,
		Value: obj,
	}

	_, err = client.Database(m.database).Collection(collection).InsertOne(ctx, mongoObj)*/

	return err
}

// posso assumir que não há mais do que um objeto com a mesma key?
func (m MongoDB) read(ctx context.Context, collection string, key string) (AntiObj, error) {

	filter := bson.D{{"key", key}}

	var result Document
	err := m.mongoClient.Database(m.database).Collection(collection).FindOne(context.Background(), filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return AntiObj{}, ErrNotFound
	} else if err != nil {
		return AntiObj{}, err
	}

	return result.Value, nil
}

func (m MongoDB) consume(context.Context, string, string, chan struct{}) (<-chan AntiObj, error) {
	return nil, nil
}

func (m MongoDB) barrier(ctx context.Context, lineage []WriteIdentifier, datastoreID string) error {

	fmt.Println("start barrier")

	for _, writeIdentifier := range lineage {
		fmt.Println("key after for: ", writeIdentifier.Key)
		if writeIdentifier.Dtstid == datastoreID {
			for {
				filter := bson.D{{"key", writeIdentifier.Key}}

				fmt.Println("key: ", writeIdentifier.Key)

				var result Document
				err := m.mongoClient.Database(m.database).Collection(writeIdentifier.TableId).FindOne(context.Background(), filter).Decode(&result)

				if !errors.Is(err, mongo.ErrNoDocuments) && err != nil {
					return err
				} else if errors.Is(err, mongo.ErrNoDocuments) { //the version replication process is not yet completed
					fmt.Println("replication in progress")
					continue
				} else {
					if result.Value.Version == writeIdentifier.Version { //the version replication process is already completed
						fmt.Println("replication done: ", result.Value.Version)
						break
					} else { //the version replication process is not yet completed
						fmt.Println("replication of the new version in progress! Old Version: ", result.Value.Version, " new version: ", writeIdentifier.Version)
						continue
					}
				}
			}
		}
	}

	return nil
}
