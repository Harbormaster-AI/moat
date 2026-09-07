package controller

import (
    StorageLocationDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to StorageLocationDAO for database creation
//----------------------------------------------------------------------------
func CreateStorageLocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StorageLocation model
	//----------------------------------------------------------------------------
	data := model.StorageLocation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StorageLocation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation data access object to create
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.CreateStorageLocation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to StorageLocationDAO to find the relevant StorageLocation
//----------------------------------------------------------------------------
func GetStorageLocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StorageLocation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.GetStorageLocation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to StorageLocationDAO for database read of all StorageLocations
//----------------------------------------------------------------------------
func GetAllStorageLocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.GetAllStorageLocation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to StorageLocationDAO for database save
//----------------------------------------------------------------------------
func UpdateStorageLocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StorageLocation model
	//----------------------------------------------------------------------------
	var data = model.StorageLocation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StorageLocation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.UpdateStorageLocation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to StorageLocationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteStorageLocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StorageLocation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := StorageLocationDAO.DeleteStorageLocation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a StorageLocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToStorageLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.AssignWarehouseToStorageLocation(storageLocationId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a StorageLocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromStorageLocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.UnassignWarehouseFromStorageLocation(storageLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ParentLocation on a StorageLocation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignParentLocationToStorageLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	parentLocationId,_ := strconv.ParseUint( vars["parentLocationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.AssignParentLocationToStorageLocation(storageLocationId, parentLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ParentLocation on a StorageLocation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignParentLocationFromStorageLocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.UnassignParentLocationFromStorageLocation(storageLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more childLocationsIds as a ChildLocations to a StorageLocation
	//----------------------------------------------------------------------------
func AddChildLocationsToStorageLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childLocationsIds,_ := vars["childLocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.AddChildLocationsToStorageLocation(storageLocationId, childLocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more childLocationsIds as a ChildLocations from a StorageLocation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveChildLocationsFromStorageLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childLocationsIds,_ := vars["childLocationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.RemoveChildLocationsFromStorageLocation(storageLocationId, childLocationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more inventoryItemsIds as a InventoryItems to a StorageLocation
	//----------------------------------------------------------------------------
func AddInventoryItemsToStorageLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.AddInventoryItemsToStorageLocation(storageLocationId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventoryItemsIds as a InventoryItems from a StorageLocation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventoryItemsFromStorageLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	storageLocationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StorageLocation DAO
	//----------------------------------------------------------------------------
	requestResult := StorageLocationDAO.RemoveInventoryItemsFromStorageLocation(storageLocationId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
