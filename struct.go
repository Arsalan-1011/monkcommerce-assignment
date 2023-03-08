package main

import "time"

type WeekDays struct {
	USER_ID    int       `json:"USER_ID"`
	Day        string    `json:"Day"`
	START_TIME time.Time `json:"START_TIME"`
	END_TIME   time.Time `json:"END_TIME"`
}

type WeekEnd struct {
	USER_ID    int       `json:"USER_ID"`
	Day        string    `json:"Day"`
	START_TIME time.Time `json:"START_TIME"`
	END_TIME   time.Time `json:"END_TIME"`
}

type DbUserConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type DbConfig struct {
	UserConfig DbUserConfig `json:"database"`
	Host       string       `json:"host"`
	Port       string       `json:"port"`
}

type Message struct {
	Msg string `json:"msg"`
}
