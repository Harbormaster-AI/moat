package controller

import (
    PositionDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PositionDAO for database creation
//----------------------------------------------------------------------------
func CreatePosition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Position model
	//----------------------------------------------------------------------------
	data := model.Position{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Position model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Position data access object to create
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.CreatePosition( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PositionDAO to find the relevant Position
//----------------------------------------------------------------------------
func GetPosition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Position data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.GetPosition(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PositionDAO for database read of all Positions
//----------------------------------------------------------------------------
func GetAllPosition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Position data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.GetAllPosition()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PositionDAO for database save
//----------------------------------------------------------------------------
func UpdatePosition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Position model
	//----------------------------------------------------------------------------
	var data = model.Position{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Position model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Position data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.UpdatePosition(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PositionDAO for database deletion
//----------------------------------------------------------------------------
func DeletePosition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Position data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PositionDAO.DeletePosition(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Department on a Position
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDepartmentToPosition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentId,_ := strconv.ParseUint( vars["departmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.AssignDepartmentToPosition(positionId, departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Department on a Position
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDepartmentFromPosition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.UnassignDepartmentFromPosition(positionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a JobProfile on a Position
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignJobProfileToPosition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfileId,_ := strconv.ParseUint( vars["jobProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.AssignJobProfileToPosition(positionId, jobProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a JobProfile on a Position
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignJobProfileFromPosition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.UnassignJobProfileFromPosition(positionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CostCenter on a Position
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCostCenterToPosition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	costCenterId,_ := strconv.ParseUint( vars["costCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.AssignCostCenterToPosition(positionId, costCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CostCenter on a Position
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCostCenterFromPosition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.UnassignCostCenterFromPosition(positionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a Position
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToPosition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.AssignLocationToPosition(positionId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a Position
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromPosition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.UnassignLocationFromPosition(positionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ManagerPosition on a Position
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignManagerPositionToPosition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	managerPositionId,_ := strconv.ParseUint( vars["managerPositionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.AssignManagerPositionToPosition(positionId, managerPositionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ManagerPosition on a Position
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignManagerPositionFromPosition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.UnassignManagerPositionFromPosition(positionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more directReportsIds as a DirectReports to a Position
	//----------------------------------------------------------------------------
func AddDirectReportsToPosition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	directReportsIds,_ := vars["directReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.AddDirectReportsToPosition(positionId, directReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more directReportsIds as a DirectReports from a Position
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDirectReportsFromPosition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	directReportsIds,_ := vars["directReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.RemoveDirectReportsFromPosition(positionId, directReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more assignmentsIds as a Assignments to a Position
	//----------------------------------------------------------------------------
func AddAssignmentsToPosition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignmentsIds,_ := vars["assignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.AddAssignmentsToPosition(positionId, assignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more assignmentsIds as a Assignments from a Position
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAssignmentsFromPosition(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	positionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assignmentsIds,_ := vars["assignmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Position DAO
	//----------------------------------------------------------------------------
	requestResult := PositionDAO.RemoveAssignmentsFromPosition(positionId, assignmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
