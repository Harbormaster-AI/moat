package controller

import (
    OrderDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OrderDAO for database creation
//----------------------------------------------------------------------------
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Order model
	//----------------------------------------------------------------------------
	data := model.Order{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Order model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Order data access object to create
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.CreateOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OrderDAO to find the relevant Order
//----------------------------------------------------------------------------
func GetOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Order data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.GetOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OrderDAO for database read of all Orders
//----------------------------------------------------------------------------
func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Order data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.GetAllOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OrderDAO for database save
//----------------------------------------------------------------------------
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Order model
	//----------------------------------------------------------------------------
	var data = model.Order{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Order model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Order data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UpdateOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Order data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OrderDAO.DeleteOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Order
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AssignOrganizationToOrder(orderId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Order
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UnassignOrganizationFromOrder(orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a Order
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AssignAccountToOrder(orderId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a Order
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UnassignAccountFromOrder(orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Opportunity on a Order
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOpportunityToOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityId,_ := strconv.ParseUint( vars["opportunityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AssignOpportunityToOrder(orderId, opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Opportunity on a Order
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOpportunityFromOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UnassignOpportunityFromOrder(orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Quote on a Order
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignQuoteToOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quoteId,_ := strconv.ParseUint( vars["quoteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AssignQuoteToOrder(orderId, quoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Quote on a Order
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignQuoteFromOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UnassignQuoteFromOrder(orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a Order
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AssignOwnerToOrder(orderId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a Order
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UnassignOwnerFromOrder(orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Contract on a Order
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContractToOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractId,_ := strconv.ParseUint( vars["contractId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AssignContractToOrder(orderId, contractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contract on a Order
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContractFromOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UnassignContractFromOrder(orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PriceBook on a Order
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPriceBookToOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBookId,_ := strconv.ParseUint( vars["priceBookId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AssignPriceBookToOrder(orderId, priceBookId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PriceBook on a Order
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPriceBookFromOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.UnassignPriceBookFromOrder(orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more itemsIds as a Items to a Order
	//----------------------------------------------------------------------------
func AddItemsToOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.AddItemsToOrder(orderId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more itemsIds as a Items from a Order
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveItemsFromOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	orderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Order DAO
	//----------------------------------------------------------------------------
	requestResult := OrderDAO.RemoveItemsFromOrder(orderId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
