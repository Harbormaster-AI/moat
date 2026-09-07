package controller

import (
    MaintenanceWorkOrderDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MaintenanceWorkOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceWorkOrder model
	//----------------------------------------------------------------------------
	data := model.MaintenanceWorkOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceWorkOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.CreateMaintenanceWorkOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MaintenanceWorkOrderDAO to find the relevant MaintenanceWorkOrder
//----------------------------------------------------------------------------
func GetMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceWorkOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.GetMaintenanceWorkOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MaintenanceWorkOrderDAO for database read of all MaintenanceWorkOrders
//----------------------------------------------------------------------------
func GetAllMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.GetAllMaintenanceWorkOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MaintenanceWorkOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceWorkOrder model
	//----------------------------------------------------------------------------
	var data = model.MaintenanceWorkOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceWorkOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.UpdateMaintenanceWorkOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MaintenanceWorkOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceWorkOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MaintenanceWorkOrderDAO.DeleteMaintenanceWorkOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Aircraft on a MaintenanceWorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAircraftToMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceWorkOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftId,_ := strconv.ParseUint( vars["aircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.AssignAircraftToMaintenanceWorkOrder(maintenanceWorkOrderId, aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Aircraft on a MaintenanceWorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAircraftFromMaintenanceWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceWorkOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.UnassignAircraftFromMaintenanceWorkOrder(maintenanceWorkOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AirworthinessDirective on a MaintenanceWorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAirworthinessDirectiveToMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceWorkOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	airworthinessDirectiveId,_ := strconv.ParseUint( vars["airworthinessDirectiveId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.AssignAirworthinessDirectiveToMaintenanceWorkOrder(maintenanceWorkOrderId, airworthinessDirectiveId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AirworthinessDirective on a MaintenanceWorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAirworthinessDirectiveFromMaintenanceWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceWorkOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.UnassignAirworthinessDirectiveFromMaintenanceWorkOrder(maintenanceWorkOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ServiceBulletin on a MaintenanceWorkOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignServiceBulletinToMaintenanceWorkOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceWorkOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serviceBulletinId,_ := strconv.ParseUint( vars["serviceBulletinId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.AssignServiceBulletinToMaintenanceWorkOrder(maintenanceWorkOrderId, serviceBulletinId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ServiceBulletin on a MaintenanceWorkOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignServiceBulletinFromMaintenanceWorkOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceWorkOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceWorkOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceWorkOrderDAO.UnassignServiceBulletinFromMaintenanceWorkOrder(maintenanceWorkOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


