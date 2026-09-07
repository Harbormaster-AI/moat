package controller

import (
    RunParameterDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RunParameterDAO for database creation
//----------------------------------------------------------------------------
func CreateRunParameter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RunParameter model
	//----------------------------------------------------------------------------
	data := model.RunParameter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RunParameter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RunParameter data access object to create
	//----------------------------------------------------------------------------
	requestResult := RunParameterDAO.CreateRunParameter( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RunParameterDAO to find the relevant RunParameter
//----------------------------------------------------------------------------
func GetRunParameter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RunParameter data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RunParameterDAO.GetRunParameter(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RunParameterDAO for database read of all RunParameters
//----------------------------------------------------------------------------
func GetAllRunParameter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RunParameter data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RunParameterDAO.GetAllRunParameter()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RunParameterDAO for database save
//----------------------------------------------------------------------------
func UpdateRunParameter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RunParameter model
	//----------------------------------------------------------------------------
	var data = model.RunParameter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RunParameter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RunParameter data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RunParameterDAO.UpdateRunParameter(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RunParameterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRunParameter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RunParameter data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RunParameterDAO.DeleteRunParameter(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a TrainingRun on a RunParameter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTrainingRunToRunParameter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runParameterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRunId,_ := strconv.ParseUint( vars["trainingRunId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunParameter DAO
	//----------------------------------------------------------------------------
	requestResult := RunParameterDAO.AssignTrainingRunToRunParameter(runParameterId, trainingRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TrainingRun on a RunParameter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTrainingRunFromRunParameter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	runParameterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RunParameter DAO
	//----------------------------------------------------------------------------
	requestResult := RunParameterDAO.UnassignTrainingRunFromRunParameter(runParameterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


