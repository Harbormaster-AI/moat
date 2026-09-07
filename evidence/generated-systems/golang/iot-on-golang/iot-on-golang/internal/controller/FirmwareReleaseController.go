package controller

import (
    FirmwareReleaseDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FirmwareReleaseDAO for database creation
//----------------------------------------------------------------------------
func CreateFirmwareRelease(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FirmwareRelease model
	//----------------------------------------------------------------------------
	data := model.FirmwareRelease{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FirmwareRelease model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FirmwareRelease data access object to create
	//----------------------------------------------------------------------------
	requestResult := FirmwareReleaseDAO.CreateFirmwareRelease( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FirmwareReleaseDAO to find the relevant FirmwareRelease
//----------------------------------------------------------------------------
func GetFirmwareRelease(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FirmwareRelease data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FirmwareReleaseDAO.GetFirmwareRelease(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FirmwareReleaseDAO for database read of all FirmwareReleases
//----------------------------------------------------------------------------
func GetAllFirmwareRelease(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FirmwareRelease data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FirmwareReleaseDAO.GetAllFirmwareRelease()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FirmwareReleaseDAO for database save
//----------------------------------------------------------------------------
func UpdateFirmwareRelease(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FirmwareRelease model
	//----------------------------------------------------------------------------
	var data = model.FirmwareRelease{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FirmwareRelease model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FirmwareRelease data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FirmwareReleaseDAO.UpdateFirmwareRelease(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FirmwareReleaseDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFirmwareRelease(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FirmwareRelease data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FirmwareReleaseDAO.DeleteFirmwareRelease(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a DeviceModel on a FirmwareRelease
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceModelToFirmwareRelease(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	firmwareReleaseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceModelId,_ := strconv.ParseUint( vars["deviceModelId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FirmwareRelease DAO
	//----------------------------------------------------------------------------
	requestResult := FirmwareReleaseDAO.AssignDeviceModelToFirmwareRelease(firmwareReleaseId, deviceModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DeviceModel on a FirmwareRelease
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceModelFromFirmwareRelease( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	firmwareReleaseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FirmwareRelease DAO
	//----------------------------------------------------------------------------
	requestResult := FirmwareReleaseDAO.UnassignDeviceModelFromFirmwareRelease(firmwareReleaseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


