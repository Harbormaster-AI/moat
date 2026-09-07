package controller

import (
    EmploymentAssignmentDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EmploymentAssignmentDAO for database creation
//----------------------------------------------------------------------------
func CreateEmploymentAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EmploymentAssignment model
	//----------------------------------------------------------------------------
	data := model.EmploymentAssignment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EmploymentAssignment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment data access object to create
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.CreateEmploymentAssignment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EmploymentAssignmentDAO to find the relevant EmploymentAssignment
//----------------------------------------------------------------------------
func GetEmploymentAssignment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EmploymentAssignment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.GetEmploymentAssignment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EmploymentAssignmentDAO for database read of all EmploymentAssignments
//----------------------------------------------------------------------------
func GetAllEmploymentAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.GetAllEmploymentAssignment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EmploymentAssignmentDAO for database save
//----------------------------------------------------------------------------
func UpdateEmploymentAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EmploymentAssignment model
	//----------------------------------------------------------------------------
	var data = model.EmploymentAssignment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EmploymentAssignment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.UpdateEmploymentAssignment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EmploymentAssignmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEmploymentAssignment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EmploymentAssignment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EmploymentAssignmentDAO.DeleteEmploymentAssignment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a EmploymentAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToEmploymentAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.AssignEmployeeToEmploymentAssignment(employmentAssignmentId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a EmploymentAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromEmploymentAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.UnassignEmployeeFromEmploymentAssignment(employmentAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Position on a EmploymentAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPositionToEmploymentAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionId,_ := strconv.ParseUint( vars["positionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.AssignPositionToEmploymentAssignment(employmentAssignmentId, positionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Position on a EmploymentAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPositionFromEmploymentAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.UnassignPositionFromEmploymentAssignment(employmentAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Supervisor on a EmploymentAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSupervisorToEmploymentAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	supervisorId,_ := strconv.ParseUint( vars["supervisorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.AssignSupervisorToEmploymentAssignment(employmentAssignmentId, supervisorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Supervisor on a EmploymentAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSupervisorFromEmploymentAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employmentAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EmploymentAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := EmploymentAssignmentDAO.UnassignSupervisorFromEmploymentAssignment(employmentAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


