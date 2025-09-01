package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"kratos-demo/internal/conf"
	"kratos-demo/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	mysql *ent.Client
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	driver, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Errorf("open db error, %s", err)
	}
	mysqlClient := ent.NewClient(ent.Driver(driver))

	if err := mysqlClient.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	mysqlClient.Article.Create().SetID(1).SetTitle("mysql").SetContext("sdada").Save(context.Background())
	article, _ := mysqlClient.Article.Get(context.Background(), 1)
	println(article.String())

	redisClient := redis.NewClient(&redis.Options{})

	d := &Data{
		mysql: mysqlClient,
		redis: redisClient,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := mysqlClient.Close(); err != nil {
			log.Error(err)
		}
		if err := redisClient.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
