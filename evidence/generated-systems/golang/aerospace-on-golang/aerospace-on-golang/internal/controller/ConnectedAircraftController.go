package controller

import (
    ConnectedAircraftDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ConnectedAircraftDAO for database creation
//----------------------------------------------------------------------------
func CreateConnectedAircraft(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ConnectedAircraft model
	//----------------------------------------------------------------------------
	data := model.ConnectedAircraft{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ConnectedAircraft model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft data access object to create
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.CreateConnectedAircraft( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ConnectedAircraftDAO to find the relevant ConnectedAircraft
//----------------------------------------------------------------------------
func GetConnectedAircraft(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ConnectedAircraft data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.GetConnectedAircraft(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ConnectedAircraftDAO for database read of all ConnectedAircrafts
//----------------------------------------------------------------------------
func GetAllConnectedAircraft(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.GetAllConnectedAircraft()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ConnectedAircraftDAO for database save
//----------------------------------------------------------------------------
func UpdateConnectedAircraft(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ConnectedAircraft model
	//----------------------------------------------------------------------------
	var data = model.ConnectedAircraft{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ConnectedAircraft model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.UpdateConnectedAircraft(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ConnectedAircraftDAO for database deletion
//----------------------------------------------------------------------------
func DeleteConnectedAircraft(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ConnectedAircraft data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ConnectedAircraftDAO.DeleteConnectedAircraft(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Aircraft on a ConnectedAircraft
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAircraftToConnectedAircraft(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	connectedAircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftId,_ := strconv.ParseUint( vars["aircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft DAO
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.AssignAircraftToConnectedAircraft(connectedAircraftId, aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Aircraft on a ConnectedAircraft
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAircraftFromConnectedAircraft( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	connectedAircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft DAO
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.UnassignAircraftFromConnectedAircraft(connectedAircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more flightHealthEventsIds as a FlightHealthEvents to a ConnectedAircraft
	//----------------------------------------------------------------------------
func AddFlightHealthEventsToConnectedAircraft(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	connectedAircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	flightHealthEventsIds,_ := vars["flightHealthEventsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft DAO
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.AddFlightHealthEventsToConnectedAircraft(connectedAircraftId, flightHealthEventsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more flightHealthEventsIds as a FlightHealthEvents from a ConnectedAircraft
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFlightHealthEventsFromConnectedAircraft(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	connectedAircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	flightHealthEventsIds,_ := vars["flightHealthEventsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft DAO
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.RemoveFlightHealthEventsFromConnectedAircraft(connectedAircraftId, flightHealthEventsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more softwareLoadsIds as a SoftwareLoads to a ConnectedAircraft
	//----------------------------------------------------------------------------
func AddSoftwareLoadsToConnectedAircraft(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	connectedAircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	softwareLoadsIds,_ := vars["softwareLoadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft DAO
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.AddSoftwareLoadsToConnectedAircraft(connectedAircraftId, softwareLoadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more softwareLoadsIds as a SoftwareLoads from a ConnectedAircraft
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSoftwareLoadsFromConnectedAircraft(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	connectedAircraftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	softwareLoadsIds,_ := vars["softwareLoadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ConnectedAircraft DAO
	//----------------------------------------------------------------------------
	requestResult := ConnectedAircraftDAO.RemoveSoftwareLoadsFromConnectedAircraft(connectedAircraftId, softwareLoadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
