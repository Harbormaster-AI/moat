package controller

import (
    PaymentCardDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PaymentCardDAO for database creation
//----------------------------------------------------------------------------
func CreatePaymentCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentCard model
	//----------------------------------------------------------------------------
	data := model.PaymentCard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentCard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object to create
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.CreatePaymentCard( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PaymentCardDAO to find the relevant PaymentCard
//----------------------------------------------------------------------------
func GetPaymentCard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentCard data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.GetPaymentCard(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PaymentCardDAO for database read of all PaymentCards
//----------------------------------------------------------------------------
func GetAllPaymentCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.GetAllPaymentCard()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PaymentCardDAO for database save
//----------------------------------------------------------------------------
func UpdatePaymentCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PaymentCard model
	//----------------------------------------------------------------------------
	var data = model.PaymentCard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PaymentCard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.UpdatePaymentCard(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PaymentCardDAO for database deletion
//----------------------------------------------------------------------------
func DeletePaymentCard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PaymentCard data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PaymentCardDAO.DeletePaymentCard(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a PaymentCard
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToPaymentCard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AssignCustomerToPaymentCard(paymentCardId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a PaymentCard
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromPaymentCard( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.UnassignCustomerFromPaymentCard(paymentCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Account on a PaymentCard
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAccountToPaymentCard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountId,_ := strconv.ParseUint( vars["accountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AssignAccountToPaymentCard(paymentCardId, accountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Account on a PaymentCard
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAccountFromPaymentCard( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.UnassignAccountFromPaymentCard(paymentCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more tokenizationsIds as a Tokenizations to a PaymentCard
	//----------------------------------------------------------------------------
func AddTokenizationsToPaymentCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tokenizationsIds,_ := vars["tokenizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AddTokenizationsToPaymentCard(paymentCardId, tokenizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tokenizationsIds as a Tokenizations from a PaymentCard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTokenizationsFromPaymentCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tokenizationsIds,_ := vars["tokenizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.RemoveTokenizationsFromPaymentCard(paymentCardId, tokenizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more disputesIds as a Disputes to a PaymentCard
	//----------------------------------------------------------------------------
func AddDisputesToPaymentCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.AddDisputesToPaymentCard(paymentCardId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more disputesIds as a Disputes from a PaymentCard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDisputesFromPaymentCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	paymentCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	disputesIds,_ := vars["disputesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PaymentCard DAO
	//----------------------------------------------------------------------------
	requestResult := PaymentCardDAO.RemoveDisputesFromPaymentCard(paymentCardId, disputesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
