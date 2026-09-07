package controller

import (
    BIQueryDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BIQueryDAO for database creation
//----------------------------------------------------------------------------
func CreateBIQuery(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BIQuery model
	//----------------------------------------------------------------------------
	data := model.BIQuery{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BIQuery model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery data access object to create
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.CreateBIQuery( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BIQueryDAO to find the relevant BIQuery
//----------------------------------------------------------------------------
func GetBIQuery(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BIQuery data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.GetBIQuery(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BIQueryDAO for database read of all BIQuerys
//----------------------------------------------------------------------------
func GetAllBIQuery(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BIQuery data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.GetAllBIQuery()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BIQueryDAO for database save
//----------------------------------------------------------------------------
func UpdateBIQuery(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BIQuery model
	//----------------------------------------------------------------------------
	var data = model.BIQuery{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BIQuery model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.UpdateBIQuery(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BIQueryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBIQuery(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BIQuery data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BIQueryDAO.DeleteBIQuery(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a BIQuery
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToBIQuery(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.AssignWorkspaceToBIQuery(bIQueryId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a BIQuery
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromBIQuery( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.UnassignWorkspaceFromBIQuery(bIQueryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a BIQuery
	//----------------------------------------------------------------------------
func AddDatasetsToBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.AddDatasetsToBIQuery(bIQueryId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a BIQuery
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.RemoveDatasetsFromBIQuery(bIQueryId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a BIQuery
	//----------------------------------------------------------------------------
func AddReportsToBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.AddReportsToBIQuery(bIQueryId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a BIQuery
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.RemoveReportsFromBIQuery(bIQueryId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dashboardsIds as a Dashboards to a BIQuery
	//----------------------------------------------------------------------------
func AddDashboardsToBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.AddDashboardsToBIQuery(bIQueryId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dashboardsIds as a Dashboards from a BIQuery
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDashboardsFromBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.RemoveDashboardsFromBIQuery(bIQueryId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more notebooksIds as a Notebooks to a BIQuery
	//----------------------------------------------------------------------------
func AddNotebooksToBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notebooksIds,_ := vars["notebooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.AddNotebooksToBIQuery(bIQueryId, notebooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more notebooksIds as a Notebooks from a BIQuery
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNotebooksFromBIQuery(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bIQueryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	notebooksIds,_ := vars["notebooksIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BIQuery DAO
	//----------------------------------------------------------------------------
	requestResult := BIQueryDAO.RemoveNotebooksFromBIQuery(bIQueryId, notebooksIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
