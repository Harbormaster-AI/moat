package controller

import (
    JobRequisitionDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to JobRequisitionDAO for database creation
//----------------------------------------------------------------------------
func CreateJobRequisition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobRequisition model
	//----------------------------------------------------------------------------
	data := model.JobRequisition{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobRequisition model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition data access object to create
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.CreateJobRequisition( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to JobRequisitionDAO to find the relevant JobRequisition
//----------------------------------------------------------------------------
func GetJobRequisition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobRequisition data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.GetJobRequisition(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to JobRequisitionDAO for database read of all JobRequisitions
//----------------------------------------------------------------------------
func GetAllJobRequisition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition data access object to get all
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.GetAllJobRequisition()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to JobRequisitionDAO for database save
//----------------------------------------------------------------------------
func UpdateJobRequisition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobRequisition model
	//----------------------------------------------------------------------------
	var data = model.JobRequisition{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobRequisition model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.UpdateJobRequisition(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to JobRequisitionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteJobRequisition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobRequisition data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := JobRequisitionDAO.DeleteJobRequisition(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Department on a JobRequisition
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDepartmentToJobRequisition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentId,_ := strconv.ParseUint( vars["departmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.AssignDepartmentToJobRequisition(jobRequisitionId, departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Department on a JobRequisition
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDepartmentFromJobRequisition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.UnassignDepartmentFromJobRequisition(jobRequisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a HiringManager on a JobRequisition
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignHiringManagerToJobRequisition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	hiringManagerId,_ := strconv.ParseUint( vars["hiringManagerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.AssignHiringManagerToJobRequisition(jobRequisitionId, hiringManagerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a HiringManager on a JobRequisition
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignHiringManagerFromJobRequisition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.UnassignHiringManagerFromJobRequisition(jobRequisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Recruiter on a JobRequisition
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRecruiterToJobRequisition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recruiterId,_ := strconv.ParseUint( vars["recruiterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.AssignRecruiterToJobRequisition(jobRequisitionId, recruiterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Recruiter on a JobRequisition
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRecruiterFromJobRequisition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.UnassignRecruiterFromJobRequisition(jobRequisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a JobProfile on a JobRequisition
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignJobProfileToJobRequisition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfileId,_ := strconv.ParseUint( vars["jobProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.AssignJobProfileToJobRequisition(jobRequisitionId, jobProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a JobProfile on a JobRequisition
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignJobProfileFromJobRequisition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.UnassignJobProfileFromJobRequisition(jobRequisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more candidatesIds as a Candidates to a JobRequisition
	//----------------------------------------------------------------------------
func AddCandidatesToJobRequisition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	candidatesIds,_ := vars["candidatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.AddCandidatesToJobRequisition(jobRequisitionId, candidatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more candidatesIds as a Candidates from a JobRequisition
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCandidatesFromJobRequisition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	candidatesIds,_ := vars["candidatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.RemoveCandidatesFromJobRequisition(jobRequisitionId, candidatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more interviewsIds as a Interviews to a JobRequisition
	//----------------------------------------------------------------------------
func AddInterviewsToJobRequisition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	interviewsIds,_ := vars["interviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.AddInterviewsToJobRequisition(jobRequisitionId, interviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more interviewsIds as a Interviews from a JobRequisition
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInterviewsFromJobRequisition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	interviewsIds,_ := vars["interviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.RemoveInterviewsFromJobRequisition(jobRequisitionId, interviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more offersIds as a Offers to a JobRequisition
	//----------------------------------------------------------------------------
func AddOffersToJobRequisition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	offersIds,_ := vars["offersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.AddOffersToJobRequisition(jobRequisitionId, offersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more offersIds as a Offers from a JobRequisition
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOffersFromJobRequisition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobRequisitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	offersIds,_ := vars["offersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobRequisition DAO
	//----------------------------------------------------------------------------
	requestResult := JobRequisitionDAO.RemoveOffersFromJobRequisition(jobRequisitionId, offersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
