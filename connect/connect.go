package db

import (
	"employee/utils"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resource struct {
	DB *mongo.Database
}

// Close use this method to close database connection
func (r *Resource) Close() {
	logrus.Warning("Closing all db connections")
}

func InitResource() (*Resource, error) {
	host := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	ctx, cancel := utils.InitContext()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		return nil, err
	}
	defer cancel()
	return &Resource{DB: mongoClient.Database(dbName)}, nil
}
