package controller

import (
    ScheduleExceptionDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ScheduleExceptionDAO for database creation
//----------------------------------------------------------------------------
func CreateScheduleException(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ScheduleException model
	//----------------------------------------------------------------------------
	data := model.ScheduleException{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ScheduleException model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ScheduleException data access object to create
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.CreateScheduleException( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ScheduleExceptionDAO to find the relevant ScheduleException
//----------------------------------------------------------------------------
func GetScheduleException(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ScheduleException data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.GetScheduleException(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ScheduleExceptionDAO for database read of all ScheduleExceptions
//----------------------------------------------------------------------------
func GetAllScheduleException(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ScheduleException data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.GetAllScheduleException()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ScheduleExceptionDAO for database save
//----------------------------------------------------------------------------
func UpdateScheduleException(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ScheduleException model
	//----------------------------------------------------------------------------
	var data = model.ScheduleException{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ScheduleException model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ScheduleException data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.UpdateScheduleException(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ScheduleExceptionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteScheduleException(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ScheduleException data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ScheduleExceptionDAO.DeleteScheduleException(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a WorkSchedule on a ScheduleException
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkScheduleToScheduleException(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	scheduleExceptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workScheduleId,_ := strconv.ParseUint( vars["workScheduleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ScheduleException DAO
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.AssignWorkScheduleToScheduleException(scheduleExceptionId, workScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkSchedule on a ScheduleException
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkScheduleFromScheduleException( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	scheduleExceptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ScheduleException DAO
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.UnassignWorkScheduleFromScheduleException(scheduleExceptionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a ScheduleException
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToScheduleException(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	scheduleExceptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ScheduleException DAO
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.AssignEmployeeToScheduleException(scheduleExceptionId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a ScheduleException
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromScheduleException( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	scheduleExceptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ScheduleException DAO
	//----------------------------------------------------------------------------
	requestResult := ScheduleExceptionDAO.UnassignEmployeeFromScheduleException(scheduleExceptionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


