package controller

import (
    InventoryItemDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
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
	// assigns a Item on a InventoryItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToInventoryItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.AssignItemToInventoryItem(inventoryItemId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a InventoryItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromInventoryItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventoryItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventoryItem DAO
	//----------------------------------------------------------------------------
	requestResult := InventoryItemDAO.UnassignItemFromInventoryItem(inventoryItemId)

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


