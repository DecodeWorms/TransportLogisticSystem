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

type RoadHandlers struct {
	store *storage.DataStore
}

func NewRoadHandlers(s *storage.DataStore) RoadHandlers {
	return RoadHandlers{
		store: s,
	}
}

func (road RoadHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var data types.Road
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		e := errors.New(fmt.Sprintf("There is an error in decoding json %v", err))
		json.NewEncoder(w).Encode(e)
	}

	log.Println("Datas are being passed...")

	err = road.store.Create(&data)
	if err != nil {
		json.NewEncoder(w).Encode(err)

	}
	json.NewEncoder(w).Encode(true)
}

func (road RoadHandlers) Roads(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing datas...")
	//var data []*types.Road
	data, err := road.store.Roads()

	if err != nil {
		e := errors.New(fmt.Sprintf("Irregular parameters %v", err))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(data)

}

func (road RoadHandlers) Road(w http.ResponseWriter, r *http.Request) {
	var tracker = mux.Vars(r)
	trac := tracker["tracker"]

	//var tracker types.Road
	//json.NewDecoder(r.Body).Decode(&tracker)
	result, err := road.store.Road(trac)

	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(result)

}

func (road RoadHandlers) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var data types.Road
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		e := errors.New(fmt.Sprintf("An error occured  while decoding json ..%v", err))
		json.NewEncoder(w).Encode(e)
	}

	err = road.store.UpdateLocation(data)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(true)
	json.NewEncoder(w).Encode(data)

}

func (road RoadHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	var data types.Road

	params := mux.Vars(r)
	trac := params["tracker"]

	err := road.store.Delete(trac, data)

	if err != nil {
		e := errors.New(fmt.Sprintf("An error has occured %v", err))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(true)
	json.NewEncoder(w).Encode(data)
}

func (road RoadHandlers) SenderNames(w http.ResponseWriter, r *http.Request) {
	var err error

	var data []types.Road
	data, err = road.store.SenderNames()

	if err != nil {
		e := errors.New(fmt.Sprintf("No data to returned"))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(data)

}
