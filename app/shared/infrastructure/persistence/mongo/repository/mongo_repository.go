package repository

import (
	"context"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/infrastructure/persistence/mongo/config"
	"go-course/demo/app/shared/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	ctx = context.TODO()
)

type MongoRepository struct {
	collectionName  string
	connection      *config.DbConnection
	mongoCollection *mongo.Collection
}

type OptionsRepository struct {
	Limit int64
	Skip  int64
}

func NewMongoRepository(collection string,
	connection *config.DbConnection) *MongoRepository {
	mongoRepository := &MongoRepository{
		collectionName: collection,
		connection:     connection,
	}
	client, err := mongoRepository.connection.GetConnection()
	if err != nil {
		panic(`cannot connect to mongo database`)
	}
	mongoRepository.mongoCollection = client.
		Database(os.Getenv(constants.MONGODB_DATABASE)).
		Collection(mongoRepository.collectionName)
	return mongoRepository
}

func (b *MongoRepository) Find(query interface{}) ([]interface{}, error) {
	cursor, err := b.mongoCollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	var response []interface{}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var c map[string]interface{}
		err := cursor.Decode(&c)
		if err != nil {
			return nil, err
		}
		response = append(response, c)
	}
	return response, nil
}

func (b *MongoRepository) FindOne(query interface{}) (interface{}, error) {
	cursor := b.mongoCollection.FindOne(ctx, query)
	var entity interface{}
	err := cursor.Decode(entity)
	if err != nil {
		if err.Error() != "" {
			return nil, nil
		}
		return nil, err
	}
	return entity, nil
}
func (b *MongoRepository) FindPageable(limit, page int64, query interface{}) ([]interface{}, error) {
	opts := options.Find()
	skip := (page - 1) * limit
	opts.Skip = &skip
	opts.Limit = &limit
	cursor, err := b.mongoCollection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	var response []interface{}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var c map[string]interface{}
		err := cursor.Decode(&c)
		if err != nil {
			return nil, err
		}
		response = append(response, c)
	}
	return response, nil
}

func (b *MongoRepository) Save(body interface{}) (string, error) {

	result, err := b.mongoCollection.InsertOne(context.TODO(), body)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	log.Info("new object was created successfully: %s", id)
	return id, nil
}

func (b *MongoRepository) SaveMany(bodies []interface{}) ([]interface{}, error) {
	result, err := b.mongoCollection.InsertMany(context.TODO(), bodies)
	if err != nil {
		return nil, err
	}
	ids := result.InsertedIDs
	log.Info("a massive object was saved successfully")
	return ids, nil
}

func (b *MongoRepository) Update(query map[string]interface{}, body interface{}) (interface{}, error) {
	update := bson.M{
		"$set": body,
	}
	result, err := b.mongoCollection.UpdateOne(context.TODO(), query, update)
	if err != nil {
		return nil, err
	}
	response := result.UpsertedID
	log.Info("object %q was updated successfully: ", query)
	return response, err
}
