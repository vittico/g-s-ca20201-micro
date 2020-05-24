package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)


// Contact
type Contact struct {
	gorm.Model
	Contact   int `gorm:"AUTO_INCREMENT"`
	Message   string
	TimeStamp string
}

// ConnectDB
func  ConnectDB(ctx string) *gorm.DB {

	dbctx, errDBConn := gorm.Open("postgres", ctx )
	if errDBConn != nil {
		log.Panic(errDBConn)
	}

	dbctx.AutoMigrate(&Contact{})
	return dbctx

}

// Recordit
func   RecordIt ( ctx *gorm.DB) {

	contact := Contact{Message: "Hit!", TimeStamp: time.Now().String()}
	ctx.Create(&contact)

}
