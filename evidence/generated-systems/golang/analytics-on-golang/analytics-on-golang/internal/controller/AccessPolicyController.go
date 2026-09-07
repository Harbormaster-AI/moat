package controller

import (
    AccessPolicyDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AccessPolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateAccessPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AccessPolicy model
	//----------------------------------------------------------------------------
	data := model.AccessPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AccessPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.CreateAccessPolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AccessPolicyDAO to find the relevant AccessPolicy
//----------------------------------------------------------------------------
func GetAccessPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AccessPolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.GetAccessPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AccessPolicyDAO for database read of all AccessPolicys
//----------------------------------------------------------------------------
func GetAllAccessPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.GetAllAccessPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AccessPolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateAccessPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AccessPolicy model
	//----------------------------------------------------------------------------
	var data = model.AccessPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AccessPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.UpdateAccessPolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AccessPolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAccessPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AccessPolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AccessPolicyDAO.DeleteAccessPolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a AccessPolicy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToAccessPolicy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AssignWorkspaceToAccessPolicy(accessPolicyId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a AccessPolicy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromAccessPolicy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.UnassignWorkspaceFromAccessPolicy(accessPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a AccessPolicy
	//----------------------------------------------------------------------------
func AddDatasetsToAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AddDatasetsToAccessPolicy(accessPolicyId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a AccessPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.RemoveDatasetsFromAccessPolicy(accessPolicyId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dashboardsIds as a Dashboards to a AccessPolicy
	//----------------------------------------------------------------------------
func AddDashboardsToAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AddDashboardsToAccessPolicy(accessPolicyId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dashboardsIds as a Dashboards from a AccessPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDashboardsFromAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dashboardsIds,_ := vars["dashboardsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.RemoveDashboardsFromAccessPolicy(accessPolicyId, dashboardsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more reportsIds as a Reports to a AccessPolicy
	//----------------------------------------------------------------------------
func AddReportsToAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AddReportsToAccessPolicy(accessPolicyId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reportsIds as a Reports from a AccessPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReportsFromAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reportsIds,_ := vars["reportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.RemoveReportsFromAccessPolicy(accessPolicyId, reportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a AccessPolicy
	//----------------------------------------------------------------------------
func AddModelsToAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AddModelsToAccessPolicy(accessPolicyId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a AccessPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.RemoveModelsFromAccessPolicy(accessPolicyId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more featureSetsIds as a FeatureSets to a AccessPolicy
	//----------------------------------------------------------------------------
func AddFeatureSetsToAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.AddFeatureSetsToAccessPolicy(accessPolicyId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featureSetsIds as a FeatureSets from a AccessPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeatureSetsFromAccessPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	accessPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AccessPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := AccessPolicyDAO.RemoveFeatureSetsFromAccessPolicy(accessPolicyId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
