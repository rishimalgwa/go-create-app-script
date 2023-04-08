package filedata

import "fmt"

func GetDBFileData(projectName string) string {
	uri := "mongodb://root:password@mongo:27017/" + projectName + "?authSource=admin"
	return fmt.Sprintf(`package Database
	import (
		"context"
		"os"
		"github.com/go-redis/redis/v8"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	)
	func NewDatabase() controller.Database {
		client, err := mongo.NewClient(options.Client().ApplyURI("%s"))
		if err != nil {
			panic(err.Error())
		}
		err = client.Connect(context.TODO())
		if err != nil {
			panic(err.Error())
		}
	
		rdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_DB_ADDR"),
			Password: os.Getenv("REDIS_DB_PASS"),
			DB:       1,
		})
		return controller.Database{
			MongoClient: client,
			RedisClient: rdb,
		}
	
	}
	`, uri)
}
