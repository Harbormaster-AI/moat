package controller

import (
    DataTaskDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataTaskDAO for database creation
//----------------------------------------------------------------------------
func CreateDataTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataTask model
	//----------------------------------------------------------------------------
	data := model.DataTask{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataTask model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataTask data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.CreateDataTask( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataTaskDAO to find the relevant DataTask
//----------------------------------------------------------------------------
func GetDataTask(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataTask data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.GetDataTask(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataTaskDAO for database read of all DataTasks
//----------------------------------------------------------------------------
func GetAllDataTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataTask data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.GetAllDataTask()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataTaskDAO for database save
//----------------------------------------------------------------------------
func UpdateDataTask(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataTask model
	//----------------------------------------------------------------------------
	var data = model.DataTask{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataTask model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataTask data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.UpdateDataTask(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataTaskDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataTask(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataTask data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataTaskDAO.DeleteDataTask(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Pipeline on a DataTask
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPipelineToDataTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pipelineId,_ := strconv.ParseUint( vars["pipelineId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataTask DAO
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.AssignPipelineToDataTask(dataTaskId, pipelineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Pipeline on a DataTask
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPipelineFromDataTask( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataTask DAO
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.UnassignPipelineFromDataTask(dataTaskId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more inputDatasetsIds as a InputDatasets to a DataTask
	//----------------------------------------------------------------------------
func AddInputDatasetsToDataTask(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inputDatasetsIds,_ := vars["inputDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataTask DAO
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.AddInputDatasetsToDataTask(dataTaskId, inputDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inputDatasetsIds as a InputDatasets from a DataTask
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInputDatasetsFromDataTask(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inputDatasetsIds,_ := vars["inputDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataTask DAO
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.RemoveInputDatasetsFromDataTask(dataTaskId, inputDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more outputDatasetsIds as a OutputDatasets to a DataTask
	//----------------------------------------------------------------------------
func AddOutputDatasetsToDataTask(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outputDatasetsIds,_ := vars["outputDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataTask DAO
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.AddOutputDatasetsToDataTask(dataTaskId, outputDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more outputDatasetsIds as a OutputDatasets from a DataTask
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOutputDatasetsFromDataTask(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataTaskId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	outputDatasetsIds,_ := vars["outputDatasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataTask DAO
	//----------------------------------------------------------------------------
	requestResult := DataTaskDAO.RemoveOutputDatasetsFromDataTask(dataTaskId, outputDatasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
