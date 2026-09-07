package controller

import (
    InvestmentPortfolioDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InvestmentPortfolioDAO for database creation
//----------------------------------------------------------------------------
func CreateInvestmentPortfolio(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InvestmentPortfolio model
	//----------------------------------------------------------------------------
	data := model.InvestmentPortfolio{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InvestmentPortfolio model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio data access object to create
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.CreateInvestmentPortfolio( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InvestmentPortfolioDAO to find the relevant InvestmentPortfolio
//----------------------------------------------------------------------------
func GetInvestmentPortfolio(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InvestmentPortfolio data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.GetInvestmentPortfolio(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InvestmentPortfolioDAO for database read of all InvestmentPortfolios
//----------------------------------------------------------------------------
func GetAllInvestmentPortfolio(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.GetAllInvestmentPortfolio()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InvestmentPortfolioDAO for database save
//----------------------------------------------------------------------------
func UpdateInvestmentPortfolio(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InvestmentPortfolio model
	//----------------------------------------------------------------------------
	var data = model.InvestmentPortfolio{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InvestmentPortfolio model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.UpdateInvestmentPortfolio(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InvestmentPortfolioDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInvestmentPortfolio(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InvestmentPortfolio data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InvestmentPortfolioDAO.DeleteInvestmentPortfolio(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a InvestmentPortfolio
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToInvestmentPortfolio(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.AssignCustomerToInvestmentPortfolio(investmentPortfolioId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a InvestmentPortfolio
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromInvestmentPortfolio( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.UnassignCustomerFromInvestmentPortfolio(investmentPortfolioId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more accountsIds as a Accounts to a InvestmentPortfolio
	//----------------------------------------------------------------------------
func AddAccountsToInvestmentPortfolio(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.AddAccountsToInvestmentPortfolio(investmentPortfolioId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more accountsIds as a Accounts from a InvestmentPortfolio
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAccountsFromInvestmentPortfolio(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	accountsIds,_ := vars["accountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.RemoveAccountsFromInvestmentPortfolio(investmentPortfolioId, accountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a InvestmentPortfolio
	//----------------------------------------------------------------------------
func AddOrdersToInvestmentPortfolio(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.AddOrdersToInvestmentPortfolio(investmentPortfolioId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a InvestmentPortfolio
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromInvestmentPortfolio(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.RemoveOrdersFromInvestmentPortfolio(investmentPortfolioId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more holdingsIds as a Holdings to a InvestmentPortfolio
	//----------------------------------------------------------------------------
func AddHoldingsToInvestmentPortfolio(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	holdingsIds,_ := vars["holdingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.AddHoldingsToInvestmentPortfolio(investmentPortfolioId, holdingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more holdingsIds as a Holdings from a InvestmentPortfolio
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveHoldingsFromInvestmentPortfolio(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	investmentPortfolioId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	holdingsIds,_ := vars["holdingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InvestmentPortfolio DAO
	//----------------------------------------------------------------------------
	requestResult := InvestmentPortfolioDAO.RemoveHoldingsFromInvestmentPortfolio(investmentPortfolioId, holdingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
