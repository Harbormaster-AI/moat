package controller

import (
    LoanApplicationDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LoanApplicationDAO for database creation
//----------------------------------------------------------------------------
func CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanApplication model
	//----------------------------------------------------------------------------
	data := model.LoanApplication{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanApplication model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication data access object to create
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.CreateLoanApplication( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LoanApplicationDAO to find the relevant LoanApplication
//----------------------------------------------------------------------------
func GetLoanApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanApplication data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.GetLoanApplication(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LoanApplicationDAO for database read of all LoanApplications
//----------------------------------------------------------------------------
func GetAllLoanApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.GetAllLoanApplication()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LoanApplicationDAO for database save
//----------------------------------------------------------------------------
func UpdateLoanApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LoanApplication model
	//----------------------------------------------------------------------------
	var data = model.LoanApplication{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LoanApplication model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.UpdateLoanApplication(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LoanApplicationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLoanApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LoanApplication data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LoanApplicationDAO.DeleteLoanApplication(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a LoanApplication
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToLoanApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication DAO
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.AssignCustomerToLoanApplication(loanApplicationId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a LoanApplication
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromLoanApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication DAO
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.UnassignCustomerFromLoanApplication(loanApplicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a RiskAssessment on a LoanApplication
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRiskAssessmentToLoanApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	riskAssessmentId,_ := strconv.ParseUint( vars["riskAssessmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication DAO
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.AssignRiskAssessmentToLoanApplication(loanApplicationId, riskAssessmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a RiskAssessment on a LoanApplication
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRiskAssessmentFromLoanApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication DAO
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.UnassignRiskAssessmentFromLoanApplication(loanApplicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Loan on a LoanApplication
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLoanToLoanApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanId,_ := strconv.ParseUint( vars["loanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication DAO
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.AssignLoanToLoanApplication(loanApplicationId, loanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Loan on a LoanApplication
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLoanFromLoanApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	loanApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LoanApplication DAO
	//----------------------------------------------------------------------------
	requestResult := LoanApplicationDAO.UnassignLoanFromLoanApplication(loanApplicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


