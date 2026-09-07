package controller

import (
    PayrollItemDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PayrollItemDAO for database creation
//----------------------------------------------------------------------------
func CreatePayrollItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PayrollItem model
	//----------------------------------------------------------------------------
	data := model.PayrollItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PayrollItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollItem data access object to create
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.CreatePayrollItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PayrollItemDAO to find the relevant PayrollItem
//----------------------------------------------------------------------------
func GetPayrollItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PayrollItem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.GetPayrollItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PayrollItemDAO for database read of all PayrollItems
//----------------------------------------------------------------------------
func GetAllPayrollItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PayrollItem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.GetAllPayrollItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PayrollItemDAO for database save
//----------------------------------------------------------------------------
func UpdatePayrollItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PayrollItem model
	//----------------------------------------------------------------------------
	var data = model.PayrollItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PayrollItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollItem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.UpdatePayrollItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PayrollItemDAO for database deletion
//----------------------------------------------------------------------------
func DeletePayrollItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PayrollItem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PayrollItemDAO.DeletePayrollItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PayrollRun on a PayrollItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPayrollRunToPayrollItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollRunId,_ := strconv.ParseUint( vars["payrollRunId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollItem DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.AssignPayrollRunToPayrollItem(payrollItemId, payrollRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PayrollRun on a PayrollItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPayrollRunFromPayrollItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollItem DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.UnassignPayrollRunFromPayrollItem(payrollItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a PayrollItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToPayrollItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollItem DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.AssignEmployeeToPayrollItem(payrollItemId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a PayrollItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromPayrollItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollItem DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollItemDAO.UnassignEmployeeFromPayrollItem(payrollItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


