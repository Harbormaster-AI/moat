package controller

import (
    LeavePolicyDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LeavePolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateLeavePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LeavePolicy model
	//----------------------------------------------------------------------------
	data := model.LeavePolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LeavePolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LeavePolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.CreateLeavePolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LeavePolicyDAO to find the relevant LeavePolicy
//----------------------------------------------------------------------------
func GetLeavePolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LeavePolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.GetLeavePolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LeavePolicyDAO for database read of all LeavePolicys
//----------------------------------------------------------------------------
func GetAllLeavePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LeavePolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.GetAllLeavePolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LeavePolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateLeavePolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LeavePolicy model
	//----------------------------------------------------------------------------
	var data = model.LeavePolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LeavePolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LeavePolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.UpdateLeavePolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LeavePolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLeavePolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LeavePolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LeavePolicyDAO.DeleteLeavePolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a LeavePolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToLeavePolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leavePolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LeavePolicy DAO
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.AssignOrganizationToLeavePolicy(leavePolicyId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a LeavePolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromLeavePolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	leavePolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LeavePolicy DAO
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.UnassignOrganizationFromLeavePolicy(leavePolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more leaveRequestsIds as a LeaveRequests to a LeavePolicy
	//----------------------------------------------------------------------------
func AddLeaveRequestsToLeavePolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leavePolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leaveRequestsIds,_ := vars["leaveRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LeavePolicy DAO
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.AddLeaveRequestsToLeavePolicy(leavePolicyId, leaveRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more leaveRequestsIds as a LeaveRequests from a LeavePolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLeaveRequestsFromLeavePolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	leavePolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leaveRequestsIds,_ := vars["leaveRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LeavePolicy DAO
	//----------------------------------------------------------------------------
	requestResult := LeavePolicyDAO.RemoveLeaveRequestsFromLeavePolicy(leavePolicyId, leaveRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
