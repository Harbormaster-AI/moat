package controller

import (
    CycleCountDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CycleCountDAO for database creation
//----------------------------------------------------------------------------
func CreateCycleCount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CycleCount model
	//----------------------------------------------------------------------------
	data := model.CycleCount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CycleCount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount data access object to create
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.CreateCycleCount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CycleCountDAO to find the relevant CycleCount
//----------------------------------------------------------------------------
func GetCycleCount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CycleCount data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.GetCycleCount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CycleCountDAO for database read of all CycleCounts
//----------------------------------------------------------------------------
func GetAllCycleCount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CycleCount data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.GetAllCycleCount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CycleCountDAO for database save
//----------------------------------------------------------------------------
func UpdateCycleCount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CycleCount model
	//----------------------------------------------------------------------------
	var data = model.CycleCount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CycleCount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.UpdateCycleCount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CycleCountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCycleCount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CycleCount data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CycleCountDAO.DeleteCycleCount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a CycleCount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToCycleCount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.AssignWarehouseToCycleCount(cycleCountId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a CycleCount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromCycleCount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.UnassignWarehouseFromCycleCount(cycleCountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more locationsIds as a Locations to a CycleCount
	//----------------------------------------------------------------------------
func AddLocationsToCycleCount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationsIds,_ := vars["locationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.AddLocationsToCycleCount(cycleCountId, locationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more locationsIds as a Locations from a CycleCount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLocationsFromCycleCount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationsIds,_ := vars["locationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.RemoveLocationsFromCycleCount(cycleCountId, locationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more entriesIds as a Entries to a CycleCount
	//----------------------------------------------------------------------------
func AddEntriesToCycleCount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	entriesIds,_ := vars["entriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.AddEntriesToCycleCount(cycleCountId, entriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more entriesIds as a Entries from a CycleCount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEntriesFromCycleCount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	entriesIds,_ := vars["entriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.RemoveEntriesFromCycleCount(cycleCountId, entriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a CycleCount
	//----------------------------------------------------------------------------
func AddTransactionsToCycleCount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.AddTransactionsToCycleCount(cycleCountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a CycleCount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromCycleCount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCount DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountDAO.RemoveTransactionsFromCycleCount(cycleCountId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
