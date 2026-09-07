package controller

import (
    LegalHoldDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LegalHoldDAO for database creation
//----------------------------------------------------------------------------
func CreateLegalHold(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LegalHold model
	//----------------------------------------------------------------------------
	data := model.LegalHold{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LegalHold model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold data access object to create
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.CreateLegalHold( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LegalHoldDAO to find the relevant LegalHold
//----------------------------------------------------------------------------
func GetLegalHold(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LegalHold data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.GetLegalHold(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LegalHoldDAO for database read of all LegalHolds
//----------------------------------------------------------------------------
func GetAllLegalHold(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LegalHold data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.GetAllLegalHold()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LegalHoldDAO for database save
//----------------------------------------------------------------------------
func UpdateLegalHold(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LegalHold model
	//----------------------------------------------------------------------------
	var data = model.LegalHold{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LegalHold model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.UpdateLegalHold(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LegalHoldDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLegalHold(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LegalHold data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LegalHoldDAO.DeleteLegalHold(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Matter on a LegalHold
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMatterToLegalHold(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	legalHoldId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	matterId,_ := strconv.ParseUint( vars["matterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold DAO
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.AssignMatterToLegalHold(legalHoldId, matterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Matter on a LegalHold
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMatterFromLegalHold( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	legalHoldId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold DAO
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.UnassignMatterFromLegalHold(legalHoldId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more repositoriesIds as a Repositories to a LegalHold
	//----------------------------------------------------------------------------
func AddRepositoriesToLegalHold(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	legalHoldId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	repositoriesIds,_ := vars["repositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold DAO
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.AddRepositoriesToLegalHold(legalHoldId, repositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more repositoriesIds as a Repositories from a LegalHold
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRepositoriesFromLegalHold(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	legalHoldId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	repositoriesIds,_ := vars["repositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold DAO
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.RemoveRepositoriesFromLegalHold(legalHoldId, repositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more recordsIds as a Records to a LegalHold
	//----------------------------------------------------------------------------
func AddRecordsToLegalHold(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	legalHoldId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold DAO
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.AddRecordsToLegalHold(legalHoldId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsIds as a Records from a LegalHold
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsFromLegalHold(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	legalHoldId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LegalHold DAO
	//----------------------------------------------------------------------------
	requestResult := LegalHoldDAO.RemoveRecordsFromLegalHold(legalHoldId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
