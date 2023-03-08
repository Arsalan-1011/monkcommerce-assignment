package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getWeekDays(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["USER_ID"]

	res, err := dbm.Query("SELECT `USER_ID`, `Day`, `START_TIME`, `END_TIME` FROM `weekdays` WHERE `USER_ID`=?", id)

	if err != nil {
		msg := Message{Msg: "User Not Found"}
		json.NewEncoder(w).Encode(msg)
	}

	for res.Next() {
		var c WeekDays
		err := res.Scan(&c.USER_ID, &c.Day, &c.START_TIME, &c.END_TIME)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(c)
	}
	defer res.Close()

}

func createWeekDays(w http.ResponseWriter, r *http.Request) {
	var c WeekDays

	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	USER_ID := c.USER_ID
	Day := c.Day
	START_TIME := c.START_TIME
	END_TIME := c.END_TIME

	res, err := dbm.Query("INSERT INTO `weekdays`(`USER_ID`,`Day`, `START_TIME`, `END_TIME`) VALUES (?,?,?,?)", USER_ID, Day, START_TIME, END_TIME)

	if err != nil {
		msg := Message{Msg: "This slot is Already Scheduled for All Week Days..."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	msg := Message{Msg: "Created Successfully..."}

	json.NewEncoder(w).Encode(msg)
	defer res.Close()
}

func deleteWeekDays(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["USER_ID"]

	res, err := dbm.Query("DELETE FROM `weekdays` WHERE `USER_ID`=?", id)
	if err != nil {
		msg := Message{Msg: "User Not Found"}
		json.NewEncoder(w).Encode(msg)
	}
	for res.Next() {
		var c WeekDays
		err := res.Scan(&c.USER_ID, &c.Day, &c.START_TIME, &c.END_TIME)

		if err != nil {
			log.Fatal(err)
			msg := Message{Msg: "Not Deleted"}
			json.NewEncoder(w).Encode(msg)
			return
		}
		msg := Message{Msg: "Deleted Successfully..."}
		json.NewEncoder(w).Encode(msg)
		defer res.Close()

	}
}

// Week End--------------------------------------------------------------------------------------------

func getWeekEnd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["USER_ID"]

	res, err := dbm.Query("SELECT `USER_ID`, `Day`, `START_TIME`, `END_TIME` FROM `weekend` WHERE `USER_ID`=?", id)

	if err != nil {
		msg := Message{Msg: "User Not Found"}
		json.NewEncoder(w).Encode(msg)
	}

	for res.Next() {
		var c WeekEnd
		err := res.Scan(&c.USER_ID, &c.Day, &c.START_TIME, &c.END_TIME)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(c)
	}
	defer res.Close()

}

func createWeekEnd(w http.ResponseWriter, r *http.Request) {
	var c WeekEnd

	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	START_TIME := c.START_TIME
	Day := c.Day
	id := c.USER_ID
	END_TIME := c.END_TIME

	res, err := dbm.Query("INSERT INTO `weekend`(`USER_ID`,`Day`, `START_TIME`, `END_TIME`) VALUES (?,?,?,?)", id, Day, START_TIME, END_TIME)

	if err != nil {
		msg := Message{Msg: "This slot is Already Scheduled for Weekend..."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	msg := Message{Msg: "Created Successfully..."}

	json.NewEncoder(w).Encode(msg)
	defer res.Close()
}

func deleteWeekEnd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["USER_ID"]

	res, err := dbm.Query("DELETE FROM `weekend` WHERE `USER_ID`=?", id)
	if err != nil {
		msg := Message{Msg: "User Not Found"}
		json.NewEncoder(w).Encode(msg)
	}
	for res.Next() {
		var c WeekEnd
		err := res.Scan(&c.USER_ID, &c.Day, &c.START_TIME, &c.END_TIME)

		if err != nil {
			log.Fatal(err)
			msg := Message{Msg: "Not Deleted"}
			json.NewEncoder(w).Encode(msg)
			return
		}
		msg := Message{Msg: "Deleted Successfully..."}
		json.NewEncoder(w).Encode(msg)
		defer res.Close()

	}
}
