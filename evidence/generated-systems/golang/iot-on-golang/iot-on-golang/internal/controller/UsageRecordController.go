package controller

import (
    UsageRecordDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to UsageRecordDAO for database creation
//----------------------------------------------------------------------------
func CreateUsageRecord(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UsageRecord model
	//----------------------------------------------------------------------------
	data := model.UsageRecord{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UsageRecord model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord data access object to create
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.CreateUsageRecord( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to UsageRecordDAO to find the relevant UsageRecord
//----------------------------------------------------------------------------
func GetUsageRecord(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UsageRecord data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.GetUsageRecord(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to UsageRecordDAO for database read of all UsageRecords
//----------------------------------------------------------------------------
func GetAllUsageRecord(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord data access object to get all
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.GetAllUsageRecord()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to UsageRecordDAO for database save
//----------------------------------------------------------------------------
func UpdateUsageRecord(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UsageRecord model
	//----------------------------------------------------------------------------
	var data = model.UsageRecord{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UsageRecord model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.UpdateUsageRecord(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to UsageRecordDAO for database deletion
//----------------------------------------------------------------------------
func DeleteUsageRecord(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UsageRecord data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := UsageRecordDAO.DeleteUsageRecord(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a UsageRecord
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToUsageRecord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord DAO
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.AssignTenantToUsageRecord(usageRecordId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a UsageRecord
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromUsageRecord( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord DAO
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.UnassignTenantFromUsageRecord(usageRecordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Device on a UsageRecord
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToUsageRecord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord DAO
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.AssignDeviceToUsageRecord(usageRecordId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a UsageRecord
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromUsageRecord( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord DAO
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.UnassignDeviceFromUsageRecord(usageRecordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ConnectivityPlan on a UsageRecord
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConnectivityPlanToUsageRecord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	connectivityPlanId,_ := strconv.ParseUint( vars["connectivityPlanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord DAO
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.AssignConnectivityPlanToUsageRecord(usageRecordId, connectivityPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConnectivityPlan on a UsageRecord
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConnectivityPlanFromUsageRecord( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	usageRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UsageRecord DAO
	//----------------------------------------------------------------------------
	requestResult := UsageRecordDAO.UnassignConnectivityPlanFromUsageRecord(usageRecordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


