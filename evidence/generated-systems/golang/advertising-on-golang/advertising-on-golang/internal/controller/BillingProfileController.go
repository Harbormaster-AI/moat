package controller

import (
    BillingProfileDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BillingProfileDAO for database creation
//----------------------------------------------------------------------------
func CreateBillingProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BillingProfile model
	//----------------------------------------------------------------------------
	data := model.BillingProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BillingProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile data access object to create
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.CreateBillingProfile( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BillingProfileDAO to find the relevant BillingProfile
//----------------------------------------------------------------------------
func GetBillingProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BillingProfile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.GetBillingProfile(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BillingProfileDAO for database read of all BillingProfiles
//----------------------------------------------------------------------------
func GetAllBillingProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.GetAllBillingProfile()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BillingProfileDAO for database save
//----------------------------------------------------------------------------
func UpdateBillingProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BillingProfile model
	//----------------------------------------------------------------------------
	var data = model.BillingProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BillingProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.UpdateBillingProfile(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BillingProfileDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBillingProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BillingProfile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BillingProfileDAO.DeleteBillingProfile(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Advertiser on a BillingProfile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdvertiserToBillingProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	billingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	advertiserId,_ := strconv.ParseUint( vars["advertiserId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.AssignAdvertiserToBillingProfile(billingProfileId, advertiserId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Advertiser on a BillingProfile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdvertiserFromBillingProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	billingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.UnassignAdvertiserFromBillingProfile(billingProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more paymentMethodsIds as a PaymentMethods to a BillingProfile
	//----------------------------------------------------------------------------
func AddPaymentMethodsToBillingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentMethodsIds,_ := vars["paymentMethodsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.AddPaymentMethodsToBillingProfile(billingProfileId, paymentMethodsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentMethodsIds as a PaymentMethods from a BillingProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentMethodsFromBillingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentMethodsIds,_ := vars["paymentMethodsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.RemovePaymentMethodsFromBillingProfile(billingProfileId, paymentMethodsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more adAccountsIds as a AdAccounts to a BillingProfile
	//----------------------------------------------------------------------------
func AddAdAccountsToBillingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.AddAdAccountsToBillingProfile(billingProfileId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more adAccountsIds as a AdAccounts from a BillingProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAdAccountsFromBillingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	billingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BillingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := BillingProfileDAO.RemoveAdAccountsFromBillingProfile(billingProfileId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
