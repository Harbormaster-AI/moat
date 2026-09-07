package controller

import (
    NotebookDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to NotebookDAO for database creation
//----------------------------------------------------------------------------
func CreateNotebook(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Notebook model
	//----------------------------------------------------------------------------
	data := model.Notebook{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Notebook model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Notebook data access object to create
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.CreateNotebook( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to NotebookDAO to find the relevant Notebook
//----------------------------------------------------------------------------
func GetNotebook(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Notebook data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.GetNotebook(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to NotebookDAO for database read of all Notebooks
//----------------------------------------------------------------------------
func GetAllNotebook(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Notebook data access object to get all
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.GetAllNotebook()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to NotebookDAO for database save
//----------------------------------------------------------------------------
func UpdateNotebook(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Notebook model
	//----------------------------------------------------------------------------
	var data = model.Notebook{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Notebook model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Notebook data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.UpdateNotebook(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to NotebookDAO for database deletion
//----------------------------------------------------------------------------
func DeleteNotebook(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Notebook data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := NotebookDAO.DeleteNotebook(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a Notebook
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToNotebook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.AssignWorkspaceToNotebook(notebookId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a Notebook
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromNotebook( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.UnassignWorkspaceFromNotebook(notebookId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Notebook
	//----------------------------------------------------------------------------
func AddDatasetsToNotebook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.AddDatasetsToNotebook(notebookId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Notebook
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromNotebook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.RemoveDatasetsFromNotebook(notebookId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more experimentsIds as a Experiments to a Notebook
	//----------------------------------------------------------------------------
func AddExperimentsToNotebook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.AddExperimentsToNotebook(notebookId, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more experimentsIds as a Experiments from a Notebook
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExperimentsFromNotebook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.RemoveExperimentsFromNotebook(notebookId, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more queriesIds as a Queries to a Notebook
	//----------------------------------------------------------------------------
func AddQueriesToNotebook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	queriesIds,_ := vars["queriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.AddQueriesToNotebook(notebookId, queriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more queriesIds as a Queries from a Notebook
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQueriesFromNotebook(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	notebookId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	queriesIds,_ := vars["queriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Notebook DAO
	//----------------------------------------------------------------------------
	requestResult := NotebookDAO.RemoveQueriesFromNotebook(notebookId, queriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
