package controller

import (
    AircraftOrderDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraftOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftOrder model
	//----------------------------------------------------------------------------
	data := model.AircraftOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.CreateAircraftOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftOrderDAO to find the relevant AircraftOrder
//----------------------------------------------------------------------------
func GetAircraftOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.GetAircraftOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftOrderDAO for database read of all AircraftOrders
//----------------------------------------------------------------------------
func GetAllAircraftOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.GetAllAircraftOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraftOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftOrder model
	//----------------------------------------------------------------------------
	var data = model.AircraftOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.UpdateAircraftOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraftOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftOrderDAO.DeleteAircraftOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Operator on a AircraftOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOperatorToAircraftOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operatorId,_ := strconv.ParseUint( vars["operatorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.AssignOperatorToAircraftOrder(aircraftOrderId, operatorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Operator on a AircraftOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOperatorFromAircraftOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.UnassignOperatorFromAircraftOrder(aircraftOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Variant on a AircraftOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignVariantToAircraftOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantId,_ := strconv.ParseUint( vars["variantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.AssignVariantToAircraftOrder(aircraftOrderId, variantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Variant on a AircraftOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignVariantFromAircraftOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.UnassignVariantFromAircraftOrder(aircraftOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Quote on a AircraftOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignQuoteToAircraftOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quoteId,_ := strconv.ParseUint( vars["quoteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.AssignQuoteToAircraftOrder(aircraftOrderId, quoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Quote on a AircraftOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignQuoteFromAircraftOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.UnassignQuoteFromAircraftOrder(aircraftOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PurchaseAgreement on a AircraftOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPurchaseAgreementToAircraftOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	purchaseAgreementId,_ := strconv.ParseUint( vars["purchaseAgreementId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.AssignPurchaseAgreementToAircraftOrder(aircraftOrderId, purchaseAgreementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PurchaseAgreement on a AircraftOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPurchaseAgreementFromAircraftOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOrder DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOrderDAO.UnassignPurchaseAgreementFromAircraftOrder(aircraftOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


