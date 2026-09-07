package controller

import (
    AppliedFeeDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AppliedFeeDAO for database creation
//----------------------------------------------------------------------------
func CreateAppliedFee(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AppliedFee model
	//----------------------------------------------------------------------------
	data := model.AppliedFee{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AppliedFee model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AppliedFee data access object to create
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.CreateAppliedFee( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AppliedFeeDAO to find the relevant AppliedFee
//----------------------------------------------------------------------------
func GetAppliedFee(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AppliedFee data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.GetAppliedFee(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AppliedFeeDAO for database read of all AppliedFees
//----------------------------------------------------------------------------
func GetAllAppliedFee(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AppliedFee data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.GetAllAppliedFee()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AppliedFeeDAO for database save
//----------------------------------------------------------------------------
func UpdateAppliedFee(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AppliedFee model
	//----------------------------------------------------------------------------
	var data = model.AppliedFee{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AppliedFee model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AppliedFee data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.UpdateAppliedFee(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AppliedFeeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAppliedFee(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AppliedFee data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AppliedFeeDAO.DeleteAppliedFee(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PaymentOrder on a AppliedFee
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPaymentOrderToAppliedFee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appliedFeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentOrderId,_ := strconv.ParseUint( vars["paymentOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AppliedFee DAO
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.AssignPaymentOrderToAppliedFee(appliedFeeId, paymentOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PaymentOrder on a AppliedFee
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPaymentOrderFromAppliedFee( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appliedFeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AppliedFee DAO
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.UnassignPaymentOrderFromAppliedFee(appliedFeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Transaction on a AppliedFee
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTransactionToAppliedFee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appliedFeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionId,_ := strconv.ParseUint( vars["transactionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AppliedFee DAO
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.AssignTransactionToAppliedFee(appliedFeeId, transactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Transaction on a AppliedFee
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTransactionFromAppliedFee( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appliedFeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AppliedFee DAO
	//----------------------------------------------------------------------------
	requestResult := AppliedFeeDAO.UnassignTransactionFromAppliedFee(appliedFeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


