package controller

import (
    AirworthinessDirectiveDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AirworthinessDirectiveDAO for database creation
//----------------------------------------------------------------------------
func CreateAirworthinessDirective(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AirworthinessDirective model
	//----------------------------------------------------------------------------
	data := model.AirworthinessDirective{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AirworthinessDirective model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AirworthinessDirective data access object to create
	//----------------------------------------------------------------------------
	requestResult := AirworthinessDirectiveDAO.CreateAirworthinessDirective( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AirworthinessDirectiveDAO to find the relevant AirworthinessDirective
//----------------------------------------------------------------------------
func GetAirworthinessDirective(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AirworthinessDirective data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AirworthinessDirectiveDAO.GetAirworthinessDirective(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AirworthinessDirectiveDAO for database read of all AirworthinessDirectives
//----------------------------------------------------------------------------
func GetAllAirworthinessDirective(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AirworthinessDirective data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AirworthinessDirectiveDAO.GetAllAirworthinessDirective()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AirworthinessDirectiveDAO for database save
//----------------------------------------------------------------------------
func UpdateAirworthinessDirective(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AirworthinessDirective model
	//----------------------------------------------------------------------------
	var data = model.AirworthinessDirective{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AirworthinessDirective model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AirworthinessDirective data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AirworthinessDirectiveDAO.UpdateAirworthinessDirective(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AirworthinessDirectiveDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAirworthinessDirective(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AirworthinessDirective data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AirworthinessDirectiveDAO.DeleteAirworthinessDirective(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more workOrdersIds as a WorkOrders to a AirworthinessDirective
	//----------------------------------------------------------------------------
func AddWorkOrdersToAirworthinessDirective(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	airworthinessDirectiveId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AirworthinessDirective DAO
	//----------------------------------------------------------------------------
	requestResult := AirworthinessDirectiveDAO.AddWorkOrdersToAirworthinessDirective(airworthinessDirectiveId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workOrdersIds as a WorkOrders from a AirworthinessDirective
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkOrdersFromAirworthinessDirective(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	airworthinessDirectiveId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrdersIds,_ := vars["workOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AirworthinessDirective DAO
	//----------------------------------------------------------------------------
	requestResult := AirworthinessDirectiveDAO.RemoveWorkOrdersFromAirworthinessDirective(airworthinessDirectiveId, workOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
