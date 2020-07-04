package lib

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)


func DBConnection() *gorm.DB{
	DSN := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Postgres_config.User, Postgres_config.Password, Postgres_config.Db_Name)
	db, _ := gorm.Open("postgres", DSN)
	db.AutoMigrate(&PradhanData{})
	return db
}
