package controller

import (
    FXQuoteDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FXQuoteDAO for database creation
//----------------------------------------------------------------------------
func CreateFXQuote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FXQuote model
	//----------------------------------------------------------------------------
	data := model.FXQuote{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FXQuote model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FXQuote data access object to create
	//----------------------------------------------------------------------------
	requestResult := FXQuoteDAO.CreateFXQuote( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FXQuoteDAO to find the relevant FXQuote
//----------------------------------------------------------------------------
func GetFXQuote(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FXQuote data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FXQuoteDAO.GetFXQuote(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FXQuoteDAO for database read of all FXQuotes
//----------------------------------------------------------------------------
func GetAllFXQuote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FXQuote data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FXQuoteDAO.GetAllFXQuote()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FXQuoteDAO for database save
//----------------------------------------------------------------------------
func UpdateFXQuote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FXQuote model
	//----------------------------------------------------------------------------
	var data = model.FXQuote{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FXQuote model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FXQuote data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FXQuoteDAO.UpdateFXQuote(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FXQuoteDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFXQuote(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FXQuote data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FXQuoteDAO.DeleteFXQuote(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a RequestedBy on a FXQuote
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRequestedByToFXQuote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXQuoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	requestedById,_ := strconv.ParseUint( vars["requestedById"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXQuote DAO
	//----------------------------------------------------------------------------
	requestResult := FXQuoteDAO.AssignRequestedByToFXQuote(fXQuoteId, requestedById)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RequestedBy on a FXQuote
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRequestedByFromFXQuote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXQuoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXQuote DAO
	//----------------------------------------------------------------------------
	requestResult := FXQuoteDAO.UnassignRequestedByFromFXQuote(fXQuoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


