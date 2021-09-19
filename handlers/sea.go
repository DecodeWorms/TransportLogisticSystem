package handlers

import (
	"Tsystem/storage"
	"Tsystem/types"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SeaHandler struct {
	sea *storage.Store
}

func NewSea(s *storage.Store) SeaHandler {
	return SeaHandler{
		sea: s,
	}

}

func (sea SeaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data types.Sea
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		e := errors.New(fmt.Sprintf("json can not be decode %v", err))
		json.NewEncoder(w).Encode(e)
	}
	err = sea.sea.Create(&data)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(true)
	json.NewEncoder(w).Encode(data)
}

func (sea SeaHandler) Seas(w http.ResponseWriter, r *http.Request) {

	res, err := sea.sea.Seas()

	if err != nil {
		e := errors.New(fmt.Sprintf("An error occured %v", err))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(true)
	json.NewEncoder(w).Encode(res)
}

func (sea SeaHandler) Sea(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	trac := params["tracker"]

	res, err := sea.sea.Sea(trac)

	if err != nil {
		e := errors.New(fmt.Sprintf("An error occured %v", err))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(true)
	json.NewEncoder(w).Encode(res)
}

func (sea SeaHandler) Update(w http.ResponseWriter, r *http.Request) {
	var data types.Sea
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to decode json %v", err))
		json.NewEncoder(w).Encode(e)
	}
	err = sea.sea.Update(data)

	json.NewEncoder(w).Encode(err)
	json.NewEncoder(w).Encode(data)
}

func (sea SeaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var data types.Sea
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to decode json %v", err))
		json.NewEncoder(w).Encode(e)
	}

	err = sea.sea.Delete(data)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(data)

}
