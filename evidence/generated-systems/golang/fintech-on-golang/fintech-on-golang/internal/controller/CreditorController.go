package controller

import (
    CreditorDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CreditorDAO for database creation
//----------------------------------------------------------------------------
func CreateCreditor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Creditor model
	//----------------------------------------------------------------------------
	data := model.Creditor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Creditor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Creditor data access object to create
	//----------------------------------------------------------------------------
	requestResult := CreditorDAO.CreateCreditor( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CreditorDAO to find the relevant Creditor
//----------------------------------------------------------------------------
func GetCreditor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Creditor data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreditorDAO.GetCreditor(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CreditorDAO for database read of all Creditors
//----------------------------------------------------------------------------
func GetAllCreditor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Creditor data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CreditorDAO.GetAllCreditor()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CreditorDAO for database save
//----------------------------------------------------------------------------
func UpdateCreditor(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Creditor model
	//----------------------------------------------------------------------------
	var data = model.Creditor{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Creditor model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Creditor data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreditorDAO.UpdateCreditor(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CreditorDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCreditor(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Creditor data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CreditorDAO.DeleteCreditor(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more mandatesIds as a Mandates to a Creditor
	//----------------------------------------------------------------------------
func AddMandatesToCreditor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creditorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mandatesIds,_ := vars["mandatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Creditor DAO
	//----------------------------------------------------------------------------
	requestResult := CreditorDAO.AddMandatesToCreditor(creditorId, mandatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more mandatesIds as a Mandates from a Creditor
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMandatesFromCreditor(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creditorId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	mandatesIds,_ := vars["mandatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Creditor DAO
	//----------------------------------------------------------------------------
	requestResult := CreditorDAO.RemoveMandatesFromCreditor(creditorId, mandatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
