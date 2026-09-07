package controller

import (
    FloorDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FloorDAO for database creation
//----------------------------------------------------------------------------
func CreateFloor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Floor model
	//----------------------------------------------------------------------------
	data := model.Floor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Floor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Floor data access object to create
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.CreateFloor( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FloorDAO to find the relevant Floor
//----------------------------------------------------------------------------
func GetFloor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]
	
	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	//----------------------------------------------------------------------------
	// Delegate to the Floor data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.GetFloor(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FloorDAO for database read of all Floors
//----------------------------------------------------------------------------
func GetAllFloor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Floor data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.GetAllFloor()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FloorDAO for database save
//----------------------------------------------------------------------------
func UpdateFloor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Floor model
	//----------------------------------------------------------------------------
	var data = model.Floor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Floor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Floor data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.UpdateFloor(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FloorDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFloor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]

	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	//----------------------------------------------------------------------------
	// Delegate to the Floor data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FloorDAO.DeleteFloor(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Building on a Floor
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBuildingToFloor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	floorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	buildingId,_ := strconv.ParseUint( vars["buildingId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Floor DAO
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.AssignBuildingToFloor(floorId, buildingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Building on a Floor
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBuildingFromFloor( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	floorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Floor DAO
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.UnassignBuildingFromFloor(floorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more roomsIds as a Rooms to a Floor
	//----------------------------------------------------------------------------
func AddRoomsToFloor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	floorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roomsIds,_ := vars["roomsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Floor DAO
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.AddRoomsToFloor(floorId, roomsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more roomsIds as a Rooms from a Floor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRoomsFromFloor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	floorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roomsIds,_ := vars["roomsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Floor DAO
	//----------------------------------------------------------------------------
	requestResult := FloorDAO.RemoveRoomsFromFloor(floorId, roomsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
