package controller

import (
    LoanDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LoanDAO for database creation
//----------------------------------------------------------------------------
func CreateLoan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Loan model
	//----------------------------------------------------------------------------
	data := model.Loan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Loan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Loan data access object to create
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.CreateLoan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LoanDAO to find the relevant Loan
//----------------------------------------------------------------------------
func GetLoan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Loan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.GetLoan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LoanDAO for database read of all Loans
//----------------------------------------------------------------------------
func GetAllLoan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Loan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.GetAllLoan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LoanDAO for database save
//----------------------------------------------------------------------------
func UpdateLoan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Loan model
	//----------------------------------------------------------------------------
	var data = model.Loan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Loan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Loan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.UpdateLoan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LoanDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLoan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Loan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LoanDAO.DeleteLoan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Loan
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToLoan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.AssignCustomerToLoan(loanId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Loan
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromLoan( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.UnassignCustomerFromLoan(loanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more scheduleIds as a Schedule to a Loan
	//----------------------------------------------------------------------------
func AddScheduleToLoan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	scheduleIds,_ := vars["scheduleIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.AddScheduleToLoan(loanId, scheduleIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more scheduleIds as a Schedule from a Loan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveScheduleFromLoan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	scheduleIds,_ := vars["scheduleIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.RemoveScheduleFromLoan(loanId, scheduleIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more collateralIds as a Collateral to a Loan
	//----------------------------------------------------------------------------
func AddCollateralToLoan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	collateralIds,_ := vars["collateralIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.AddCollateralToLoan(loanId, collateralIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more collateralIds as a Collateral from a Loan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCollateralFromLoan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	collateralIds,_ := vars["collateralIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.RemoveCollateralFromLoan(loanId, collateralIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more transactionsIds as a Transactions to a Loan
	//----------------------------------------------------------------------------
func AddTransactionsToLoan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.AddTransactionsToLoan(loanId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more transactionsIds as a Transactions from a Loan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTransactionsFromLoan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	loanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	transactionsIds,_ := vars["transactionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Loan DAO
	//----------------------------------------------------------------------------
	requestResult := LoanDAO.RemoveTransactionsFromLoan(loanId, transactionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
