package database

import (
	"bolo/go-utils/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var databases = make(map[string]*gorm.DB)

func Init() {
	conf := config.Instance()
	dbConfigs := conf.GetStringMap("database")
	if len(dbConfigs) == 0 {
		log.Fatal("数据库配置获取失败！")
	}

	for key, value := range dbConfigs {
		dbConfig := value.(map[string]interface{})

		var host = dbConfig["host"].(string)
		var port = dbConfig["port"].(string)
		var dbname = dbConfig["dbname"].(string)
		var username = dbConfig["username"].(string)
		var password = dbConfig["password"].(string)
		var prefix = dbConfig["prefix"].(string)

		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: prefix,
			},
		})
		if err != nil {
			log.Fatal("数据库连接失败！")
		}
		databases[key] = db
	}
}

func Instance(key string) *gorm.DB {
	db, ok := databases[key]
	if !ok {
		log.Fatal("数据库配置不存在")
	}
	return db
}
