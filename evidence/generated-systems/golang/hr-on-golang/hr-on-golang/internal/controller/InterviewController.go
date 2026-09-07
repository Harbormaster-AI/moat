package controller

import (
    InterviewDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InterviewDAO for database creation
//----------------------------------------------------------------------------
func CreateInterview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Interview model
	//----------------------------------------------------------------------------
	data := model.Interview{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Interview model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Interview data access object to create
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.CreateInterview( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InterviewDAO to find the relevant Interview
//----------------------------------------------------------------------------
func GetInterview(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Interview data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.GetInterview(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InterviewDAO for database read of all Interviews
//----------------------------------------------------------------------------
func GetAllInterview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Interview data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.GetAllInterview()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InterviewDAO for database save
//----------------------------------------------------------------------------
func UpdateInterview(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Interview model
	//----------------------------------------------------------------------------
	var data = model.Interview{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Interview model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Interview data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.UpdateInterview(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InterviewDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInterview(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Interview data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InterviewDAO.DeleteInterview(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Requisition on a Interview
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRequisitionToInterview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	interviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	requisitionId,_ := strconv.ParseUint( vars["requisitionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Interview DAO
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.AssignRequisitionToInterview(interviewId, requisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Requisition on a Interview
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRequisitionFromInterview( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	interviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Interview DAO
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.UnassignRequisitionFromInterview(interviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Candidate on a Interview
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCandidateToInterview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	interviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	candidateId,_ := strconv.ParseUint( vars["candidateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Interview DAO
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.AssignCandidateToInterview(interviewId, candidateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Candidate on a Interview
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCandidateFromInterview( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	interviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Interview DAO
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.UnassignCandidateFromInterview(interviewId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more interviewersIds as a Interviewers to a Interview
	//----------------------------------------------------------------------------
func AddInterviewersToInterview(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	interviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	interviewersIds,_ := vars["interviewersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Interview DAO
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.AddInterviewersToInterview(interviewId, interviewersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more interviewersIds as a Interviewers from a Interview
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInterviewersFromInterview(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	interviewId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	interviewersIds,_ := vars["interviewersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Interview DAO
	//----------------------------------------------------------------------------
	requestResult := InterviewDAO.RemoveInterviewersFromInterview(interviewId, interviewersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
