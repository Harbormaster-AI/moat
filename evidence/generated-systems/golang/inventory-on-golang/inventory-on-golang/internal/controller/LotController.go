package controller

import (
    LotDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LotDAO for database creation
//----------------------------------------------------------------------------
func CreateLot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Lot model
	//----------------------------------------------------------------------------
	data := model.Lot{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Lot model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Lot data access object to create
	//----------------------------------------------------------------------------
	requestResult := LotDAO.CreateLot( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LotDAO to find the relevant Lot
//----------------------------------------------------------------------------
func GetLot(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Lot data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LotDAO.GetLot(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LotDAO for database read of all Lots
//----------------------------------------------------------------------------
func GetAllLot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Lot data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LotDAO.GetAllLot()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LotDAO for database save
//----------------------------------------------------------------------------
func UpdateLot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Lot model
	//----------------------------------------------------------------------------
	var data = model.Lot{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Lot model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Lot data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LotDAO.UpdateLot(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LotDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLot(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Lot data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LotDAO.DeleteLot(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a Lot
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToLot(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lot DAO
	//----------------------------------------------------------------------------
	requestResult := LotDAO.AssignSkuToLot(lotId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a Lot
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromLot( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Lot DAO
	//----------------------------------------------------------------------------
	requestResult := LotDAO.UnassignSkuFromLot(lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more inventoryItemsIds as a InventoryItems to a Lot
	//----------------------------------------------------------------------------
func AddInventoryItemsToLot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lot DAO
	//----------------------------------------------------------------------------
	requestResult := LotDAO.AddInventoryItemsToLot(lotId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventoryItemsIds as a InventoryItems from a Lot
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventoryItemsFromLot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Lot DAO
	//----------------------------------------------------------------------------
	requestResult := LotDAO.RemoveInventoryItemsFromLot(lotId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
