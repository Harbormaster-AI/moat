package controller

import (
    ShiftDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ShiftDAO for database creation
//----------------------------------------------------------------------------
func CreateShift(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Shift model
	//----------------------------------------------------------------------------
	data := model.Shift{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Shift model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Shift data access object to create
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.CreateShift( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ShiftDAO to find the relevant Shift
//----------------------------------------------------------------------------
func GetShift(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Shift data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.GetShift(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ShiftDAO for database read of all Shifts
//----------------------------------------------------------------------------
func GetAllShift(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Shift data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.GetAllShift()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ShiftDAO for database save
//----------------------------------------------------------------------------
func UpdateShift(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Shift model
	//----------------------------------------------------------------------------
	var data = model.Shift{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Shift model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Shift data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.UpdateShift(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ShiftDAO for database deletion
//----------------------------------------------------------------------------
func DeleteShift(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Shift data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ShiftDAO.DeleteShift(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Plant on a Shift
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToShift(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Shift DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.AssignPlantToShift(shiftId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a Shift
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromShift( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	shiftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Shift DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.UnassignPlantFromShift(shiftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more assignmentsIds as a Assignments to a Shift
	//----------------------------------------------------------------------------
func AddAssignmentsToShift(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	shiftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignmentsIds,_ := vars["assignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Shift DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.AddAssignmentsToShift(shiftId, assignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more assignmentsIds as a Assignments from a Shift
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAssignmentsFromShift(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	shiftId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignmentsIds,_ := vars["assignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Shift DAO
	//----------------------------------------------------------------------------
	requestResult := ShiftDAO.RemoveAssignmentsFromShift(shiftId, assignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
