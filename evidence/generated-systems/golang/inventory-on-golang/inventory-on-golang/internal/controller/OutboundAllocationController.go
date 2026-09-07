package controller

import (
    OutboundAllocationDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OutboundAllocationDAO for database creation
//----------------------------------------------------------------------------
func CreateOutboundAllocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OutboundAllocation model
	//----------------------------------------------------------------------------
	data := model.OutboundAllocation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OutboundAllocation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation data access object to create
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.CreateOutboundAllocation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OutboundAllocationDAO to find the relevant OutboundAllocation
//----------------------------------------------------------------------------
func GetOutboundAllocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OutboundAllocation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.GetOutboundAllocation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OutboundAllocationDAO for database read of all OutboundAllocations
//----------------------------------------------------------------------------
func GetAllOutboundAllocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.GetAllOutboundAllocation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OutboundAllocationDAO for database save
//----------------------------------------------------------------------------
func UpdateOutboundAllocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OutboundAllocation model
	//----------------------------------------------------------------------------
	var data = model.OutboundAllocation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OutboundAllocation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.UpdateOutboundAllocation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OutboundAllocationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOutboundAllocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OutboundAllocation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OutboundAllocationDAO.DeleteOutboundAllocation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a OutboundAllocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToOutboundAllocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.AssignWarehouseToOutboundAllocation(outboundAllocationId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a OutboundAllocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromOutboundAllocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.UnassignWarehouseFromOutboundAllocation(outboundAllocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Sku on a OutboundAllocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToOutboundAllocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.AssignSkuToOutboundAllocation(outboundAllocationId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a OutboundAllocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromOutboundAllocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.UnassignSkuFromOutboundAllocation(outboundAllocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InventoryItem on a OutboundAllocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInventoryItemToOutboundAllocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemId,_ := strconv.ParseUint( vars["inventoryItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.AssignInventoryItemToOutboundAllocation(outboundAllocationId, inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InventoryItem on a OutboundAllocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInventoryItemFromOutboundAllocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.UnassignInventoryItemFromOutboundAllocation(outboundAllocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Reservation on a OutboundAllocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignReservationToOutboundAllocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservationId,_ := strconv.ParseUint( vars["reservationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.AssignReservationToOutboundAllocation(outboundAllocationId, reservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Reservation on a OutboundAllocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignReservationFromOutboundAllocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.UnassignReservationFromOutboundAllocation(outboundAllocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a OutboundAllocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToOutboundAllocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.AssignLotToOutboundAllocation(outboundAllocationId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a OutboundAllocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromOutboundAllocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.UnassignLotFromOutboundAllocation(outboundAllocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a SourceLocation on a OutboundAllocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSourceLocationToOutboundAllocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourceLocationId,_ := strconv.ParseUint( vars["sourceLocationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.AssignSourceLocationToOutboundAllocation(outboundAllocationId, sourceLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SourceLocation on a OutboundAllocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSourceLocationFromOutboundAllocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.UnassignSourceLocationFromOutboundAllocation(outboundAllocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a OutboundAllocation
	//----------------------------------------------------------------------------
func AddSerialNumbersToOutboundAllocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.AddSerialNumbersToOutboundAllocation(outboundAllocationId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a OutboundAllocation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromOutboundAllocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	outboundAllocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the OutboundAllocation DAO
	//----------------------------------------------------------------------------
	requestResult := OutboundAllocationDAO.RemoveSerialNumbersFromOutboundAllocation(outboundAllocationId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
