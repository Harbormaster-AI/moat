package controller

import (
    RateCardDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RateCardDAO for database creation
//----------------------------------------------------------------------------
func CreateRateCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RateCard model
	//----------------------------------------------------------------------------
	data := model.RateCard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RateCard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RateCard data access object to create
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.CreateRateCard( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RateCardDAO to find the relevant RateCard
//----------------------------------------------------------------------------
func GetRateCard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RateCard data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.GetRateCard(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RateCardDAO for database read of all RateCards
//----------------------------------------------------------------------------
func GetAllRateCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RateCard data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.GetAllRateCard()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RateCardDAO for database save
//----------------------------------------------------------------------------
func UpdateRateCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RateCard model
	//----------------------------------------------------------------------------
	var data = model.RateCard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RateCard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RateCard data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.UpdateRateCard(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RateCardDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRateCard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RateCard data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RateCardDAO.DeleteRateCard(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Publisher on a RateCard
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPublisherToRateCard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	rateCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	publisherId,_ := strconv.ParseUint( vars["publisherId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RateCard DAO
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.AssignPublisherToRateCard(rateCardId, publisherId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Publisher on a RateCard
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPublisherFromRateCard( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	rateCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RateCard DAO
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.UnassignPublisherFromRateCard(rateCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more ratesIds as a Rates to a RateCard
	//----------------------------------------------------------------------------
func AddRatesToRateCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	rateCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ratesIds,_ := vars["ratesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RateCard DAO
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.AddRatesToRateCard(rateCardId, ratesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ratesIds as a Rates from a RateCard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRatesFromRateCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	rateCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ratesIds,_ := vars["ratesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RateCard DAO
	//----------------------------------------------------------------------------
	requestResult := RateCardDAO.RemoveRatesFromRateCard(rateCardId, ratesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
