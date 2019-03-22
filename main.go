package main

import (
	"os"
	"fmt"
	"time"
	"log"
	_ "test/routers"
	_ "test/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	// Load enviornment variables
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Database connection
	orm.RegisterDriver("mysql", orm.DRMySQL)
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	maxIdle := 30
	maxConn := 30
	conn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8", db_user, db_pass, db_name)
	orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)
	// set up the timezone
	orm.DefaultTimeLoc = time.UTC
}

func main() {
	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := false
	verbose := true
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
    // Logs set up
    beego.SetLogger("file", `{"filename":"logs/test.log"}`)
  	beego.Run()
}

