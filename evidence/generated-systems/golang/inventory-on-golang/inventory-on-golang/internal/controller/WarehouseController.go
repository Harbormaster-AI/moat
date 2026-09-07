package controller

import (
    WarehouseDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to WarehouseDAO for database creation
//----------------------------------------------------------------------------
func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Warehouse model
	//----------------------------------------------------------------------------
	data := model.Warehouse{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Warehouse model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse data access object to create
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.CreateWarehouse( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to WarehouseDAO to find the relevant Warehouse
//----------------------------------------------------------------------------
func GetWarehouse(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Warehouse data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.GetWarehouse(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to WarehouseDAO for database read of all Warehouses
//----------------------------------------------------------------------------
func GetAllWarehouse(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Warehouse data access object to get all
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.GetAllWarehouse()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to WarehouseDAO for database save
//----------------------------------------------------------------------------
func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Warehouse model
	//----------------------------------------------------------------------------
	var data = model.Warehouse{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Warehouse model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.UpdateWarehouse(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to WarehouseDAO for database deletion
//----------------------------------------------------------------------------
func DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Warehouse data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := WarehouseDAO.DeleteWarehouse(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more storageLocationsIds as a StorageLocations to a Warehouse
	//----------------------------------------------------------------------------
func AddStorageLocationsToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	storageLocationsIds,_ := vars["storageLocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddStorageLocationsToWarehouse(warehouseId, storageLocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more storageLocationsIds as a StorageLocations from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStorageLocationsFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	storageLocationsIds,_ := vars["storageLocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveStorageLocationsFromWarehouse(warehouseId, storageLocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more inventoryItemsIds as a InventoryItems to a Warehouse
	//----------------------------------------------------------------------------
func AddInventoryItemsToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddInventoryItemsToWarehouse(warehouseId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventoryItemsIds as a InventoryItems from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventoryItemsFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveInventoryItemsFromWarehouse(warehouseId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more inboundShipmentsIds as a InboundShipments to a Warehouse
	//----------------------------------------------------------------------------
func AddInboundShipmentsToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inboundShipmentsIds,_ := vars["inboundShipmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddInboundShipmentsToWarehouse(warehouseId, inboundShipmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inboundShipmentsIds as a InboundShipments from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInboundShipmentsFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inboundShipmentsIds,_ := vars["inboundShipmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveInboundShipmentsFromWarehouse(warehouseId, inboundShipmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more outboundAllocationsIds as a OutboundAllocations to a Warehouse
	//----------------------------------------------------------------------------
func AddOutboundAllocationsToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outboundAllocationsIds,_ := vars["outboundAllocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddOutboundAllocationsToWarehouse(warehouseId, outboundAllocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more outboundAllocationsIds as a OutboundAllocations from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOutboundAllocationsFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outboundAllocationsIds,_ := vars["outboundAllocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveOutboundAllocationsFromWarehouse(warehouseId, outboundAllocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more originTransfersIds as a OriginTransfers to a Warehouse
	//----------------------------------------------------------------------------
func AddOriginTransfersToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	originTransfersIds,_ := vars["originTransfersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddOriginTransfersToWarehouse(warehouseId, originTransfersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more originTransfersIds as a OriginTransfers from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOriginTransfersFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	originTransfersIds,_ := vars["originTransfersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveOriginTransfersFromWarehouse(warehouseId, originTransfersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more destinationTransfersIds as a DestinationTransfers to a Warehouse
	//----------------------------------------------------------------------------
func AddDestinationTransfersToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationTransfersIds,_ := vars["destinationTransfersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddDestinationTransfersToWarehouse(warehouseId, destinationTransfersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more destinationTransfersIds as a DestinationTransfers from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDestinationTransfersFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationTransfersIds,_ := vars["destinationTransfersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveDestinationTransfersFromWarehouse(warehouseId, destinationTransfersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more cycleCountsIds as a CycleCounts to a Warehouse
	//----------------------------------------------------------------------------
func AddCycleCountsToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cycleCountsIds,_ := vars["cycleCountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddCycleCountsToWarehouse(warehouseId, cycleCountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more cycleCountsIds as a CycleCounts from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCycleCountsFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cycleCountsIds,_ := vars["cycleCountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveCycleCountsFromWarehouse(warehouseId, cycleCountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
