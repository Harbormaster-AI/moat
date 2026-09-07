package controller

import (
    ProductDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProductDAO for database creation
//----------------------------------------------------------------------------
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Product model
	//----------------------------------------------------------------------------
	data := model.Product{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Product model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Product data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.CreateProduct( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProductDAO to find the relevant Product
//----------------------------------------------------------------------------
func GetProduct(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Product data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.GetProduct(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProductDAO for database read of all Products
//----------------------------------------------------------------------------
func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Product data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.GetAllProduct()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProductDAO for database save
//----------------------------------------------------------------------------
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Product model
	//----------------------------------------------------------------------------
	var data = model.Product{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Product model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Product data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.UpdateProduct(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProductDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Product data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProductDAO.DeleteProduct(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Product
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.AssignOrganizationToProduct(productId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Product
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromProduct( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.UnassignOrganizationFromProduct(productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more priceBookEntriesIds as a PriceBookEntries to a Product
	//----------------------------------------------------------------------------
func AddPriceBookEntriesToProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBookEntriesIds,_ := vars["priceBookEntriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.AddPriceBookEntriesToProduct(productId, priceBookEntriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more priceBookEntriesIds as a PriceBookEntries from a Product
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePriceBookEntriesFromProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBookEntriesIds,_ := vars["priceBookEntriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.RemovePriceBookEntriesFromProduct(productId, priceBookEntriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more opportunityLineItemsIds as a OpportunityLineItems to a Product
	//----------------------------------------------------------------------------
func AddOpportunityLineItemsToProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityLineItemsIds,_ := vars["opportunityLineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.AddOpportunityLineItemsToProduct(productId, opportunityLineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more opportunityLineItemsIds as a OpportunityLineItems from a Product
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOpportunityLineItemsFromProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityLineItemsIds,_ := vars["opportunityLineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.RemoveOpportunityLineItemsFromProduct(productId, opportunityLineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more quoteLineItemsIds as a QuoteLineItems to a Product
	//----------------------------------------------------------------------------
func AddQuoteLineItemsToProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quoteLineItemsIds,_ := vars["quoteLineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.AddQuoteLineItemsToProduct(productId, quoteLineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more quoteLineItemsIds as a QuoteLineItems from a Product
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQuoteLineItemsFromProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quoteLineItemsIds,_ := vars["quoteLineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.RemoveQuoteLineItemsFromProduct(productId, quoteLineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more orderItemsIds as a OrderItems to a Product
	//----------------------------------------------------------------------------
func AddOrderItemsToProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderItemsIds,_ := vars["orderItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.AddOrderItemsToProduct(productId, orderItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more orderItemsIds as a OrderItems from a Product
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrderItemsFromProduct(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderItemsIds,_ := vars["orderItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Product DAO
	//----------------------------------------------------------------------------
	requestResult := ProductDAO.RemoveOrderItemsFromProduct(productId, orderItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
