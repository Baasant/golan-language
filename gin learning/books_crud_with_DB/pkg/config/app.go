// the aim of the file is to return a variable called db which is
// which will help other variables to interact with database
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// this function help us to open a connection with database
func Connect() {
	//if every thing okey that database that open will cpme inside d if there is error in the connection it come inside err
	d, err := gorm.Open("mysql", "bassant:Bassant@12@/simplerest?charset=utf8&parseTime=True&loc=Local") //open a connection with database.
	if err != nil {                                                                                      // id there is an error
		panic(err)
	} //if
	db = d
} // connect

func GetDB() *gorm.DB {
	return db
}
