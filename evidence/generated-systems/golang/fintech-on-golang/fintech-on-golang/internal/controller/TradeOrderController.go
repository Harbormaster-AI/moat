package controller

import (
    TradeOrderDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TradeOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateTradeOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TradeOrder model
	//----------------------------------------------------------------------------
	data := model.TradeOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TradeOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.CreateTradeOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TradeOrderDAO to find the relevant TradeOrder
//----------------------------------------------------------------------------
func GetTradeOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TradeOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.GetTradeOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TradeOrderDAO for database read of all TradeOrders
//----------------------------------------------------------------------------
func GetAllTradeOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.GetAllTradeOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TradeOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateTradeOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TradeOrder model
	//----------------------------------------------------------------------------
	var data = model.TradeOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TradeOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.UpdateTradeOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TradeOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTradeOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TradeOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TradeOrderDAO.DeleteTradeOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Portfolio on a TradeOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPortfolioToTradeOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	tradeOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	portfolioId,_ := strconv.ParseUint( vars["portfolioId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.AssignPortfolioToTradeOrder(tradeOrderId, portfolioId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Portfolio on a TradeOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPortfolioFromTradeOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	tradeOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.UnassignPortfolioFromTradeOrder(tradeOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Security on a TradeOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSecurityToTradeOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	tradeOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	securityId,_ := strconv.ParseUint( vars["securityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.AssignSecurityToTradeOrder(tradeOrderId, securityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Security on a TradeOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSecurityFromTradeOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	tradeOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.UnassignSecurityFromTradeOrder(tradeOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more tradesIds as a Trades to a TradeOrder
	//----------------------------------------------------------------------------
func AddTradesToTradeOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tradeOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tradesIds,_ := vars["tradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.AddTradesToTradeOrder(tradeOrderId, tradesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tradesIds as a Trades from a TradeOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTradesFromTradeOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	tradeOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tradesIds,_ := vars["tradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TradeOrder DAO
	//----------------------------------------------------------------------------
	requestResult := TradeOrderDAO.RemoveTradesFromTradeOrder(tradeOrderId, tradesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
