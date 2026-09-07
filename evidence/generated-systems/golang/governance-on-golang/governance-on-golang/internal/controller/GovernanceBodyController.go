package controller

import (
    GovernanceBodyDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to GovernanceBodyDAO for database creation
//----------------------------------------------------------------------------
func CreateGovernanceBody(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GovernanceBody model
	//----------------------------------------------------------------------------
	data := model.GovernanceBody{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GovernanceBody model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody data access object to create
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.CreateGovernanceBody( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to GovernanceBodyDAO to find the relevant GovernanceBody
//----------------------------------------------------------------------------
func GetGovernanceBody(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GovernanceBody data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.GetGovernanceBody(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to GovernanceBodyDAO for database read of all GovernanceBodys
//----------------------------------------------------------------------------
func GetAllGovernanceBody(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody data access object to get all
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.GetAllGovernanceBody()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to GovernanceBodyDAO for database save
//----------------------------------------------------------------------------
func UpdateGovernanceBody(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GovernanceBody model
	//----------------------------------------------------------------------------
	var data = model.GovernanceBody{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GovernanceBody model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.UpdateGovernanceBody(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to GovernanceBodyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteGovernanceBody(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GovernanceBody data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := GovernanceBodyDAO.DeleteGovernanceBody(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a GovernanceBody
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToGovernanceBody(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	governanceBodyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody DAO
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.AssignOrganizationToGovernanceBody(governanceBodyId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a GovernanceBody
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromGovernanceBody( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	governanceBodyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody DAO
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.UnassignOrganizationFromGovernanceBody(governanceBodyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more roleAssignmentsIds as a RoleAssignments to a GovernanceBody
	//----------------------------------------------------------------------------
func AddRoleAssignmentsToGovernanceBody(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	governanceBodyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roleAssignmentsIds,_ := vars["roleAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody DAO
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.AddRoleAssignmentsToGovernanceBody(governanceBodyId, roleAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more roleAssignmentsIds as a RoleAssignments from a GovernanceBody
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRoleAssignmentsFromGovernanceBody(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	governanceBodyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roleAssignmentsIds,_ := vars["roleAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody DAO
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.RemoveRoleAssignmentsFromGovernanceBody(governanceBodyId, roleAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a GovernanceBody
	//----------------------------------------------------------------------------
func AddPoliciesToGovernanceBody(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	governanceBodyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody DAO
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.AddPoliciesToGovernanceBody(governanceBodyId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a GovernanceBody
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromGovernanceBody(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	governanceBodyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GovernanceBody DAO
	//----------------------------------------------------------------------------
	requestResult := GovernanceBodyDAO.RemovePoliciesFromGovernanceBody(governanceBodyId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
