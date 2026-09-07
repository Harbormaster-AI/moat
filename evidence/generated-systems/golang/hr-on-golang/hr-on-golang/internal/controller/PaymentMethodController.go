package controller

import (
    PaymentMethodDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PaymentMethodDAO for database creation
//----------------------------------------------------------------------------
func CreatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentMethod model
	//----------------------------------------------------------------------------
	data := model.PaymentMethod{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentMethod model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod data access object to create
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.CreatePaymentMethod( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PaymentMethodDAO to find the relevant PaymentMethod
//----------------------------------------------------------------------------
func GetPaymentMethod(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentMethod data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.GetPaymentMethod(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PaymentMethodDAO for database read of all PaymentMethods
//----------------------------------------------------------------------------
func GetAllPaymentMethod(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.GetAllPaymentMethod()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PaymentMethodDAO for database save
//----------------------------------------------------------------------------
func UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentMethod model
	//----------------------------------------------------------------------------
	var data = model.PaymentMethod{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentMethod model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.UpdatePaymentMethod(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PaymentMethodDAO for database deletion
//----------------------------------------------------------------------------
func DeletePaymentMethod(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentMethod data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PaymentMethodDAO.DeletePaymentMethod(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a PaymentMethod
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToPaymentMethod(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentMethodId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.AssignEmployeeToPaymentMethod(paymentMethodId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a PaymentMethod
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromPaymentMethod( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentMethodId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.UnassignEmployeeFromPaymentMethod(paymentMethodId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a BankAccount on a PaymentMethod
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBankAccountToPaymentMethod(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentMethodId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bankAccountId,_ := strconv.ParseUint( vars["bankAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.AssignBankAccountToPaymentMethod(paymentMethodId, bankAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BankAccount on a PaymentMethod
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBankAccountFromPaymentMethod( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentMethodId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.UnassignBankAccountFromPaymentMethod(paymentMethodId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


