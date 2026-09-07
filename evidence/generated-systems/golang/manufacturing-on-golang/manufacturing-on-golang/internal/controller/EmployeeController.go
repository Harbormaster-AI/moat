package controller

import (
    EmployeeDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EmployeeDAO for database creation
//----------------------------------------------------------------------------
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Employee model
	//----------------------------------------------------------------------------
	data := model.Employee{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Employee model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Employee data access object to create
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.CreateEmployee( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EmployeeDAO to find the relevant Employee
//----------------------------------------------------------------------------
func GetEmployee(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Employee data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.GetEmployee(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EmployeeDAO for database read of all Employees
//----------------------------------------------------------------------------
func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Employee data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.GetAllEmployee()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EmployeeDAO for database save
//----------------------------------------------------------------------------
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Employee model
	//----------------------------------------------------------------------------
	var data = model.Employee{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Employee model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Employee data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.UpdateEmployee(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EmployeeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Employee data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EmployeeDAO.DeleteEmployee(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a WorkCenter on a Employee
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkCenterToEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCenterId,_ := strconv.ParseUint( vars["workCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AssignWorkCenterToEmployee(employeeId, workCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkCenter on a Employee
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkCenterFromEmployee( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.UnassignWorkCenterFromEmployee(employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more shiftAssignmentsIds as a ShiftAssignments to a Employee
	//----------------------------------------------------------------------------
func AddShiftAssignmentsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	shiftAssignmentsIds,_ := vars["shiftAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddShiftAssignmentsToEmployee(employeeId, shiftAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more shiftAssignmentsIds as a ShiftAssignments from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveShiftAssignmentsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	shiftAssignmentsIds,_ := vars["shiftAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveShiftAssignmentsFromEmployee(employeeId, shiftAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more correctiveActionsIds as a CorrectiveActions to a Employee
	//----------------------------------------------------------------------------
func AddCorrectiveActionsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddCorrectiveActionsToEmployee(employeeId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more correctiveActionsIds as a CorrectiveActions from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	correctiveActionsIds,_ := vars["correctiveActionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveCorrectiveActionsFromEmployee(employeeId, correctiveActionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
