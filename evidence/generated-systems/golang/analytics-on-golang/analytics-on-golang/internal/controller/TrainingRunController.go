package controller

import (
    TrainingRunDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TrainingRunDAO for database creation
//----------------------------------------------------------------------------
func CreateTrainingRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrainingRun model
	//----------------------------------------------------------------------------
	data := model.TrainingRun{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrainingRun model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun data access object to create
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.CreateTrainingRun( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TrainingRunDAO to find the relevant TrainingRun
//----------------------------------------------------------------------------
func GetTrainingRun(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrainingRun data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.GetTrainingRun(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TrainingRunDAO for database read of all TrainingRuns
//----------------------------------------------------------------------------
func GetAllTrainingRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.GetAllTrainingRun()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TrainingRunDAO for database save
//----------------------------------------------------------------------------
func UpdateTrainingRun(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TrainingRun model
	//----------------------------------------------------------------------------
	var data = model.TrainingRun{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TrainingRun model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.UpdateTrainingRun(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TrainingRunDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTrainingRun(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TrainingRun data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TrainingRunDAO.DeleteTrainingRun(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Experiment on a TrainingRun
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignExperimentToTrainingRun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentId,_ := strconv.ParseUint( vars["experimentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.AssignExperimentToTrainingRun(trainingRunId, experimentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Experiment on a TrainingRun
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignExperimentFromTrainingRun( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.UnassignExperimentFromTrainingRun(trainingRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ModelVersion on a TrainingRun
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModelVersionToTrainingRun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionId,_ := strconv.ParseUint( vars["modelVersionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.AssignModelVersionToTrainingRun(trainingRunId, modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ModelVersion on a TrainingRun
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModelVersionFromTrainingRun( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.UnassignModelVersionFromTrainingRun(trainingRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more inputDatasetsIds as a InputDatasets to a TrainingRun
	//----------------------------------------------------------------------------
func AddInputDatasetsToTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inputDatasetsIds,_ := vars["inputDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.AddInputDatasetsToTrainingRun(trainingRunId, inputDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inputDatasetsIds as a InputDatasets from a TrainingRun
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInputDatasetsFromTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inputDatasetsIds,_ := vars["inputDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.RemoveInputDatasetsFromTrainingRun(trainingRunId, inputDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more featuresIds as a Features to a TrainingRun
	//----------------------------------------------------------------------------
func AddFeaturesToTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featuresIds,_ := vars["featuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.AddFeaturesToTrainingRun(trainingRunId, featuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featuresIds as a Features from a TrainingRun
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeaturesFromTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featuresIds,_ := vars["featuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.RemoveFeaturesFromTrainingRun(trainingRunId, featuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more runMetricsIds as a RunMetrics to a TrainingRun
	//----------------------------------------------------------------------------
func AddRunMetricsToTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	runMetricsIds,_ := vars["runMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.AddRunMetricsToTrainingRun(trainingRunId, runMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more runMetricsIds as a RunMetrics from a TrainingRun
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRunMetricsFromTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	runMetricsIds,_ := vars["runMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.RemoveRunMetricsFromTrainingRun(trainingRunId, runMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more runParametersIds as a RunParameters to a TrainingRun
	//----------------------------------------------------------------------------
func AddRunParametersToTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	runParametersIds,_ := vars["runParametersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.AddRunParametersToTrainingRun(trainingRunId, runParametersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more runParametersIds as a RunParameters from a TrainingRun
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRunParametersFromTrainingRun(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	trainingRunId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	runParametersIds,_ := vars["runParametersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TrainingRun DAO
	//----------------------------------------------------------------------------
	requestResult := TrainingRunDAO.RemoveRunParametersFromTrainingRun(trainingRunId, runParametersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
