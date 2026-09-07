package controller

import (
    ShiftAssignmentDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ShiftAssignmentDAO for database creation
//----------------------------------------------------------------------------
func CreateShiftAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ShiftAssignment model
	//----------------------------------------------------------------------------
	data := model.ShiftAssignment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ShiftAssignment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment data access object to create
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.CreateShiftAssignment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ShiftAssignmentDAO to find the relevant ShiftAssignment
//----------------------------------------------------------------------------
func GetShiftAssignment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ShiftAssignment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.GetShiftAssignment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ShiftAssignmentDAO for database read of all ShiftAssignments
//----------------------------------------------------------------------------
func GetAllShiftAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.GetAllShiftAssignment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ShiftAssignmentDAO for database save
//----------------------------------------------------------------------------
func UpdateShiftAssignment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ShiftAssignment model
	//----------------------------------------------------------------------------
	var data = model.ShiftAssignment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ShiftAssignment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.UpdateShiftAssignment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ShiftAssignmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteShiftAssignment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ShiftAssignment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ShiftAssignmentDAO.DeleteShiftAssignment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Shift on a ShiftAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignShiftToShiftAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	shiftId,_ := strconv.ParseUint( vars["shiftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.AssignShiftToShiftAssignment(shiftAssignmentId, shiftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Shift on a ShiftAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignShiftFromShiftAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.UnassignShiftFromShiftAssignment(shiftAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a ShiftAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToShiftAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.AssignEmployeeToShiftAssignment(shiftAssignmentId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a ShiftAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromShiftAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.UnassignEmployeeFromShiftAssignment(shiftAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a WorkCenter on a ShiftAssignment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkCenterToShiftAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCenterId,_ := strconv.ParseUint( vars["workCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.AssignWorkCenterToShiftAssignment(shiftAssignmentId, workCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkCenter on a ShiftAssignment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkCenterFromShiftAssignment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftAssignmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ShiftAssignment DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftAssignmentDAO.UnassignWorkCenterFromShiftAssignment(shiftAssignmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


