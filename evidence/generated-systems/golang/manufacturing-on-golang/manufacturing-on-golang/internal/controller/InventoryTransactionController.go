package controller

import (
    InventoryTransactionDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
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
	// assigns a Item on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignItemToInventoryTransaction(inventoryTransactionId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignItemFromInventoryTransaction(inventoryTransactionId)

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
	// assigns a WorkOrder on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkOrderToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrderId,_ := strconv.ParseUint( vars["workOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignWorkOrderToInventoryTransaction(inventoryTransactionId, workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkOrder on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkOrderFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignWorkOrderFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PurchaseOrder on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPurchaseOrderToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	purchaseOrderId,_ := strconv.ParseUint( vars["purchaseOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignPurchaseOrderToInventoryTransaction(inventoryTransactionId, purchaseOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PurchaseOrder on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPurchaseOrderFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignPurchaseOrderFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a SalesOrder on a InventoryTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSalesOrderToInventoryTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesOrderId,_ := strconv.ParseUint( vars["salesOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.AssignSalesOrderToInventoryTransaction(inventoryTransactionId, salesOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SalesOrder on a InventoryTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSalesOrderFromInventoryTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryTransactionDAO.UnassignSalesOrderFromInventoryTransaction(inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


