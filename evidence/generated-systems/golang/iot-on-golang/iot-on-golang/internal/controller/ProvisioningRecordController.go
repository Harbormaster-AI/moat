package controller

import (
    ProvisioningRecordDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProvisioningRecordDAO for database creation
//----------------------------------------------------------------------------
func CreateProvisioningRecord(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProvisioningRecord model
	//----------------------------------------------------------------------------
	data := model.ProvisioningRecord{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProvisioningRecord model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.CreateProvisioningRecord( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProvisioningRecordDAO to find the relevant ProvisioningRecord
//----------------------------------------------------------------------------
func GetProvisioningRecord(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProvisioningRecord data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.GetProvisioningRecord(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProvisioningRecordDAO for database read of all ProvisioningRecords
//----------------------------------------------------------------------------
func GetAllProvisioningRecord(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.GetAllProvisioningRecord()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProvisioningRecordDAO for database save
//----------------------------------------------------------------------------
func UpdateProvisioningRecord(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProvisioningRecord model
	//----------------------------------------------------------------------------
	var data = model.ProvisioningRecord{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProvisioningRecord model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.UpdateProvisioningRecord(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProvisioningRecordDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProvisioningRecord(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProvisioningRecord data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProvisioningRecordDAO.DeleteProvisioningRecord(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Device on a ProvisioningRecord
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToProvisioningRecord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	provisioningRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord DAO
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.AssignDeviceToProvisioningRecord(provisioningRecordId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a ProvisioningRecord
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromProvisioningRecord( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	provisioningRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord DAO
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.UnassignDeviceFromProvisioningRecord(provisioningRecordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Certificate on a ProvisioningRecord
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCertificateToProvisioningRecord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	provisioningRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	certificateId,_ := strconv.ParseUint( vars["certificateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord DAO
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.AssignCertificateToProvisioningRecord(provisioningRecordId, certificateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Certificate on a ProvisioningRecord
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCertificateFromProvisioningRecord( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	provisioningRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord DAO
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.UnassignCertificateFromProvisioningRecord(provisioningRecordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a ProvisioningRecord
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToProvisioningRecord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	provisioningRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord DAO
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.AssignTenantToProvisioningRecord(provisioningRecordId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a ProvisioningRecord
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromProvisioningRecord( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	provisioningRecordId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProvisioningRecord DAO
	//----------------------------------------------------------------------------
	requestResult := ProvisioningRecordDAO.UnassignTenantFromProvisioningRecord(provisioningRecordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


