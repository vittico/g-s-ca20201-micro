package dbhelper

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connection
type Connection struct {
	CTX string
}

// Recordit
func (c Connection) RecordIt() {

	type Contact struct {
		gorm.Model
		Contact   int `gorm:"AUTO_INCREMENT"`
		Message   string
		TimeStamp string
	}

	db, errDBConn := gorm.Open("postgres", c.CTX)
	if errDBConn != nil {
		log.Panic(errDBConn)
	}
	db.AutoMigrate(&Contact{})

	contact := Contact{Message: "Hit!", TimeStamp: time.Now().String()}
	db.Create(&contact)
	defer db.Close()

}
