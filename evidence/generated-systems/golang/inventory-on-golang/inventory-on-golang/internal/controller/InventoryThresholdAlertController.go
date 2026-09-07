package controller

import (
    InventoryThresholdAlertDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InventoryThresholdAlertDAO for database creation
//----------------------------------------------------------------------------
func CreateInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventoryThresholdAlert model
	//----------------------------------------------------------------------------
	data := model.InventoryThresholdAlert{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventoryThresholdAlert model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert data access object to create
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.CreateInventoryThresholdAlert( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InventoryThresholdAlertDAO to find the relevant InventoryThresholdAlert
//----------------------------------------------------------------------------
func GetInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventoryThresholdAlert data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.GetInventoryThresholdAlert(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InventoryThresholdAlertDAO for database read of all InventoryThresholdAlerts
//----------------------------------------------------------------------------
func GetAllInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.GetAllInventoryThresholdAlert()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InventoryThresholdAlertDAO for database save
//----------------------------------------------------------------------------
func UpdateInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventoryThresholdAlert model
	//----------------------------------------------------------------------------
	var data = model.InventoryThresholdAlert{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventoryThresholdAlert model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.UpdateInventoryThresholdAlert(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InventoryThresholdAlertDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventoryThresholdAlert data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InventoryThresholdAlertDAO.DeleteInventoryThresholdAlert(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a InventoryThresholdAlert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.AssignSkuToInventoryThresholdAlert(inventoryThresholdAlertId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a InventoryThresholdAlert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromInventoryThresholdAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.UnassignSkuFromInventoryThresholdAlert(inventoryThresholdAlertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a InventoryThresholdAlert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.AssignWarehouseToInventoryThresholdAlert(inventoryThresholdAlertId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a InventoryThresholdAlert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromInventoryThresholdAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.UnassignWarehouseFromInventoryThresholdAlert(inventoryThresholdAlertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a InventoryThresholdAlert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.AssignLocationToInventoryThresholdAlert(inventoryThresholdAlertId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a InventoryThresholdAlert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromInventoryThresholdAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.UnassignLocationFromInventoryThresholdAlert(inventoryThresholdAlertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a RelatedPolicy on a InventoryThresholdAlert
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRelatedPolicyToInventoryThresholdAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedPolicyId,_ := strconv.ParseUint( vars["relatedPolicyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.AssignRelatedPolicyToInventoryThresholdAlert(inventoryThresholdAlertId, relatedPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RelatedPolicy on a InventoryThresholdAlert
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRelatedPolicyFromInventoryThresholdAlert( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryThresholdAlertId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryThresholdAlert DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryThresholdAlertDAO.UnassignRelatedPolicyFromInventoryThresholdAlert(inventoryThresholdAlertId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


