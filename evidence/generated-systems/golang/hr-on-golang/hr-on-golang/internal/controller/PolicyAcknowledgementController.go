package controller

import (
    PolicyAcknowledgementDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PolicyAcknowledgementDAO for database creation
//----------------------------------------------------------------------------
func CreatePolicyAcknowledgement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PolicyAcknowledgement model
	//----------------------------------------------------------------------------
	data := model.PolicyAcknowledgement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PolicyAcknowledgement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyAcknowledgement data access object to create
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.CreatePolicyAcknowledgement( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PolicyAcknowledgementDAO to find the relevant PolicyAcknowledgement
//----------------------------------------------------------------------------
func GetPolicyAcknowledgement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PolicyAcknowledgement data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.GetPolicyAcknowledgement(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PolicyAcknowledgementDAO for database read of all PolicyAcknowledgements
//----------------------------------------------------------------------------
func GetAllPolicyAcknowledgement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PolicyAcknowledgement data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.GetAllPolicyAcknowledgement()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PolicyAcknowledgementDAO for database save
//----------------------------------------------------------------------------
func UpdatePolicyAcknowledgement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PolicyAcknowledgement model
	//----------------------------------------------------------------------------
	var data = model.PolicyAcknowledgement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PolicyAcknowledgement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyAcknowledgement data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.UpdatePolicyAcknowledgement(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PolicyAcknowledgementDAO for database deletion
//----------------------------------------------------------------------------
func DeletePolicyAcknowledgement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PolicyAcknowledgement data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PolicyAcknowledgementDAO.DeletePolicyAcknowledgement(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Policy on a PolicyAcknowledgement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToPolicyAcknowledgement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyAcknowledgementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyAcknowledgement DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.AssignPolicyToPolicyAcknowledgement(policyAcknowledgementId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a PolicyAcknowledgement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromPolicyAcknowledgement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyAcknowledgementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyAcknowledgement DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.UnassignPolicyFromPolicyAcknowledgement(policyAcknowledgementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a PolicyAcknowledgement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToPolicyAcknowledgement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyAcknowledgementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyAcknowledgement DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.AssignEmployeeToPolicyAcknowledgement(policyAcknowledgementId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a PolicyAcknowledgement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromPolicyAcknowledgement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	policyAcknowledgementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PolicyAcknowledgement DAO
	//----------------------------------------------------------------------------
	requestResult := PolicyAcknowledgementDAO.UnassignEmployeeFromPolicyAcknowledgement(policyAcknowledgementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


