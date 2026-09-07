package controller

import (
    PaymentDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PaymentDAO for database creation
//----------------------------------------------------------------------------
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Payment model
	//----------------------------------------------------------------------------
	data := model.Payment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Payment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Payment data access object to create
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.CreatePayment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PaymentDAO to find the relevant Payment
//----------------------------------------------------------------------------
func GetPayment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Payment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.GetPayment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PaymentDAO for database read of all Payments
//----------------------------------------------------------------------------
func GetAllPayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Payment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.GetAllPayment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PaymentDAO for database save
//----------------------------------------------------------------------------
func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Payment model
	//----------------------------------------------------------------------------
	var data = model.Payment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Payment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Payment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.UpdatePayment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PaymentDAO for database deletion
//----------------------------------------------------------------------------
func DeletePayment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Payment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PaymentDAO.DeletePayment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Invoice on a Payment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInvoiceToPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	invoiceId,_ := strconv.ParseUint( vars["invoiceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payment DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.AssignInvoiceToPayment(paymentId, invoiceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Invoice on a Payment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInvoiceFromPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payment DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.UnassignInvoiceFromPayment(paymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a BillingAccount on a Payment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBillingAccountToPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	billingAccountId,_ := strconv.ParseUint( vars["billingAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payment DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.AssignBillingAccountToPayment(paymentId, billingAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BillingAccount on a Payment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBillingAccountFromPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payment DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.UnassignBillingAccountFromPayment(paymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Policy on a Payment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payment DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.AssignPolicyToPayment(paymentId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a Payment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromPayment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Payment DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentDAO.UnassignPolicyFromPayment(paymentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


