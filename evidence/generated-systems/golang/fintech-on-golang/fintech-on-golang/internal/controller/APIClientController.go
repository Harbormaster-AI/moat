package controller

import (
    APIClientDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to APIClientDAO for database creation
//----------------------------------------------------------------------------
func CreateAPIClient(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty APIClient model
	//----------------------------------------------------------------------------
	data := model.APIClient{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a APIClient model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the APIClient data access object to create
	//----------------------------------------------------------------------------
	requestResult := APIClientDAO.CreateAPIClient( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to APIClientDAO to find the relevant APIClient
//----------------------------------------------------------------------------
func GetAPIClient(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the APIClient data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := APIClientDAO.GetAPIClient(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to APIClientDAO for database read of all APIClients
//----------------------------------------------------------------------------
func GetAllAPIClient(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the APIClient data access object to get all
	//----------------------------------------------------------------------------
	requestResult := APIClientDAO.GetAllAPIClient()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to APIClientDAO for database save
//----------------------------------------------------------------------------
func UpdateAPIClient(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty APIClient model
	//----------------------------------------------------------------------------
	var data = model.APIClient{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a APIClient model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the APIClient data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := APIClientDAO.UpdateAPIClient(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to APIClientDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAPIClient(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the APIClient data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := APIClientDAO.DeleteAPIClient(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a APIClient
	//----------------------------------------------------------------------------
func AddConsentsToAPIClient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aPIClientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the APIClient DAO
	//----------------------------------------------------------------------------
	requestResult := APIClientDAO.AddConsentsToAPIClient(aPIClientId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a APIClient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConsentsFromAPIClient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aPIClientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the APIClient DAO
	//----------------------------------------------------------------------------
	requestResult := APIClientDAO.RemoveConsentsFromAPIClient(aPIClientId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
