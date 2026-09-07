package controller

import (
    Exception_DAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to Exception_DAO for database creation
//----------------------------------------------------------------------------
func CreateException_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Exception_ model
	//----------------------------------------------------------------------------
	data := model.Exception_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Exception_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ data access object to create
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.CreateException_( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to Exception_DAO to find the relevant Exception_
//----------------------------------------------------------------------------
func GetException_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Exception_ data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.GetException_(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to Exception_DAO for database read of all Exception_s
//----------------------------------------------------------------------------
func GetAllException_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Exception_ data access object to get all
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.GetAllException_()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to Exception_DAO for database save
//----------------------------------------------------------------------------
func UpdateException_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Exception_ model
	//----------------------------------------------------------------------------
	var data = model.Exception_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Exception_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.UpdateException_(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to Exception_DAO for database deletion
//----------------------------------------------------------------------------
func DeleteException_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Exception_ data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := Exception_DAO.DeleteException_(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a RetentionSchedule on a Exception_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRetentionScheduleToException_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	retentionScheduleId,_ := strconv.ParseUint( vars["retentionScheduleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.AssignRetentionScheduleToException_(exception_Id, retentionScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RetentionSchedule on a Exception_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRetentionScheduleFromException_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.UnassignRetentionScheduleFromException_(exception_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Policy on a Exception_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToException_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.AssignPolicyToException_(exception_Id, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a Exception_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromException_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.UnassignPolicyFromException_(exception_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Control on a Exception_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignControlToException_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlId,_ := strconv.ParseUint( vars["controlId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.AssignControlToException_(exception_Id, controlId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Control on a Exception_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignControlFromException_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.UnassignControlFromException_(exception_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Risk on a Exception_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRiskToException_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	riskId,_ := strconv.ParseUint( vars["riskId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.AssignRiskToException_(exception_Id, riskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Risk on a Exception_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRiskFromException_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	exception_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Exception_ DAO
	//----------------------------------------------------------------------------
	requestResult := Exception_DAO.UnassignRiskFromException_(exception_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


