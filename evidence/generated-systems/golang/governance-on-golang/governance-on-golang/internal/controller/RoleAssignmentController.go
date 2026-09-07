package controller

import (
    RoleAssignmentDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RoleAssignmentDAO for database creation
//----------------------------------------------------------------------------
func CreateRoleAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RoleAssignment model
	//----------------------------------------------------------------------------
	data := model.RoleAssignment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RoleAssignment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment data access object to create
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.CreateRoleAssignment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RoleAssignmentDAO to find the relevant RoleAssignment
//----------------------------------------------------------------------------
func GetRoleAssignment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RoleAssignment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.GetRoleAssignment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RoleAssignmentDAO for database read of all RoleAssignments
//----------------------------------------------------------------------------
func GetAllRoleAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.GetAllRoleAssignment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RoleAssignmentDAO for database save
//----------------------------------------------------------------------------
func UpdateRoleAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RoleAssignment model
	//----------------------------------------------------------------------------
	var data = model.RoleAssignment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RoleAssignment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.UpdateRoleAssignment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RoleAssignmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRoleAssignment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RoleAssignment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RoleAssignmentDAO.DeleteRoleAssignment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Person on a RoleAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPersonToRoleAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	personId,_ := strconv.ParseUint( vars["personId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.AssignPersonToRoleAssignment(roleAssignmentId, personId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Person on a RoleAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPersonFromRoleAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.UnassignPersonFromRoleAssignment(roleAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Role on a RoleAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRoleToRoleAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roleId,_ := strconv.ParseUint( vars["roleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.AssignRoleToRoleAssignment(roleAssignmentId, roleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Role on a RoleAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRoleFromRoleAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.UnassignRoleFromRoleAssignment(roleAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a GovernanceBody on a RoleAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignGovernanceBodyToRoleAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	governanceBodyId,_ := strconv.ParseUint( vars["governanceBodyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.AssignGovernanceBodyToRoleAssignment(roleAssignmentId, governanceBodyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a GovernanceBody on a RoleAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignGovernanceBodyFromRoleAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.UnassignGovernanceBodyFromRoleAssignment(roleAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Organization on a RoleAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToRoleAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.AssignOrganizationToRoleAssignment(roleAssignmentId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a RoleAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromRoleAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	roleAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RoleAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := RoleAssignmentDAO.UnassignOrganizationFromRoleAssignment(roleAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


