package controller

import (
    TransferOrderLineDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TransferOrderLineDAO for database creation
//----------------------------------------------------------------------------
func CreateTransferOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TransferOrderLine model
	//----------------------------------------------------------------------------
	data := model.TransferOrderLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TransferOrderLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.CreateTransferOrderLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TransferOrderLineDAO to find the relevant TransferOrderLine
//----------------------------------------------------------------------------
func GetTransferOrderLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TransferOrderLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.GetTransferOrderLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TransferOrderLineDAO for database read of all TransferOrderLines
//----------------------------------------------------------------------------
func GetAllTransferOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.GetAllTransferOrderLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TransferOrderLineDAO for database save
//----------------------------------------------------------------------------
func UpdateTransferOrderLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TransferOrderLine model
	//----------------------------------------------------------------------------
	var data = model.TransferOrderLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TransferOrderLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.UpdateTransferOrderLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TransferOrderLineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTransferOrderLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TransferOrderLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TransferOrderLineDAO.DeleteTransferOrderLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a TransferOrder on a TransferOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransferOrderToTransferOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transferOrderId,_ := strconv.ParseUint( vars["transferOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.AssignTransferOrderToTransferOrderLine(transferOrderLineId, transferOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TransferOrder on a TransferOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransferOrderFromTransferOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.UnassignTransferOrderFromTransferOrderLine(transferOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Sku on a TransferOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToTransferOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.AssignSkuToTransferOrderLine(transferOrderLineId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a TransferOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromTransferOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.UnassignSkuFromTransferOrderLine(transferOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a TransferOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToTransferOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.AssignLotToTransferOrderLine(transferOrderLineId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a TransferOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromTransferOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.UnassignLotFromTransferOrderLine(transferOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a FromLocation on a TransferOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFromLocationToTransferOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	fromLocationId,_ := strconv.ParseUint( vars["fromLocationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.AssignFromLocationToTransferOrderLine(transferOrderLineId, fromLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a FromLocation on a TransferOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFromLocationFromTransferOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.UnassignFromLocationFromTransferOrderLine(transferOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ToLocation on a TransferOrderLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignToLocationToTransferOrderLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	toLocationId,_ := strconv.ParseUint( vars["toLocationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.AssignToLocationToTransferOrderLine(transferOrderLineId, toLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ToLocation on a TransferOrderLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignToLocationFromTransferOrderLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.UnassignToLocationFromTransferOrderLine(transferOrderLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a TransferOrderLine
	//----------------------------------------------------------------------------
func AddSerialNumbersToTransferOrderLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.AddSerialNumbersToTransferOrderLine(transferOrderLineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a TransferOrderLine
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromTransferOrderLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	transferOrderLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TransferOrderLine DAO
	//----------------------------------------------------------------------------
	requestResult := TransferOrderLineDAO.RemoveSerialNumbersFromTransferOrderLine(transferOrderLineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
