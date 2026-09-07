package controller

import (
    InventoryTransactionDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InventoryTransactionDAO for database creation
//----------------------------------------------------------------------------
func CreateInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventoryTransaction model
	//----------------------------------------------------------------------------
	data := model.InventoryTransaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventoryTransaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction data access object to create
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.CreateInventoryTransaction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InventoryTransactionDAO to find the relevant InventoryTransaction
//----------------------------------------------------------------------------
func GetInventoryTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventoryTransaction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.GetInventoryTransaction(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InventoryTransactionDAO for database read of all InventoryTransactions
//----------------------------------------------------------------------------
func GetAllInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.GetAllInventoryTransaction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InventoryTransactionDAO for database save
//----------------------------------------------------------------------------
func UpdateInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventoryTransaction model
	//----------------------------------------------------------------------------
	var data = model.InventoryTransaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventoryTransaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UpdateInventoryTransaction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InventoryTransactionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInventoryTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventoryTransaction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InventoryTransactionDAO.DeleteInventoryTransaction(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignSkuToInventoryTransaction(inventoryTransactionId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignSkuFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignWarehouseToInventoryTransaction(inventoryTransactionId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignWarehouseFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignLocationToInventoryTransaction(inventoryTransactionId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignLocationFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignLotToInventoryTransaction(inventoryTransactionId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignLotFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a RelatedReservation on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRelatedReservationToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	relatedReservationId,_ := strconv.ParseUint( vars["relatedReservationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignRelatedReservationToInventoryTransaction(inventoryTransactionId, relatedReservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RelatedReservation on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRelatedReservationFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignRelatedReservationFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a TransferOrder on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransferOrderToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transferOrderId,_ := strconv.ParseUint( vars["transferOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignTransferOrderToInventoryTransaction(inventoryTransactionId, transferOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TransferOrder on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransferOrderFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignTransferOrderFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Adjustment on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdjustmentToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adjustmentId,_ := strconv.ParseUint( vars["adjustmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignAdjustmentToInventoryTransaction(inventoryTransactionId, adjustmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Adjustment on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdjustmentFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignAdjustmentFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CycleCount on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCycleCountToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cycleCountId,_ := strconv.ParseUint( vars["cycleCountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignCycleCountToInventoryTransaction(inventoryTransactionId, cycleCountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CycleCount on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCycleCountFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignCycleCountFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a InventoryTransaction
	//----------------------------------------------------------------------------
func AddSerialNumbersToInventoryTransaction(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AddSerialNumbersToInventoryTransaction(inventoryTransactionId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a InventoryTransaction
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromInventoryTransaction(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.RemoveSerialNumbersFromInventoryTransaction(inventoryTransactionId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
