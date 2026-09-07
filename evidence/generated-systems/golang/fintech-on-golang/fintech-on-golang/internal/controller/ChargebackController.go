package controller

import (
    ChargebackDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ChargebackDAO for database creation
//----------------------------------------------------------------------------
func CreateChargeback(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Chargeback model
	//----------------------------------------------------------------------------
	data := model.Chargeback{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Chargeback model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Chargeback data access object to create
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.CreateChargeback( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ChargebackDAO to find the relevant Chargeback
//----------------------------------------------------------------------------
func GetChargeback(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Chargeback data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.GetChargeback(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ChargebackDAO for database read of all Chargebacks
//----------------------------------------------------------------------------
func GetAllChargeback(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Chargeback data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.GetAllChargeback()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ChargebackDAO for database save
//----------------------------------------------------------------------------
func UpdateChargeback(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Chargeback model
	//----------------------------------------------------------------------------
	var data = model.Chargeback{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Chargeback model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Chargeback data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.UpdateChargeback(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ChargebackDAO for database deletion
//----------------------------------------------------------------------------
func DeleteChargeback(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Chargeback data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ChargebackDAO.DeleteChargeback(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Dispute on a Chargeback
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDisputeToChargeback(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	chargebackId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputeId,_ := strconv.ParseUint( vars["disputeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Chargeback DAO
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.AssignDisputeToChargeback(chargebackId, disputeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Dispute on a Chargeback
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDisputeFromChargeback( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	chargebackId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Chargeback DAO
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.UnassignDisputeFromChargeback(chargebackId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Transaction on a Chargeback
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransactionToChargeback(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	chargebackId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionId,_ := strconv.ParseUint( vars["transactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Chargeback DAO
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.AssignTransactionToChargeback(chargebackId, transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Transaction on a Chargeback
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransactionFromChargeback( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	chargebackId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Chargeback DAO
	//----------------------------------------------------------------------------
	requestResult := ChargebackDAO.UnassignTransactionFromChargeback(chargebackId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


