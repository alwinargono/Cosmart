package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Genre string `json:"genre"`
}

type Response struct {
	Genre       string      `json:"name"`
	ListOfBooks []BooksList `json:"works"`
}

type BooksList struct {
	TitleOfBook string          `json:"title"`
	Edition     int             `json:"edition_count"`
	Authors     []ListOfAuthors `json:"authors"`
}

type ListOfAuthors struct {
	Name string `json:"name"`
}

type PickUpSched struct {
	Title   string `json:"title"`
	Edition int    `json:"edition_number"`
	Date    string `json:"date"`
}

var temp []PickUpSched

func GetList(w http.ResponseWriter, r *http.Request) {
	var req Request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter JSON Data for genre of the book you want to look for")
	}

	json.Unmarshal(reqBody, &req)

	template := "https://openlibrary.org/subjects/" + req.Genre + ".json"
	response, err := http.Get(template)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resp Response
	json.Unmarshal(responseData, &resp)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}

func PickUp(w http.ResponseWriter, r *http.Request) {
	var req PickUpSched
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter memberID and data that wants to be updated")
	}

	json.Unmarshal(reqBody, &req)
	w.WriteHeader(http.StatusCreated)
	for i := range temp {
		if req.Date == temp[i].Date && req.Edition == temp[i].Edition {
			json.NewEncoder(w).Encode("Book is Currently not available at date")
			return
		}
	}

	temp = append(temp, PickUpSched{
		Title:   req.Title,
		Edition: req.Edition,
		Date:    req.Date,
	})

	json.NewEncoder(w).Encode(temp)
}
