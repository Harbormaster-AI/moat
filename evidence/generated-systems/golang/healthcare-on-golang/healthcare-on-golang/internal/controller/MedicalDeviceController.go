package controller

import (
    MedicalDeviceDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MedicalDeviceDAO for database creation
//----------------------------------------------------------------------------
func CreateMedicalDevice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicalDevice model
	//----------------------------------------------------------------------------
	data := model.MedicalDevice{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicalDevice model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice data access object to create
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.CreateMedicalDevice( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MedicalDeviceDAO to find the relevant MedicalDevice
//----------------------------------------------------------------------------
func GetMedicalDevice(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicalDevice data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.GetMedicalDevice(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MedicalDeviceDAO for database read of all MedicalDevices
//----------------------------------------------------------------------------
func GetAllMedicalDevice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.GetAllMedicalDevice()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MedicalDeviceDAO for database save
//----------------------------------------------------------------------------
func UpdateMedicalDevice(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicalDevice model
	//----------------------------------------------------------------------------
	var data = model.MedicalDevice{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicalDevice model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.UpdateMedicalDevice(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MedicalDeviceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMedicalDevice(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicalDevice data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MedicalDeviceDAO.DeleteMedicalDevice(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a MedicalDevice
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToMedicalDevice(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicalDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.AssignPatientToMedicalDevice(medicalDeviceId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a MedicalDevice
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromMedicalDevice( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicalDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.UnassignPatientFromMedicalDevice(medicalDeviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more observationsIds as a Observations to a MedicalDevice
	//----------------------------------------------------------------------------
func AddObservationsToMedicalDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.AddObservationsToMedicalDevice(medicalDeviceId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more observationsIds as a Observations from a MedicalDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObservationsFromMedicalDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.RemoveObservationsFromMedicalDevice(medicalDeviceId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more softwareUpdatesIds as a SoftwareUpdates to a MedicalDevice
	//----------------------------------------------------------------------------
func AddSoftwareUpdatesToMedicalDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	softwareUpdatesIds,_ := vars["softwareUpdatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.AddSoftwareUpdatesToMedicalDevice(medicalDeviceId, softwareUpdatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more softwareUpdatesIds as a SoftwareUpdates from a MedicalDevice
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSoftwareUpdatesFromMedicalDevice(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalDeviceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	softwareUpdatesIds,_ := vars["softwareUpdatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalDevice DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalDeviceDAO.RemoveSoftwareUpdatesFromMedicalDevice(medicalDeviceId, softwareUpdatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
