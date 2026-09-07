package controller

import (
    EmployeeDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
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
	// assigns a Manager on a Employee
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignManagerToEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	managerId,_ := strconv.ParseUint( vars["managerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AssignManagerToEmployee(employeeId, managerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Manager on a Employee
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignManagerFromEmployee( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.UnassignManagerFromEmployee(employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Department on a Employee
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDepartmentToEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentId,_ := strconv.ParseUint( vars["departmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AssignDepartmentToEmployee(employeeId, departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Department on a Employee
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDepartmentFromEmployee( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.UnassignDepartmentFromEmployee(employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a PrimaryLocation on a Employee
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPrimaryLocationToEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	primaryLocationId,_ := strconv.ParseUint( vars["primaryLocationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AssignPrimaryLocationToEmployee(employeeId, primaryLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a PrimaryLocation on a Employee
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPrimaryLocationFromEmployee( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.UnassignPrimaryLocationFromEmployee(employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CostCenter on a Employee
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCostCenterToEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	costCenterId,_ := strconv.ParseUint( vars["costCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AssignCostCenterToEmployee(employeeId, costCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CostCenter on a Employee
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCostCenterFromEmployee( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.UnassignCostCenterFromEmployee(employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more directReportsIds as a DirectReports to a Employee
	//----------------------------------------------------------------------------
func AddDirectReportsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	directReportsIds,_ := vars["directReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddDirectReportsToEmployee(employeeId, directReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more directReportsIds as a DirectReports from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDirectReportsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	directReportsIds,_ := vars["directReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveDirectReportsFromEmployee(employeeId, directReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more employmentAssignmentsIds as a EmploymentAssignments to a Employee
	//----------------------------------------------------------------------------
func AddEmploymentAssignmentsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employmentAssignmentsIds,_ := vars["employmentAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddEmploymentAssignmentsToEmployee(employeeId, employmentAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more employmentAssignmentsIds as a EmploymentAssignments from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmploymentAssignmentsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employmentAssignmentsIds,_ := vars["employmentAssignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveEmploymentAssignmentsFromEmployee(employeeId, employmentAssignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a Employee
	//----------------------------------------------------------------------------
func AddContractsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddContractsToEmployee(employeeId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveContractsFromEmployee(employeeId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more benefitEnrollmentsIds as a BenefitEnrollments to a Employee
	//----------------------------------------------------------------------------
func AddBenefitEnrollmentsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	benefitEnrollmentsIds,_ := vars["benefitEnrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddBenefitEnrollmentsToEmployee(employeeId, benefitEnrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more benefitEnrollmentsIds as a BenefitEnrollments from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBenefitEnrollmentsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	benefitEnrollmentsIds,_ := vars["benefitEnrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveBenefitEnrollmentsFromEmployee(employeeId, benefitEnrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more timesheetsIds as a Timesheets to a Employee
	//----------------------------------------------------------------------------
func AddTimesheetsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timesheetsIds,_ := vars["timesheetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddTimesheetsToEmployee(employeeId, timesheetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more timesheetsIds as a Timesheets from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTimesheetsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timesheetsIds,_ := vars["timesheetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveTimesheetsFromEmployee(employeeId, timesheetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more leaveRequestsIds as a LeaveRequests to a Employee
	//----------------------------------------------------------------------------
func AddLeaveRequestsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leaveRequestsIds,_ := vars["leaveRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddLeaveRequestsToEmployee(employeeId, leaveRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more leaveRequestsIds as a LeaveRequests from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLeaveRequestsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leaveRequestsIds,_ := vars["leaveRequestsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveLeaveRequestsFromEmployee(employeeId, leaveRequestsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more performanceReviewsIds as a PerformanceReviews to a Employee
	//----------------------------------------------------------------------------
func AddPerformanceReviewsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	performanceReviewsIds,_ := vars["performanceReviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddPerformanceReviewsToEmployee(employeeId, performanceReviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more performanceReviewsIds as a PerformanceReviews from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePerformanceReviewsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	performanceReviewsIds,_ := vars["performanceReviewsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemovePerformanceReviewsFromEmployee(employeeId, performanceReviewsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more trainingEnrollmentsIds as a TrainingEnrollments to a Employee
	//----------------------------------------------------------------------------
func AddTrainingEnrollmentsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingEnrollmentsIds,_ := vars["trainingEnrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddTrainingEnrollmentsToEmployee(employeeId, trainingEnrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more trainingEnrollmentsIds as a TrainingEnrollments from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTrainingEnrollmentsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingEnrollmentsIds,_ := vars["trainingEnrollmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveTrainingEnrollmentsFromEmployee(employeeId, trainingEnrollmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more workAuthorizationsIds as a WorkAuthorizations to a Employee
	//----------------------------------------------------------------------------
func AddWorkAuthorizationsToEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workAuthorizationsIds,_ := vars["workAuthorizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.AddWorkAuthorizationsToEmployee(employeeId, workAuthorizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workAuthorizationsIds as a WorkAuthorizations from a Employee
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkAuthorizationsFromEmployee(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	employeeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workAuthorizationsIds,_ := vars["workAuthorizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Employee DAO
	//----------------------------------------------------------------------------
	requestResult := EmployeeDAO.RemoveWorkAuthorizationsFromEmployee(employeeId, workAuthorizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
