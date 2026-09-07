package controller

import (
    InferenceEndpointDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InferenceEndpointDAO for database creation
//----------------------------------------------------------------------------
func CreateInferenceEndpoint(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InferenceEndpoint model
	//----------------------------------------------------------------------------
	data := model.InferenceEndpoint{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InferenceEndpoint model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint data access object to create
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.CreateInferenceEndpoint( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InferenceEndpointDAO to find the relevant InferenceEndpoint
//----------------------------------------------------------------------------
func GetInferenceEndpoint(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InferenceEndpoint data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.GetInferenceEndpoint(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InferenceEndpointDAO for database read of all InferenceEndpoints
//----------------------------------------------------------------------------
func GetAllInferenceEndpoint(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.GetAllInferenceEndpoint()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InferenceEndpointDAO for database save
//----------------------------------------------------------------------------
func UpdateInferenceEndpoint(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InferenceEndpoint model
	//----------------------------------------------------------------------------
	var data = model.InferenceEndpoint{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InferenceEndpoint model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.UpdateInferenceEndpoint(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InferenceEndpointDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInferenceEndpoint(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InferenceEndpoint data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InferenceEndpointDAO.DeleteInferenceEndpoint(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ModelVersion on a InferenceEndpoint
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModelVersionToInferenceEndpoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inferenceEndpointId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionId,_ := strconv.ParseUint( vars["modelVersionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint DAO
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.AssignModelVersionToInferenceEndpoint(inferenceEndpointId, modelVersionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ModelVersion on a InferenceEndpoint
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModelVersionFromInferenceEndpoint( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inferenceEndpointId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint DAO
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.UnassignModelVersionFromInferenceEndpoint(inferenceEndpointId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a InferenceEndpoint
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToInferenceEndpoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inferenceEndpointId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint DAO
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.AssignWorkspaceToInferenceEndpoint(inferenceEndpointId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a InferenceEndpoint
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromInferenceEndpoint( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inferenceEndpointId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint DAO
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.UnassignWorkspaceFromInferenceEndpoint(inferenceEndpointId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more predictionsIds as a Predictions to a InferenceEndpoint
	//----------------------------------------------------------------------------
func AddPredictionsToInferenceEndpoint(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inferenceEndpointId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	predictionsIds,_ := vars["predictionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint DAO
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.AddPredictionsToInferenceEndpoint(inferenceEndpointId, predictionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more predictionsIds as a Predictions from a InferenceEndpoint
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePredictionsFromInferenceEndpoint(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inferenceEndpointId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	predictionsIds,_ := vars["predictionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InferenceEndpoint DAO
	//----------------------------------------------------------------------------
	requestResult := InferenceEndpointDAO.RemovePredictionsFromInferenceEndpoint(inferenceEndpointId, predictionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
