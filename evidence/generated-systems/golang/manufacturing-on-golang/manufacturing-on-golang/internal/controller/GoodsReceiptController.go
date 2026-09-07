package controller

import (
    GoodsReceiptDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to GoodsReceiptDAO for database creation
//----------------------------------------------------------------------------
func CreateGoodsReceipt(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GoodsReceipt model
	//----------------------------------------------------------------------------
	data := model.GoodsReceipt{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GoodsReceipt model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt data access object to create
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.CreateGoodsReceipt( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to GoodsReceiptDAO to find the relevant GoodsReceipt
//----------------------------------------------------------------------------
func GetGoodsReceipt(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GoodsReceipt data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.GetGoodsReceipt(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to GoodsReceiptDAO for database read of all GoodsReceipts
//----------------------------------------------------------------------------
func GetAllGoodsReceipt(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt data access object to get all
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.GetAllGoodsReceipt()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to GoodsReceiptDAO for database save
//----------------------------------------------------------------------------
func UpdateGoodsReceipt(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GoodsReceipt model
	//----------------------------------------------------------------------------
	var data = model.GoodsReceipt{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GoodsReceipt model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.UpdateGoodsReceipt(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to GoodsReceiptDAO for database deletion
//----------------------------------------------------------------------------
func DeleteGoodsReceipt(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GoodsReceipt data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := GoodsReceiptDAO.DeleteGoodsReceipt(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PurchaseOrder on a GoodsReceipt
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPurchaseOrderToGoodsReceipt(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	purchaseOrderId,_ := strconv.ParseUint( vars["purchaseOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.AssignPurchaseOrderToGoodsReceipt(goodsReceiptId, purchaseOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PurchaseOrder on a GoodsReceipt
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPurchaseOrderFromGoodsReceipt( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.UnassignPurchaseOrderFromGoodsReceipt(goodsReceiptId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a GoodsReceipt
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToGoodsReceipt(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.AssignWarehouseToGoodsReceipt(goodsReceiptId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a GoodsReceipt
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromGoodsReceipt( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.UnassignWarehouseFromGoodsReceipt(goodsReceiptId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more linesIds as a Lines to a GoodsReceipt
	//----------------------------------------------------------------------------
func AddLinesToGoodsReceipt(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	goodsReceiptId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.AddLinesToGoodsReceipt(goodsReceiptId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more linesIds as a Lines from a GoodsReceipt
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLinesFromGoodsReceipt(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	goodsReceiptId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceipt DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptDAO.RemoveLinesFromGoodsReceipt(goodsReceiptId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
