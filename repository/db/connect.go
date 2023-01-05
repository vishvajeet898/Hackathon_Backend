package repository

import (
	"Hackathon_Backend/model"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Database struct {
	dialector gorm.Dialector
}

func NewDatabse(dialector gorm.Dialector) *Database {
	return &Database{

		dialector: dialector,
	}
}

func (d *Database) Connect() {

	db, err := gorm.Open(d.dialector, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
	pgdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	// Config connection pool
	pgdb.SetMaxIdleConns(10)
	pgdb.SetMaxOpenConns(100)

	err = db.AutoMigrate(&model.Basic_Info{}, &model.Aws_File{}, &model.Drug_Donation{}, &model.Health_insurance{}, &model.Measurement{}, &model.Public_Page{}, &model.SharedProfile{}, &model.User{}, &model.Visit{})
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	// Set DB connection as global
	DB = db

}
