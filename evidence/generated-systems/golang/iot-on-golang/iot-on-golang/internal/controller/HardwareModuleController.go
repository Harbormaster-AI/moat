package controller

import (
    HardwareModuleDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to HardwareModuleDAO for database creation
//----------------------------------------------------------------------------
func CreateHardwareModule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty HardwareModule model
	//----------------------------------------------------------------------------
	data := model.HardwareModule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a HardwareModule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the HardwareModule data access object to create
	//----------------------------------------------------------------------------
	requestResult := HardwareModuleDAO.CreateHardwareModule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to HardwareModuleDAO to find the relevant HardwareModule
//----------------------------------------------------------------------------
func GetHardwareModule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the HardwareModule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := HardwareModuleDAO.GetHardwareModule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to HardwareModuleDAO for database read of all HardwareModules
//----------------------------------------------------------------------------
func GetAllHardwareModule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the HardwareModule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := HardwareModuleDAO.GetAllHardwareModule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to HardwareModuleDAO for database save
//----------------------------------------------------------------------------
func UpdateHardwareModule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty HardwareModule model
	//----------------------------------------------------------------------------
	var data = model.HardwareModule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a HardwareModule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the HardwareModule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := HardwareModuleDAO.UpdateHardwareModule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to HardwareModuleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteHardwareModule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the HardwareModule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := HardwareModuleDAO.DeleteHardwareModule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Vendor on a HardwareModule
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignVendorToHardwareModule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	hardwareModuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	vendorId,_ := strconv.ParseUint( vars["vendorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the HardwareModule DAO
	//----------------------------------------------------------------------------
	requestResult := HardwareModuleDAO.AssignVendorToHardwareModule(hardwareModuleId, vendorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Vendor on a HardwareModule
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignVendorFromHardwareModule( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	hardwareModuleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the HardwareModule DAO
	//----------------------------------------------------------------------------
	requestResult := HardwareModuleDAO.UnassignVendorFromHardwareModule(hardwareModuleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


