package controller

import (
    CardTokenizationDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CardTokenizationDAO for database creation
//----------------------------------------------------------------------------
func CreateCardTokenization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CardTokenization model
	//----------------------------------------------------------------------------
	data := model.CardTokenization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CardTokenization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CardTokenization data access object to create
	//----------------------------------------------------------------------------
	requestResult := CardTokenizationDAO.CreateCardTokenization( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CardTokenizationDAO to find the relevant CardTokenization
//----------------------------------------------------------------------------
func GetCardTokenization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CardTokenization data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CardTokenizationDAO.GetCardTokenization(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CardTokenizationDAO for database read of all CardTokenizations
//----------------------------------------------------------------------------
func GetAllCardTokenization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CardTokenization data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CardTokenizationDAO.GetAllCardTokenization()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CardTokenizationDAO for database save
//----------------------------------------------------------------------------
func UpdateCardTokenization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CardTokenization model
	//----------------------------------------------------------------------------
	var data = model.CardTokenization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CardTokenization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CardTokenization data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CardTokenizationDAO.UpdateCardTokenization(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CardTokenizationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCardTokenization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CardTokenization data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CardTokenizationDAO.DeleteCardTokenization(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Card on a CardTokenization
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCardToCardTokenization(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cardTokenizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cardId,_ := strconv.ParseUint( vars["cardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CardTokenization DAO
	//----------------------------------------------------------------------------
	requestResult := CardTokenizationDAO.AssignCardToCardTokenization(cardTokenizationId, cardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Card on a CardTokenization
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCardFromCardTokenization( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cardTokenizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CardTokenization DAO
	//----------------------------------------------------------------------------
	requestResult := CardTokenizationDAO.UnassignCardFromCardTokenization(cardTokenizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


