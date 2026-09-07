package controller

import (
    AccessPolicyDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AccessPolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateAccessPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AccessPolicy model
	//----------------------------------------------------------------------------
	data := model.AccessPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AccessPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.CreateAccessPolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AccessPolicyDAO to find the relevant AccessPolicy
//----------------------------------------------------------------------------
func GetAccessPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AccessPolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.GetAccessPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AccessPolicyDAO for database read of all AccessPolicys
//----------------------------------------------------------------------------
func GetAllAccessPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.GetAllAccessPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AccessPolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateAccessPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AccessPolicy model
	//----------------------------------------------------------------------------
	var data = model.AccessPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AccessPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.UpdateAccessPolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AccessPolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAccessPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AccessPolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AccessPolicyDAO.DeleteAccessPolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a AccessPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToAccessPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AssignTenantToAccessPolicy(accessPolicyId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a AccessPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromAccessPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.UnassignTenantFromAccessPolicy(accessPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more apiKeysIds as a ApiKeys to a AccessPolicy
	//----------------------------------------------------------------------------
func AddApiKeysToAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	apiKeysIds,_ := vars["apiKeysIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AddApiKeysToAccessPolicy(accessPolicyId, apiKeysIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more apiKeysIds as a ApiKeys from a AccessPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveApiKeysFromAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	apiKeysIds,_ := vars["apiKeysIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.RemoveApiKeysFromAccessPolicy(accessPolicyId, apiKeysIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more usersIds as a Users to a AccessPolicy
	//----------------------------------------------------------------------------
func AddUsersToAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AddUsersToAccessPolicy(accessPolicyId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usersIds as a Users from a AccessPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsersFromAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.RemoveUsersFromAccessPolicy(accessPolicyId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
