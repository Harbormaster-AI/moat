package controller

import (
    ObligationDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ObligationDAO for database creation
//----------------------------------------------------------------------------
func CreateObligation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Obligation model
	//----------------------------------------------------------------------------
	data := model.Obligation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Obligation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Obligation data access object to create
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.CreateObligation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ObligationDAO to find the relevant Obligation
//----------------------------------------------------------------------------
func GetObligation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Obligation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.GetObligation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ObligationDAO for database read of all Obligations
//----------------------------------------------------------------------------
func GetAllObligation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Obligation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.GetAllObligation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ObligationDAO for database save
//----------------------------------------------------------------------------
func UpdateObligation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Obligation model
	//----------------------------------------------------------------------------
	var data = model.Obligation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Obligation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Obligation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.UpdateObligation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ObligationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteObligation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Obligation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ObligationDAO.DeleteObligation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Regulation on a Obligation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRegulationToObligation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	regulationId,_ := strconv.ParseUint( vars["regulationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.AssignRegulationToObligation(obligationId, regulationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Regulation on a Obligation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRegulationFromObligation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.UnassignRegulationFromObligation(obligationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more controlsIds as a Controls to a Obligation
	//----------------------------------------------------------------------------
func AddControlsToObligation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.AddControlsToObligation(obligationId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more controlsIds as a Controls from a Obligation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveControlsFromObligation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlsIds,_ := vars["controlsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.RemoveControlsFromObligation(obligationId, controlsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a Obligation
	//----------------------------------------------------------------------------
func AddPoliciesToObligation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.AddPoliciesToObligation(obligationId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a Obligation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromObligation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.RemovePoliciesFromObligation(obligationId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a Obligation
	//----------------------------------------------------------------------------
func AddContractsToObligation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.AddContractsToObligation(obligationId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a Obligation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromObligation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	obligationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Obligation DAO
	//----------------------------------------------------------------------------
	requestResult := ObligationDAO.RemoveContractsFromObligation(obligationId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
