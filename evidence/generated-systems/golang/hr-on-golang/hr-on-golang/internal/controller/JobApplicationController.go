package controller

import (
    JobApplicationDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to JobApplicationDAO for database creation
//----------------------------------------------------------------------------
func CreateJobApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobApplication model
	//----------------------------------------------------------------------------
	data := model.JobApplication{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobApplication model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication data access object to create
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.CreateJobApplication( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to JobApplicationDAO to find the relevant JobApplication
//----------------------------------------------------------------------------
func GetJobApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobApplication data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.GetJobApplication(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to JobApplicationDAO for database read of all JobApplications
//----------------------------------------------------------------------------
func GetAllJobApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the JobApplication data access object to get all
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.GetAllJobApplication()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to JobApplicationDAO for database save
//----------------------------------------------------------------------------
func UpdateJobApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobApplication model
	//----------------------------------------------------------------------------
	var data = model.JobApplication{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobApplication model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.UpdateJobApplication(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to JobApplicationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteJobApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobApplication data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := JobApplicationDAO.DeleteJobApplication(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Candidate on a JobApplication
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCandidateToJobApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	candidateId,_ := strconv.ParseUint( vars["candidateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication DAO
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.AssignCandidateToJobApplication(jobApplicationId, candidateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Candidate on a JobApplication
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCandidateFromJobApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication DAO
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.UnassignCandidateFromJobApplication(jobApplicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Requisition on a JobApplication
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRequisitionToJobApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	requisitionId,_ := strconv.ParseUint( vars["requisitionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication DAO
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.AssignRequisitionToJobApplication(jobApplicationId, requisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Requisition on a JobApplication
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRequisitionFromJobApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication DAO
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.UnassignRequisitionFromJobApplication(jobApplicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more screeningsIds as a Screenings to a JobApplication
	//----------------------------------------------------------------------------
func AddScreeningsToJobApplication(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	screeningsIds,_ := vars["screeningsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication DAO
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.AddScreeningsToJobApplication(jobApplicationId, screeningsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more screeningsIds as a Screenings from a JobApplication
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveScreeningsFromJobApplication(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	screeningsIds,_ := vars["screeningsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobApplication DAO
	//----------------------------------------------------------------------------
	requestResult := JobApplicationDAO.RemoveScreeningsFromJobApplication(jobApplicationId, screeningsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
