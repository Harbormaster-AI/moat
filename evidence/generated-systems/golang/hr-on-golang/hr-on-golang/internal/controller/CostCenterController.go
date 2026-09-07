package controller

import (
    CostCenterDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CostCenterDAO for database creation
//----------------------------------------------------------------------------
func CreateCostCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CostCenter model
	//----------------------------------------------------------------------------
	data := model.CostCenter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CostCenter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter data access object to create
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.CreateCostCenter( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CostCenterDAO to find the relevant CostCenter
//----------------------------------------------------------------------------
func GetCostCenter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CostCenter data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.GetCostCenter(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CostCenterDAO for database read of all CostCenters
//----------------------------------------------------------------------------
func GetAllCostCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CostCenter data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.GetAllCostCenter()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CostCenterDAO for database save
//----------------------------------------------------------------------------
func UpdateCostCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CostCenter model
	//----------------------------------------------------------------------------
	var data = model.CostCenter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CostCenter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.UpdateCostCenter(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CostCenterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCostCenter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CostCenter data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CostCenterDAO.DeleteCostCenter(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a CostCenter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToCostCenter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.AssignOrganizationToCostCenter(costCenterId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a CostCenter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromCostCenter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.UnassignOrganizationFromCostCenter(costCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more departmentsIds as a Departments to a CostCenter
	//----------------------------------------------------------------------------
func AddDepartmentsToCostCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.AddDepartmentsToCostCenter(costCenterId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more departmentsIds as a Departments from a CostCenter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDepartmentsFromCostCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.RemoveDepartmentsFromCostCenter(costCenterId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more positionsIds as a Positions to a CostCenter
	//----------------------------------------------------------------------------
func AddPositionsToCostCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.AddPositionsToCostCenter(costCenterId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more positionsIds as a Positions from a CostCenter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePositionsFromCostCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.RemovePositionsFromCostCenter(costCenterId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more employeesIds as a Employees to a CostCenter
	//----------------------------------------------------------------------------
func AddEmployeesToCostCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.AddEmployeesToCostCenter(costCenterId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more employeesIds as a Employees from a CostCenter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmployeesFromCostCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	costCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CostCenter DAO
	//----------------------------------------------------------------------------
	requestResult := CostCenterDAO.RemoveEmployeesFromCostCenter(costCenterId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
