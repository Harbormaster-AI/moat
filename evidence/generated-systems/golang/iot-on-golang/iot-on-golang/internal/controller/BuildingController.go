package controller

import (
    BuildingDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BuildingDAO for database creation
//----------------------------------------------------------------------------
func CreateBuilding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Building model
	//----------------------------------------------------------------------------
	data := model.Building{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Building model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Building data access object to create
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.CreateBuilding( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BuildingDAO to find the relevant Building
//----------------------------------------------------------------------------
func GetBuilding(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Building data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.GetBuilding(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BuildingDAO for database read of all Buildings
//----------------------------------------------------------------------------
func GetAllBuilding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Building data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.GetAllBuilding()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BuildingDAO for database save
//----------------------------------------------------------------------------
func UpdateBuilding(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Building model
	//----------------------------------------------------------------------------
	var data = model.Building{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Building model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Building data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.UpdateBuilding(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BuildingDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBuilding(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Building data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BuildingDAO.DeleteBuilding(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Site on a Building
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSiteToBuilding(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	buildingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	siteId,_ := strconv.ParseUint( vars["siteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Building DAO
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.AssignSiteToBuilding(buildingId, siteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Site on a Building
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSiteFromBuilding( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	buildingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Building DAO
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.UnassignSiteFromBuilding(buildingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more floorsIds as a Floors to a Building
	//----------------------------------------------------------------------------
func AddFloorsToBuilding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	buildingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	floorsIds,_ := vars["floorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Building DAO
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.AddFloorsToBuilding(buildingId, floorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more floorsIds as a Floors from a Building
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFloorsFromBuilding(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	buildingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	floorsIds,_ := vars["floorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Building DAO
	//----------------------------------------------------------------------------
	requestResult := BuildingDAO.RemoveFloorsFromBuilding(buildingId, floorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
