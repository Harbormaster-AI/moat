package controller

import (
    DeviceVendorDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DeviceVendorDAO for database creation
//----------------------------------------------------------------------------
func CreateDeviceVendor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeviceVendor model
	//----------------------------------------------------------------------------
	data := model.DeviceVendor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DeviceVendor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor data access object to create
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.CreateDeviceVendor( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DeviceVendorDAO to find the relevant DeviceVendor
//----------------------------------------------------------------------------
func GetDeviceVendor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DeviceVendor data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.GetDeviceVendor(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DeviceVendorDAO for database read of all DeviceVendors
//----------------------------------------------------------------------------
func GetAllDeviceVendor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.GetAllDeviceVendor()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DeviceVendorDAO for database save
//----------------------------------------------------------------------------
func UpdateDeviceVendor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DeviceVendor model
	//----------------------------------------------------------------------------
	var data = model.DeviceVendor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DeviceVendor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.UpdateDeviceVendor(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DeviceVendorDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDeviceVendor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DeviceVendor data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DeviceVendorDAO.DeleteDeviceVendor(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more deviceModelsIds as a DeviceModels to a DeviceVendor
	//----------------------------------------------------------------------------
func AddDeviceModelsToDeviceVendor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceVendorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceModelsIds,_ := vars["deviceModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.AddDeviceModelsToDeviceVendor(deviceVendorId, deviceModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more deviceModelsIds as a DeviceModels from a DeviceVendor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDeviceModelsFromDeviceVendor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceVendorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceModelsIds,_ := vars["deviceModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.RemoveDeviceModelsFromDeviceVendor(deviceVendorId, deviceModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more firmwareReleasesIds as a FirmwareReleases to a DeviceVendor
	//----------------------------------------------------------------------------
func AddFirmwareReleasesToDeviceVendor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceVendorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	firmwareReleasesIds,_ := vars["firmwareReleasesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.AddFirmwareReleasesToDeviceVendor(deviceVendorId, firmwareReleasesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more firmwareReleasesIds as a FirmwareReleases from a DeviceVendor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFirmwareReleasesFromDeviceVendor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceVendorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	firmwareReleasesIds,_ := vars["firmwareReleasesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.RemoveFirmwareReleasesFromDeviceVendor(deviceVendorId, firmwareReleasesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more hardwareModulesIds as a HardwareModules to a DeviceVendor
	//----------------------------------------------------------------------------
func AddHardwareModulesToDeviceVendor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceVendorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	hardwareModulesIds,_ := vars["hardwareModulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.AddHardwareModulesToDeviceVendor(deviceVendorId, hardwareModulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more hardwareModulesIds as a HardwareModules from a DeviceVendor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveHardwareModulesFromDeviceVendor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	deviceVendorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	hardwareModulesIds,_ := vars["hardwareModulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DeviceVendor DAO
	//----------------------------------------------------------------------------
	requestResult := DeviceVendorDAO.RemoveHardwareModulesFromDeviceVendor(deviceVendorId, hardwareModulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
