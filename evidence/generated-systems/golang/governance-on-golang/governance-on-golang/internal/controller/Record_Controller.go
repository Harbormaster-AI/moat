package controller

import (
    Record_DAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to Record_DAO for database creation
//----------------------------------------------------------------------------
func CreateRecord_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Record_ model
	//----------------------------------------------------------------------------
	data := model.Record_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Record_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Record_ data access object to create
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.CreateRecord_( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to Record_DAO to find the relevant Record_
//----------------------------------------------------------------------------
func GetRecord_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Record_ data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.GetRecord_(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to Record_DAO for database read of all Record_s
//----------------------------------------------------------------------------
func GetAllRecord_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Record_ data access object to get all
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.GetAllRecord_()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to Record_DAO for database save
//----------------------------------------------------------------------------
func UpdateRecord_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Record_ model
	//----------------------------------------------------------------------------
	var data = model.Record_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Record_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Record_ data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.UpdateRecord_(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to Record_DAO for database deletion
//----------------------------------------------------------------------------
func DeleteRecord_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Record_ data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := Record_DAO.DeleteRecord_(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Repository on a Record_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRepositoryToRecord_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	repositoryId,_ := strconv.ParseUint( vars["repositoryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.AssignRepositoryToRecord_(record_Id, repositoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Repository on a Record_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRepositoryFromRecord_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.UnassignRepositoryFromRecord_(record_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a RetentionSchedule on a Record_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRetentionScheduleToRecord_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	retentionScheduleId,_ := strconv.ParseUint( vars["retentionScheduleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.AssignRetentionScheduleToRecord_(record_Id, retentionScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RetentionSchedule on a Record_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRetentionScheduleFromRecord_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.UnassignRetentionScheduleFromRecord_(record_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a Record_
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.AddProcessingActivitiesToRecord_(record_Id, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a Record_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.RemoveProcessingActivitiesFromRecord_(record_Id, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataCategoriesIds as a DataCategories to a Record_
	//----------------------------------------------------------------------------
func AddDataCategoriesToRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataCategoriesIds,_ := vars["dataCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.AddDataCategoriesToRecord_(record_Id, dataCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataCategoriesIds as a DataCategories from a Record_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataCategoriesFromRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataCategoriesIds,_ := vars["dataCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.RemoveDataCategoriesFromRecord_(record_Id, dataCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more legalHoldsIds as a LegalHolds to a Record_
	//----------------------------------------------------------------------------
func AddLegalHoldsToRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	legalHoldsIds,_ := vars["legalHoldsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.AddLegalHoldsToRecord_(record_Id, legalHoldsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more legalHoldsIds as a LegalHolds from a Record_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLegalHoldsFromRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	legalHoldsIds,_ := vars["legalHoldsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.RemoveLegalHoldsFromRecord_(record_Id, legalHoldsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataSubjectRequestsIds as a DataSubjectRequests to a Record_
	//----------------------------------------------------------------------------
func AddDataSubjectRequestsToRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataSubjectRequestsIds,_ := vars["dataSubjectRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.AddDataSubjectRequestsToRecord_(record_Id, dataSubjectRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataSubjectRequestsIds as a DataSubjectRequests from a Record_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataSubjectRequestsFromRecord_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	record_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataSubjectRequestsIds,_ := vars["dataSubjectRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Record_ DAO
	//----------------------------------------------------------------------------
	requestResult := Record_DAO.RemoveDataSubjectRequestsFromRecord_(record_Id, dataSubjectRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
