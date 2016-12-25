package models

import(
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

type DBconfig struct {
	ConnectString string
}

func (dbc DBconfig) Connect() *gorm.DB {
	db, err := gorm.Open("mysql", dbc.ConnectString + "?parseTime=true")

	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	if err != nil{
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	migration(db)

	return db
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Profile{})

	db.Model(&User{}).Related(&Image{})
	db.Model(&User{}).Related(&Profile{})

	if !db.HasTable(&Image{}) || !db.HasTable(&User{}) {
		panic("Error auto Migrate!")
	}
}