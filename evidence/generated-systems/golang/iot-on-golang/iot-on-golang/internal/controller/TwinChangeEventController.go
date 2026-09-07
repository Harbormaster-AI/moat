package controller

import (
    TwinChangeEventDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TwinChangeEventDAO for database creation
//----------------------------------------------------------------------------
func CreateTwinChangeEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TwinChangeEvent model
	//----------------------------------------------------------------------------
	data := model.TwinChangeEvent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TwinChangeEvent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TwinChangeEvent data access object to create
	//----------------------------------------------------------------------------
	requestResult := TwinChangeEventDAO.CreateTwinChangeEvent( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TwinChangeEventDAO to find the relevant TwinChangeEvent
//----------------------------------------------------------------------------
func GetTwinChangeEvent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TwinChangeEvent data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TwinChangeEventDAO.GetTwinChangeEvent(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TwinChangeEventDAO for database read of all TwinChangeEvents
//----------------------------------------------------------------------------
func GetAllTwinChangeEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TwinChangeEvent data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TwinChangeEventDAO.GetAllTwinChangeEvent()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TwinChangeEventDAO for database save
//----------------------------------------------------------------------------
func UpdateTwinChangeEvent(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TwinChangeEvent model
	//----------------------------------------------------------------------------
	var data = model.TwinChangeEvent{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TwinChangeEvent model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TwinChangeEvent data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TwinChangeEventDAO.UpdateTwinChangeEvent(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TwinChangeEventDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTwinChangeEvent(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TwinChangeEvent data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TwinChangeEventDAO.DeleteTwinChangeEvent(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Twin on a TwinChangeEvent
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTwinToTwinChangeEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	twinChangeEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	twinId,_ := strconv.ParseUint( vars["twinId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TwinChangeEvent DAO
	//----------------------------------------------------------------------------
	requestResult := TwinChangeEventDAO.AssignTwinToTwinChangeEvent(twinChangeEventId, twinId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Twin on a TwinChangeEvent
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTwinFromTwinChangeEvent( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	twinChangeEventId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TwinChangeEvent DAO
	//----------------------------------------------------------------------------
	requestResult := TwinChangeEventDAO.UnassignTwinFromTwinChangeEvent(twinChangeEventId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


