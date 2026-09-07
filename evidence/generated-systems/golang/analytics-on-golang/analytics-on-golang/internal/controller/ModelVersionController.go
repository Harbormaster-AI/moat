package controller

import (
    ModelVersionDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ModelVersionDAO for database creation
//----------------------------------------------------------------------------
func CreateModelVersion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ModelVersion model
	//----------------------------------------------------------------------------
	data := model.ModelVersion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ModelVersion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion data access object to create
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.CreateModelVersion( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ModelVersionDAO to find the relevant ModelVersion
//----------------------------------------------------------------------------
func GetModelVersion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ModelVersion data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.GetModelVersion(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ModelVersionDAO for database read of all ModelVersions
//----------------------------------------------------------------------------
func GetAllModelVersion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.GetAllModelVersion()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ModelVersionDAO for database save
//----------------------------------------------------------------------------
func UpdateModelVersion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ModelVersion model
	//----------------------------------------------------------------------------
	var data = model.ModelVersion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ModelVersion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.UpdateModelVersion(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ModelVersionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteModelVersion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ModelVersion data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ModelVersionDAO.DeleteModelVersion(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Model_ on a ModelVersion
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModel_ToModelVersion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	model_Id,_ := strconv.ParseUint( vars["model_Id"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.AssignModel_ToModelVersion(modelVersionId, model_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Model_ on a ModelVersion
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModel_FromModelVersion( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.UnassignModel_FromModelVersion(modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a TrainingRun on a ModelVersion
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTrainingRunToModelVersion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRunId,_ := strconv.ParseUint( vars["trainingRunId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.AssignTrainingRunToModelVersion(modelVersionId, trainingRunId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TrainingRun on a ModelVersion
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTrainingRunFromModelVersion( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.UnassignTrainingRunFromModelVersion(modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more evaluationMetricsIds as a EvaluationMetrics to a ModelVersion
	//----------------------------------------------------------------------------
func AddEvaluationMetricsToModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evaluationMetricsIds,_ := vars["evaluationMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.AddEvaluationMetricsToModelVersion(modelVersionId, evaluationMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more evaluationMetricsIds as a EvaluationMetrics from a ModelVersion
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEvaluationMetricsFromModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	evaluationMetricsIds,_ := vars["evaluationMetricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.RemoveEvaluationMetricsFromModelVersion(modelVersionId, evaluationMetricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more deploymentsIds as a Deployments to a ModelVersion
	//----------------------------------------------------------------------------
func AddDeploymentsToModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deploymentsIds,_ := vars["deploymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.AddDeploymentsToModelVersion(modelVersionId, deploymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more deploymentsIds as a Deployments from a ModelVersion
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDeploymentsFromModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deploymentsIds,_ := vars["deploymentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.RemoveDeploymentsFromModelVersion(modelVersionId, deploymentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more featureSetsIds as a FeatureSets to a ModelVersion
	//----------------------------------------------------------------------------
func AddFeatureSetsToModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.AddFeatureSetsToModelVersion(modelVersionId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featureSetsIds as a FeatureSets from a ModelVersion
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeatureSetsFromModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.RemoveFeatureSetsFromModelVersion(modelVersionId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a ModelVersion
	//----------------------------------------------------------------------------
func AddDatasetsToModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.AddDatasetsToModelVersion(modelVersionId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a ModelVersion
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromModelVersion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelVersionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ModelVersion DAO
	//----------------------------------------------------------------------------
	requestResult := ModelVersionDAO.RemoveDatasetsFromModelVersion(modelVersionId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
