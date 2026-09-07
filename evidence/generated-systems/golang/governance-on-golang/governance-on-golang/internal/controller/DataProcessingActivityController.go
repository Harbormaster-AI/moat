package controller

import (
    DataProcessingActivityDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataProcessingActivityDAO for database creation
//----------------------------------------------------------------------------
func CreateDataProcessingActivity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataProcessingActivity model
	//----------------------------------------------------------------------------
	data := model.DataProcessingActivity{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataProcessingActivity model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.CreateDataProcessingActivity( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataProcessingActivityDAO to find the relevant DataProcessingActivity
//----------------------------------------------------------------------------
func GetDataProcessingActivity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataProcessingActivity data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.GetDataProcessingActivity(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataProcessingActivityDAO for database read of all DataProcessingActivitys
//----------------------------------------------------------------------------
func GetAllDataProcessingActivity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.GetAllDataProcessingActivity()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataProcessingActivityDAO for database save
//----------------------------------------------------------------------------
func UpdateDataProcessingActivity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataProcessingActivity model
	//----------------------------------------------------------------------------
	var data = model.DataProcessingActivity{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataProcessingActivity model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.UpdateDataProcessingActivity(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataProcessingActivityDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataProcessingActivity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataProcessingActivity data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataProcessingActivityDAO.DeleteDataProcessingActivity(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a DataProcessingActivity
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToDataProcessingActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AssignOrganizationToDataProcessingActivity(dataProcessingActivityId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a DataProcessingActivity
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromDataProcessingActivity( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.UnassignOrganizationFromDataProcessingActivity(dataProcessingActivityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more dataCategoriesIds as a DataCategories to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddDataCategoriesToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataCategoriesIds,_ := vars["dataCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddDataCategoriesToDataProcessingActivity(dataProcessingActivityId, dataCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataCategoriesIds as a DataCategories from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataCategoriesFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataCategoriesIds,_ := vars["dataCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemoveDataCategoriesFromDataProcessingActivity(dataProcessingActivityId, dataCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more systemsIds as a Systems to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddSystemsToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	systemsIds,_ := vars["systemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddSystemsToDataProcessingActivity(dataProcessingActivityId, systemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more systemsIds as a Systems from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSystemsFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	systemsIds,_ := vars["systemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemoveSystemsFromDataProcessingActivity(dataProcessingActivityId, systemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more recordsIds as a Records to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddRecordsToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddRecordsToDataProcessingActivity(dataProcessingActivityId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsIds as a Records from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemoveRecordsFromDataProcessingActivity(dataProcessingActivityId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more privacyNoticesIds as a PrivacyNotices to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddPrivacyNoticesToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	privacyNoticesIds,_ := vars["privacyNoticesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddPrivacyNoticesToDataProcessingActivity(dataProcessingActivityId, privacyNoticesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more privacyNoticesIds as a PrivacyNotices from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePrivacyNoticesFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	privacyNoticesIds,_ := vars["privacyNoticesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemovePrivacyNoticesFromDataProcessingActivity(dataProcessingActivityId, privacyNoticesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more thirdPartiesIds as a ThirdParties to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddThirdPartiesToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartiesIds,_ := vars["thirdPartiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddThirdPartiesToDataProcessingActivity(dataProcessingActivityId, thirdPartiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more thirdPartiesIds as a ThirdParties from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveThirdPartiesFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartiesIds,_ := vars["thirdPartiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemoveThirdPartiesFromDataProcessingActivity(dataProcessingActivityId, thirdPartiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more consentsIds as a Consents to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddConsentsToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddConsentsToDataProcessingActivity(dataProcessingActivityId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more consentsIds as a Consents from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConsentsFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	consentsIds,_ := vars["consentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemoveConsentsFromDataProcessingActivity(dataProcessingActivityId, consentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataBreachesIds as a DataBreaches to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddDataBreachesToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddDataBreachesToDataProcessingActivity(dataProcessingActivityId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataBreachesIds as a DataBreaches from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataBreachesFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemoveDataBreachesFromDataProcessingActivity(dataProcessingActivityId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataSubjectRequestsIds as a DataSubjectRequests to a DataProcessingActivity
	//----------------------------------------------------------------------------
func AddDataSubjectRequestsToDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataSubjectRequestsIds,_ := vars["dataSubjectRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.AddDataSubjectRequestsToDataProcessingActivity(dataProcessingActivityId, dataSubjectRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataSubjectRequestsIds as a DataSubjectRequests from a DataProcessingActivity
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataSubjectRequestsFromDataProcessingActivity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataProcessingActivityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataSubjectRequestsIds,_ := vars["dataSubjectRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataProcessingActivity DAO
	//----------------------------------------------------------------------------
	requestResult := DataProcessingActivityDAO.RemoveDataSubjectRequestsFromDataProcessingActivity(dataProcessingActivityId, dataSubjectRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
