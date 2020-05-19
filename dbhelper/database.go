package dbhelper

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
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
	log.Println("Asegurandome que la tabla Test tiene la estructura correcta...")
	db.AutoMigrate(&Contact{})

	contact := Contact{Message: "Restart Contact!", TimeStamp: time.Now().String()}
	db.Create(&contact)

}
