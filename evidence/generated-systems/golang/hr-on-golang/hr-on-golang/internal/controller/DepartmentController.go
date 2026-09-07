package controller

import (
    DepartmentDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DepartmentDAO for database creation
//----------------------------------------------------------------------------
func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Department model
	//----------------------------------------------------------------------------
	data := model.Department{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Department model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Department data access object to create
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.CreateDepartment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DepartmentDAO to find the relevant Department
//----------------------------------------------------------------------------
func GetDepartment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Department data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.GetDepartment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DepartmentDAO for database read of all Departments
//----------------------------------------------------------------------------
func GetAllDepartment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Department data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.GetAllDepartment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DepartmentDAO for database save
//----------------------------------------------------------------------------
func UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Department model
	//----------------------------------------------------------------------------
	var data = model.Department{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Department model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Department data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.UpdateDepartment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DepartmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Department data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DepartmentDAO.DeleteDepartment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Department
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToDepartment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.AssignOrganizationToDepartment(departmentId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Department
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromDepartment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.UnassignOrganizationFromDepartment(departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Manager on a Department
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignManagerToDepartment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	managerId,_ := strconv.ParseUint( vars["managerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.AssignManagerToDepartment(departmentId, managerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Manager on a Department
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignManagerFromDepartment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.UnassignManagerFromDepartment(departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CostCenter on a Department
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCostCenterToDepartment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	costCenterId,_ := strconv.ParseUint( vars["costCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.AssignCostCenterToDepartment(departmentId, costCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CostCenter on a Department
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCostCenterFromDepartment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.UnassignCostCenterFromDepartment(departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more positionsIds as a Positions to a Department
	//----------------------------------------------------------------------------
func AddPositionsToDepartment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.AddPositionsToDepartment(departmentId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more positionsIds as a Positions from a Department
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePositionsFromDepartment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.RemovePositionsFromDepartment(departmentId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more employeesIds as a Employees to a Department
	//----------------------------------------------------------------------------
func AddEmployeesToDepartment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.AddEmployeesToDepartment(departmentId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more employeesIds as a Employees from a Department
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmployeesFromDepartment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.RemoveEmployeesFromDepartment(departmentId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
