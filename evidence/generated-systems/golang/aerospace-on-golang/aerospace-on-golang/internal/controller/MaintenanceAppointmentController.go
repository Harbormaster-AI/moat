package controller

import (
    MaintenanceAppointmentDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MaintenanceAppointmentDAO for database creation
//----------------------------------------------------------------------------
func CreateMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceAppointment model
	//----------------------------------------------------------------------------
	data := model.MaintenanceAppointment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceAppointment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment data access object to create
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.CreateMaintenanceAppointment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MaintenanceAppointmentDAO to find the relevant MaintenanceAppointment
//----------------------------------------------------------------------------
func GetMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceAppointment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.GetMaintenanceAppointment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MaintenanceAppointmentDAO for database read of all MaintenanceAppointments
//----------------------------------------------------------------------------
func GetAllMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.GetAllMaintenanceAppointment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MaintenanceAppointmentDAO for database save
//----------------------------------------------------------------------------
func UpdateMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceAppointment model
	//----------------------------------------------------------------------------
	var data = model.MaintenanceAppointment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceAppointment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.UpdateMaintenanceAppointment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MaintenanceAppointmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceAppointment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MaintenanceAppointmentDAO.DeleteMaintenanceAppointment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Aircraft on a MaintenanceAppointment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAircraftToMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceAppointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftId,_ := strconv.ParseUint( vars["aircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.AssignAircraftToMaintenanceAppointment(maintenanceAppointmentId, aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Aircraft on a MaintenanceAppointment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAircraftFromMaintenanceAppointment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceAppointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.UnassignAircraftFromMaintenanceAppointment(maintenanceAppointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a MroFacility on a MaintenanceAppointment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMroFacilityToMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceAppointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mroFacilityId,_ := strconv.ParseUint( vars["mroFacilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.AssignMroFacilityToMaintenanceAppointment(maintenanceAppointmentId, mroFacilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a MroFacility on a MaintenanceAppointment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMroFacilityFromMaintenanceAppointment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceAppointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.UnassignMroFacilityFromMaintenanceAppointment(maintenanceAppointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkOrder on a MaintenanceAppointment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkOrderToMaintenanceAppointment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceAppointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrderId,_ := strconv.ParseUint( vars["workOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.AssignWorkOrderToMaintenanceAppointment(maintenanceAppointmentId, workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkOrder on a MaintenanceAppointment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkOrderFromMaintenanceAppointment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceAppointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceAppointment DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceAppointmentDAO.UnassignWorkOrderFromMaintenanceAppointment(maintenanceAppointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


