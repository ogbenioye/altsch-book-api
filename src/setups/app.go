package setup

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connection() {
	// "@tcp(mysql:3306)" is a way of connecting two dockerized services together
	// "root" is the user and "secret" is the password has set in the docker-compose file
	// database is hosted on "mysql" on port 3306
	d, err := gorm.Open("mysql", "root:secret@tcp(mysql:3306)/altsch-go?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
