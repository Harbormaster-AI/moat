package controller

import (
    ControlTest_DAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ControlTest_DAO for database creation
//----------------------------------------------------------------------------
func CreateControlTest_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ControlTest_ model
	//----------------------------------------------------------------------------
	data := model.ControlTest_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ControlTest_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ data access object to create
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.CreateControlTest_( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ControlTest_DAO to find the relevant ControlTest_
//----------------------------------------------------------------------------
func GetControlTest_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ControlTest_ data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.GetControlTest_(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ControlTest_DAO for database read of all ControlTest_s
//----------------------------------------------------------------------------
func GetAllControlTest_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.GetAllControlTest_()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ControlTest_DAO for database save
//----------------------------------------------------------------------------
func UpdateControlTest_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ControlTest_ model
	//----------------------------------------------------------------------------
	var data = model.ControlTest_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ControlTest_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.UpdateControlTest_(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ControlTest_DAO for database deletion
//----------------------------------------------------------------------------
func DeleteControlTest_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ControlTest_ data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ControlTest_DAO.DeleteControlTest_(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Control on a ControlTest_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignControlToControlTest_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	controlTest_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlId,_ := strconv.ParseUint( vars["controlId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ DAO
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.AssignControlToControlTest_(controlTest_Id, controlId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Control on a ControlTest_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignControlFromControlTest_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	controlTest_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ DAO
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.UnassignControlFromControlTest_(controlTest_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Engagement on a ControlTest_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEngagementToControlTest_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	controlTest_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engagementId,_ := strconv.ParseUint( vars["engagementId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ DAO
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.AssignEngagementToControlTest_(controlTest_Id, engagementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Engagement on a ControlTest_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEngagementFromControlTest_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	controlTest_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ DAO
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.UnassignEngagementFromControlTest_(controlTest_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more evidenceIds as a Evidence to a ControlTest_
	//----------------------------------------------------------------------------
func AddEvidenceToControlTest_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlTest_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evidenceIds,_ := vars["evidenceIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ DAO
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.AddEvidenceToControlTest_(controlTest_Id, evidenceIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more evidenceIds as a Evidence from a ControlTest_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEvidenceFromControlTest_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	controlTest_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evidenceIds,_ := vars["evidenceIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ControlTest_ DAO
	//----------------------------------------------------------------------------
	requestResult := ControlTest_DAO.RemoveEvidenceFromControlTest_(controlTest_Id, evidenceIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
