package controller

import (
    ImagingOrderDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ImagingOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateImagingOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ImagingOrder model
	//----------------------------------------------------------------------------
	data := model.ImagingOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ImagingOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.CreateImagingOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ImagingOrderDAO to find the relevant ImagingOrder
//----------------------------------------------------------------------------
func GetImagingOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ImagingOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.GetImagingOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ImagingOrderDAO for database read of all ImagingOrders
//----------------------------------------------------------------------------
func GetAllImagingOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.GetAllImagingOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ImagingOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateImagingOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ImagingOrder model
	//----------------------------------------------------------------------------
	var data = model.ImagingOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ImagingOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.UpdateImagingOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ImagingOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteImagingOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ImagingOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ImagingOrderDAO.DeleteImagingOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Order on a ImagingOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrderToImagingOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderId,_ := strconv.ParseUint( vars["orderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.AssignOrderToImagingOrder(imagingOrderId, orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Order on a ImagingOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrderFromImagingOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.UnassignOrderFromImagingOrder(imagingOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ImagingCenter on a ImagingOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignImagingCenterToImagingOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingCenterId,_ := strconv.ParseUint( vars["imagingCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.AssignImagingCenterToImagingOrder(imagingOrderId, imagingCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ImagingCenter on a ImagingOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignImagingCenterFromImagingOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.UnassignImagingCenterFromImagingOrder(imagingOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a ImagingOrder
	//----------------------------------------------------------------------------
func AddReportsToImagingOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	imagingOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.AddReportsToImagingOrder(imagingOrderId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a ImagingOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromImagingOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	imagingOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ImagingOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingOrderDAO.RemoveReportsFromImagingOrder(imagingOrderId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
