package controller

import (
    DepartmentDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
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
	// assigns a Facility on a Department
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToDepartment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.AssignFacilityToDepartment(departmentId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a Department
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromDepartment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.UnassignFacilityFromDepartment(departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more careTeamsIds as a CareTeams to a Department
	//----------------------------------------------------------------------------
func AddCareTeamsToDepartment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	careTeamsIds,_ := vars["careTeamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.AddCareTeamsToDepartment(departmentId, careTeamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more careTeamsIds as a CareTeams from a Department
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCareTeamsFromDepartment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	departmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	careTeamsIds,_ := vars["careTeamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Department DAO
	//----------------------------------------------------------------------------
	requestResult := DepartmentDAO.RemoveCareTeamsFromDepartment(departmentId, careTeamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
