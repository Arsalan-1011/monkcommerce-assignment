package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

func connect() {

	//get the config values
	file, err := ioutil.ReadFile("Config.json")
	if nil != err {
		//("Error %s", err.Error())
		return
	}
	var config DbConfig
	err = json.Unmarshal(file, &config)

	if nil != err {
		//("Error %s", err.Error())
		return
	}

	//use the values now from config
	log.Println(config)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.UserConfig.User, config.UserConfig.Password, config.Host, config.Port, config.UserConfig.Database))
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("connected")
	}
	dbm = db

}
