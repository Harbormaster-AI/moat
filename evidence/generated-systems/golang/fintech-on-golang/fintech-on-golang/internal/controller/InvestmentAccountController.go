package controller

import (
    InvestmentAccountDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InvestmentAccountDAO for database creation
//----------------------------------------------------------------------------
func CreateInvestmentAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InvestmentAccount model
	//----------------------------------------------------------------------------
	data := model.InvestmentAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InvestmentAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount data access object to create
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.CreateInvestmentAccount( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InvestmentAccountDAO to find the relevant InvestmentAccount
//----------------------------------------------------------------------------
func GetInvestmentAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InvestmentAccount data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.GetInvestmentAccount(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InvestmentAccountDAO for database read of all InvestmentAccounts
//----------------------------------------------------------------------------
func GetAllInvestmentAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.GetAllInvestmentAccount()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InvestmentAccountDAO for database save
//----------------------------------------------------------------------------
func UpdateInvestmentAccount(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InvestmentAccount model
	//----------------------------------------------------------------------------
	var data = model.InvestmentAccount{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InvestmentAccount model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.UpdateInvestmentAccount(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InvestmentAccountDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInvestmentAccount(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InvestmentAccount data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InvestmentAccountDAO.DeleteInvestmentAccount(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Portfolio on a InvestmentAccount
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPortfolioToInvestmentAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	investmentAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	portfolioId,_ := strconv.ParseUint( vars["portfolioId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.AssignPortfolioToInvestmentAccount(investmentAccountId, portfolioId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Portfolio on a InvestmentAccount
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPortfolioFromInvestmentAccount( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	investmentAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.UnassignPortfolioFromInvestmentAccount(investmentAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more tradesIds as a Trades to a InvestmentAccount
	//----------------------------------------------------------------------------
func AddTradesToInvestmentAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tradesIds,_ := vars["tradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.AddTradesToInvestmentAccount(investmentAccountId, tradesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tradesIds as a Trades from a InvestmentAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTradesFromInvestmentAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tradesIds,_ := vars["tradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.RemoveTradesFromInvestmentAccount(investmentAccountId, tradesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a InvestmentAccount
	//----------------------------------------------------------------------------
func AddOrdersToInvestmentAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.AddOrdersToInvestmentAccount(investmentAccountId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a InvestmentAccount
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromInvestmentAccount(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentAccountId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentAccount DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentAccountDAO.RemoveOrdersFromInvestmentAccount(investmentAccountId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
