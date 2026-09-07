package controller

import (
    QuoteLineItemDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to QuoteLineItemDAO for database creation
//----------------------------------------------------------------------------
func CreateQuoteLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QuoteLineItem model
	//----------------------------------------------------------------------------
	data := model.QuoteLineItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QuoteLineItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem data access object to create
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.CreateQuoteLineItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to QuoteLineItemDAO to find the relevant QuoteLineItem
//----------------------------------------------------------------------------
func GetQuoteLineItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QuoteLineItem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.GetQuoteLineItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to QuoteLineItemDAO for database read of all QuoteLineItems
//----------------------------------------------------------------------------
func GetAllQuoteLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.GetAllQuoteLineItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to QuoteLineItemDAO for database save
//----------------------------------------------------------------------------
func UpdateQuoteLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QuoteLineItem model
	//----------------------------------------------------------------------------
	var data = model.QuoteLineItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QuoteLineItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.UpdateQuoteLineItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to QuoteLineItemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteQuoteLineItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QuoteLineItem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := QuoteLineItemDAO.DeleteQuoteLineItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Quote on a QuoteLineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignQuoteToQuoteLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quoteId,_ := strconv.ParseUint( vars["quoteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.AssignQuoteToQuoteLineItem(quoteLineItemId, quoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Quote on a QuoteLineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignQuoteFromQuoteLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.UnassignQuoteFromQuoteLineItem(quoteLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a QuoteLineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToQuoteLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.AssignProductToQuoteLineItem(quoteLineItemId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a QuoteLineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromQuoteLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.UnassignProductFromQuoteLineItem(quoteLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PriceBookEntry on a QuoteLineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPriceBookEntryToQuoteLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBookEntryId,_ := strconv.ParseUint( vars["priceBookEntryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.AssignPriceBookEntryToQuoteLineItem(quoteLineItemId, priceBookEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PriceBookEntry on a QuoteLineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPriceBookEntryFromQuoteLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.UnassignPriceBookEntryFromQuoteLineItem(quoteLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a OpportunityLineItem on a QuoteLineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOpportunityLineItemToQuoteLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityLineItemId,_ := strconv.ParseUint( vars["opportunityLineItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.AssignOpportunityLineItemToQuoteLineItem(quoteLineItemId, opportunityLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a OpportunityLineItem on a QuoteLineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOpportunityLineItemFromQuoteLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QuoteLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteLineItemDAO.UnassignOpportunityLineItemFromQuoteLineItem(quoteLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


