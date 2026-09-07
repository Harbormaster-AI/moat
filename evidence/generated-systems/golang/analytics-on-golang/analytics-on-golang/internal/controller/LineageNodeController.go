package controller

import (
    LineageNodeDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LineageNodeDAO for database creation
//----------------------------------------------------------------------------
func CreateLineageNode(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LineageNode model
	//----------------------------------------------------------------------------
	data := model.LineageNode{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LineageNode model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode data access object to create
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.CreateLineageNode( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LineageNodeDAO to find the relevant LineageNode
//----------------------------------------------------------------------------
func GetLineageNode(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LineageNode data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.GetLineageNode(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LineageNodeDAO for database read of all LineageNodes
//----------------------------------------------------------------------------
func GetAllLineageNode(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LineageNode data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.GetAllLineageNode()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LineageNodeDAO for database save
//----------------------------------------------------------------------------
func UpdateLineageNode(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LineageNode model
	//----------------------------------------------------------------------------
	var data = model.LineageNode{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LineageNode model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.UpdateLineageNode(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LineageNodeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLineageNode(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LineageNode data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LineageNodeDAO.DeleteLineageNode(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a LineageNode
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToLineageNode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AssignWorkspaceToLineageNode(lineageNodeId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a LineageNode
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromLineageNode( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.UnassignWorkspaceFromLineageNode(lineageNodeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more inputsIds as a Inputs to a LineageNode
	//----------------------------------------------------------------------------
func AddInputsToLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inputsIds,_ := vars["inputsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AddInputsToLineageNode(lineageNodeId, inputsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inputsIds as a Inputs from a LineageNode
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInputsFromLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inputsIds,_ := vars["inputsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.RemoveInputsFromLineageNode(lineageNodeId, inputsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more outputsIds as a Outputs to a LineageNode
	//----------------------------------------------------------------------------
func AddOutputsToLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outputsIds,_ := vars["outputsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AddOutputsToLineageNode(lineageNodeId, outputsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more outputsIds as a Outputs from a LineageNode
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOutputsFromLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outputsIds,_ := vars["outputsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.RemoveOutputsFromLineageNode(lineageNodeId, outputsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a LineageNode
	//----------------------------------------------------------------------------
func AddDatasetsToLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AddDatasetsToLineageNode(lineageNodeId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a LineageNode
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.RemoveDatasetsFromLineageNode(lineageNodeId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a LineageNode
	//----------------------------------------------------------------------------
func AddModelsToLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AddModelsToLineageNode(lineageNodeId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a LineageNode
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.RemoveModelsFromLineageNode(lineageNodeId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more pipelinesIds as a Pipelines to a LineageNode
	//----------------------------------------------------------------------------
func AddPipelinesToLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AddPipelinesToLineageNode(lineageNodeId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more pipelinesIds as a Pipelines from a LineageNode
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePipelinesFromLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.RemovePipelinesFromLineageNode(lineageNodeId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dashboardsIds as a Dashboards to a LineageNode
	//----------------------------------------------------------------------------
func AddDashboardsToLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AddDashboardsToLineageNode(lineageNodeId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dashboardsIds as a Dashboards from a LineageNode
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDashboardsFromLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.RemoveDashboardsFromLineageNode(lineageNodeId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a LineageNode
	//----------------------------------------------------------------------------
func AddReportsToLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.AddReportsToLineageNode(lineageNodeId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a LineageNode
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromLineageNode(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	lineageNodeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LineageNode DAO
	//----------------------------------------------------------------------------
	requestResult := LineageNodeDAO.RemoveReportsFromLineageNode(lineageNodeId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
