package controller

import (
    GoalDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to GoalDAO for database creation
//----------------------------------------------------------------------------
func CreateGoal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Goal model
	//----------------------------------------------------------------------------
	data := model.Goal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Goal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Goal data access object to create
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.CreateGoal( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to GoalDAO to find the relevant Goal
//----------------------------------------------------------------------------
func GetGoal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Goal data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.GetGoal(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to GoalDAO for database read of all Goals
//----------------------------------------------------------------------------
func GetAllGoal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Goal data access object to get all
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.GetAllGoal()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to GoalDAO for database save
//----------------------------------------------------------------------------
func UpdateGoal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Goal model
	//----------------------------------------------------------------------------
	var data = model.Goal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Goal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Goal data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.UpdateGoal(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to GoalDAO for database deletion
//----------------------------------------------------------------------------
func DeleteGoal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Goal data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := GoalDAO.DeleteGoal(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a Goal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToGoal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.AssignEmployeeToGoal(goalId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a Goal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromGoal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.UnassignEmployeeFromGoal(goalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Cycle on a Goal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCycleToGoal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cycleId,_ := strconv.ParseUint( vars["cycleId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.AssignCycleToGoal(goalId, cycleId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Cycle on a Goal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCycleFromGoal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.UnassignCycleFromGoal(goalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ParentGoal on a Goal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignParentGoalToGoal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	parentGoalId,_ := strconv.ParseUint( vars["parentGoalId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.AssignParentGoalToGoal(goalId, parentGoalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ParentGoal on a Goal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignParentGoalFromGoal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.UnassignParentGoalFromGoal(goalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more childGoalsIds as a ChildGoals to a Goal
	//----------------------------------------------------------------------------
func AddChildGoalsToGoal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childGoalsIds,_ := vars["childGoalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.AddChildGoalsToGoal(goalId, childGoalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more childGoalsIds as a ChildGoals from a Goal
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveChildGoalsFromGoal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	goalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childGoalsIds,_ := vars["childGoalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Goal DAO
	//----------------------------------------------------------------------------
	requestResult := GoalDAO.RemoveChildGoalsFromGoal(goalId, childGoalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
