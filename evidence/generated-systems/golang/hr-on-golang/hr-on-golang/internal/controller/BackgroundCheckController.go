package controller

import (
    BackgroundCheckDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BackgroundCheckDAO for database creation
//----------------------------------------------------------------------------
func CreateBackgroundCheck(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BackgroundCheck model
	//----------------------------------------------------------------------------
	data := model.BackgroundCheck{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BackgroundCheck model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck data access object to create
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.CreateBackgroundCheck( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BackgroundCheckDAO to find the relevant BackgroundCheck
//----------------------------------------------------------------------------
func GetBackgroundCheck(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BackgroundCheck data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.GetBackgroundCheck(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BackgroundCheckDAO for database read of all BackgroundChecks
//----------------------------------------------------------------------------
func GetAllBackgroundCheck(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.GetAllBackgroundCheck()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BackgroundCheckDAO for database save
//----------------------------------------------------------------------------
func UpdateBackgroundCheck(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BackgroundCheck model
	//----------------------------------------------------------------------------
	var data = model.BackgroundCheck{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BackgroundCheck model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.UpdateBackgroundCheck(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BackgroundCheckDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBackgroundCheck(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BackgroundCheck data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BackgroundCheckDAO.DeleteBackgroundCheck(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Candidate on a BackgroundCheck
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCandidateToBackgroundCheck(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	backgroundCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	candidateId,_ := strconv.ParseUint( vars["candidateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck DAO
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.AssignCandidateToBackgroundCheck(backgroundCheckId, candidateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Candidate on a BackgroundCheck
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCandidateFromBackgroundCheck( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	backgroundCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck DAO
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.UnassignCandidateFromBackgroundCheck(backgroundCheckId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Requisition on a BackgroundCheck
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRequisitionToBackgroundCheck(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	backgroundCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	requisitionId,_ := strconv.ParseUint( vars["requisitionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck DAO
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.AssignRequisitionToBackgroundCheck(backgroundCheckId, requisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Requisition on a BackgroundCheck
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRequisitionFromBackgroundCheck( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	backgroundCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck DAO
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.UnassignRequisitionFromBackgroundCheck(backgroundCheckId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Report on a BackgroundCheck
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignReportToBackgroundCheck(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	backgroundCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportId,_ := strconv.ParseUint( vars["reportId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck DAO
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.AssignReportToBackgroundCheck(backgroundCheckId, reportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Report on a BackgroundCheck
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignReportFromBackgroundCheck( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	backgroundCheckId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BackgroundCheck DAO
	//----------------------------------------------------------------------------
	requestResult := BackgroundCheckDAO.UnassignReportFromBackgroundCheck(backgroundCheckId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


