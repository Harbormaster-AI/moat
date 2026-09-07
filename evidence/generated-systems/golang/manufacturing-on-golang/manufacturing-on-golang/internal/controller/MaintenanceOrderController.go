package controller

import (
    MaintenanceOrderDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MaintenanceOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateMaintenanceOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceOrder model
	//----------------------------------------------------------------------------
	data := model.MaintenanceOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.CreateMaintenanceOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MaintenanceOrderDAO to find the relevant MaintenanceOrder
//----------------------------------------------------------------------------
func GetMaintenanceOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.GetMaintenanceOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MaintenanceOrderDAO for database read of all MaintenanceOrders
//----------------------------------------------------------------------------
func GetAllMaintenanceOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.GetAllMaintenanceOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MaintenanceOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateMaintenanceOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MaintenanceOrder model
	//----------------------------------------------------------------------------
	var data = model.MaintenanceOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MaintenanceOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.UpdateMaintenanceOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MaintenanceOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMaintenanceOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MaintenanceOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MaintenanceOrderDAO.DeleteMaintenanceOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Asset on a MaintenanceOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAssetToMaintenanceOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assetId,_ := strconv.ParseUint( vars["assetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.AssignAssetToMaintenanceOrder(maintenanceOrderId, assetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Asset on a MaintenanceOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAssetFromMaintenanceOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.UnassignAssetFromMaintenanceOrder(maintenanceOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Plan on a MaintenanceOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlanToMaintenanceOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	planId,_ := strconv.ParseUint( vars["planId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.AssignPlanToMaintenanceOrder(maintenanceOrderId, planId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plan on a MaintenanceOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlanFromMaintenanceOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.UnassignPlanFromMaintenanceOrder(maintenanceOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkCenter on a MaintenanceOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkCenterToMaintenanceOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCenterId,_ := strconv.ParseUint( vars["workCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.AssignWorkCenterToMaintenanceOrder(maintenanceOrderId, workCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkCenter on a MaintenanceOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkCenterFromMaintenanceOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	maintenanceOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MaintenanceOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MaintenanceOrderDAO.UnassignWorkCenterFromMaintenanceOrder(maintenanceOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


