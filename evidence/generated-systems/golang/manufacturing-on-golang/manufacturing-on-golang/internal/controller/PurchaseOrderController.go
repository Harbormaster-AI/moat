package controller

import (
    PurchaseOrderDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PurchaseOrderDAO for database creation
//----------------------------------------------------------------------------
func CreatePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PurchaseOrder model
	//----------------------------------------------------------------------------
	data := model.PurchaseOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PurchaseOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.CreatePurchaseOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PurchaseOrderDAO to find the relevant PurchaseOrder
//----------------------------------------------------------------------------
func GetPurchaseOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PurchaseOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.GetPurchaseOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PurchaseOrderDAO for database read of all PurchaseOrders
//----------------------------------------------------------------------------
func GetAllPurchaseOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.GetAllPurchaseOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PurchaseOrderDAO for database save
//----------------------------------------------------------------------------
func UpdatePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PurchaseOrder model
	//----------------------------------------------------------------------------
	var data = model.PurchaseOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PurchaseOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.UpdatePurchaseOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PurchaseOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeletePurchaseOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PurchaseOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PurchaseOrderDAO.DeletePurchaseOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Supplier on a PurchaseOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSupplierToPurchaseOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	supplierId,_ := strconv.ParseUint( vars["supplierId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.AssignSupplierToPurchaseOrder(purchaseOrderId, supplierId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Supplier on a PurchaseOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSupplierFromPurchaseOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.UnassignSupplierFromPurchaseOrder(purchaseOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Plant on a PurchaseOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToPurchaseOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.AssignPlantToPurchaseOrder(purchaseOrderId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a PurchaseOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromPurchaseOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.UnassignPlantFromPurchaseOrder(purchaseOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more linesIds as a Lines to a PurchaseOrder
	//----------------------------------------------------------------------------
func AddLinesToPurchaseOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.AddLinesToPurchaseOrder(purchaseOrderId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more linesIds as a Lines from a PurchaseOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLinesFromPurchaseOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	linesIds,_ := vars["linesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.RemoveLinesFromPurchaseOrder(purchaseOrderId, linesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more goodsReceiptsIds as a GoodsReceipts to a PurchaseOrder
	//----------------------------------------------------------------------------
func AddGoodsReceiptsToPurchaseOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goodsReceiptsIds,_ := vars["goodsReceiptsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.AddGoodsReceiptsToPurchaseOrder(purchaseOrderId, goodsReceiptsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more goodsReceiptsIds as a GoodsReceipts from a PurchaseOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGoodsReceiptsFromPurchaseOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	purchaseOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goodsReceiptsIds,_ := vars["goodsReceiptsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseOrder DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseOrderDAO.RemoveGoodsReceiptsFromPurchaseOrder(purchaseOrderId, goodsReceiptsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
