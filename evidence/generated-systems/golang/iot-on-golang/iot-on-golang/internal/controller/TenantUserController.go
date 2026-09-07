package controller

import (
    TenantUserDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TenantUserDAO for database creation
//----------------------------------------------------------------------------
func CreateTenantUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TenantUser model
	//----------------------------------------------------------------------------
	data := model.TenantUser{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TenantUser model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TenantUser data access object to create
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.CreateTenantUser( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TenantUserDAO to find the relevant TenantUser
//----------------------------------------------------------------------------
func GetTenantUser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TenantUser data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.GetTenantUser(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TenantUserDAO for database read of all TenantUsers
//----------------------------------------------------------------------------
func GetAllTenantUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TenantUser data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.GetAllTenantUser()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TenantUserDAO for database save
//----------------------------------------------------------------------------
func UpdateTenantUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TenantUser model
	//----------------------------------------------------------------------------
	var data = model.TenantUser{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TenantUser model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TenantUser data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.UpdateTenantUser(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TenantUserDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTenantUser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TenantUser data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TenantUserDAO.DeleteTenantUser(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a TenantUser
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToTenantUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	tenantUserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TenantUser DAO
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.AssignTenantToTenantUser(tenantUserId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a TenantUser
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromTenantUser( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	tenantUserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TenantUser DAO
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.UnassignTenantFromTenantUser(tenantUserId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more commandInvocationsIds as a CommandInvocations to a TenantUser
	//----------------------------------------------------------------------------
func AddCommandInvocationsToTenantUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantUserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	commandInvocationsIds,_ := vars["commandInvocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TenantUser DAO
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.AddCommandInvocationsToTenantUser(tenantUserId, commandInvocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more commandInvocationsIds as a CommandInvocations from a TenantUser
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCommandInvocationsFromTenantUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tenantUserId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	commandInvocationsIds,_ := vars["commandInvocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TenantUser DAO
	//----------------------------------------------------------------------------
	requestResult := TenantUserDAO.RemoveCommandInvocationsFromTenantUser(tenantUserId, commandInvocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
