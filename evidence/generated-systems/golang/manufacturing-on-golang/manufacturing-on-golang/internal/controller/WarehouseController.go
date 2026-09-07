package controller

import (
    WarehouseDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
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
	// assigns a Plant on a Warehouse
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToWarehouse(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AssignPlantToWarehouse(warehouseId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a Warehouse
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromWarehouse( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.UnassignPlantFromWarehouse(warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more locationsIds as a Locations to a Warehouse
	//----------------------------------------------------------------------------
func AddLocationsToWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationsIds,_ := vars["locationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.AddLocationsToWarehouse(warehouseId, locationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more locationsIds as a Locations from a Warehouse
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLocationsFromWarehouse(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	warehouseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationsIds,_ := vars["locationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Warehouse DAO
	//----------------------------------------------------------------------------
	requestResult := WarehouseDAO.RemoveLocationsFromWarehouse(warehouseId, locationsIds)

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
		
