package controller

import (
    ApprovalDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ApprovalDAO for database creation
//----------------------------------------------------------------------------
func CreateApproval(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Approval model
	//----------------------------------------------------------------------------
	data := model.Approval{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Approval model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Approval data access object to create
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.CreateApproval( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ApprovalDAO to find the relevant Approval
//----------------------------------------------------------------------------
func GetApproval(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Approval data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.GetApproval(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ApprovalDAO for database read of all Approvals
//----------------------------------------------------------------------------
func GetAllApproval(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Approval data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.GetAllApproval()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ApprovalDAO for database save
//----------------------------------------------------------------------------
func UpdateApproval(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Approval model
	//----------------------------------------------------------------------------
	var data = model.Approval{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Approval model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Approval data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.UpdateApproval(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ApprovalDAO for database deletion
//----------------------------------------------------------------------------
func DeleteApproval(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Approval data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ApprovalDAO.DeleteApproval(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Approver on a Approval
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignApproverToApproval(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	approvalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approverId,_ := strconv.ParseUint( vars["approverId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Approval DAO
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.AssignApproverToApproval(approvalId, approverId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Approver on a Approval
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignApproverFromApproval( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	approvalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Approval DAO
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.UnassignApproverFromApproval(approvalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Timesheet on a Approval
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTimesheetToApproval(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	approvalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timesheetId,_ := strconv.ParseUint( vars["timesheetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Approval DAO
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.AssignTimesheetToApproval(approvalId, timesheetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Timesheet on a Approval
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTimesheetFromApproval( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	approvalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Approval DAO
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.UnassignTimesheetFromApproval(approvalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LeaveRequest on a Approval
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLeaveRequestToApproval(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	approvalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leaveRequestId,_ := strconv.ParseUint( vars["leaveRequestId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Approval DAO
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.AssignLeaveRequestToApproval(approvalId, leaveRequestId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LeaveRequest on a Approval
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLeaveRequestFromApproval( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	approvalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Approval DAO
	//----------------------------------------------------------------------------
	requestResult := ApprovalDAO.UnassignLeaveRequestFromApproval(approvalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


