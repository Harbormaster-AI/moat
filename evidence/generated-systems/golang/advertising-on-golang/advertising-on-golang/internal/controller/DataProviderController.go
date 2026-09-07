package controller

import (
    DataProviderDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataProviderDAO for database creation
//----------------------------------------------------------------------------
func CreateDataProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataProvider model
	//----------------------------------------------------------------------------
	data := model.DataProvider{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataProvider model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataProvider data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataProviderDAO.CreateDataProvider( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataProviderDAO to find the relevant DataProvider
//----------------------------------------------------------------------------
func GetDataProvider(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataProvider data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataProviderDAO.GetDataProvider(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataProviderDAO for database read of all DataProviders
//----------------------------------------------------------------------------
func GetAllDataProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataProvider data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataProviderDAO.GetAllDataProvider()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataProviderDAO for database save
//----------------------------------------------------------------------------
func UpdateDataProvider(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataProvider model
	//----------------------------------------------------------------------------
	var data = model.DataProvider{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataProvider model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataProvider data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataProviderDAO.UpdateDataProvider(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataProviderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataProvider(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataProvider data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataProviderDAO.DeleteDataProvider(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more audienceSegmentsIds as a AudienceSegments to a DataProvider
	//----------------------------------------------------------------------------
func AddAudienceSegmentsToDataProvider(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	audienceSegmentsIds,_ := vars["audienceSegmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProvider DAO
	//----------------------------------------------------------------------------
	requestResult := DataProviderDAO.AddAudienceSegmentsToDataProvider(dataProviderId, audienceSegmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more audienceSegmentsIds as a AudienceSegments from a DataProvider
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAudienceSegmentsFromDataProvider(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProviderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	audienceSegmentsIds,_ := vars["audienceSegmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProvider DAO
	//----------------------------------------------------------------------------
	requestResult := DataProviderDAO.RemoveAudienceSegmentsFromDataProvider(dataProviderId, audienceSegmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
