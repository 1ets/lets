package drivers

import (
	"context"

	"github.com/1ets/lets"
	"github.com/1ets/lets/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBConfig types.IMongoDB

type mongodbProvider struct {
	dsn      string
	database string
	mongodb  *mongo.Client
	DB       *mongo.Database
}

func (m *mongodbProvider) Connect() {
	clientOptions := options.Client()
	clientOptions.ApplyURI(m.dsn)

	var err error
	m.mongodb, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		lets.LogE("MongoDB: %v", err)
		return
	}

	m.DB = m.mongodb.Database(m.database)
}

// Define MySQL service host and port
func MongoDB() {
	if MongoDBConfig == nil {
		return
	}

	lets.LogI("MongoDB Client Starting ...")

	mongodb := mongodbProvider{
		dsn:      MongoDBConfig.GetDsn(),
		database: MongoDBConfig.GetDatabase(),
	}
	mongodb.Connect()

	// Inject Gorm into repository
	for _, repository := range MongoDBConfig.GetRepositories() {
		repository.SetDriver(mongodb.DB)
	}
}
