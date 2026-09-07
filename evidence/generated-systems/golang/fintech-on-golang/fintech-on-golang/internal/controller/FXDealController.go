package controller

import (
    FXDealDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FXDealDAO for database creation
//----------------------------------------------------------------------------
func CreateFXDeal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FXDeal model
	//----------------------------------------------------------------------------
	data := model.FXDeal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FXDeal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FXDeal data access object to create
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.CreateFXDeal( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FXDealDAO to find the relevant FXDeal
//----------------------------------------------------------------------------
func GetFXDeal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FXDeal data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.GetFXDeal(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FXDealDAO for database read of all FXDeals
//----------------------------------------------------------------------------
func GetAllFXDeal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FXDeal data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.GetAllFXDeal()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FXDealDAO for database save
//----------------------------------------------------------------------------
func UpdateFXDeal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FXDeal model
	//----------------------------------------------------------------------------
	var data = model.FXDeal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FXDeal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FXDeal data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.UpdateFXDeal(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FXDealDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFXDeal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FXDeal data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FXDealDAO.DeleteFXDeal(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Quote on a FXDeal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignQuoteToFXDeal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXDealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quoteId,_ := strconv.ParseUint( vars["quoteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXDeal DAO
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.AssignQuoteToFXDeal(fXDealId, quoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Quote on a FXDeal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignQuoteFromFXDeal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	fXDealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FXDeal DAO
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.UnassignQuoteFromFXDeal(fXDealId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more paymentOrdersIds as a PaymentOrders to a FXDeal
	//----------------------------------------------------------------------------
func AddPaymentOrdersToFXDeal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fXDealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentOrdersIds,_ := vars["paymentOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FXDeal DAO
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.AddPaymentOrdersToFXDeal(fXDealId, paymentOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentOrdersIds as a PaymentOrders from a FXDeal
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentOrdersFromFXDeal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	fXDealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentOrdersIds,_ := vars["paymentOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FXDeal DAO
	//----------------------------------------------------------------------------
	requestResult := FXDealDAO.RemovePaymentOrdersFromFXDeal(fXDealId, paymentOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
