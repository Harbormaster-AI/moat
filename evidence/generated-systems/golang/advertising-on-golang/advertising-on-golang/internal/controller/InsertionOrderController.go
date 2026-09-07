package controller

import (
    InsertionOrderDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InsertionOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateInsertionOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsertionOrder model
	//----------------------------------------------------------------------------
	data := model.InsertionOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsertionOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.CreateInsertionOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InsertionOrderDAO to find the relevant InsertionOrder
//----------------------------------------------------------------------------
func GetInsertionOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsertionOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.GetInsertionOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InsertionOrderDAO for database read of all InsertionOrders
//----------------------------------------------------------------------------
func GetAllInsertionOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.GetAllInsertionOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InsertionOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateInsertionOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsertionOrder model
	//----------------------------------------------------------------------------
	var data = model.InsertionOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsertionOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.UpdateInsertionOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InsertionOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInsertionOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsertionOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InsertionOrderDAO.DeleteInsertionOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Advertiser on a InsertionOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdvertiserToInsertionOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	advertiserId,_ := strconv.ParseUint( vars["advertiserId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.AssignAdvertiserToInsertionOrder(insertionOrderId, advertiserId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Advertiser on a InsertionOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdvertiserFromInsertionOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.UnassignAdvertiserFromInsertionOrder(insertionOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Agency on a InsertionOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAgencyToInsertionOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agencyId,_ := strconv.ParseUint( vars["agencyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.AssignAgencyToInsertionOrder(insertionOrderId, agencyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Agency on a InsertionOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAgencyFromInsertionOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.UnassignAgencyFromInsertionOrder(insertionOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Publisher on a InsertionOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPublisherToInsertionOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	publisherId,_ := strconv.ParseUint( vars["publisherId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.AssignPublisherToInsertionOrder(insertionOrderId, publisherId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Publisher on a InsertionOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPublisherFromInsertionOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.UnassignPublisherFromInsertionOrder(insertionOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more campaignsIds as a Campaigns to a InsertionOrder
	//----------------------------------------------------------------------------
func AddCampaignsToInsertionOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.AddCampaignsToInsertionOrder(insertionOrderId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more campaignsIds as a Campaigns from a InsertionOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCampaignsFromInsertionOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insertionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignsIds,_ := vars["campaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsertionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := InsertionOrderDAO.RemoveCampaignsFromInsertionOrder(insertionOrderId, campaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
