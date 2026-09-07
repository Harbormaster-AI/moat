package controller

import (
    DataSourceDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataSourceDAO for database creation
//----------------------------------------------------------------------------
func CreateDataSource(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataSource model
	//----------------------------------------------------------------------------
	data := model.DataSource{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataSource model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataSource data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.CreateDataSource( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataSourceDAO to find the relevant DataSource
//----------------------------------------------------------------------------
func GetDataSource(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataSource data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.GetDataSource(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataSourceDAO for database read of all DataSources
//----------------------------------------------------------------------------
func GetAllDataSource(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataSource data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.GetAllDataSource()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataSourceDAO for database save
//----------------------------------------------------------------------------
func UpdateDataSource(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataSource model
	//----------------------------------------------------------------------------
	var data = model.DataSource{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataSource model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataSource data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.UpdateDataSource(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataSourceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataSource(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataSource data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataSourceDAO.DeleteDataSource(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a DataSource
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToDataSource(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSource DAO
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.AssignWorkspaceToDataSource(dataSourceId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a DataSource
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromDataSource( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataSourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataSource DAO
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.UnassignWorkspaceFromDataSource(dataSourceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more producedDatasetsIds as a ProducedDatasets to a DataSource
	//----------------------------------------------------------------------------
func AddProducedDatasetsToDataSource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	producedDatasetsIds,_ := vars["producedDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSource DAO
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.AddProducedDatasetsToDataSource(dataSourceId, producedDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more producedDatasetsIds as a ProducedDatasets from a DataSource
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProducedDatasetsFromDataSource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	producedDatasetsIds,_ := vars["producedDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSource DAO
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.RemoveProducedDatasetsFromDataSource(dataSourceId, producedDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more pipelinesIds as a Pipelines to a DataSource
	//----------------------------------------------------------------------------
func AddPipelinesToDataSource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSource DAO
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.AddPipelinesToDataSource(dataSourceId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more pipelinesIds as a Pipelines from a DataSource
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePipelinesFromDataSource(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataSourceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelinesIds,_ := vars["pipelinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataSource DAO
	//----------------------------------------------------------------------------
	requestResult := DataSourceDAO.RemovePipelinesFromDataSource(dataSourceId, pipelinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
