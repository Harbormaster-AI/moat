package controller

import (
    PublisherDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PublisherDAO for database creation
//----------------------------------------------------------------------------
func CreatePublisher(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Publisher model
	//----------------------------------------------------------------------------
	data := model.Publisher{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Publisher model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Publisher data access object to create
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.CreatePublisher( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PublisherDAO to find the relevant Publisher
//----------------------------------------------------------------------------
func GetPublisher(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Publisher data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.GetPublisher(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PublisherDAO for database read of all Publishers
//----------------------------------------------------------------------------
func GetAllPublisher(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Publisher data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.GetAllPublisher()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PublisherDAO for database save
//----------------------------------------------------------------------------
func UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Publisher model
	//----------------------------------------------------------------------------
	var data = model.Publisher{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Publisher model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Publisher data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.UpdatePublisher(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PublisherDAO for database deletion
//----------------------------------------------------------------------------
func DeletePublisher(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Publisher data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PublisherDAO.DeletePublisher(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more inventorySourcesIds as a InventorySources to a Publisher
	//----------------------------------------------------------------------------
func AddInventorySourcesToPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventorySourcesIds,_ := vars["inventorySourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.AddInventorySourcesToPublisher(publisherId, inventorySourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventorySourcesIds as a InventorySources from a Publisher
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventorySourcesFromPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventorySourcesIds,_ := vars["inventorySourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.RemoveInventorySourcesFromPublisher(publisherId, inventorySourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dealsIds as a Deals to a Publisher
	//----------------------------------------------------------------------------
func AddDealsToPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dealsIds,_ := vars["dealsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.AddDealsToPublisher(publisherId, dealsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dealsIds as a Deals from a Publisher
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDealsFromPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dealsIds,_ := vars["dealsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.RemoveDealsFromPublisher(publisherId, dealsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more creativeApprovalsIds as a CreativeApprovals to a Publisher
	//----------------------------------------------------------------------------
func AddCreativeApprovalsToPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativeApprovalsIds,_ := vars["creativeApprovalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.AddCreativeApprovalsToPublisher(publisherId, creativeApprovalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more creativeApprovalsIds as a CreativeApprovals from a Publisher
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCreativeApprovalsFromPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativeApprovalsIds,_ := vars["creativeApprovalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.RemoveCreativeApprovalsFromPublisher(publisherId, creativeApprovalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more insertionOrdersIds as a InsertionOrders to a Publisher
	//----------------------------------------------------------------------------
func AddInsertionOrdersToPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insertionOrdersIds,_ := vars["insertionOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.AddInsertionOrdersToPublisher(publisherId, insertionOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more insertionOrdersIds as a InsertionOrders from a Publisher
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInsertionOrdersFromPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insertionOrdersIds,_ := vars["insertionOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.RemoveInsertionOrdersFromPublisher(publisherId, insertionOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more rateCardsIds as a RateCards to a Publisher
	//----------------------------------------------------------------------------
func AddRateCardsToPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	rateCardsIds,_ := vars["rateCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.AddRateCardsToPublisher(publisherId, rateCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more rateCardsIds as a RateCards from a Publisher
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRateCardsFromPublisher(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	publisherId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	rateCardsIds,_ := vars["rateCardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Publisher DAO
	//----------------------------------------------------------------------------
	requestResult := PublisherDAO.RemoveRateCardsFromPublisher(publisherId, rateCardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
