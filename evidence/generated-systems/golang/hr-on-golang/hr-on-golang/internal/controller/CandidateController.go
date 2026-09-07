package controller

import (
    CandidateDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CandidateDAO for database creation
//----------------------------------------------------------------------------
func CreateCandidate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Candidate model
	//----------------------------------------------------------------------------
	data := model.Candidate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Candidate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Candidate data access object to create
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.CreateCandidate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CandidateDAO to find the relevant Candidate
//----------------------------------------------------------------------------
func GetCandidate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Candidate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.GetCandidate(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CandidateDAO for database read of all Candidates
//----------------------------------------------------------------------------
func GetAllCandidate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Candidate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.GetAllCandidate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CandidateDAO for database save
//----------------------------------------------------------------------------
func UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Candidate model
	//----------------------------------------------------------------------------
	var data = model.Candidate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Candidate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Candidate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.UpdateCandidate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CandidateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCandidate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Candidate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CandidateDAO.DeleteCandidate(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more applicationsIds as a Applications to a Candidate
	//----------------------------------------------------------------------------
func AddApplicationsToCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	applicationsIds,_ := vars["applicationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.AddApplicationsToCandidate(candidateId, applicationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more applicationsIds as a Applications from a Candidate
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveApplicationsFromCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	applicationsIds,_ := vars["applicationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.RemoveApplicationsFromCandidate(candidateId, applicationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more interviewsIds as a Interviews to a Candidate
	//----------------------------------------------------------------------------
func AddInterviewsToCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	interviewsIds,_ := vars["interviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.AddInterviewsToCandidate(candidateId, interviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more interviewsIds as a Interviews from a Candidate
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInterviewsFromCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	interviewsIds,_ := vars["interviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.RemoveInterviewsFromCandidate(candidateId, interviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more offersIds as a Offers to a Candidate
	//----------------------------------------------------------------------------
func AddOffersToCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	offersIds,_ := vars["offersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.AddOffersToCandidate(candidateId, offersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more offersIds as a Offers from a Candidate
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOffersFromCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	offersIds,_ := vars["offersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.RemoveOffersFromCandidate(candidateId, offersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more documentsIds as a Documents to a Candidate
	//----------------------------------------------------------------------------
func AddDocumentsToCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	documentsIds,_ := vars["documentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.AddDocumentsToCandidate(candidateId, documentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more documentsIds as a Documents from a Candidate
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDocumentsFromCandidate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	candidateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	documentsIds,_ := vars["documentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Candidate DAO
	//----------------------------------------------------------------------------
	requestResult := CandidateDAO.RemoveDocumentsFromCandidate(candidateId, documentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
