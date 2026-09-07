package controller

import (
    TimesheetDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TimesheetDAO for database creation
//----------------------------------------------------------------------------
func CreateTimesheet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Timesheet model
	//----------------------------------------------------------------------------
	data := model.Timesheet{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Timesheet model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet data access object to create
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.CreateTimesheet( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TimesheetDAO to find the relevant Timesheet
//----------------------------------------------------------------------------
func GetTimesheet(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Timesheet data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.GetTimesheet(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TimesheetDAO for database read of all Timesheets
//----------------------------------------------------------------------------
func GetAllTimesheet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Timesheet data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.GetAllTimesheet()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TimesheetDAO for database save
//----------------------------------------------------------------------------
func UpdateTimesheet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Timesheet model
	//----------------------------------------------------------------------------
	var data = model.Timesheet{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Timesheet model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.UpdateTimesheet(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TimesheetDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTimesheet(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Timesheet data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TimesheetDAO.DeleteTimesheet(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a Timesheet
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToTimesheet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timesheetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet DAO
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.AssignEmployeeToTimesheet(timesheetId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a Timesheet
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromTimesheet( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	timesheetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet DAO
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.UnassignEmployeeFromTimesheet(timesheetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more timeEntriesIds as a TimeEntries to a Timesheet
	//----------------------------------------------------------------------------
func AddTimeEntriesToTimesheet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timesheetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timeEntriesIds,_ := vars["timeEntriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet DAO
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.AddTimeEntriesToTimesheet(timesheetId, timeEntriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more timeEntriesIds as a TimeEntries from a Timesheet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTimeEntriesFromTimesheet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timesheetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	timeEntriesIds,_ := vars["timeEntriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet DAO
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.RemoveTimeEntriesFromTimesheet(timesheetId, timeEntriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more approvalsIds as a Approvals to a Timesheet
	//----------------------------------------------------------------------------
func AddApprovalsToTimesheet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timesheetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approvalsIds,_ := vars["approvalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet DAO
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.AddApprovalsToTimesheet(timesheetId, approvalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more approvalsIds as a Approvals from a Timesheet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveApprovalsFromTimesheet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	timesheetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approvalsIds,_ := vars["approvalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Timesheet DAO
	//----------------------------------------------------------------------------
	requestResult := TimesheetDAO.RemoveApprovalsFromTimesheet(timesheetId, approvalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
