package controller

import (
    TransferOrderDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TransferOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateTransferOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TransferOrder model
	//----------------------------------------------------------------------------
	data := model.TransferOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TransferOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.CreateTransferOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TransferOrderDAO to find the relevant TransferOrder
//----------------------------------------------------------------------------
func GetTransferOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TransferOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.GetTransferOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TransferOrderDAO for database read of all TransferOrders
//----------------------------------------------------------------------------
func GetAllTransferOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.GetAllTransferOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TransferOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateTransferOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TransferOrder model
	//----------------------------------------------------------------------------
	var data = model.TransferOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TransferOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.UpdateTransferOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TransferOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTransferOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TransferOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TransferOrderDAO.DeleteTransferOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a OriginWarehouse on a TransferOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOriginWarehouseToTransferOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	originWarehouseId,_ := strconv.ParseUint( vars["originWarehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.AssignOriginWarehouseToTransferOrder(transferOrderId, originWarehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a OriginWarehouse on a TransferOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOriginWarehouseFromTransferOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.UnassignOriginWarehouseFromTransferOrder(transferOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DestinationWarehouse on a TransferOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDestinationWarehouseToTransferOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationWarehouseId,_ := strconv.ParseUint( vars["destinationWarehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.AssignDestinationWarehouseToTransferOrder(transferOrderId, destinationWarehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DestinationWarehouse on a TransferOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDestinationWarehouseFromTransferOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.UnassignDestinationWarehouseFromTransferOrder(transferOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more linesIds as a Lines to a TransferOrder
	//----------------------------------------------------------------------------
func AddLinesToTransferOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.AddLinesToTransferOrder(transferOrderId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more linesIds as a Lines from a TransferOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLinesFromTransferOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.RemoveLinesFromTransferOrder(transferOrderId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a TransferOrder
	//----------------------------------------------------------------------------
func AddTransactionsToTransferOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.AddTransactionsToTransferOrder(transferOrderId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a TransferOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromTransferOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transferOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderDAO.RemoveTransactionsFromTransferOrder(transferOrderId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
