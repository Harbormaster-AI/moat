package controller

import (
    OpportunityLineItemDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OpportunityLineItemDAO for database creation
//----------------------------------------------------------------------------
func CreateOpportunityLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OpportunityLineItem model
	//----------------------------------------------------------------------------
	data := model.OpportunityLineItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OpportunityLineItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem data access object to create
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.CreateOpportunityLineItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OpportunityLineItemDAO to find the relevant OpportunityLineItem
//----------------------------------------------------------------------------
func GetOpportunityLineItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OpportunityLineItem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.GetOpportunityLineItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OpportunityLineItemDAO for database read of all OpportunityLineItems
//----------------------------------------------------------------------------
func GetAllOpportunityLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.GetAllOpportunityLineItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OpportunityLineItemDAO for database save
//----------------------------------------------------------------------------
func UpdateOpportunityLineItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OpportunityLineItem model
	//----------------------------------------------------------------------------
	var data = model.OpportunityLineItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OpportunityLineItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.UpdateOpportunityLineItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OpportunityLineItemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOpportunityLineItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OpportunityLineItem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OpportunityLineItemDAO.DeleteOpportunityLineItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Opportunity on a OpportunityLineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOpportunityToOpportunityLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityId,_ := strconv.ParseUint( vars["opportunityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.AssignOpportunityToOpportunityLineItem(opportunityLineItemId, opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Opportunity on a OpportunityLineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOpportunityFromOpportunityLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.UnassignOpportunityFromOpportunityLineItem(opportunityLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a OpportunityLineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToOpportunityLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.AssignProductToOpportunityLineItem(opportunityLineItemId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a OpportunityLineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromOpportunityLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.UnassignProductFromOpportunityLineItem(opportunityLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PriceBookEntry on a OpportunityLineItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPriceBookEntryToOpportunityLineItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBookEntryId,_ := strconv.ParseUint( vars["priceBookEntryId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.AssignPriceBookEntryToOpportunityLineItem(opportunityLineItemId, priceBookEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PriceBookEntry on a OpportunityLineItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPriceBookEntryFromOpportunityLineItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityLineItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityLineItem DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityLineItemDAO.UnassignPriceBookEntryFromOpportunityLineItem(opportunityLineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


