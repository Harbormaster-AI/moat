package controller

import (
    RepaymentScheduleDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RepaymentScheduleDAO for database creation
//----------------------------------------------------------------------------
func CreateRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RepaymentSchedule model
	//----------------------------------------------------------------------------
	data := model.RepaymentSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RepaymentSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object to create
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.CreateRepaymentSchedule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RepaymentScheduleDAO to find the relevant RepaymentSchedule
//----------------------------------------------------------------------------
func GetRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RepaymentSchedule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.GetRepaymentSchedule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RepaymentScheduleDAO for database read of all RepaymentSchedules
//----------------------------------------------------------------------------
func GetAllRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.GetAllRepaymentSchedule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RepaymentScheduleDAO for database save
//----------------------------------------------------------------------------
func UpdateRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RepaymentSchedule model
	//----------------------------------------------------------------------------
	var data = model.RepaymentSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RepaymentSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UpdateRepaymentSchedule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RepaymentScheduleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRepaymentSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RepaymentSchedule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RepaymentScheduleDAO.DeleteRepaymentSchedule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Loan on a RepaymentSchedule
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLoanToRepaymentSchedule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	loanId,_ := strconv.ParseUint( vars["loanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.AssignLoanToRepaymentSchedule(repaymentScheduleId, loanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Loan on a RepaymentSchedule
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLoanFromRepaymentSchedule( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.UnassignLoanFromRepaymentSchedule(repaymentScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more paymentsIds as a Payments to a RepaymentSchedule
	//----------------------------------------------------------------------------
func AddPaymentsToRepaymentSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.AddPaymentsToRepaymentSchedule(repaymentScheduleId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more paymentsIds as a Payments from a RepaymentSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePaymentsFromRepaymentSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	repaymentScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	paymentsIds,_ := vars["paymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RepaymentSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := RepaymentScheduleDAO.RemovePaymentsFromRepaymentSchedule(repaymentScheduleId, paymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
