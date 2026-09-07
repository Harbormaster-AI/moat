package controller

import (
    MaintenanceTicketDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MaintenanceTicketDAO for database creation
//----------------------------------------------------------------------------
func CreateMaintenanceTicket(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceTicket model
	//----------------------------------------------------------------------------
	data := model.MaintenanceTicket{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceTicket model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceTicket data access object to create
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.CreateMaintenanceTicket( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MaintenanceTicketDAO to find the relevant MaintenanceTicket
//----------------------------------------------------------------------------
func GetMaintenanceTicket(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceTicket data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.GetMaintenanceTicket(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MaintenanceTicketDAO for database read of all MaintenanceTickets
//----------------------------------------------------------------------------
func GetAllMaintenanceTicket(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceTicket data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.GetAllMaintenanceTicket()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MaintenanceTicketDAO for database save
//----------------------------------------------------------------------------
func UpdateMaintenanceTicket(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceTicket model
	//----------------------------------------------------------------------------
	var data = model.MaintenanceTicket{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceTicket model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceTicket data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.UpdateMaintenanceTicket(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MaintenanceTicketDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMaintenanceTicket(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceTicket data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MaintenanceTicketDAO.DeleteMaintenanceTicket(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Device on a MaintenanceTicket
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToMaintenanceTicket(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceTicketId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceTicket DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.AssignDeviceToMaintenanceTicket(maintenanceTicketId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a MaintenanceTicket
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromMaintenanceTicket( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceTicketId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceTicket DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.UnassignDeviceFromMaintenanceTicket(maintenanceTicketId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a MaintenanceTicket
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToMaintenanceTicket(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceTicketId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceTicket DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.AssignTenantToMaintenanceTicket(maintenanceTicketId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a MaintenanceTicket
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromMaintenanceTicket( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceTicketId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceTicket DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceTicketDAO.UnassignTenantFromMaintenanceTicket(maintenanceTicketId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


