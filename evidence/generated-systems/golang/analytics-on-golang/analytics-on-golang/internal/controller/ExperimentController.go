package controller

import (
    ExperimentDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExperimentDAO for database creation
//----------------------------------------------------------------------------
func CreateExperiment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Experiment model
	//----------------------------------------------------------------------------
	data := model.Experiment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Experiment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.CreateExperiment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExperimentDAO to find the relevant Experiment
//----------------------------------------------------------------------------
func GetExperiment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Experiment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.GetExperiment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExperimentDAO for database read of all Experiments
//----------------------------------------------------------------------------
func GetAllExperiment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Experiment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.GetAllExperiment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExperimentDAO for database save
//----------------------------------------------------------------------------
func UpdateExperiment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Experiment model
	//----------------------------------------------------------------------------
	var data = model.Experiment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Experiment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.UpdateExperiment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExperimentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExperiment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Experiment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExperimentDAO.DeleteExperiment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a Experiment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToExperiment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.AssignWorkspaceToExperiment(experimentId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a Experiment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromExperiment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.UnassignWorkspaceFromExperiment(experimentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more trainingRunsIds as a TrainingRuns to a Experiment
	//----------------------------------------------------------------------------
func AddTrainingRunsToExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRunsIds,_ := vars["trainingRunsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.AddTrainingRunsToExperiment(experimentId, trainingRunsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more trainingRunsIds as a TrainingRuns from a Experiment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTrainingRunsFromExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRunsIds,_ := vars["trainingRunsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.RemoveTrainingRunsFromExperiment(experimentId, trainingRunsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a Experiment
	//----------------------------------------------------------------------------
func AddModelsToExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.AddModelsToExperiment(experimentId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a Experiment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.RemoveModelsFromExperiment(experimentId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more notebooksIds as a Notebooks to a Experiment
	//----------------------------------------------------------------------------
func AddNotebooksToExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notebooksIds,_ := vars["notebooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.AddNotebooksToExperiment(experimentId, notebooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more notebooksIds as a Notebooks from a Experiment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNotebooksFromExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notebooksIds,_ := vars["notebooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.RemoveNotebooksFromExperiment(experimentId, notebooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
