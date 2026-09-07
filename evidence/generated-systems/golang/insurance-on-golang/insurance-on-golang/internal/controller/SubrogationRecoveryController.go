package controller

import (
    SubrogationRecoveryDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SubrogationRecoveryDAO for database creation
//----------------------------------------------------------------------------
func CreateSubrogationRecovery(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SubrogationRecovery model
	//----------------------------------------------------------------------------
	data := model.SubrogationRecovery{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SubrogationRecovery model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery data access object to create
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.CreateSubrogationRecovery( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SubrogationRecoveryDAO to find the relevant SubrogationRecovery
//----------------------------------------------------------------------------
func GetSubrogationRecovery(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SubrogationRecovery data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.GetSubrogationRecovery(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SubrogationRecoveryDAO for database read of all SubrogationRecoverys
//----------------------------------------------------------------------------
func GetAllSubrogationRecovery(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.GetAllSubrogationRecovery()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SubrogationRecoveryDAO for database save
//----------------------------------------------------------------------------
func UpdateSubrogationRecovery(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SubrogationRecovery model
	//----------------------------------------------------------------------------
	var data = model.SubrogationRecovery{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SubrogationRecovery model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.UpdateSubrogationRecovery(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SubrogationRecoveryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSubrogationRecovery(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SubrogationRecovery data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SubrogationRecoveryDAO.DeleteSubrogationRecovery(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Claim on a SubrogationRecovery
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClaimToSubrogationRecovery(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	subrogationRecoveryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimId,_ := strconv.ParseUint( vars["claimId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery DAO
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.AssignClaimToSubrogationRecovery(subrogationRecoveryId, claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Claim on a SubrogationRecovery
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClaimFromSubrogationRecovery( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	subrogationRecoveryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery DAO
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.UnassignClaimFromSubrogationRecovery(subrogationRecoveryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Exposure on a SubrogationRecovery
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignExposureToSubrogationRecovery(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	subrogationRecoveryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exposureId,_ := strconv.ParseUint( vars["exposureId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery DAO
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.AssignExposureToSubrogationRecovery(subrogationRecoveryId, exposureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Exposure on a SubrogationRecovery
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignExposureFromSubrogationRecovery( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	subrogationRecoveryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery DAO
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.UnassignExposureFromSubrogationRecovery(subrogationRecoveryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Counterparty on a SubrogationRecovery
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCounterpartyToSubrogationRecovery(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	subrogationRecoveryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	counterpartyId,_ := strconv.ParseUint( vars["counterpartyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery DAO
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.AssignCounterpartyToSubrogationRecovery(subrogationRecoveryId, counterpartyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Counterparty on a SubrogationRecovery
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCounterpartyFromSubrogationRecovery( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	subrogationRecoveryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SubrogationRecovery DAO
	//----------------------------------------------------------------------------
	requestResult := SubrogationRecoveryDAO.UnassignCounterpartyFromSubrogationRecovery(subrogationRecoveryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


