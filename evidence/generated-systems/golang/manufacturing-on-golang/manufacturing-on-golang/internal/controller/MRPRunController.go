package controller

import (
    MRPRunDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MRPRunDAO for database creation
//----------------------------------------------------------------------------
func CreateMRPRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MRPRun model
	//----------------------------------------------------------------------------
	data := model.MRPRun{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MRPRun model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MRPRun data access object to create
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.CreateMRPRun( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MRPRunDAO to find the relevant MRPRun
//----------------------------------------------------------------------------
func GetMRPRun(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MRPRun data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.GetMRPRun(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MRPRunDAO for database read of all MRPRuns
//----------------------------------------------------------------------------
func GetAllMRPRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MRPRun data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.GetAllMRPRun()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MRPRunDAO for database save
//----------------------------------------------------------------------------
func UpdateMRPRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MRPRun model
	//----------------------------------------------------------------------------
	var data = model.MRPRun{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MRPRun model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MRPRun data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.UpdateMRPRun(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MRPRunDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMRPRun(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MRPRun data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MRPRunDAO.DeleteMRPRun(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Plant on a MRPRun
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToMRPRun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	mRPRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MRPRun DAO
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.AssignPlantToMRPRun(mRPRunId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a MRPRun
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromMRPRun( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	mRPRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MRPRun DAO
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.UnassignPlantFromMRPRun(mRPRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more plannedOrdersIds as a PlannedOrders to a MRPRun
	//----------------------------------------------------------------------------
func AddPlannedOrdersToMRPRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	mRPRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plannedOrdersIds,_ := vars["plannedOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MRPRun DAO
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.AddPlannedOrdersToMRPRun(mRPRunId, plannedOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more plannedOrdersIds as a PlannedOrders from a MRPRun
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlannedOrdersFromMRPRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	mRPRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plannedOrdersIds,_ := vars["plannedOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MRPRun DAO
	//----------------------------------------------------------------------------
	requestResult := MRPRunDAO.RemovePlannedOrdersFromMRPRun(mRPRunId, plannedOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
