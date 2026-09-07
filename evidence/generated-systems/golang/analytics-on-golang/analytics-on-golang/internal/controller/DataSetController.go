package controller

import (
    DataSetDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataSetDAO for database creation
//----------------------------------------------------------------------------
func CreateDataSet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataSet model
	//----------------------------------------------------------------------------
	data := model.DataSet{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataSet model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataSet data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.CreateDataSet( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataSetDAO to find the relevant DataSet
//----------------------------------------------------------------------------
func GetDataSet(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataSet data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.GetDataSet(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataSetDAO for database read of all DataSets
//----------------------------------------------------------------------------
func GetAllDataSet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataSet data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.GetAllDataSet()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataSetDAO for database save
//----------------------------------------------------------------------------
func UpdateDataSet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataSet model
	//----------------------------------------------------------------------------
	var data = model.DataSet{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataSet model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataSet data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.UpdateDataSet(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataSetDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataSet(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataSet data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataSetDAO.DeleteDataSet(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a DataSet
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToDataSet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AssignWorkspaceToDataSet(dataSetId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a DataSet
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromDataSet( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.UnassignWorkspaceFromDataSet(dataSetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LineageNode on a DataSet
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLineageNodeToDataSet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineageNodeId,_ := strconv.ParseUint( vars["lineageNodeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AssignLineageNodeToDataSet(dataSetId, lineageNodeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LineageNode on a DataSet
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLineageNodeFromDataSet( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.UnassignLineageNodeFromDataSet(dataSetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more sourcesIds as a Sources to a DataSet
	//----------------------------------------------------------------------------
func AddSourcesToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourcesIds,_ := vars["sourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddSourcesToDataSet(dataSetId, sourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more sourcesIds as a Sources from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSourcesFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourcesIds,_ := vars["sourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemoveSourcesFromDataSet(dataSetId, sourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more pipelinesIds as a Pipelines to a DataSet
	//----------------------------------------------------------------------------
func AddPipelinesToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddPipelinesToDataSet(dataSetId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more pipelinesIds as a Pipelines from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePipelinesFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemovePipelinesFromDataSet(dataSetId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more semanticModelsIds as a SemanticModels to a DataSet
	//----------------------------------------------------------------------------
func AddSemanticModelsToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	semanticModelsIds,_ := vars["semanticModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddSemanticModelsToDataSet(dataSetId, semanticModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more semanticModelsIds as a SemanticModels from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSemanticModelsFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	semanticModelsIds,_ := vars["semanticModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemoveSemanticModelsFromDataSet(dataSetId, semanticModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dimensionsIds as a Dimensions to a DataSet
	//----------------------------------------------------------------------------
func AddDimensionsToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddDimensionsToDataSet(dataSetId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dimensionsIds as a Dimensions from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDimensionsFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dimensionsIds,_ := vars["dimensionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemoveDimensionsFromDataSet(dataSetId, dimensionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more measuresIds as a Measures to a DataSet
	//----------------------------------------------------------------------------
func AddMeasuresToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	measuresIds,_ := vars["measuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddMeasuresToDataSet(dataSetId, measuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more measuresIds as a Measures from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMeasuresFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	measuresIds,_ := vars["measuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemoveMeasuresFromDataSet(dataSetId, measuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more metricsIds as a Metrics to a DataSet
	//----------------------------------------------------------------------------
func AddMetricsToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddMetricsToDataSet(dataSetId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more metricsIds as a Metrics from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMetricsFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	metricsIds,_ := vars["metricsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemoveMetricsFromDataSet(dataSetId, metricsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more qualityRulesIds as a QualityRules to a DataSet
	//----------------------------------------------------------------------------
func AddQualityRulesToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	qualityRulesIds,_ := vars["qualityRulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddQualityRulesToDataSet(dataSetId, qualityRulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more qualityRulesIds as a QualityRules from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQualityRulesFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	qualityRulesIds,_ := vars["qualityRulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemoveQualityRulesFromDataSet(dataSetId, qualityRulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tagsIds as a Tags to a DataSet
	//----------------------------------------------------------------------------
func AddTagsToDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.AddTagsToDataSet(dataSetId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tagsIds as a Tags from a DataSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTagsFromDataSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSet DAO
	//----------------------------------------------------------------------------
	requestResult := DataSetDAO.RemoveTagsFromDataSet(dataSetId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
