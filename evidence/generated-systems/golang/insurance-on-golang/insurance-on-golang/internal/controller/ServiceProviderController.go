package controller

import (
    ServiceProviderDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ServiceProviderDAO for database creation
//----------------------------------------------------------------------------
func CreateServiceProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ServiceProvider model
	//----------------------------------------------------------------------------
	data := model.ServiceProvider{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ServiceProvider model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ServiceProvider data access object to create
	//----------------------------------------------------------------------------
	requestResult := ServiceProviderDAO.CreateServiceProvider( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ServiceProviderDAO to find the relevant ServiceProvider
//----------------------------------------------------------------------------
func GetServiceProvider(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ServiceProvider data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ServiceProviderDAO.GetServiceProvider(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ServiceProviderDAO for database read of all ServiceProviders
//----------------------------------------------------------------------------
func GetAllServiceProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ServiceProvider data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ServiceProviderDAO.GetAllServiceProvider()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ServiceProviderDAO for database save
//----------------------------------------------------------------------------
func UpdateServiceProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ServiceProvider model
	//----------------------------------------------------------------------------
	var data = model.ServiceProvider{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ServiceProvider model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ServiceProvider data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ServiceProviderDAO.UpdateServiceProvider(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ServiceProviderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteServiceProvider(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ServiceProvider data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ServiceProviderDAO.DeleteServiceProvider(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a ServiceProvider
	//----------------------------------------------------------------------------
func AddClaimsToServiceProvider(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	serviceProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ServiceProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ServiceProviderDAO.AddClaimsToServiceProvider(serviceProviderId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a ServiceProvider
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromServiceProvider(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	serviceProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ServiceProvider DAO
	//----------------------------------------------------------------------------
	requestResult := ServiceProviderDAO.RemoveClaimsFromServiceProvider(serviceProviderId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
