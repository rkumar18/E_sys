package lib

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)


func DBConnection() *gorm.DB{
	DSN := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Postgres_config.User, Postgres_config.Password, Postgres_config.Db_Name)
	db, _ := gorm.Open("postgres", DSN)
	db.AutoMigrate(&Users{})
	return db
}

func dbsignin(input *Users)Users{
	var result Users
	fmt.Printf("o")
	db := DBConnection()
	db.Where("Email=?",input.Email).Find(&result)
	if result.Email==""{
		db.Create(&input)
	}
	defer db.Close()
	return result
}