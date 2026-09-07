package controller

import (
    WorkScheduleDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to WorkScheduleDAO for database creation
//----------------------------------------------------------------------------
func CreateWorkSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkSchedule model
	//----------------------------------------------------------------------------
	data := model.WorkSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule data access object to create
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.CreateWorkSchedule( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to WorkScheduleDAO to find the relevant WorkSchedule
//----------------------------------------------------------------------------
func GetWorkSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkSchedule data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.GetWorkSchedule(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to WorkScheduleDAO for database read of all WorkSchedules
//----------------------------------------------------------------------------
func GetAllWorkSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule data access object to get all
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.GetAllWorkSchedule()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to WorkScheduleDAO for database save
//----------------------------------------------------------------------------
func UpdateWorkSchedule(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkSchedule model
	//----------------------------------------------------------------------------
	var data = model.WorkSchedule{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkSchedule model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.UpdateWorkSchedule(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to WorkScheduleDAO for database deletion
//----------------------------------------------------------------------------
func DeleteWorkSchedule(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkSchedule data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := WorkScheduleDAO.DeleteWorkSchedule(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a WorkSchedule
	//----------------------------------------------------------------------------
func AddContractsToWorkSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.AddContractsToWorkSchedule(workScheduleId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a WorkSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromWorkSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.RemoveContractsFromWorkSchedule(workScheduleId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more shiftsIds as a Shifts to a WorkSchedule
	//----------------------------------------------------------------------------
func AddShiftsToWorkSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	shiftsIds,_ := vars["shiftsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.AddShiftsToWorkSchedule(workScheduleId, shiftsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more shiftsIds as a Shifts from a WorkSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveShiftsFromWorkSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	shiftsIds,_ := vars["shiftsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.RemoveShiftsFromWorkSchedule(workScheduleId, shiftsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more exceptionsIds as a Exceptions to a WorkSchedule
	//----------------------------------------------------------------------------
func AddExceptionsToWorkSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exceptionsIds,_ := vars["exceptionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.AddExceptionsToWorkSchedule(workScheduleId, exceptionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more exceptionsIds as a Exceptions from a WorkSchedule
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExceptionsFromWorkSchedule(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workScheduleId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exceptionsIds,_ := vars["exceptionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkSchedule DAO
	//----------------------------------------------------------------------------
	requestResult := WorkScheduleDAO.RemoveExceptionsFromWorkSchedule(workScheduleId, exceptionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
