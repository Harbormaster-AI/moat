package controller

import (
    IoTDeviceDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to IoTDeviceDAO for database creation
//----------------------------------------------------------------------------
func CreateIoTDevice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty IoTDevice model
	//----------------------------------------------------------------------------
	data := model.IoTDevice{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a IoTDevice model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice data access object to create
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.CreateIoTDevice( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to IoTDeviceDAO to find the relevant IoTDevice
//----------------------------------------------------------------------------
func GetIoTDevice(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the IoTDevice data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.GetIoTDevice(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to IoTDeviceDAO for database read of all IoTDevices
//----------------------------------------------------------------------------
func GetAllIoTDevice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice data access object to get all
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.GetAllIoTDevice()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to IoTDeviceDAO for database save
//----------------------------------------------------------------------------
func UpdateIoTDevice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty IoTDevice model
	//----------------------------------------------------------------------------
	var data = model.IoTDevice{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a IoTDevice model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UpdateIoTDevice(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to IoTDeviceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteIoTDevice(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the IoTDevice data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := IoTDeviceDAO.DeleteIoTDevice(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a DeviceModel on a IoTDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceModelToIoTDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceModelId,_ := strconv.ParseUint( vars["deviceModelId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AssignDeviceModelToIoTDevice(ioTDeviceId, deviceModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DeviceModel on a IoTDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceModelFromIoTDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UnassignDeviceModelFromIoTDevice(ioTDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a IoTDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToIoTDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AssignTenantToIoTDevice(ioTDeviceId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a IoTDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromIoTDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UnassignTenantFromIoTDevice(ioTDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Site on a IoTDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSiteToIoTDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	siteId,_ := strconv.ParseUint( vars["siteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AssignSiteToIoTDevice(ioTDeviceId, siteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Site on a IoTDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSiteFromIoTDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UnassignSiteFromIoTDevice(ioTDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Room on a IoTDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRoomToIoTDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	roomId,_ := strconv.ParseUint( vars["roomId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AssignRoomToIoTDevice(ioTDeviceId, roomId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Room on a IoTDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRoomFromIoTDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UnassignRoomFromIoTDevice(ioTDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Gateway on a IoTDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignGatewayToIoTDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	gatewayId,_ := strconv.ParseUint( vars["gatewayId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AssignGatewayToIoTDevice(ioTDeviceId, gatewayId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Gateway on a IoTDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignGatewayFromIoTDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UnassignGatewayFromIoTDevice(ioTDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DigitalTwin on a IoTDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDigitalTwinToIoTDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	digitalTwinId,_ := strconv.ParseUint( vars["digitalTwinId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AssignDigitalTwinToIoTDevice(ioTDeviceId, digitalTwinId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DigitalTwin on a IoTDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDigitalTwinFromIoTDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UnassignDigitalTwinFromIoTDevice(ioTDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ProvisioningRecord on a IoTDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProvisioningRecordToIoTDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	provisioningRecordId,_ := strconv.ParseUint( vars["provisioningRecordId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AssignProvisioningRecordToIoTDevice(ioTDeviceId, provisioningRecordId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ProvisioningRecord on a IoTDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProvisioningRecordFromIoTDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.UnassignProvisioningRecordFromIoTDevice(ioTDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more sensorsIds as a Sensors to a IoTDevice
	//----------------------------------------------------------------------------
func AddSensorsToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sensorsIds,_ := vars["sensorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddSensorsToIoTDevice(ioTDeviceId, sensorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more sensorsIds as a Sensors from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSensorsFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sensorsIds,_ := vars["sensorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveSensorsFromIoTDevice(ioTDeviceId, sensorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more actuatorsIds as a Actuators to a IoTDevice
	//----------------------------------------------------------------------------
func AddActuatorsToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	actuatorsIds,_ := vars["actuatorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddActuatorsToIoTDevice(ioTDeviceId, actuatorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more actuatorsIds as a Actuators from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveActuatorsFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	actuatorsIds,_ := vars["actuatorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveActuatorsFromIoTDevice(ioTDeviceId, actuatorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more certificatesIds as a Certificates to a IoTDevice
	//----------------------------------------------------------------------------
func AddCertificatesToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	certificatesIds,_ := vars["certificatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddCertificatesToIoTDevice(ioTDeviceId, certificatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more certificatesIds as a Certificates from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCertificatesFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	certificatesIds,_ := vars["certificatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveCertificatesFromIoTDevice(ioTDeviceId, certificatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more telemetryStreamsIds as a TelemetryStreams to a IoTDevice
	//----------------------------------------------------------------------------
func AddTelemetryStreamsToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	telemetryStreamsIds,_ := vars["telemetryStreamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddTelemetryStreamsToIoTDevice(ioTDeviceId, telemetryStreamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more telemetryStreamsIds as a TelemetryStreams from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTelemetryStreamsFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	telemetryStreamsIds,_ := vars["telemetryStreamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveTelemetryStreamsFromIoTDevice(ioTDeviceId, telemetryStreamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more commandInvocationsIds as a CommandInvocations to a IoTDevice
	//----------------------------------------------------------------------------
func AddCommandInvocationsToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	commandInvocationsIds,_ := vars["commandInvocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddCommandInvocationsToIoTDevice(ioTDeviceId, commandInvocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more commandInvocationsIds as a CommandInvocations from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCommandInvocationsFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	commandInvocationsIds,_ := vars["commandInvocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveCommandInvocationsFromIoTDevice(ioTDeviceId, commandInvocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more alertsIds as a Alerts to a IoTDevice
	//----------------------------------------------------------------------------
func AddAlertsToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddAlertsToIoTDevice(ioTDeviceId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more alertsIds as a Alerts from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAlertsFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	alertsIds,_ := vars["alertsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveAlertsFromIoTDevice(ioTDeviceId, alertsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more deviceGroupsIds as a DeviceGroups to a IoTDevice
	//----------------------------------------------------------------------------
func AddDeviceGroupsToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceGroupsIds,_ := vars["deviceGroupsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddDeviceGroupsToIoTDevice(ioTDeviceId, deviceGroupsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more deviceGroupsIds as a DeviceGroups from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDeviceGroupsFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceGroupsIds,_ := vars["deviceGroupsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveDeviceGroupsFromIoTDevice(ioTDeviceId, deviceGroupsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more networkProfilesIds as a NetworkProfiles to a IoTDevice
	//----------------------------------------------------------------------------
func AddNetworkProfilesToIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	networkProfilesIds,_ := vars["networkProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.AddNetworkProfilesToIoTDevice(ioTDeviceId, networkProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more networkProfilesIds as a NetworkProfiles from a IoTDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNetworkProfilesFromIoTDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	ioTDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	networkProfilesIds,_ := vars["networkProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the IoTDevice DAO
	//----------------------------------------------------------------------------
	requestResult := IoTDeviceDAO.RemoveNetworkProfilesFromIoTDevice(ioTDeviceId, networkProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
