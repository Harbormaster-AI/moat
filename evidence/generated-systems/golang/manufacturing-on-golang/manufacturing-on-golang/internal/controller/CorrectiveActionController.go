package controller

import (
    CorrectiveActionDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CorrectiveActionDAO for database creation
//----------------------------------------------------------------------------
func CreateCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CorrectiveAction model
	//----------------------------------------------------------------------------
	data := model.CorrectiveAction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CorrectiveAction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction data access object to create
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.CreateCorrectiveAction( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CorrectiveActionDAO to find the relevant CorrectiveAction
//----------------------------------------------------------------------------
func GetCorrectiveAction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CorrectiveAction data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.GetCorrectiveAction(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CorrectiveActionDAO for database read of all CorrectiveActions
//----------------------------------------------------------------------------
func GetAllCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.GetAllCorrectiveAction()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CorrectiveActionDAO for database save
//----------------------------------------------------------------------------
func UpdateCorrectiveAction(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CorrectiveAction model
	//----------------------------------------------------------------------------
	var data = model.CorrectiveAction{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CorrectiveAction model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.UpdateCorrectiveAction(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CorrectiveActionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCorrectiveAction(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CorrectiveAction data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CorrectiveActionDAO.DeleteCorrectiveAction(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Nonconformance on a CorrectiveAction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignNonconformanceToCorrectiveAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	nonconformanceId,_ := strconv.ParseUint( vars["nonconformanceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.AssignNonconformanceToCorrectiveAction(correctiveActionId, nonconformanceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Nonconformance on a CorrectiveAction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignNonconformanceFromCorrectiveAction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.UnassignNonconformanceFromCorrectiveAction(correctiveActionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Owner on a CorrectiveAction
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOwnerToCorrectiveAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ownerId,_ := strconv.ParseUint( vars["ownerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.AssignOwnerToCorrectiveAction(correctiveActionId, ownerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Owner on a CorrectiveAction
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOwnerFromCorrectiveAction( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	correctiveActionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CorrectiveAction DAO
	//----------------------------------------------------------------------------
	requestResult := CorrectiveActionDAO.UnassignOwnerFromCorrectiveAction(correctiveActionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


