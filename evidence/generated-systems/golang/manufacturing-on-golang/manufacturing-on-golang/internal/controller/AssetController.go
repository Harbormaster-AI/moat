package controller

import (
    AssetDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AssetDAO for database creation
//----------------------------------------------------------------------------
func CreateAsset(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Asset model
	//----------------------------------------------------------------------------
	data := model.Asset{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Asset model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Asset data access object to create
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.CreateAsset( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AssetDAO to find the relevant Asset
//----------------------------------------------------------------------------
func GetAsset(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Asset data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.GetAsset(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AssetDAO for database read of all Assets
//----------------------------------------------------------------------------
func GetAllAsset(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Asset data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.GetAllAsset()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AssetDAO for database save
//----------------------------------------------------------------------------
func UpdateAsset(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Asset model
	//----------------------------------------------------------------------------
	var data = model.Asset{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Asset model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Asset data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.UpdateAsset(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AssetDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAsset(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Asset data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AssetDAO.DeleteAsset(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Plant on a Asset
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToAsset(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.AssignPlantToAsset(assetId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a Asset
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromAsset( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.UnassignPlantFromAsset(assetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkCenter on a Asset
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkCenterToAsset(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCenterId,_ := strconv.ParseUint( vars["workCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.AssignWorkCenterToAsset(assetId, workCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkCenter on a Asset
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkCenterFromAsset( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.UnassignWorkCenterFromAsset(assetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more maintenanceOrdersIds as a MaintenanceOrders to a Asset
	//----------------------------------------------------------------------------
func AddMaintenanceOrdersToAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceOrdersIds,_ := vars["maintenanceOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.AddMaintenanceOrdersToAsset(assetId, maintenanceOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more maintenanceOrdersIds as a MaintenanceOrders from a Asset
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMaintenanceOrdersFromAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenanceOrdersIds,_ := vars["maintenanceOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.RemoveMaintenanceOrdersFromAsset(assetId, maintenanceOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more maintenancePlansIds as a MaintenancePlans to a Asset
	//----------------------------------------------------------------------------
func AddMaintenancePlansToAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenancePlansIds,_ := vars["maintenancePlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.AddMaintenancePlansToAsset(assetId, maintenancePlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more maintenancePlansIds as a MaintenancePlans from a Asset
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMaintenancePlansFromAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	assetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	maintenancePlansIds,_ := vars["maintenancePlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Asset DAO
	//----------------------------------------------------------------------------
	requestResult := AssetDAO.RemoveMaintenancePlansFromAsset(assetId, maintenancePlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
