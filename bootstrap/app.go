package bootstrap

import (
	// "database/sql"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	// "github.com/go-redis/redis/v9"
)

var (
	App *Application
)

type Application struct {
	Config *viper.Viper
	Mongo  *mongo.Client
	Mysql  *gorm.DB
	Log    *logrus.Logger
	// Maria  *sql.DB
	// Redis  *redis.Client
}

func init() {
	AppInit()
}

func AppInit() {
	App = &Application{}
	App.Config = InitConfig()
	App.Mongo = InitMongoDatabase()
	App.Mysql = InitMysqlDatabase()
	App.Log = InitLogger()
	// App.Maria = InitMariaDatabase()
	// App.Redis = InitRedis()

}
