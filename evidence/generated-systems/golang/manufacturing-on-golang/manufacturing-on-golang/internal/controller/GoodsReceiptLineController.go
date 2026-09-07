package controller

import (
    GoodsReceiptLineDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to GoodsReceiptLineDAO for database creation
//----------------------------------------------------------------------------
func CreateGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GoodsReceiptLine model
	//----------------------------------------------------------------------------
	data := model.GoodsReceiptLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GoodsReceiptLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.CreateGoodsReceiptLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to GoodsReceiptLineDAO to find the relevant GoodsReceiptLine
//----------------------------------------------------------------------------
func GetGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GoodsReceiptLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.GetGoodsReceiptLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to GoodsReceiptLineDAO for database read of all GoodsReceiptLines
//----------------------------------------------------------------------------
func GetAllGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.GetAllGoodsReceiptLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to GoodsReceiptLineDAO for database save
//----------------------------------------------------------------------------
func UpdateGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GoodsReceiptLine model
	//----------------------------------------------------------------------------
	var data = model.GoodsReceiptLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GoodsReceiptLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.UpdateGoodsReceiptLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to GoodsReceiptLineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GoodsReceiptLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := GoodsReceiptLineDAO.DeleteGoodsReceiptLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a GoodsReceipt on a GoodsReceiptLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignGoodsReceiptToGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goodsReceiptId,_ := strconv.ParseUint( vars["goodsReceiptId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.AssignGoodsReceiptToGoodsReceiptLine(goodsReceiptLineId, goodsReceiptId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a GoodsReceipt on a GoodsReceiptLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignGoodsReceiptFromGoodsReceiptLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.UnassignGoodsReceiptFromGoodsReceiptLine(goodsReceiptLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Item on a GoodsReceiptLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.AssignItemToGoodsReceiptLine(goodsReceiptLineId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a GoodsReceiptLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromGoodsReceiptLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.UnassignItemFromGoodsReceiptLine(goodsReceiptLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InventoryTransaction on a GoodsReceiptLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInventoryTransactionToGoodsReceiptLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryTransactionId,_ := strconv.ParseUint( vars["inventoryTransactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.AssignInventoryTransactionToGoodsReceiptLine(goodsReceiptLineId, inventoryTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InventoryTransaction on a GoodsReceiptLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInventoryTransactionFromGoodsReceiptLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goodsReceiptLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GoodsReceiptLine DAO
	//----------------------------------------------------------------------------
	requestResult := GoodsReceiptLineDAO.UnassignInventoryTransactionFromGoodsReceiptLine(goodsReceiptLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


