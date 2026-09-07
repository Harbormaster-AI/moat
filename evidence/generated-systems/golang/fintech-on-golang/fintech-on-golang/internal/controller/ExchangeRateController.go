package controller

import (
    ExchangeRateDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExchangeRateDAO for database creation
//----------------------------------------------------------------------------
func CreateExchangeRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExchangeRate model
	//----------------------------------------------------------------------------
	data := model.ExchangeRate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExchangeRate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.CreateExchangeRate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExchangeRateDAO to find the relevant ExchangeRate
//----------------------------------------------------------------------------
func GetExchangeRate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExchangeRate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.GetExchangeRate(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExchangeRateDAO for database read of all ExchangeRates
//----------------------------------------------------------------------------
func GetAllExchangeRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.GetAllExchangeRate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExchangeRateDAO for database save
//----------------------------------------------------------------------------
func UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExchangeRate model
	//----------------------------------------------------------------------------
	var data = model.ExchangeRate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExchangeRate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.UpdateExchangeRate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExchangeRateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExchangeRate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExchangeRate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExchangeRateDAO.DeleteExchangeRate(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more usedByQuotesIds as a UsedByQuotes to a ExchangeRate
	//----------------------------------------------------------------------------
func AddUsedByQuotesToExchangeRate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exchangeRateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usedByQuotesIds,_ := vars["usedByQuotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.AddUsedByQuotesToExchangeRate(exchangeRateId, usedByQuotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usedByQuotesIds as a UsedByQuotes from a ExchangeRate
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsedByQuotesFromExchangeRate(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	exchangeRateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usedByQuotesIds,_ := vars["usedByQuotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ExchangeRate DAO
	//----------------------------------------------------------------------------
	requestResult := ExchangeRateDAO.RemoveUsedByQuotesFromExchangeRate(exchangeRateId, usedByQuotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
