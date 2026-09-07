package controller

import (
    PayrollRunDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PayrollRunDAO for database creation
//----------------------------------------------------------------------------
func CreatePayrollRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PayrollRun model
	//----------------------------------------------------------------------------
	data := model.PayrollRun{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PayrollRun model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollRun data access object to create
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.CreatePayrollRun( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PayrollRunDAO to find the relevant PayrollRun
//----------------------------------------------------------------------------
func GetPayrollRun(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PayrollRun data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.GetPayrollRun(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PayrollRunDAO for database read of all PayrollRuns
//----------------------------------------------------------------------------
func GetAllPayrollRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PayrollRun data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.GetAllPayrollRun()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PayrollRunDAO for database save
//----------------------------------------------------------------------------
func UpdatePayrollRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PayrollRun model
	//----------------------------------------------------------------------------
	var data = model.PayrollRun{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PayrollRun model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollRun data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.UpdatePayrollRun(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PayrollRunDAO for database deletion
//----------------------------------------------------------------------------
func DeletePayrollRun(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PayrollRun data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PayrollRunDAO.DeletePayrollRun(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a PayrollCalendar on a PayrollRun
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPayrollCalendarToPayrollRun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollCalendarId,_ := strconv.ParseUint( vars["payrollCalendarId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollRun DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.AssignPayrollCalendarToPayrollRun(payrollRunId, payrollCalendarId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PayrollCalendar on a PayrollRun
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPayrollCalendarFromPayrollRun( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollRun DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.UnassignPayrollCalendarFromPayrollRun(payrollRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more payrollItemsIds as a PayrollItems to a PayrollRun
	//----------------------------------------------------------------------------
func AddPayrollItemsToPayrollRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	payrollRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollItemsIds,_ := vars["payrollItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PayrollRun DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.AddPayrollItemsToPayrollRun(payrollRunId, payrollItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more payrollItemsIds as a PayrollItems from a PayrollRun
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePayrollItemsFromPayrollRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	payrollRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollItemsIds,_ := vars["payrollItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PayrollRun DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollRunDAO.RemovePayrollItemsFromPayrollRun(payrollRunId, payrollItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
