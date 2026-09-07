package controller

import (
    IssueDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to IssueDAO for database creation
//----------------------------------------------------------------------------
func CreateIssue(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Issue model
	//----------------------------------------------------------------------------
	data := model.Issue{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Issue model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Issue data access object to create
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.CreateIssue( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to IssueDAO to find the relevant Issue
//----------------------------------------------------------------------------
func GetIssue(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Issue data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.GetIssue(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to IssueDAO for database read of all Issues
//----------------------------------------------------------------------------
func GetAllIssue(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Issue data access object to get all
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.GetAllIssue()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to IssueDAO for database save
//----------------------------------------------------------------------------
func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Issue model
	//----------------------------------------------------------------------------
	var data = model.Issue{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Issue model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Issue data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.UpdateIssue(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to IssueDAO for database deletion
//----------------------------------------------------------------------------
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Issue data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := IssueDAO.DeleteIssue(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Risk on a Issue
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRiskToIssue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	riskId,_ := strconv.ParseUint( vars["riskId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.AssignRiskToIssue(issueId, riskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Risk on a Issue
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRiskFromIssue( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.UnassignRiskFromIssue(issueId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Finding on a Issue
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFindingToIssue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingId,_ := strconv.ParseUint( vars["findingId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.AssignFindingToIssue(issueId, findingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Finding on a Issue
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFindingFromIssue( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.UnassignFindingFromIssue(issueId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Control on a Issue
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignControlToIssue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlId,_ := strconv.ParseUint( vars["controlId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.AssignControlToIssue(issueId, controlId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Control on a Issue
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignControlFromIssue( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.UnassignControlFromIssue(issueId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more correctiveActionsIds as a CorrectiveActions to a Issue
	//----------------------------------------------------------------------------
func AddCorrectiveActionsToIssue(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.AddCorrectiveActionsToIssue(issueId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more correctiveActionsIds as a CorrectiveActions from a Issue
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromIssue(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	issueId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Issue DAO
	//----------------------------------------------------------------------------
	requestResult := IssueDAO.RemoveCorrectiveActionsFromIssue(issueId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
