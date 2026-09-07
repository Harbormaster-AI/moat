package controller

import (
    MaintenancePlanDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MaintenancePlanDAO for database creation
//----------------------------------------------------------------------------
func CreateMaintenancePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenancePlan model
	//----------------------------------------------------------------------------
	data := model.MaintenancePlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenancePlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenancePlan data access object to create
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.CreateMaintenancePlan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MaintenancePlanDAO to find the relevant MaintenancePlan
//----------------------------------------------------------------------------
func GetMaintenancePlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenancePlan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.GetMaintenancePlan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MaintenancePlanDAO for database read of all MaintenancePlans
//----------------------------------------------------------------------------
func GetAllMaintenancePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MaintenancePlan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.GetAllMaintenancePlan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MaintenancePlanDAO for database save
//----------------------------------------------------------------------------
func UpdateMaintenancePlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenancePlan model
	//----------------------------------------------------------------------------
	var data = model.MaintenancePlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenancePlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenancePlan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.UpdateMaintenancePlan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MaintenancePlanDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMaintenancePlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenancePlan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MaintenancePlanDAO.DeleteMaintenancePlan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Asset on a MaintenancePlan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAssetToMaintenancePlan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assetId,_ := strconv.ParseUint( vars["assetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.AssignAssetToMaintenancePlan(maintenancePlanId, assetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Asset on a MaintenancePlan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAssetFromMaintenancePlan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.UnassignAssetFromMaintenancePlan(maintenancePlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more maintenanceOrdersIds as a MaintenanceOrders to a MaintenancePlan
	//----------------------------------------------------------------------------
func AddMaintenanceOrdersToMaintenancePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	maintenancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceOrdersIds,_ := vars["maintenanceOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MaintenancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.AddMaintenanceOrdersToMaintenancePlan(maintenancePlanId, maintenanceOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more maintenanceOrdersIds as a MaintenanceOrders from a MaintenancePlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMaintenanceOrdersFromMaintenancePlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	maintenancePlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceOrdersIds,_ := vars["maintenanceOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MaintenancePlan DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenancePlanDAO.RemoveMaintenanceOrdersFromMaintenancePlan(maintenancePlanId, maintenanceOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
