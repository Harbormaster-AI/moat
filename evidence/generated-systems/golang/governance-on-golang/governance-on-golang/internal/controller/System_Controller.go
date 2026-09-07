package controller

import (
    System_DAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to System_DAO for database creation
//----------------------------------------------------------------------------
func CreateSystem_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty System_ model
	//----------------------------------------------------------------------------
	data := model.System_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a System_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the System_ data access object to create
	//----------------------------------------------------------------------------
	requestResult := System_DAO.CreateSystem_( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to System_DAO to find the relevant System_
//----------------------------------------------------------------------------
func GetSystem_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the System_ data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := System_DAO.GetSystem_(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to System_DAO for database read of all System_s
//----------------------------------------------------------------------------
func GetAllSystem_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the System_ data access object to get all
	//----------------------------------------------------------------------------
	requestResult := System_DAO.GetAllSystem_()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to System_DAO for database save
//----------------------------------------------------------------------------
func UpdateSystem_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty System_ model
	//----------------------------------------------------------------------------
	var data = model.System_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a System_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the System_ data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := System_DAO.UpdateSystem_(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to System_DAO for database deletion
//----------------------------------------------------------------------------
func DeleteSystem_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the System_ data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := System_DAO.DeleteSystem_(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a System_
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToSystem_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	system_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the System_ DAO
	//----------------------------------------------------------------------------
	requestResult := System_DAO.AddProcessingActivitiesToSystem_(system_Id, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a System_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromSystem_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	system_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the System_ DAO
	//----------------------------------------------------------------------------
	requestResult := System_DAO.RemoveProcessingActivitiesFromSystem_(system_Id, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more recordsRepositoriesIds as a RecordsRepositories to a System_
	//----------------------------------------------------------------------------
func AddRecordsRepositoriesToSystem_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	system_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsRepositoriesIds,_ := vars["recordsRepositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the System_ DAO
	//----------------------------------------------------------------------------
	requestResult := System_DAO.AddRecordsRepositoriesToSystem_(system_Id, recordsRepositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsRepositoriesIds as a RecordsRepositories from a System_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsRepositoriesFromSystem_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	system_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsRepositoriesIds,_ := vars["recordsRepositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the System_ DAO
	//----------------------------------------------------------------------------
	requestResult := System_DAO.RemoveRecordsRepositoriesFromSystem_(system_Id, recordsRepositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
