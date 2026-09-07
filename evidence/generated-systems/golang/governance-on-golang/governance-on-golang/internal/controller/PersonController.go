package controller

import (
    PersonDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PersonDAO for database creation
//----------------------------------------------------------------------------
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Person model
	//----------------------------------------------------------------------------
	data := model.Person{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Person model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Person data access object to create
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.CreatePerson( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PersonDAO to find the relevant Person
//----------------------------------------------------------------------------
func GetPerson(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Person data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.GetPerson(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PersonDAO for database read of all Persons
//----------------------------------------------------------------------------
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Person data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.GetAllPerson()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PersonDAO for database save
//----------------------------------------------------------------------------
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Person model
	//----------------------------------------------------------------------------
	var data = model.Person{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Person model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Person data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.UpdatePerson(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PersonDAO for database deletion
//----------------------------------------------------------------------------
func DeletePerson(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Person data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PersonDAO.DeletePerson(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more roleAssignmentsIds as a RoleAssignments to a Person
	//----------------------------------------------------------------------------
func AddRoleAssignmentsToPerson(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	personId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roleAssignmentsIds,_ := vars["roleAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Person DAO
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.AddRoleAssignmentsToPerson(personId, roleAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more roleAssignmentsIds as a RoleAssignments from a Person
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRoleAssignmentsFromPerson(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	personId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roleAssignmentsIds,_ := vars["roleAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Person DAO
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.RemoveRoleAssignmentsFromPerson(personId, roleAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ownedPoliciesIds as a OwnedPolicies to a Person
	//----------------------------------------------------------------------------
func AddOwnedPoliciesToPerson(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	personId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedPoliciesIds,_ := vars["ownedPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Person DAO
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.AddOwnedPoliciesToPerson(personId, ownedPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ownedPoliciesIds as a OwnedPolicies from a Person
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOwnedPoliciesFromPerson(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	personId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownedPoliciesIds,_ := vars["ownedPoliciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Person DAO
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.RemoveOwnedPoliciesFromPerson(personId, ownedPoliciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more correctiveActionsIds as a CorrectiveActions to a Person
	//----------------------------------------------------------------------------
func AddCorrectiveActionsToPerson(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	personId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Person DAO
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.AddCorrectiveActionsToPerson(personId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more correctiveActionsIds as a CorrectiveActions from a Person
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromPerson(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	personId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Person DAO
	//----------------------------------------------------------------------------
	requestResult := PersonDAO.RemoveCorrectiveActionsFromPerson(personId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
