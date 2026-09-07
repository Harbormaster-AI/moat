package controller

import (
    FlightHealthEventDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FlightHealthEventDAO for database creation
//----------------------------------------------------------------------------
func CreateFlightHealthEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FlightHealthEvent model
	//----------------------------------------------------------------------------
	data := model.FlightHealthEvent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FlightHealthEvent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FlightHealthEvent data access object to create
	//----------------------------------------------------------------------------
	requestResult := FlightHealthEventDAO.CreateFlightHealthEvent( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FlightHealthEventDAO to find the relevant FlightHealthEvent
//----------------------------------------------------------------------------
func GetFlightHealthEvent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FlightHealthEvent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FlightHealthEventDAO.GetFlightHealthEvent(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FlightHealthEventDAO for database read of all FlightHealthEvents
//----------------------------------------------------------------------------
func GetAllFlightHealthEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FlightHealthEvent data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FlightHealthEventDAO.GetAllFlightHealthEvent()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FlightHealthEventDAO for database save
//----------------------------------------------------------------------------
func UpdateFlightHealthEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FlightHealthEvent model
	//----------------------------------------------------------------------------
	var data = model.FlightHealthEvent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FlightHealthEvent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FlightHealthEvent data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FlightHealthEventDAO.UpdateFlightHealthEvent(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FlightHealthEventDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFlightHealthEvent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FlightHealthEvent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FlightHealthEventDAO.DeleteFlightHealthEvent(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ConnectedAircraft on a FlightHealthEvent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConnectedAircraftToFlightHealthEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	flightHealthEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	connectedAircraftId,_ := strconv.ParseUint( vars["connectedAircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FlightHealthEvent DAO
	//----------------------------------------------------------------------------
	requestResult := FlightHealthEventDAO.AssignConnectedAircraftToFlightHealthEvent(flightHealthEventId, connectedAircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConnectedAircraft on a FlightHealthEvent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConnectedAircraftFromFlightHealthEvent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	flightHealthEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FlightHealthEvent DAO
	//----------------------------------------------------------------------------
	requestResult := FlightHealthEventDAO.UnassignConnectedAircraftFromFlightHealthEvent(flightHealthEventId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


