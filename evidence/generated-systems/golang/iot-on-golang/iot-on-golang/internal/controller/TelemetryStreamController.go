package controller

import (
    TelemetryStreamDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TelemetryStreamDAO for database creation
//----------------------------------------------------------------------------
func CreateTelemetryStream(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TelemetryStream model
	//----------------------------------------------------------------------------
	data := model.TelemetryStream{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TelemetryStream model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream data access object to create
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.CreateTelemetryStream( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TelemetryStreamDAO to find the relevant TelemetryStream
//----------------------------------------------------------------------------
func GetTelemetryStream(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TelemetryStream data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.GetTelemetryStream(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TelemetryStreamDAO for database read of all TelemetryStreams
//----------------------------------------------------------------------------
func GetAllTelemetryStream(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.GetAllTelemetryStream()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TelemetryStreamDAO for database save
//----------------------------------------------------------------------------
func UpdateTelemetryStream(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TelemetryStream model
	//----------------------------------------------------------------------------
	var data = model.TelemetryStream{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TelemetryStream model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.UpdateTelemetryStream(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TelemetryStreamDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTelemetryStream(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TelemetryStream data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TelemetryStreamDAO.DeleteTelemetryStream(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Device on a TelemetryStream
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToTelemetryStream(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.AssignDeviceToTelemetryStream(telemetryStreamId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a TelemetryStream
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromTelemetryStream( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.UnassignDeviceFromTelemetryStream(telemetryStreamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Sensor on a TelemetryStream
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSensorToTelemetryStream(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sensorId,_ := strconv.ParseUint( vars["sensorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.AssignSensorToTelemetryStream(telemetryStreamId, sensorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sensor on a TelemetryStream
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSensorFromTelemetryStream( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.UnassignSensorFromTelemetryStream(telemetryStreamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Schema on a TelemetryStream
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSchemaToTelemetryStream(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	schemaId,_ := strconv.ParseUint( vars["schemaId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.AssignSchemaToTelemetryStream(telemetryStreamId, schemaId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Schema on a TelemetryStream
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSchemaFromTelemetryStream( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.UnassignSchemaFromTelemetryStream(telemetryStreamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a MessagingEndpoint on a TelemetryStream
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMessagingEndpointToTelemetryStream(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	messagingEndpointId,_ := strconv.ParseUint( vars["messagingEndpointId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.AssignMessagingEndpointToTelemetryStream(telemetryStreamId, messagingEndpointId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a MessagingEndpoint on a TelemetryStream
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMessagingEndpointFromTelemetryStream( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.UnassignMessagingEndpointFromTelemetryStream(telemetryStreamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a RetentionPolicy on a TelemetryStream
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRetentionPolicyToTelemetryStream(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	retentionPolicyId,_ := strconv.ParseUint( vars["retentionPolicyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.AssignRetentionPolicyToTelemetryStream(telemetryStreamId, retentionPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RetentionPolicy on a TelemetryStream
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRetentionPolicyFromTelemetryStream( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	telemetryStreamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetryStream DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetryStreamDAO.UnassignRetentionPolicyFromTelemetryStream(telemetryStreamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


