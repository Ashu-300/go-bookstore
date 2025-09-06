package config

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    db *gorm.DB
)

func Connect() {
    // Connect without specifying a database
    d, err := gorm.Open("mysql", "root:ashupassword@tcp(127.0.0.1:3306)/")
    if err != nil {
        panic(err)
    }

    // Create the database if it doesn't exist
    d.Exec("CREATE DATABASE IF NOT EXISTS simplerest")

    // Connect to the actual database
    d.Close() // close the first connection
    d, err = gorm.Open("mysql", "root:ashupassword@tcp(127.0.0.1:3306)/simplerest?parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }

    db = d
    fmt.Println("Database connected successfully")
}

func GetDB() *gorm.DB {
    return db
}
