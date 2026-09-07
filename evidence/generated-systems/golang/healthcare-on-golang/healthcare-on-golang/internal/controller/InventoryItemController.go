package controller

import (
    InventoryItemDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
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
	// assigns a Facility on a InventoryItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToInventoryItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AssignFacilityToInventoryItem(inventoryItemId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a InventoryItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromInventoryItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UnassignFacilityFromInventoryItem(inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Supplier on a InventoryItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSupplierToInventoryItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	supplierId,_ := strconv.ParseUint( vars["supplierId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AssignSupplierToInventoryItem(inventoryItemId, supplierId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Supplier on a InventoryItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSupplierFromInventoryItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UnassignSupplierFromInventoryItem(inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


