package controller

import (
    RoleDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RoleDAO for database creation
//----------------------------------------------------------------------------
func CreateRole(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Role model
	//----------------------------------------------------------------------------
	data := model.Role{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Role model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Role data access object to create
	//----------------------------------------------------------------------------
	requestResult := RoleDAO.CreateRole( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RoleDAO to find the relevant Role
//----------------------------------------------------------------------------
func GetRole(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Role data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RoleDAO.GetRole(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RoleDAO for database read of all Roles
//----------------------------------------------------------------------------
func GetAllRole(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Role data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RoleDAO.GetAllRole()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RoleDAO for database save
//----------------------------------------------------------------------------
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Role model
	//----------------------------------------------------------------------------
	var data = model.Role{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Role model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Role data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RoleDAO.UpdateRole(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RoleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRole(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Role data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RoleDAO.DeleteRole(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more assignmentsIds as a Assignments to a Role
	//----------------------------------------------------------------------------
func AddAssignmentsToRole(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	roleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignmentsIds,_ := vars["assignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Role DAO
	//----------------------------------------------------------------------------
	requestResult := RoleDAO.AddAssignmentsToRole(roleId, assignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more assignmentsIds as a Assignments from a Role
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAssignmentsFromRole(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	roleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignmentsIds,_ := vars["assignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Role DAO
	//----------------------------------------------------------------------------
	requestResult := RoleDAO.RemoveAssignmentsFromRole(roleId, assignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
