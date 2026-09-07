package controller

import (
    LocationDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LocationDAO for database creation
//----------------------------------------------------------------------------
func CreateLocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Location model
	//----------------------------------------------------------------------------
	data := model.Location{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Location model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Location data access object to create
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.CreateLocation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LocationDAO to find the relevant Location
//----------------------------------------------------------------------------
func GetLocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Location data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.GetLocation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LocationDAO for database read of all Locations
//----------------------------------------------------------------------------
func GetAllLocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Location data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.GetAllLocation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LocationDAO for database save
//----------------------------------------------------------------------------
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Location model
	//----------------------------------------------------------------------------
	var data = model.Location{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Location model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Location data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.UpdateLocation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LocationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLocation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Location data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LocationDAO.DeleteLocation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a Location
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToLocation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.AssignOrganizationToLocation(locationId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a Location
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromLocation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.UnassignOrganizationFromLocation(locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more departmentsIds as a Departments to a Location
	//----------------------------------------------------------------------------
func AddDepartmentsToLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.AddDepartmentsToLocation(locationId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more departmentsIds as a Departments from a Location
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDepartmentsFromLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.RemoveDepartmentsFromLocation(locationId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more positionsIds as a Positions to a Location
	//----------------------------------------------------------------------------
func AddPositionsToLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.AddPositionsToLocation(locationId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more positionsIds as a Positions from a Location
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePositionsFromLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.RemovePositionsFromLocation(locationId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more employeesIds as a Employees to a Location
	//----------------------------------------------------------------------------
func AddEmployeesToLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.AddEmployeesToLocation(locationId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more employeesIds as a Employees from a Location
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEmployeesFromLocation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	locationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeesIds,_ := vars["employeesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Location DAO
	//----------------------------------------------------------------------------
	requestResult := LocationDAO.RemoveEmployeesFromLocation(locationId, employeesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
