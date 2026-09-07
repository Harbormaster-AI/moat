package controller

import (
    PaymentContractDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PaymentContractDAO for database creation
//----------------------------------------------------------------------------
func CreatePaymentContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentContract model
	//----------------------------------------------------------------------------
	data := model.PaymentContract{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentContract model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentContract data access object to create
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.CreatePaymentContract( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PaymentContractDAO to find the relevant PaymentContract
//----------------------------------------------------------------------------
func GetPaymentContract(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentContract data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.GetPaymentContract(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PaymentContractDAO for database read of all PaymentContracts
//----------------------------------------------------------------------------
func GetAllPaymentContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PaymentContract data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.GetAllPaymentContract()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PaymentContractDAO for database save
//----------------------------------------------------------------------------
func UpdatePaymentContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentContract model
	//----------------------------------------------------------------------------
	var data = model.PaymentContract{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentContract model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentContract data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.UpdatePaymentContract(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PaymentContractDAO for database deletion
//----------------------------------------------------------------------------
func DeletePaymentContract(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentContract data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PaymentContractDAO.DeletePaymentContract(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Merchant on a PaymentContract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMerchantToPaymentContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	merchantId,_ := strconv.ParseUint( vars["merchantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.AssignMerchantToPaymentContract(paymentContractId, merchantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Merchant on a PaymentContract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMerchantFromPaymentContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.UnassignMerchantFromPaymentContract(paymentContractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Acquirer on a PaymentContract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAcquirerToPaymentContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	acquirerId,_ := strconv.ParseUint( vars["acquirerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.AssignAcquirerToPaymentContract(paymentContractId, acquirerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Acquirer on a PaymentContract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAcquirerFromPaymentContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentContractDAO.UnassignAcquirerFromPaymentContract(paymentContractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


