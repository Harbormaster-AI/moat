package controller

import (
    RetentionScheduleDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RetentionScheduleDAO for database creation
//----------------------------------------------------------------------------
func CreateRetentionSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RetentionSchedule model
	//----------------------------------------------------------------------------
	data := model.RetentionSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RetentionSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule data access object to create
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.CreateRetentionSchedule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RetentionScheduleDAO to find the relevant RetentionSchedule
//----------------------------------------------------------------------------
func GetRetentionSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RetentionSchedule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.GetRetentionSchedule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RetentionScheduleDAO for database read of all RetentionSchedules
//----------------------------------------------------------------------------
func GetAllRetentionSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.GetAllRetentionSchedule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RetentionScheduleDAO for database save
//----------------------------------------------------------------------------
func UpdateRetentionSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RetentionSchedule model
	//----------------------------------------------------------------------------
	var data = model.RetentionSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RetentionSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.UpdateRetentionSchedule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RetentionScheduleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRetentionSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RetentionSchedule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RetentionScheduleDAO.DeleteRetentionSchedule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more repositoriesIds as a Repositories to a RetentionSchedule
	//----------------------------------------------------------------------------
func AddRepositoriesToRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	repositoriesIds,_ := vars["repositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.AddRepositoriesToRetentionSchedule(retentionScheduleId, repositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more repositoriesIds as a Repositories from a RetentionSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRepositoriesFromRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	repositoriesIds,_ := vars["repositoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.RemoveRepositoriesFromRetentionSchedule(retentionScheduleId, repositoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more recordsIds as a Records to a RetentionSchedule
	//----------------------------------------------------------------------------
func AddRecordsToRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.AddRecordsToRetentionSchedule(retentionScheduleId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsIds as a Records from a RetentionSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsFromRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.RemoveRecordsFromRetentionSchedule(retentionScheduleId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more exceptionsIds as a Exceptions to a RetentionSchedule
	//----------------------------------------------------------------------------
func AddExceptionsToRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exceptionsIds,_ := vars["exceptionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.AddExceptionsToRetentionSchedule(retentionScheduleId, exceptionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more exceptionsIds as a Exceptions from a RetentionSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExceptionsFromRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exceptionsIds,_ := vars["exceptionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.RemoveExceptionsFromRetentionSchedule(retentionScheduleId, exceptionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dispositionReviewsIds as a DispositionReviews to a RetentionSchedule
	//----------------------------------------------------------------------------
func AddDispositionReviewsToRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dispositionReviewsIds,_ := vars["dispositionReviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.AddDispositionReviewsToRetentionSchedule(retentionScheduleId, dispositionReviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dispositionReviewsIds as a DispositionReviews from a RetentionSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDispositionReviewsFromRetentionSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	retentionScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dispositionReviewsIds,_ := vars["dispositionReviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RetentionSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RetentionScheduleDAO.RemoveDispositionReviewsFromRetentionSchedule(retentionScheduleId, dispositionReviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
