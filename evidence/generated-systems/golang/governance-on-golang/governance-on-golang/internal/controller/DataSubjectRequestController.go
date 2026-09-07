package controller

import (
    DataSubjectRequestDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataSubjectRequestDAO for database creation
//----------------------------------------------------------------------------
func CreateDataSubjectRequest(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataSubjectRequest model
	//----------------------------------------------------------------------------
	data := model.DataSubjectRequest{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataSubjectRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.CreateDataSubjectRequest( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataSubjectRequestDAO to find the relevant DataSubjectRequest
//----------------------------------------------------------------------------
func GetDataSubjectRequest(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataSubjectRequest data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.GetDataSubjectRequest(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataSubjectRequestDAO for database read of all DataSubjectRequests
//----------------------------------------------------------------------------
func GetAllDataSubjectRequest(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.GetAllDataSubjectRequest()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataSubjectRequestDAO for database save
//----------------------------------------------------------------------------
func UpdateDataSubjectRequest(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataSubjectRequest model
	//----------------------------------------------------------------------------
	var data = model.DataSubjectRequest{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataSubjectRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.UpdateDataSubjectRequest(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataSubjectRequestDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataSubjectRequest(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataSubjectRequest data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataSubjectRequestDAO.DeleteDataSubjectRequest(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a DataSubjectRequest
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToDataSubjectRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSubjectRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest DAO
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.AssignOrganizationToDataSubjectRequest(dataSubjectRequestId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a DataSubjectRequest
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromDataSubjectRequest( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSubjectRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest DAO
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.UnassignOrganizationFromDataSubjectRequest(dataSubjectRequestId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a DataSubjectRequest
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToDataSubjectRequest(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSubjectRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest DAO
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.AddProcessingActivitiesToDataSubjectRequest(dataSubjectRequestId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a DataSubjectRequest
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromDataSubjectRequest(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSubjectRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest DAO
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.RemoveProcessingActivitiesFromDataSubjectRequest(dataSubjectRequestId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more recordsIds as a Records to a DataSubjectRequest
	//----------------------------------------------------------------------------
func AddRecordsToDataSubjectRequest(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSubjectRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest DAO
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.AddRecordsToDataSubjectRequest(dataSubjectRequestId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsIds as a Records from a DataSubjectRequest
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsFromDataSubjectRequest(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSubjectRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSubjectRequest DAO
	//----------------------------------------------------------------------------
	requestResult := DataSubjectRequestDAO.RemoveRecordsFromDataSubjectRequest(dataSubjectRequestId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
