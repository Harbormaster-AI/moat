package controller

import (
    LoanTransactionDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LoanTransactionDAO for database creation
//----------------------------------------------------------------------------
func CreateLoanTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanTransaction model
	//----------------------------------------------------------------------------
	data := model.LoanTransaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanTransaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanTransaction data access object to create
	//----------------------------------------------------------------------------
	requestResult := LoanTransactionDAO.CreateLoanTransaction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LoanTransactionDAO to find the relevant LoanTransaction
//----------------------------------------------------------------------------
func GetLoanTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanTransaction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanTransactionDAO.GetLoanTransaction(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LoanTransactionDAO for database read of all LoanTransactions
//----------------------------------------------------------------------------
func GetAllLoanTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LoanTransaction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LoanTransactionDAO.GetAllLoanTransaction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LoanTransactionDAO for database save
//----------------------------------------------------------------------------
func UpdateLoanTransaction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanTransaction model
	//----------------------------------------------------------------------------
	var data = model.LoanTransaction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanTransaction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanTransaction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanTransactionDAO.UpdateLoanTransaction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LoanTransactionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLoanTransaction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanTransaction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LoanTransactionDAO.DeleteLoanTransaction(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Loan on a LoanTransaction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLoanToLoanTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanId,_ := strconv.ParseUint( vars["loanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := LoanTransactionDAO.AssignLoanToLoanTransaction(loanTransactionId, loanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Loan on a LoanTransaction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLoanFromLoanTransaction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanTransactionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanTransaction DAO
	//----------------------------------------------------------------------------
	requestResult := LoanTransactionDAO.UnassignLoanFromLoanTransaction(loanTransactionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


