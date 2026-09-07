package controller

import (
    SensorInstanceDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SensorInstanceDAO for database creation
//----------------------------------------------------------------------------
func CreateSensorInstance(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SensorInstance model
	//----------------------------------------------------------------------------
	data := model.SensorInstance{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SensorInstance model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SensorInstance data access object to create
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.CreateSensorInstance( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SensorInstanceDAO to find the relevant SensorInstance
//----------------------------------------------------------------------------
func GetSensorInstance(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SensorInstance data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.GetSensorInstance(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SensorInstanceDAO for database read of all SensorInstances
//----------------------------------------------------------------------------
func GetAllSensorInstance(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SensorInstance data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.GetAllSensorInstance()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SensorInstanceDAO for database save
//----------------------------------------------------------------------------
func UpdateSensorInstance(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SensorInstance model
	//----------------------------------------------------------------------------
	var data = model.SensorInstance{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SensorInstance model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SensorInstance data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.UpdateSensorInstance(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SensorInstanceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSensorInstance(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SensorInstance data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SensorInstanceDAO.DeleteSensorInstance(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Device on a SensorInstance
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToSensorInstance(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	sensorInstanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SensorInstance DAO
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.AssignDeviceToSensorInstance(sensorInstanceId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a SensorInstance
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromSensorInstance( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	sensorInstanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SensorInstance DAO
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.UnassignDeviceFromSensorInstance(sensorInstanceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more telemetryStreamsIds as a TelemetryStreams to a SensorInstance
	//----------------------------------------------------------------------------
func AddTelemetryStreamsToSensorInstance(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	sensorInstanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	telemetryStreamsIds,_ := vars["telemetryStreamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SensorInstance DAO
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.AddTelemetryStreamsToSensorInstance(sensorInstanceId, telemetryStreamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more telemetryStreamsIds as a TelemetryStreams from a SensorInstance
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTelemetryStreamsFromSensorInstance(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	sensorInstanceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	telemetryStreamsIds,_ := vars["telemetryStreamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SensorInstance DAO
	//----------------------------------------------------------------------------
	requestResult := SensorInstanceDAO.RemoveTelemetryStreamsFromSensorInstance(sensorInstanceId, telemetryStreamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
