package controller

import (
    DispositionReviewDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DispositionReviewDAO for database creation
//----------------------------------------------------------------------------
func CreateDispositionReview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DispositionReview model
	//----------------------------------------------------------------------------
	data := model.DispositionReview{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DispositionReview model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DispositionReview data access object to create
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.CreateDispositionReview( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DispositionReviewDAO to find the relevant DispositionReview
//----------------------------------------------------------------------------
func GetDispositionReview(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DispositionReview data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.GetDispositionReview(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DispositionReviewDAO for database read of all DispositionReviews
//----------------------------------------------------------------------------
func GetAllDispositionReview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DispositionReview data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.GetAllDispositionReview()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DispositionReviewDAO for database save
//----------------------------------------------------------------------------
func UpdateDispositionReview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DispositionReview model
	//----------------------------------------------------------------------------
	var data = model.DispositionReview{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DispositionReview model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DispositionReview data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.UpdateDispositionReview(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DispositionReviewDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDispositionReview(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DispositionReview data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DispositionReviewDAO.DeleteDispositionReview(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Record on a DispositionReview
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRecordToDispositionReview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dispositionReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordId,_ := strconv.ParseUint( vars["recordId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DispositionReview DAO
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.AssignRecordToDispositionReview(dispositionReviewId, recordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Record on a DispositionReview
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRecordFromDispositionReview( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dispositionReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DispositionReview DAO
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.UnassignRecordFromDispositionReview(dispositionReviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a RetentionSchedule on a DispositionReview
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRetentionScheduleToDispositionReview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dispositionReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	retentionScheduleId,_ := strconv.ParseUint( vars["retentionScheduleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DispositionReview DAO
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.AssignRetentionScheduleToDispositionReview(dispositionReviewId, retentionScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RetentionSchedule on a DispositionReview
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRetentionScheduleFromDispositionReview( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dispositionReviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DispositionReview DAO
	//----------------------------------------------------------------------------
	requestResult := DispositionReviewDAO.UnassignRetentionScheduleFromDispositionReview(dispositionReviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


