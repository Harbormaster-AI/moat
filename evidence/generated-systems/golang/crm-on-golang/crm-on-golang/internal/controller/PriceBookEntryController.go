package controller

import (
    PriceBookEntryDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PriceBookEntryDAO for database creation
//----------------------------------------------------------------------------
func CreatePriceBookEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PriceBookEntry model
	//----------------------------------------------------------------------------
	data := model.PriceBookEntry{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PriceBookEntry model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBookEntry data access object to create
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.CreatePriceBookEntry( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PriceBookEntryDAO to find the relevant PriceBookEntry
//----------------------------------------------------------------------------
func GetPriceBookEntry(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PriceBookEntry data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.GetPriceBookEntry(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PriceBookEntryDAO for database read of all PriceBookEntrys
//----------------------------------------------------------------------------
func GetAllPriceBookEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PriceBookEntry data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.GetAllPriceBookEntry()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PriceBookEntryDAO for database save
//----------------------------------------------------------------------------
func UpdatePriceBookEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PriceBookEntry model
	//----------------------------------------------------------------------------
	var data = model.PriceBookEntry{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PriceBookEntry model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBookEntry data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.UpdatePriceBookEntry(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PriceBookEntryDAO for database deletion
//----------------------------------------------------------------------------
func DeletePriceBookEntry(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PriceBookEntry data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PriceBookEntryDAO.DeletePriceBookEntry(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PriceBook on a PriceBookEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPriceBookToPriceBookEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	priceBookEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	priceBookId,_ := strconv.ParseUint( vars["priceBookId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBookEntry DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.AssignPriceBookToPriceBookEntry(priceBookEntryId, priceBookId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PriceBook on a PriceBookEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPriceBookFromPriceBookEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	priceBookEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBookEntry DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.UnassignPriceBookFromPriceBookEntry(priceBookEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a PriceBookEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToPriceBookEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	priceBookEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBookEntry DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.AssignProductToPriceBookEntry(priceBookEntryId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a PriceBookEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromPriceBookEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	priceBookEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PriceBookEntry DAO
	//----------------------------------------------------------------------------
	requestResult := PriceBookEntryDAO.UnassignProductFromPriceBookEntry(priceBookEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


