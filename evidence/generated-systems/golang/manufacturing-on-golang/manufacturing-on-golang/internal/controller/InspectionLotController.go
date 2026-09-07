package controller

import (
    InspectionLotDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InspectionLotDAO for database creation
//----------------------------------------------------------------------------
func CreateInspectionLot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionLot model
	//----------------------------------------------------------------------------
	data := model.InspectionLot{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionLot model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot data access object to create
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.CreateInspectionLot( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InspectionLotDAO to find the relevant InspectionLot
//----------------------------------------------------------------------------
func GetInspectionLot(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionLot data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.GetInspectionLot(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InspectionLotDAO for database read of all InspectionLots
//----------------------------------------------------------------------------
func GetAllInspectionLot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.GetAllInspectionLot()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InspectionLotDAO for database save
//----------------------------------------------------------------------------
func UpdateInspectionLot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InspectionLot model
	//----------------------------------------------------------------------------
	var data = model.InspectionLot{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InspectionLot model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.UpdateInspectionLot(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InspectionLotDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInspectionLot(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InspectionLot data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InspectionLotDAO.DeleteInspectionLot(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Item on a InspectionLot
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToInspectionLot(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.AssignItemToInspectionLot(inspectionLotId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a InspectionLot
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromInspectionLot( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.UnassignItemFromInspectionLot(inspectionLotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkOrder on a InspectionLot
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkOrderToInspectionLot(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workOrderId,_ := strconv.ParseUint( vars["workOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.AssignWorkOrderToInspectionLot(inspectionLotId, workOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkOrder on a InspectionLot
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkOrderFromInspectionLot( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.UnassignWorkOrderFromInspectionLot(inspectionLotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a GoodsReceipt on a InspectionLot
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignGoodsReceiptToInspectionLot(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	goodsReceiptId,_ := strconv.ParseUint( vars["goodsReceiptId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.AssignGoodsReceiptToInspectionLot(inspectionLotId, goodsReceiptId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a GoodsReceipt on a InspectionLot
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignGoodsReceiptFromInspectionLot( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.UnassignGoodsReceiptFromInspectionLot(inspectionLotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more resultsIds as a Results to a InspectionLot
	//----------------------------------------------------------------------------
func AddResultsToInspectionLot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	resultsIds,_ := vars["resultsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.AddResultsToInspectionLot(inspectionLotId, resultsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more resultsIds as a Results from a InspectionLot
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveResultsFromInspectionLot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inspectionLotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	resultsIds,_ := vars["resultsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InspectionLot DAO
	//----------------------------------------------------------------------------
	requestResult := InspectionLotDAO.RemoveResultsFromInspectionLot(inspectionLotId, resultsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
