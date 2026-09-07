package controller

import (
    DirectDebitMandateDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DirectDebitMandateDAO for database creation
//----------------------------------------------------------------------------
func CreateDirectDebitMandate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DirectDebitMandate model
	//----------------------------------------------------------------------------
	data := model.DirectDebitMandate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DirectDebitMandate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DirectDebitMandate data access object to create
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.CreateDirectDebitMandate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DirectDebitMandateDAO to find the relevant DirectDebitMandate
//----------------------------------------------------------------------------
func GetDirectDebitMandate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DirectDebitMandate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.GetDirectDebitMandate(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DirectDebitMandateDAO for database read of all DirectDebitMandates
//----------------------------------------------------------------------------
func GetAllDirectDebitMandate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DirectDebitMandate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.GetAllDirectDebitMandate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DirectDebitMandateDAO for database save
//----------------------------------------------------------------------------
func UpdateDirectDebitMandate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DirectDebitMandate model
	//----------------------------------------------------------------------------
	var data = model.DirectDebitMandate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DirectDebitMandate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DirectDebitMandate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.UpdateDirectDebitMandate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DirectDebitMandateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDirectDebitMandate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DirectDebitMandate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DirectDebitMandateDAO.DeleteDirectDebitMandate(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Account on a DirectDebitMandate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToDirectDebitMandate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	directDebitMandateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DirectDebitMandate DAO
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.AssignAccountToDirectDebitMandate(directDebitMandateId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a DirectDebitMandate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromDirectDebitMandate( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	directDebitMandateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DirectDebitMandate DAO
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.UnassignAccountFromDirectDebitMandate(directDebitMandateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Creditor on a DirectDebitMandate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCreditorToDirectDebitMandate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	directDebitMandateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creditorId,_ := strconv.ParseUint( vars["creditorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DirectDebitMandate DAO
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.AssignCreditorToDirectDebitMandate(directDebitMandateId, creditorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Creditor on a DirectDebitMandate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCreditorFromDirectDebitMandate( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	directDebitMandateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DirectDebitMandate DAO
	//----------------------------------------------------------------------------
	requestResult := DirectDebitMandateDAO.UnassignCreditorFromDirectDebitMandate(directDebitMandateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


