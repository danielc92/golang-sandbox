package main

/*
interacting with postgres via gorm
*/

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// models

type Cat struct {
	// include default fields from gorm.Model
	gorm.Model
	Name string
	Breed string
	Colour string
}

func GetDSN()string {
	
	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSL_MODE"),
	)
	fmt.Println("DDDD",dsn)
	return dsn	
}


func runPostgres(){
	// Load environment
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatalf("something went wrong! failed to load environment from .env")
	}

	
	dsn := GetDSN()
	// dsn := "host=localhost user=postgres password=secret dbname=postgres port=3008 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	pgDB, _:= db.DB()

	// Pooling https://gorm.io/docs/generic_interface.html#Connection-Pool
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	pgDB.SetMaxIdleConns(8)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	pgDB.SetMaxOpenConns(40)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	pgDB.SetConnMaxLifetime(time.Minute * 30)


	// Handle connection error
	if err != nil {
		log.Fatal("failed to connect")
	}	
	fmt.Println("Connected to database successfully", db.Config)

	db.AutoMigrate(&Cat{})

	// Creating an entry (https://gorm.io/docs/create.html)
	myCat := Cat{Name:"Jerry", Breed:"Siamese", Colour: "White"}
	myCat2 := Cat{Name:"Jerry", Breed:"Siamese", Colour: "White"}
	myCat3 := Cat{Name:"Jerry", Breed:"Siamese", Colour: "Orange"}
	db.Create(&myCat)
	db.Create(&myCat2)
	db.Create(&myCat3)


	var queryCat Cat 
	
	// Query first row (https://gorm.io/docs/query.html)
	db.First(&queryCat)
	fmt.Println("Found: ",queryCat.Name)

	// Query multiple based on condition
	var colourQuery Cat
	res := db.Find(&colourQuery, "colour = ?", "Orange")
	fmt.Println("Found: ",colourQuery, res.RowsAffected, res.Error)

	// Updating a whole column based on condition (https://gorm.io/docs/update.html)
	db.Model(&Cat{}).Where("breed = ?", "Siamese").Update("colour", "RAINBOW")

	// Updating a single entry
	var catToUpdate Cat
	db.First(&catToUpdate)
	catToUpdate.Colour = "Red"
	res2 := db.Save(&catToUpdate)
	fmt.Println(res2.RowsAffected, res2.Error)


	// Deleting an entry (https://gorm.io/docs/delete.html)
	db.Delete(&myCat)
	
}