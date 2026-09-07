package controller

import (
    CycleCountEntryDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CycleCountEntryDAO for database creation
//----------------------------------------------------------------------------
func CreateCycleCountEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CycleCountEntry model
	//----------------------------------------------------------------------------
	data := model.CycleCountEntry{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CycleCountEntry model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry data access object to create
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.CreateCycleCountEntry( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CycleCountEntryDAO to find the relevant CycleCountEntry
//----------------------------------------------------------------------------
func GetCycleCountEntry(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CycleCountEntry data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.GetCycleCountEntry(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CycleCountEntryDAO for database read of all CycleCountEntrys
//----------------------------------------------------------------------------
func GetAllCycleCountEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.GetAllCycleCountEntry()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CycleCountEntryDAO for database save
//----------------------------------------------------------------------------
func UpdateCycleCountEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CycleCountEntry model
	//----------------------------------------------------------------------------
	var data = model.CycleCountEntry{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CycleCountEntry model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.UpdateCycleCountEntry(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CycleCountEntryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCycleCountEntry(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CycleCountEntry data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CycleCountEntryDAO.DeleteCycleCountEntry(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a CycleCount on a CycleCountEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCycleCountToCycleCountEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cycleCountId,_ := strconv.ParseUint( vars["cycleCountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.AssignCycleCountToCycleCountEntry(cycleCountEntryId, cycleCountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CycleCount on a CycleCountEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCycleCountFromCycleCountEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.UnassignCycleCountFromCycleCountEntry(cycleCountEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Sku on a CycleCountEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToCycleCountEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.AssignSkuToCycleCountEntry(cycleCountEntryId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a CycleCountEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromCycleCountEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.UnassignSkuFromCycleCountEntry(cycleCountEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a CycleCountEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToCycleCountEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.AssignLotToCycleCountEntry(cycleCountEntryId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a CycleCountEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromCycleCountEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.UnassignLotFromCycleCountEntry(cycleCountEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a CycleCountEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToCycleCountEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.AssignLocationToCycleCountEntry(cycleCountEntryId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a CycleCountEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromCycleCountEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.UnassignLocationFromCycleCountEntry(cycleCountEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a CycleCountEntry
	//----------------------------------------------------------------------------
func AddSerialNumbersToCycleCountEntry(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.AddSerialNumbersToCycleCountEntry(cycleCountEntryId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a CycleCountEntry
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromCycleCountEntry(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cycleCountEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CycleCountEntry DAO
	//----------------------------------------------------------------------------
	requestResult := CycleCountEntryDAO.RemoveSerialNumbersFromCycleCountEntry(cycleCountEntryId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
