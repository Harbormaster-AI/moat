package controller

import (
    StockAdjustmentLineDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to StockAdjustmentLineDAO for database creation
//----------------------------------------------------------------------------
func CreateStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StockAdjustmentLine model
	//----------------------------------------------------------------------------
	data := model.StockAdjustmentLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StockAdjustmentLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.CreateStockAdjustmentLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to StockAdjustmentLineDAO to find the relevant StockAdjustmentLine
//----------------------------------------------------------------------------
func GetStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StockAdjustmentLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.GetStockAdjustmentLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to StockAdjustmentLineDAO for database read of all StockAdjustmentLines
//----------------------------------------------------------------------------
func GetAllStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.GetAllStockAdjustmentLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to StockAdjustmentLineDAO for database save
//----------------------------------------------------------------------------
func UpdateStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty StockAdjustmentLine model
	//----------------------------------------------------------------------------
	var data = model.StockAdjustmentLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a StockAdjustmentLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.UpdateStockAdjustmentLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to StockAdjustmentLineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the StockAdjustmentLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := StockAdjustmentLineDAO.DeleteStockAdjustmentLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Adjustment on a StockAdjustmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdjustmentToStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adjustmentId,_ := strconv.ParseUint( vars["adjustmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.AssignAdjustmentToStockAdjustmentLine(stockAdjustmentLineId, adjustmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Adjustment on a StockAdjustmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdjustmentFromStockAdjustmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.UnassignAdjustmentFromStockAdjustmentLine(stockAdjustmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Sku on a StockAdjustmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.AssignSkuToStockAdjustmentLine(stockAdjustmentLineId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a StockAdjustmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromStockAdjustmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.UnassignSkuFromStockAdjustmentLine(stockAdjustmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a StockAdjustmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.AssignLotToStockAdjustmentLine(stockAdjustmentLineId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a StockAdjustmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromStockAdjustmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.UnassignLotFromStockAdjustmentLine(stockAdjustmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a StockAdjustmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToStockAdjustmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.AssignLocationToStockAdjustmentLine(stockAdjustmentLineId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a StockAdjustmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromStockAdjustmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.UnassignLocationFromStockAdjustmentLine(stockAdjustmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a StockAdjustmentLine
	//----------------------------------------------------------------------------
func AddSerialNumbersToStockAdjustmentLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.AddSerialNumbersToStockAdjustmentLine(stockAdjustmentLineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a StockAdjustmentLine
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromStockAdjustmentLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	stockAdjustmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the StockAdjustmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := StockAdjustmentLineDAO.RemoveSerialNumbersFromStockAdjustmentLine(stockAdjustmentLineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
