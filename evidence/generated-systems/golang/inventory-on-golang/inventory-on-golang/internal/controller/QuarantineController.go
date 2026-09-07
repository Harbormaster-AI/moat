package controller

import (
    QuarantineDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to QuarantineDAO for database creation
//----------------------------------------------------------------------------
func CreateQuarantine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Quarantine model
	//----------------------------------------------------------------------------
	data := model.Quarantine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Quarantine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine data access object to create
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.CreateQuarantine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to QuarantineDAO to find the relevant Quarantine
//----------------------------------------------------------------------------
func GetQuarantine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Quarantine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.GetQuarantine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to QuarantineDAO for database read of all Quarantines
//----------------------------------------------------------------------------
func GetAllQuarantine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Quarantine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.GetAllQuarantine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to QuarantineDAO for database save
//----------------------------------------------------------------------------
func UpdateQuarantine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Quarantine model
	//----------------------------------------------------------------------------
	var data = model.Quarantine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Quarantine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.UpdateQuarantine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to QuarantineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteQuarantine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Quarantine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := QuarantineDAO.DeleteQuarantine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a Quarantine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToQuarantine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.AssignWarehouseToQuarantine(quarantineId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a Quarantine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromQuarantine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.UnassignWarehouseFromQuarantine(quarantineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a Quarantine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToQuarantine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.AssignLotToQuarantine(quarantineId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a Quarantine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromQuarantine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.UnassignLotFromQuarantine(quarantineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more itemsIds as a Items to a Quarantine
	//----------------------------------------------------------------------------
func AddItemsToQuarantine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.AddItemsToQuarantine(quarantineId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more itemsIds as a Items from a Quarantine
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveItemsFromQuarantine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.RemoveItemsFromQuarantine(quarantineId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a Quarantine
	//----------------------------------------------------------------------------
func AddSerialNumbersToQuarantine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.AddSerialNumbersToQuarantine(quarantineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a Quarantine
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromQuarantine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	quarantineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Quarantine DAO
	//----------------------------------------------------------------------------
	requestResult := QuarantineDAO.RemoveSerialNumbersFromQuarantine(quarantineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
