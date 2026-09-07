package controller

import (
    WorkCenterDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to WorkCenterDAO for database creation
//----------------------------------------------------------------------------
func CreateWorkCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkCenter model
	//----------------------------------------------------------------------------
	data := model.WorkCenter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkCenter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter data access object to create
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.CreateWorkCenter( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to WorkCenterDAO to find the relevant WorkCenter
//----------------------------------------------------------------------------
func GetWorkCenter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkCenter data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.GetWorkCenter(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to WorkCenterDAO for database read of all WorkCenters
//----------------------------------------------------------------------------
func GetAllWorkCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter data access object to get all
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.GetAllWorkCenter()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to WorkCenterDAO for database save
//----------------------------------------------------------------------------
func UpdateWorkCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkCenter model
	//----------------------------------------------------------------------------
	var data = model.WorkCenter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkCenter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.UpdateWorkCenter(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to WorkCenterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteWorkCenter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkCenter data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := WorkCenterDAO.DeleteWorkCenter(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ProductionLine on a WorkCenter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductionLineToWorkCenter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionLineId,_ := strconv.ParseUint( vars["productionLineId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter DAO
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.AssignProductionLineToWorkCenter(workCenterId, productionLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ProductionLine on a WorkCenter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductionLineFromWorkCenter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter DAO
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.UnassignProductionLineFromWorkCenter(workCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more assetsIds as a Assets to a WorkCenter
	//----------------------------------------------------------------------------
func AddAssetsToWorkCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assetsIds,_ := vars["assetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter DAO
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.AddAssetsToWorkCenter(workCenterId, assetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more assetsIds as a Assets from a WorkCenter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAssetsFromWorkCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assetsIds,_ := vars["assetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter DAO
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.RemoveAssetsFromWorkCenter(workCenterId, assetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more maintenanceOrdersIds as a MaintenanceOrders to a WorkCenter
	//----------------------------------------------------------------------------
func AddMaintenanceOrdersToWorkCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceOrdersIds,_ := vars["maintenanceOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter DAO
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.AddMaintenanceOrdersToWorkCenter(workCenterId, maintenanceOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more maintenanceOrdersIds as a MaintenanceOrders from a WorkCenter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMaintenanceOrdersFromWorkCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceOrdersIds,_ := vars["maintenanceOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkCenter DAO
	//----------------------------------------------------------------------------
	requestResult := WorkCenterDAO.RemoveMaintenanceOrdersFromWorkCenter(workCenterId, maintenanceOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
