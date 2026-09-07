package controller

import (
    WorkShiftDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to WorkShiftDAO for database creation
//----------------------------------------------------------------------------
func CreateWorkShift(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkShift model
	//----------------------------------------------------------------------------
	data := model.WorkShift{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkShift model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkShift data access object to create
	//----------------------------------------------------------------------------
	requestResult := WorkShiftDAO.CreateWorkShift( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to WorkShiftDAO to find the relevant WorkShift
//----------------------------------------------------------------------------
func GetWorkShift(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkShift data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkShiftDAO.GetWorkShift(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to WorkShiftDAO for database read of all WorkShifts
//----------------------------------------------------------------------------
func GetAllWorkShift(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the WorkShift data access object to get all
	//----------------------------------------------------------------------------
	requestResult := WorkShiftDAO.GetAllWorkShift()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to WorkShiftDAO for database save
//----------------------------------------------------------------------------
func UpdateWorkShift(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkShift model
	//----------------------------------------------------------------------------
	var data = model.WorkShift{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkShift model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkShift data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkShiftDAO.UpdateWorkShift(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to WorkShiftDAO for database deletion
//----------------------------------------------------------------------------
func DeleteWorkShift(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkShift data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := WorkShiftDAO.DeleteWorkShift(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a WorkSchedule on a WorkShift
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkScheduleToWorkShift(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workShiftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workScheduleId,_ := strconv.ParseUint( vars["workScheduleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkShift DAO
	//----------------------------------------------------------------------------
	requestResult := WorkShiftDAO.AssignWorkScheduleToWorkShift(workShiftId, workScheduleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a WorkSchedule on a WorkShift
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkScheduleFromWorkShift( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workShiftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkShift DAO
	//----------------------------------------------------------------------------
	requestResult := WorkShiftDAO.UnassignWorkScheduleFromWorkShift(workShiftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


