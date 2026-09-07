package controller

import (
    DeviceModelDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DeviceModelDAO for database creation
//----------------------------------------------------------------------------
func CreateDeviceModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeviceModel model
	//----------------------------------------------------------------------------
	data := model.DeviceModel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DeviceModel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel data access object to create
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.CreateDeviceModel( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DeviceModelDAO to find the relevant DeviceModel
//----------------------------------------------------------------------------
func GetDeviceModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DeviceModel data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.GetDeviceModel(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DeviceModelDAO for database read of all DeviceModels
//----------------------------------------------------------------------------
func GetAllDeviceModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.GetAllDeviceModel()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DeviceModelDAO for database save
//----------------------------------------------------------------------------
func UpdateDeviceModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeviceModel model
	//----------------------------------------------------------------------------
	var data = model.DeviceModel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DeviceModel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.UpdateDeviceModel(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DeviceModelDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDeviceModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DeviceModel data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DeviceModelDAO.DeleteDeviceModel(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Vendor on a DeviceModel
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignVendorToDeviceModel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	vendorId,_ := strconv.ParseUint( vars["vendorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.AssignVendorToDeviceModel(deviceModelId, vendorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Vendor on a DeviceModel
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignVendorFromDeviceModel( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.UnassignVendorFromDeviceModel(deviceModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a TwinTemplate on a DeviceModel
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTwinTemplateToDeviceModel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	twinTemplateId,_ := strconv.ParseUint( vars["twinTemplateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.AssignTwinTemplateToDeviceModel(deviceModelId, twinTemplateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TwinTemplate on a DeviceModel
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTwinTemplateFromDeviceModel( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.UnassignTwinTemplateFromDeviceModel(deviceModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more hardwareModulesIds as a HardwareModules to a DeviceModel
	//----------------------------------------------------------------------------
func AddHardwareModulesToDeviceModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	hardwareModulesIds,_ := vars["hardwareModulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.AddHardwareModulesToDeviceModel(deviceModelId, hardwareModulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more hardwareModulesIds as a HardwareModules from a DeviceModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveHardwareModulesFromDeviceModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	hardwareModulesIds,_ := vars["hardwareModulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.RemoveHardwareModulesFromDeviceModel(deviceModelId, hardwareModulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more firmwareReleasesIds as a FirmwareReleases to a DeviceModel
	//----------------------------------------------------------------------------
func AddFirmwareReleasesToDeviceModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	firmwareReleasesIds,_ := vars["firmwareReleasesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.AddFirmwareReleasesToDeviceModel(deviceModelId, firmwareReleasesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more firmwareReleasesIds as a FirmwareReleases from a DeviceModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFirmwareReleasesFromDeviceModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	firmwareReleasesIds,_ := vars["firmwareReleasesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.RemoveFirmwareReleasesFromDeviceModel(deviceModelId, firmwareReleasesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more commandDefinitionsIds as a CommandDefinitions to a DeviceModel
	//----------------------------------------------------------------------------
func AddCommandDefinitionsToDeviceModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	commandDefinitionsIds,_ := vars["commandDefinitionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.AddCommandDefinitionsToDeviceModel(deviceModelId, commandDefinitionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more commandDefinitionsIds as a CommandDefinitions from a DeviceModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCommandDefinitionsFromDeviceModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	commandDefinitionsIds,_ := vars["commandDefinitionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceModel DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceModelDAO.RemoveCommandDefinitionsFromDeviceModel(deviceModelId, commandDefinitionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
