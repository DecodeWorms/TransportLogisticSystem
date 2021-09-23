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

	err = road.store.Create(data)
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

	var data []*types.Road
	data, err = road.store.SenderNames()

	if err != nil {
		e := errors.New(fmt.Sprintf("No data to returned"))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(data)

}

func (road RoadHandlers) GetLocation(w http.ResponseWriter, r *http.Request) {
	var data []types.Road
	var err error
	// param := mux.Vars(r)
	// trc := param["tracker"]

	data, err = road.store.GetLocation()

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(data)
}

func (road RoadHandlers) GetId(w http.ResponseWriter, r *http.Request) {
	var data types.Road
	var err error

	params := mux.Vars(r)
	trc := params["tracker"]

	data, err = road.store.GetId(trc)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(data)
}

func (road RoadHandlers) UpdateSenderLocation(w http.ResponseWriter, r *http.Request) {
	var data types.Road
	err := json.NewDecoder(r.Body).Decode(&data)

	params := mux.Vars(r)
	trc := params["tracker"]

	if err != nil {
		e := errors.New(fmt.Sprintf("Un able to decode json"))
		json.NewEncoder(w).Encode(e)
	}
	err = road.store.UpdateSenderLocation(trc, data)

}

func (road RoadHandlers) CreateIdentifiedField(w http.ResponseWriter, r *http.Request) {
	var data types.Road
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	road.store.CreateIdentifiedField(data)
}
