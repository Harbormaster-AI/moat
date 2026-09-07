package controller

import (
    DataRetentionPolicyDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataRetentionPolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateDataRetentionPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataRetentionPolicy model
	//----------------------------------------------------------------------------
	data := model.DataRetentionPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataRetentionPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataRetentionPolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.CreateDataRetentionPolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataRetentionPolicyDAO to find the relevant DataRetentionPolicy
//----------------------------------------------------------------------------
func GetDataRetentionPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataRetentionPolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.GetDataRetentionPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataRetentionPolicyDAO for database read of all DataRetentionPolicys
//----------------------------------------------------------------------------
func GetAllDataRetentionPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataRetentionPolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.GetAllDataRetentionPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataRetentionPolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateDataRetentionPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataRetentionPolicy model
	//----------------------------------------------------------------------------
	var data = model.DataRetentionPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataRetentionPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataRetentionPolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.UpdateDataRetentionPolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataRetentionPolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataRetentionPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataRetentionPolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataRetentionPolicyDAO.DeleteDataRetentionPolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a DataRetentionPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToDataRetentionPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataRetentionPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataRetentionPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.AssignTenantToDataRetentionPolicy(dataRetentionPolicyId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a DataRetentionPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromDataRetentionPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataRetentionPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataRetentionPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.UnassignTenantFromDataRetentionPolicy(dataRetentionPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more streamsIds as a Streams to a DataRetentionPolicy
	//----------------------------------------------------------------------------
func AddStreamsToDataRetentionPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataRetentionPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	streamsIds,_ := vars["streamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataRetentionPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.AddStreamsToDataRetentionPolicy(dataRetentionPolicyId, streamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more streamsIds as a Streams from a DataRetentionPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStreamsFromDataRetentionPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataRetentionPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	streamsIds,_ := vars["streamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataRetentionPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := DataRetentionPolicyDAO.RemoveStreamsFromDataRetentionPolicy(dataRetentionPolicyId, streamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
