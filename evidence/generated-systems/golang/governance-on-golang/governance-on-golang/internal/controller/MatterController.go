package controller

import (
    MatterDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MatterDAO for database creation
//----------------------------------------------------------------------------
func CreateMatter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Matter model
	//----------------------------------------------------------------------------
	data := model.Matter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Matter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Matter data access object to create
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.CreateMatter( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MatterDAO to find the relevant Matter
//----------------------------------------------------------------------------
func GetMatter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Matter data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.GetMatter(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MatterDAO for database read of all Matters
//----------------------------------------------------------------------------
func GetAllMatter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Matter data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.GetAllMatter()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MatterDAO for database save
//----------------------------------------------------------------------------
func UpdateMatter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Matter model
	//----------------------------------------------------------------------------
	var data = model.Matter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Matter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Matter data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.UpdateMatter(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MatterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMatter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Matter data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MatterDAO.DeleteMatter(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Matter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToMatter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.AssignOrganizationToMatter(matterId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Matter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromMatter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.UnassignOrganizationFromMatter(matterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more legalHoldsIds as a LegalHolds to a Matter
	//----------------------------------------------------------------------------
func AddLegalHoldsToMatter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	legalHoldsIds,_ := vars["legalHoldsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.AddLegalHoldsToMatter(matterId, legalHoldsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more legalHoldsIds as a LegalHolds from a Matter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLegalHoldsFromMatter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	legalHoldsIds,_ := vars["legalHoldsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.RemoveLegalHoldsFromMatter(matterId, legalHoldsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataBreachesIds as a DataBreaches to a Matter
	//----------------------------------------------------------------------------
func AddDataBreachesToMatter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.AddDataBreachesToMatter(matterId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataBreachesIds as a DataBreaches from a Matter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataBreachesFromMatter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.RemoveDataBreachesFromMatter(matterId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a Matter
	//----------------------------------------------------------------------------
func AddContractsToMatter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.AddContractsToMatter(matterId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a Matter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromMatter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	matterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Matter DAO
	//----------------------------------------------------------------------------
	requestResult := MatterDAO.RemoveContractsFromMatter(matterId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
