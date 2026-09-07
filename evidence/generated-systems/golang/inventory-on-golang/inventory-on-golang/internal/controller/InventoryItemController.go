package controller

import (
    InventoryItemDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InventoryItemDAO for database creation
//----------------------------------------------------------------------------
func CreateInventoryItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventoryItem model
	//----------------------------------------------------------------------------
	data := model.InventoryItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventoryItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem data access object to create
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.CreateInventoryItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InventoryItemDAO to find the relevant InventoryItem
//----------------------------------------------------------------------------
func GetInventoryItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventoryItem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.GetInventoryItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InventoryItemDAO for database read of all InventoryItems
//----------------------------------------------------------------------------
func GetAllInventoryItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.GetAllInventoryItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InventoryItemDAO for database save
//----------------------------------------------------------------------------
func UpdateInventoryItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventoryItem model
	//----------------------------------------------------------------------------
	var data = model.InventoryItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventoryItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UpdateInventoryItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InventoryItemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventoryItem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InventoryItemDAO.DeleteInventoryItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a InventoryItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToInventoryItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AssignSkuToInventoryItem(inventoryItemId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a InventoryItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromInventoryItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UnassignSkuFromInventoryItem(inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a InventoryItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToInventoryItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AssignWarehouseToInventoryItem(inventoryItemId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a InventoryItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromInventoryItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UnassignWarehouseFromInventoryItem(inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a InventoryItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToInventoryItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AssignLocationToInventoryItem(inventoryItemId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a InventoryItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromInventoryItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UnassignLocationFromInventoryItem(inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a InventoryItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToInventoryItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AssignLotToInventoryItem(inventoryItemId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a InventoryItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromInventoryItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UnassignLotFromInventoryItem(inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a InventoryItem
	//----------------------------------------------------------------------------
func AddSerialNumbersToInventoryItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AddSerialNumbersToInventoryItem(inventoryItemId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a InventoryItem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromInventoryItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.RemoveSerialNumbersFromInventoryItem(inventoryItemId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a InventoryItem
	//----------------------------------------------------------------------------
func AddTransactionsToInventoryItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AddTransactionsToInventoryItem(inventoryItemId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a InventoryItem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromInventoryItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.RemoveTransactionsFromInventoryItem(inventoryItemId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reservationsIds as a Reservations to a InventoryItem
	//----------------------------------------------------------------------------
func AddReservationsToInventoryItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservationsIds,_ := vars["reservationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AddReservationsToInventoryItem(inventoryItemId, reservationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reservationsIds as a Reservations from a InventoryItem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReservationsFromInventoryItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservationsIds,_ := vars["reservationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.RemoveReservationsFromInventoryItem(inventoryItemId, reservationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
