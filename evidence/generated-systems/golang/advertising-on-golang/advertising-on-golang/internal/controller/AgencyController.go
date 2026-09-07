package controller

import (
    AgencyDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AgencyDAO for database creation
//----------------------------------------------------------------------------
func CreateAgency(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Agency model
	//----------------------------------------------------------------------------
	data := model.Agency{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Agency model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Agency data access object to create
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.CreateAgency( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AgencyDAO to find the relevant Agency
//----------------------------------------------------------------------------
func GetAgency(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Agency data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.GetAgency(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AgencyDAO for database read of all Agencys
//----------------------------------------------------------------------------
func GetAllAgency(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Agency data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.GetAllAgency()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AgencyDAO for database save
//----------------------------------------------------------------------------
func UpdateAgency(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Agency model
	//----------------------------------------------------------------------------
	var data = model.Agency{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Agency model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Agency data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.UpdateAgency(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AgencyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAgency(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Agency data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AgencyDAO.DeleteAgency(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more advertisersIds as a Advertisers to a Agency
	//----------------------------------------------------------------------------
func AddAdvertisersToAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	advertisersIds,_ := vars["advertisersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.AddAdvertisersToAgency(agencyId, advertisersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more advertisersIds as a Advertisers from a Agency
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAdvertisersFromAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	advertisersIds,_ := vars["advertisersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.RemoveAdvertisersFromAgency(agencyId, advertisersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more teamsIds as a Teams to a Agency
	//----------------------------------------------------------------------------
func AddTeamsToAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.AddTeamsToAgency(agencyId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more teamsIds as a Teams from a Agency
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTeamsFromAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.RemoveTeamsFromAgency(agencyId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more usersIds as a Users to a Agency
	//----------------------------------------------------------------------------
func AddUsersToAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.AddUsersToAgency(agencyId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usersIds as a Users from a Agency
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsersFromAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.RemoveUsersFromAgency(agencyId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more insertionOrdersIds as a InsertionOrders to a Agency
	//----------------------------------------------------------------------------
func AddInsertionOrdersToAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insertionOrdersIds,_ := vars["insertionOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.AddInsertionOrdersToAgency(agencyId, insertionOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more insertionOrdersIds as a InsertionOrders from a Agency
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInsertionOrdersFromAgency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	agencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insertionOrdersIds,_ := vars["insertionOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Agency DAO
	//----------------------------------------------------------------------------
	requestResult := AgencyDAO.RemoveInsertionOrdersFromAgency(agencyId, insertionOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
