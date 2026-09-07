package controller

import (
    DataPipelineDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataPipelineDAO for database creation
//----------------------------------------------------------------------------
func CreateDataPipeline(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataPipeline model
	//----------------------------------------------------------------------------
	data := model.DataPipeline{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataPipeline model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.CreateDataPipeline( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataPipelineDAO to find the relevant DataPipeline
//----------------------------------------------------------------------------
func GetDataPipeline(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataPipeline data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.GetDataPipeline(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataPipelineDAO for database read of all DataPipelines
//----------------------------------------------------------------------------
func GetAllDataPipeline(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.GetAllDataPipeline()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataPipelineDAO for database save
//----------------------------------------------------------------------------
func UpdateDataPipeline(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataPipeline model
	//----------------------------------------------------------------------------
	var data = model.DataPipeline{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataPipeline model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.UpdateDataPipeline(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataPipelineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataPipeline(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataPipeline data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataPipelineDAO.DeleteDataPipeline(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a DataPipeline
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToDataPipeline(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.AssignWorkspaceToDataPipeline(dataPipelineId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a DataPipeline
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromDataPipeline( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.UnassignWorkspaceFromDataPipeline(dataPipelineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LineageNode on a DataPipeline
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLineageNodeToDataPipeline(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineageNodeId,_ := strconv.ParseUint( vars["lineageNodeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.AssignLineageNodeToDataPipeline(dataPipelineId, lineageNodeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LineageNode on a DataPipeline
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLineageNodeFromDataPipeline( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.UnassignLineageNodeFromDataPipeline(dataPipelineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more tasksIds as a Tasks to a DataPipeline
	//----------------------------------------------------------------------------
func AddTasksToDataPipeline(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tasksIds,_ := vars["tasksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.AddTasksToDataPipeline(dataPipelineId, tasksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tasksIds as a Tasks from a DataPipeline
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTasksFromDataPipeline(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tasksIds,_ := vars["tasksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.RemoveTasksFromDataPipeline(dataPipelineId, tasksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more sourcesIds as a Sources to a DataPipeline
	//----------------------------------------------------------------------------
func AddSourcesToDataPipeline(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourcesIds,_ := vars["sourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.AddSourcesToDataPipeline(dataPipelineId, sourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more sourcesIds as a Sources from a DataPipeline
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSourcesFromDataPipeline(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	sourcesIds,_ := vars["sourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.RemoveSourcesFromDataPipeline(dataPipelineId, sourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more outputsIds as a Outputs to a DataPipeline
	//----------------------------------------------------------------------------
func AddOutputsToDataPipeline(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outputsIds,_ := vars["outputsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.AddOutputsToDataPipeline(dataPipelineId, outputsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more outputsIds as a Outputs from a DataPipeline
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOutputsFromDataPipeline(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataPipelineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outputsIds,_ := vars["outputsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataPipeline DAO
	//----------------------------------------------------------------------------
	requestResult := DataPipelineDAO.RemoveOutputsFromDataPipeline(dataPipelineId, outputsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
