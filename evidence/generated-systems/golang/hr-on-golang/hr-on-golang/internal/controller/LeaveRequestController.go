package controller

import (
    LeaveRequestDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LeaveRequestDAO for database creation
//----------------------------------------------------------------------------
func CreateLeaveRequest(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LeaveRequest model
	//----------------------------------------------------------------------------
	data := model.LeaveRequest{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LeaveRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest data access object to create
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.CreateLeaveRequest( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LeaveRequestDAO to find the relevant LeaveRequest
//----------------------------------------------------------------------------
func GetLeaveRequest(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LeaveRequest data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.GetLeaveRequest(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LeaveRequestDAO for database read of all LeaveRequests
//----------------------------------------------------------------------------
func GetAllLeaveRequest(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.GetAllLeaveRequest()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LeaveRequestDAO for database save
//----------------------------------------------------------------------------
func UpdateLeaveRequest(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LeaveRequest model
	//----------------------------------------------------------------------------
	var data = model.LeaveRequest{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LeaveRequest model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.UpdateLeaveRequest(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LeaveRequestDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLeaveRequest(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LeaveRequest data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LeaveRequestDAO.DeleteLeaveRequest(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a LeaveRequest
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToLeaveRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leaveRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest DAO
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.AssignEmployeeToLeaveRequest(leaveRequestId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a LeaveRequest
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromLeaveRequest( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leaveRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest DAO
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.UnassignEmployeeFromLeaveRequest(leaveRequestId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LeavePolicy on a LeaveRequest
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLeavePolicyToLeaveRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leaveRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leavePolicyId,_ := strconv.ParseUint( vars["leavePolicyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest DAO
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.AssignLeavePolicyToLeaveRequest(leaveRequestId, leavePolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LeavePolicy on a LeaveRequest
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLeavePolicyFromLeaveRequest( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leaveRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest DAO
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.UnassignLeavePolicyFromLeaveRequest(leaveRequestId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more approvalsIds as a Approvals to a LeaveRequest
	//----------------------------------------------------------------------------
func AddApprovalsToLeaveRequest(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leaveRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approvalsIds,_ := vars["approvalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest DAO
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.AddApprovalsToLeaveRequest(leaveRequestId, approvalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more approvalsIds as a Approvals from a LeaveRequest
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveApprovalsFromLeaveRequest(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leaveRequestId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approvalsIds,_ := vars["approvalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LeaveRequest DAO
	//----------------------------------------------------------------------------
	requestResult := LeaveRequestDAO.RemoveApprovalsFromLeaveRequest(leaveRequestId, approvalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
