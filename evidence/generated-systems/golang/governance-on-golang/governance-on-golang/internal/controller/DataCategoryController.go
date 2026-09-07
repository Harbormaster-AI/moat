package controller

import (
    DataCategoryDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataCategoryDAO for database creation
//----------------------------------------------------------------------------
func CreateDataCategory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataCategory model
	//----------------------------------------------------------------------------
	data := model.DataCategory{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataCategory model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.CreateDataCategory( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataCategoryDAO to find the relevant DataCategory
//----------------------------------------------------------------------------
func GetDataCategory(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataCategory data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.GetDataCategory(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataCategoryDAO for database read of all DataCategorys
//----------------------------------------------------------------------------
func GetAllDataCategory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataCategory data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.GetAllDataCategory()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataCategoryDAO for database save
//----------------------------------------------------------------------------
func UpdateDataCategory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataCategory model
	//----------------------------------------------------------------------------
	var data = model.DataCategory{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataCategory model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.UpdateDataCategory(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataCategoryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataCategory(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataCategory data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataCategoryDAO.DeleteDataCategory(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a DataCategory
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToDataCategory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataCategoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory DAO
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.AddProcessingActivitiesToDataCategory(dataCategoryId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a DataCategory
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromDataCategory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataCategoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory DAO
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.RemoveProcessingActivitiesFromDataCategory(dataCategoryId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more recordsIds as a Records to a DataCategory
	//----------------------------------------------------------------------------
func AddRecordsToDataCategory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataCategoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory DAO
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.AddRecordsToDataCategory(dataCategoryId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsIds as a Records from a DataCategory
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsFromDataCategory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataCategoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory DAO
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.RemoveRecordsFromDataCategory(dataCategoryId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataBreachesIds as a DataBreaches to a DataCategory
	//----------------------------------------------------------------------------
func AddDataBreachesToDataCategory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataCategoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory DAO
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.AddDataBreachesToDataCategory(dataCategoryId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataBreachesIds as a DataBreaches from a DataCategory
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataBreachesFromDataCategory(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataCategoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataCategory DAO
	//----------------------------------------------------------------------------
	requestResult := DataCategoryDAO.RemoveDataBreachesFromDataCategory(dataCategoryId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
