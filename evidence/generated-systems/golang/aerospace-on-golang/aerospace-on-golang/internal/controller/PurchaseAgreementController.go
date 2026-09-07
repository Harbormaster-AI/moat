package controller

import (
    PurchaseAgreementDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PurchaseAgreementDAO for database creation
//----------------------------------------------------------------------------
func CreatePurchaseAgreement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PurchaseAgreement model
	//----------------------------------------------------------------------------
	data := model.PurchaseAgreement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PurchaseAgreement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseAgreement data access object to create
	//----------------------------------------------------------------------------
	requestResult := PurchaseAgreementDAO.CreatePurchaseAgreement( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PurchaseAgreementDAO to find the relevant PurchaseAgreement
//----------------------------------------------------------------------------
func GetPurchaseAgreement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PurchaseAgreement data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PurchaseAgreementDAO.GetPurchaseAgreement(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PurchaseAgreementDAO for database read of all PurchaseAgreements
//----------------------------------------------------------------------------
func GetAllPurchaseAgreement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PurchaseAgreement data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PurchaseAgreementDAO.GetAllPurchaseAgreement()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PurchaseAgreementDAO for database save
//----------------------------------------------------------------------------
func UpdatePurchaseAgreement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PurchaseAgreement model
	//----------------------------------------------------------------------------
	var data = model.PurchaseAgreement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PurchaseAgreement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseAgreement data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PurchaseAgreementDAO.UpdatePurchaseAgreement(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PurchaseAgreementDAO for database deletion
//----------------------------------------------------------------------------
func DeletePurchaseAgreement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PurchaseAgreement data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PurchaseAgreementDAO.DeletePurchaseAgreement(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a AircraftOrder on a PurchaseAgreement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAircraftOrderToPurchaseAgreement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseAgreementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftOrderId,_ := strconv.ParseUint( vars["aircraftOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseAgreement DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseAgreementDAO.AssignAircraftOrderToPurchaseAgreement(purchaseAgreementId, aircraftOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AircraftOrder on a PurchaseAgreement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAircraftOrderFromPurchaseAgreement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	purchaseAgreementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PurchaseAgreement DAO
	//----------------------------------------------------------------------------
	requestResult := PurchaseAgreementDAO.UnassignAircraftOrderFromPurchaseAgreement(purchaseAgreementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


