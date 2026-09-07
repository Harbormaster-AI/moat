package controller

import (
    RateDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RateDAO for database creation
//----------------------------------------------------------------------------
func CreateRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Rate model
	//----------------------------------------------------------------------------
	data := model.Rate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Rate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Rate data access object to create
	//----------------------------------------------------------------------------
	requestResult := RateDAO.CreateRate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RateDAO to find the relevant Rate
//----------------------------------------------------------------------------
func GetRate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Rate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RateDAO.GetRate(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RateDAO for database read of all Rates
//----------------------------------------------------------------------------
func GetAllRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Rate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RateDAO.GetAllRate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RateDAO for database save
//----------------------------------------------------------------------------
func UpdateRate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Rate model
	//----------------------------------------------------------------------------
	var data = model.Rate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Rate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Rate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RateDAO.UpdateRate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Rate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RateDAO.DeleteRate(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a RateCard on a Rate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRateCardToRate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	rateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	rateCardId,_ := strconv.ParseUint( vars["rateCardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Rate DAO
	//----------------------------------------------------------------------------
	requestResult := RateDAO.AssignRateCardToRate(rateId, rateCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RateCard on a Rate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRateCardFromRate( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	rateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Rate DAO
	//----------------------------------------------------------------------------
	requestResult := RateDAO.UnassignRateCardFromRate(rateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AdSlot on a Rate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdSlotToRate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	rateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adSlotId,_ := strconv.ParseUint( vars["adSlotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Rate DAO
	//----------------------------------------------------------------------------
	requestResult := RateDAO.AssignAdSlotToRate(rateId, adSlotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AdSlot on a Rate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdSlotFromRate( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	rateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Rate DAO
	//----------------------------------------------------------------------------
	requestResult := RateDAO.UnassignAdSlotFromRate(rateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


