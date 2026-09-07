package controller

import (
    SerialNumberDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SerialNumberDAO for database creation
//----------------------------------------------------------------------------
func CreateSerialNumber(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SerialNumber model
	//----------------------------------------------------------------------------
	data := model.SerialNumber{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SerialNumber model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber data access object to create
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.CreateSerialNumber( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SerialNumberDAO to find the relevant SerialNumber
//----------------------------------------------------------------------------
func GetSerialNumber(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SerialNumber data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.GetSerialNumber(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SerialNumberDAO for database read of all SerialNumbers
//----------------------------------------------------------------------------
func GetAllSerialNumber(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.GetAllSerialNumber()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SerialNumberDAO for database save
//----------------------------------------------------------------------------
func UpdateSerialNumber(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SerialNumber model
	//----------------------------------------------------------------------------
	var data = model.SerialNumber{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SerialNumber model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.UpdateSerialNumber(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SerialNumberDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSerialNumber(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SerialNumber data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SerialNumberDAO.DeleteSerialNumber(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a SerialNumber
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToSerialNumber(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	serialNumberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber DAO
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.AssignSkuToSerialNumber(serialNumberId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a SerialNumber
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromSerialNumber( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	serialNumberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber DAO
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.UnassignSkuFromSerialNumber(serialNumberId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CurrentInventoryItem on a SerialNumber
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCurrentInventoryItemToSerialNumber(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	serialNumberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	currentInventoryItemId,_ := strconv.ParseUint( vars["currentInventoryItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber DAO
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.AssignCurrentInventoryItemToSerialNumber(serialNumberId, currentInventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CurrentInventoryItem on a SerialNumber
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCurrentInventoryItemFromSerialNumber( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	serialNumberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber DAO
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.UnassignCurrentInventoryItemFromSerialNumber(serialNumberId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a SerialNumber
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToSerialNumber(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	serialNumberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber DAO
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.AssignLotToSerialNumber(serialNumberId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a SerialNumber
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromSerialNumber( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	serialNumberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SerialNumber DAO
	//----------------------------------------------------------------------------
	requestResult := SerialNumberDAO.UnassignLotFromSerialNumber(serialNumberId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


