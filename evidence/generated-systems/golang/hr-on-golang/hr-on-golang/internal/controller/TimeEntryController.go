package controller

import (
    TimeEntryDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TimeEntryDAO for database creation
//----------------------------------------------------------------------------
func CreateTimeEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TimeEntry model
	//----------------------------------------------------------------------------
	data := model.TimeEntry{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TimeEntry model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry data access object to create
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.CreateTimeEntry( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TimeEntryDAO to find the relevant TimeEntry
//----------------------------------------------------------------------------
func GetTimeEntry(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TimeEntry data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.GetTimeEntry(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TimeEntryDAO for database read of all TimeEntrys
//----------------------------------------------------------------------------
func GetAllTimeEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.GetAllTimeEntry()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TimeEntryDAO for database save
//----------------------------------------------------------------------------
func UpdateTimeEntry(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TimeEntry model
	//----------------------------------------------------------------------------
	var data = model.TimeEntry{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TimeEntry model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.UpdateTimeEntry(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TimeEntryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTimeEntry(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TimeEntry data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TimeEntryDAO.DeleteTimeEntry(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Timesheet on a TimeEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTimesheetToTimeEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timeEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timesheetId,_ := strconv.ParseUint( vars["timesheetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry DAO
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.AssignTimesheetToTimeEntry(timeEntryId, timesheetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Timesheet on a TimeEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTimesheetFromTimeEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timeEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry DAO
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.UnassignTimesheetFromTimeEntry(timeEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a TimeEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToTimeEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timeEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry DAO
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.AssignEmployeeToTimeEntry(timeEntryId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a TimeEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromTimeEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timeEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry DAO
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.UnassignEmployeeFromTimeEntry(timeEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CostCenter on a TimeEntry
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCostCenterToTimeEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timeEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	costCenterId,_ := strconv.ParseUint( vars["costCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry DAO
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.AssignCostCenterToTimeEntry(timeEntryId, costCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CostCenter on a TimeEntry
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCostCenterFromTimeEntry( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timeEntryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TimeEntry DAO
	//----------------------------------------------------------------------------
	requestResult := TimeEntryDAO.UnassignCostCenterFromTimeEntry(timeEntryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


