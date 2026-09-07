package controller

import (
    PriceBookDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PriceBookDAO for database creation
//----------------------------------------------------------------------------
func CreatePriceBook(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PriceBook model
	//----------------------------------------------------------------------------
	data := model.PriceBook{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PriceBook model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook data access object to create
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.CreatePriceBook( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PriceBookDAO to find the relevant PriceBook
//----------------------------------------------------------------------------
func GetPriceBook(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PriceBook data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.GetPriceBook(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PriceBookDAO for database read of all PriceBooks
//----------------------------------------------------------------------------
func GetAllPriceBook(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PriceBook data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.GetAllPriceBook()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PriceBookDAO for database save
//----------------------------------------------------------------------------
func UpdatePriceBook(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PriceBook model
	//----------------------------------------------------------------------------
	var data = model.PriceBook{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PriceBook model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.UpdatePriceBook(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PriceBookDAO for database deletion
//----------------------------------------------------------------------------
func DeletePriceBook(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PriceBook data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PriceBookDAO.DeletePriceBook(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a PriceBook
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToPriceBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.AssignOrganizationToPriceBook(priceBookId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a PriceBook
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromPriceBook( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.UnassignOrganizationFromPriceBook(priceBookId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more entriesIds as a Entries to a PriceBook
	//----------------------------------------------------------------------------
func AddEntriesToPriceBook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	entriesIds,_ := vars["entriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.AddEntriesToPriceBook(priceBookId, entriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more entriesIds as a Entries from a PriceBook
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEntriesFromPriceBook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	entriesIds,_ := vars["entriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.RemoveEntriesFromPriceBook(priceBookId, entriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more quotesIds as a Quotes to a PriceBook
	//----------------------------------------------------------------------------
func AddQuotesToPriceBook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.AddQuotesToPriceBook(priceBookId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more quotesIds as a Quotes from a PriceBook
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQuotesFromPriceBook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.RemoveQuotesFromPriceBook(priceBookId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a PriceBook
	//----------------------------------------------------------------------------
func AddOrdersToPriceBook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.AddOrdersToPriceBook(priceBookId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a PriceBook
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromPriceBook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	priceBookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PriceBook DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookDAO.RemoveOrdersFromPriceBook(priceBookId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
