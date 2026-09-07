package controller

import (
    ComplianceAlertDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ComplianceAlertDAO for database creation
//----------------------------------------------------------------------------
func CreateComplianceAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ComplianceAlert model
	//----------------------------------------------------------------------------
	data := model.ComplianceAlert{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ComplianceAlert model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceAlert data access object to create
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.CreateComplianceAlert( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ComplianceAlertDAO to find the relevant ComplianceAlert
//----------------------------------------------------------------------------
func GetComplianceAlert(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ComplianceAlert data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.GetComplianceAlert(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ComplianceAlertDAO for database read of all ComplianceAlerts
//----------------------------------------------------------------------------
func GetAllComplianceAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ComplianceAlert data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.GetAllComplianceAlert()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ComplianceAlertDAO for database save
//----------------------------------------------------------------------------
func UpdateComplianceAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ComplianceAlert model
	//----------------------------------------------------------------------------
	var data = model.ComplianceAlert{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ComplianceAlert model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceAlert data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.UpdateComplianceAlert(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ComplianceAlertDAO for database deletion
//----------------------------------------------------------------------------
func DeleteComplianceAlert(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ComplianceAlert data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ComplianceAlertDAO.DeleteComplianceAlert(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Screening on a ComplianceAlert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignScreeningToComplianceAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	screeningId,_ := strconv.ParseUint( vars["screeningId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceAlert DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.AssignScreeningToComplianceAlert(complianceAlertId, screeningId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Screening on a ComplianceAlert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignScreeningFromComplianceAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceAlert DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.UnassignScreeningFromComplianceAlert(complianceAlertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Transaction on a ComplianceAlert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransactionToComplianceAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionId,_ := strconv.ParseUint( vars["transactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceAlert DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.AssignTransactionToComplianceAlert(complianceAlertId, transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Transaction on a ComplianceAlert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransactionFromComplianceAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	complianceAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ComplianceAlert DAO
	//----------------------------------------------------------------------------
	requestResult := ComplianceAlertDAO.UnassignTransactionFromComplianceAlert(complianceAlertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


