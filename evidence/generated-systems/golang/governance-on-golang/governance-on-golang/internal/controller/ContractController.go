package controller

import (
    ContractDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ContractDAO for database creation
//----------------------------------------------------------------------------
func CreateContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Contract model
	//----------------------------------------------------------------------------
	data := model.Contract{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Contract model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Contract data access object to create
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.CreateContract( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ContractDAO to find the relevant Contract
//----------------------------------------------------------------------------
func GetContract(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Contract data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.GetContract(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ContractDAO for database read of all Contracts
//----------------------------------------------------------------------------
func GetAllContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Contract data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.GetAllContract()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ContractDAO for database save
//----------------------------------------------------------------------------
func UpdateContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Contract model
	//----------------------------------------------------------------------------
	var data = model.Contract{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Contract model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Contract data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.UpdateContract(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ContractDAO for database deletion
//----------------------------------------------------------------------------
func DeleteContract(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Contract data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ContractDAO.DeleteContract(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ThirdParty on a Contract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignThirdPartyToContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartyId,_ := strconv.ParseUint( vars["thirdPartyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.AssignThirdPartyToContract(contractId, thirdPartyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ThirdParty on a Contract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignThirdPartyFromContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.UnassignThirdPartyFromContract(contractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Matter on a Contract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMatterToContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	matterId,_ := strconv.ParseUint( vars["matterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.AssignMatterToContract(contractId, matterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Matter on a Contract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMatterFromContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.UnassignMatterFromContract(contractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more obligationsIds as a Obligations to a Contract
	//----------------------------------------------------------------------------
func AddObligationsToContract(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.AddObligationsToContract(contractId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more obligationsIds as a Obligations from a Contract
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObligationsFromContract(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.RemoveObligationsFromContract(contractId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataProcessingActivitiesIds as a DataProcessingActivities to a Contract
	//----------------------------------------------------------------------------
func AddDataProcessingActivitiesToContract(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataProcessingActivitiesIds,_ := vars["dataProcessingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.AddDataProcessingActivitiesToContract(contractId, dataProcessingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataProcessingActivitiesIds as a DataProcessingActivities from a Contract
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataProcessingActivitiesFromContract(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	contractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataProcessingActivitiesIds,_ := vars["dataProcessingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Contract DAO
	//----------------------------------------------------------------------------
	requestResult := ContractDAO.RemoveDataProcessingActivitiesFromContract(contractId, dataProcessingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
