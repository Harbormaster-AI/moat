package controller

import (
    PayrollCalendarDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PayrollCalendarDAO for database creation
//----------------------------------------------------------------------------
func CreatePayrollCalendar(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PayrollCalendar model
	//----------------------------------------------------------------------------
	data := model.PayrollCalendar{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PayrollCalendar model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar data access object to create
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.CreatePayrollCalendar( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PayrollCalendarDAO to find the relevant PayrollCalendar
//----------------------------------------------------------------------------
func GetPayrollCalendar(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PayrollCalendar data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.GetPayrollCalendar(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PayrollCalendarDAO for database read of all PayrollCalendars
//----------------------------------------------------------------------------
func GetAllPayrollCalendar(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.GetAllPayrollCalendar()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PayrollCalendarDAO for database save
//----------------------------------------------------------------------------
func UpdatePayrollCalendar(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty PayrollCalendar model
	//----------------------------------------------------------------------------
	var data = model.PayrollCalendar{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a PayrollCalendar model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.UpdatePayrollCalendar(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PayrollCalendarDAO for database deletion
//----------------------------------------------------------------------------
func DeletePayrollCalendar(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the PayrollCalendar data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PayrollCalendarDAO.DeletePayrollCalendar(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a PayrollCalendar
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToPayrollCalendar(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollCalendarId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.AssignOrganizationToPayrollCalendar(payrollCalendarId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a PayrollCalendar
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromPayrollCalendar( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	payrollCalendarId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.UnassignOrganizationFromPayrollCalendar(payrollCalendarId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more payrollRunsIds as a PayrollRuns to a PayrollCalendar
	//----------------------------------------------------------------------------
func AddPayrollRunsToPayrollCalendar(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	payrollCalendarId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollRunsIds,_ := vars["payrollRunsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.AddPayrollRunsToPayrollCalendar(payrollCalendarId, payrollRunsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more payrollRunsIds as a PayrollRuns from a PayrollCalendar
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePayrollRunsFromPayrollCalendar(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	payrollCalendarId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollRunsIds,_ := vars["payrollRunsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.RemovePayrollRunsFromPayrollCalendar(payrollCalendarId, payrollRunsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more employeesIds as a Employees to a PayrollCalendar
	//----------------------------------------------------------------------------
func AddEmployeesToPayrollCalendar(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	payrollCalendarId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.AddEmployeesToPayrollCalendar(payrollCalendarId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more employeesIds as a Employees from a PayrollCalendar
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmployeesFromPayrollCalendar(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	payrollCalendarId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the PayrollCalendar DAO
	//----------------------------------------------------------------------------
	requestResult := PayrollCalendarDAO.RemoveEmployeesFromPayrollCalendar(payrollCalendarId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
