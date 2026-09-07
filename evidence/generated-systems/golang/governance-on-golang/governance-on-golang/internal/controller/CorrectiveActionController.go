package controller

import (
    CorrectiveActionDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CorrectiveActionDAO for database creation
//----------------------------------------------------------------------------
func CreateCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CorrectiveAction model
	//----------------------------------------------------------------------------
	data := model.CorrectiveAction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CorrectiveAction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction data access object to create
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.CreateCorrectiveAction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CorrectiveActionDAO to find the relevant CorrectiveAction
//----------------------------------------------------------------------------
func GetCorrectiveAction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CorrectiveAction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.GetCorrectiveAction(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CorrectiveActionDAO for database read of all CorrectiveActions
//----------------------------------------------------------------------------
func GetAllCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.GetAllCorrectiveAction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CorrectiveActionDAO for database save
//----------------------------------------------------------------------------
func UpdateCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CorrectiveAction model
	//----------------------------------------------------------------------------
	var data = model.CorrectiveAction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CorrectiveAction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.UpdateCorrectiveAction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CorrectiveActionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCorrectiveAction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CorrectiveAction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CorrectiveActionDAO.DeleteCorrectiveAction(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Finding on a CorrectiveAction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFindingToCorrectiveAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	findingId,_ := strconv.ParseUint( vars["findingId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.AssignFindingToCorrectiveAction(correctiveActionId, findingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Finding on a CorrectiveAction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFindingFromCorrectiveAction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.UnassignFindingFromCorrectiveAction(correctiveActionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Issue on a CorrectiveAction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignIssueToCorrectiveAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	issueId,_ := strconv.ParseUint( vars["issueId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.AssignIssueToCorrectiveAction(correctiveActionId, issueId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Issue on a CorrectiveAction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignIssueFromCorrectiveAction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.UnassignIssueFromCorrectiveAction(correctiveActionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


