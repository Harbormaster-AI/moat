package controller

import (
    EmploymentContractDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EmploymentContractDAO for database creation
//----------------------------------------------------------------------------
func CreateEmploymentContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EmploymentContract model
	//----------------------------------------------------------------------------
	data := model.EmploymentContract{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EmploymentContract model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract data access object to create
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.CreateEmploymentContract( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EmploymentContractDAO to find the relevant EmploymentContract
//----------------------------------------------------------------------------
func GetEmploymentContract(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EmploymentContract data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.GetEmploymentContract(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EmploymentContractDAO for database read of all EmploymentContracts
//----------------------------------------------------------------------------
func GetAllEmploymentContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.GetAllEmploymentContract()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EmploymentContractDAO for database save
//----------------------------------------------------------------------------
func UpdateEmploymentContract(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EmploymentContract model
	//----------------------------------------------------------------------------
	var data = model.EmploymentContract{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EmploymentContract model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.UpdateEmploymentContract(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EmploymentContractDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEmploymentContract(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EmploymentContract data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EmploymentContractDAO.DeleteEmploymentContract(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a EmploymentContract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToEmploymentContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.AssignEmployeeToEmploymentContract(employmentContractId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a EmploymentContract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromEmploymentContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.UnassignEmployeeFromEmploymentContract(employmentContractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CompensationPackage on a EmploymentContract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCompensationPackageToEmploymentContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compensationPackageId,_ := strconv.ParseUint( vars["compensationPackageId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.AssignCompensationPackageToEmploymentContract(employmentContractId, compensationPackageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CompensationPackage on a EmploymentContract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCompensationPackageFromEmploymentContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.UnassignCompensationPackageFromEmploymentContract(employmentContractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkSchedule on a EmploymentContract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkScheduleToEmploymentContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workScheduleId,_ := strconv.ParseUint( vars["workScheduleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.AssignWorkScheduleToEmploymentContract(employmentContractId, workScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkSchedule on a EmploymentContract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkScheduleFromEmploymentContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.UnassignWorkScheduleFromEmploymentContract(employmentContractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a EmploymentContract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToEmploymentContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.AssignLocationToEmploymentContract(employmentContractId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a EmploymentContract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromEmploymentContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.UnassignLocationFromEmploymentContract(employmentContractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PayrollCalendar on a EmploymentContract
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPayrollCalendarToEmploymentContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	payrollCalendarId,_ := strconv.ParseUint( vars["payrollCalendarId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.AssignPayrollCalendarToEmploymentContract(employmentContractId, payrollCalendarId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PayrollCalendar on a EmploymentContract
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPayrollCalendarFromEmploymentContract( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentContractId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentContract DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentContractDAO.UnassignPayrollCalendarFromEmploymentContract(employmentContractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


