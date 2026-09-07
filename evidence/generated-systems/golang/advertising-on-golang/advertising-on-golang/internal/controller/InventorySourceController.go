package controller

import (
    InventorySourceDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InventorySourceDAO for database creation
//----------------------------------------------------------------------------
func CreateInventorySource(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventorySource model
	//----------------------------------------------------------------------------
	data := model.InventorySource{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventorySource model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource data access object to create
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.CreateInventorySource( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InventorySourceDAO to find the relevant InventorySource
//----------------------------------------------------------------------------
func GetInventorySource(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventorySource data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.GetInventorySource(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InventorySourceDAO for database read of all InventorySources
//----------------------------------------------------------------------------
func GetAllInventorySource(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InventorySource data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.GetAllInventorySource()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InventorySourceDAO for database save
//----------------------------------------------------------------------------
func UpdateInventorySource(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InventorySource model
	//----------------------------------------------------------------------------
	var data = model.InventorySource{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InventorySource model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.UpdateInventorySource(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InventorySourceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInventorySource(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InventorySource data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InventorySourceDAO.DeleteInventorySource(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Publisher on a InventorySource
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPublisherToInventorySource(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventorySourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	publisherId,_ := strconv.ParseUint( vars["publisherId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource DAO
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.AssignPublisherToInventorySource(inventorySourceId, publisherId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Publisher on a InventorySource
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPublisherFromInventorySource( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inventorySourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource DAO
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.UnassignPublisherFromInventorySource(inventorySourceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more adSlotsIds as a AdSlots to a InventorySource
	//----------------------------------------------------------------------------
func AddAdSlotsToInventorySource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventorySourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adSlotsIds,_ := vars["adSlotsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource DAO
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.AddAdSlotsToInventorySource(inventorySourceId, adSlotsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more adSlotsIds as a AdSlots from a InventorySource
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAdSlotsFromInventorySource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventorySourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adSlotsIds,_ := vars["adSlotsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource DAO
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.RemoveAdSlotsFromInventorySource(inventorySourceId, adSlotsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dealsIds as a Deals to a InventorySource
	//----------------------------------------------------------------------------
func AddDealsToInventorySource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventorySourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dealsIds,_ := vars["dealsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource DAO
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.AddDealsToInventorySource(inventorySourceId, dealsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dealsIds as a Deals from a InventorySource
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDealsFromInventorySource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inventorySourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dealsIds,_ := vars["dealsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InventorySource DAO
	//----------------------------------------------------------------------------
	requestResult := InventorySourceDAO.RemoveDealsFromInventorySource(inventorySourceId, dealsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
