package controller

import (
    SoftwareUpdateDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SoftwareUpdateDAO for database creation
//----------------------------------------------------------------------------
func CreateSoftwareUpdate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SoftwareUpdate model
	//----------------------------------------------------------------------------
	data := model.SoftwareUpdate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SoftwareUpdate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareUpdate data access object to create
	//----------------------------------------------------------------------------
	requestResult := SoftwareUpdateDAO.CreateSoftwareUpdate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SoftwareUpdateDAO to find the relevant SoftwareUpdate
//----------------------------------------------------------------------------
func GetSoftwareUpdate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SoftwareUpdate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SoftwareUpdateDAO.GetSoftwareUpdate(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SoftwareUpdateDAO for database read of all SoftwareUpdates
//----------------------------------------------------------------------------
func GetAllSoftwareUpdate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SoftwareUpdate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SoftwareUpdateDAO.GetAllSoftwareUpdate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SoftwareUpdateDAO for database save
//----------------------------------------------------------------------------
func UpdateSoftwareUpdate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SoftwareUpdate model
	//----------------------------------------------------------------------------
	var data = model.SoftwareUpdate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SoftwareUpdate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareUpdate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SoftwareUpdateDAO.UpdateSoftwareUpdate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SoftwareUpdateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSoftwareUpdate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SoftwareUpdate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SoftwareUpdateDAO.DeleteSoftwareUpdate(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Device on a SoftwareUpdate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToSoftwareUpdate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	softwareUpdateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareUpdate DAO
	//----------------------------------------------------------------------------
	requestResult := SoftwareUpdateDAO.AssignDeviceToSoftwareUpdate(softwareUpdateId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a SoftwareUpdate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromSoftwareUpdate( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	softwareUpdateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SoftwareUpdate DAO
	//----------------------------------------------------------------------------
	requestResult := SoftwareUpdateDAO.UnassignDeviceFromSoftwareUpdate(softwareUpdateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


