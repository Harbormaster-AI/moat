package controller

import (
    PaymentMethodDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
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
	// assigns a BillingProfile on a PaymentMethod
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBillingProfileToPaymentMethod(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentMethodId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	billingProfileId,_ := strconv.ParseUint( vars["billingProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.AssignBillingProfileToPaymentMethod(paymentMethodId, billingProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BillingProfile on a PaymentMethod
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBillingProfileFromPaymentMethod( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentMethodId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentMethod DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentMethodDAO.UnassignBillingProfileFromPaymentMethod(paymentMethodId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


