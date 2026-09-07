package controller

import (
    ModelDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ModelDAO for database creation
//----------------------------------------------------------------------------
func CreateModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Model model
	//----------------------------------------------------------------------------
	data := model.Model{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Model model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Model data access object to create
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.CreateModel( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ModelDAO to find the relevant Model
//----------------------------------------------------------------------------
func GetModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Model data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.GetModel(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ModelDAO for database read of all Models
//----------------------------------------------------------------------------
func GetAllModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Model data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.GetAllModel()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ModelDAO for database save
//----------------------------------------------------------------------------
func UpdateModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Model model
	//----------------------------------------------------------------------------
	var data = model.Model{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Model model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Model data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.UpdateModel(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ModelDAO for database deletion
//----------------------------------------------------------------------------
func DeleteModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Model data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ModelDAO.DeleteModel(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a Model
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToModel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.AssignWorkspaceToModel(modelId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a Model
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromModel( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.UnassignWorkspaceFromModel(modelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more versionsIds as a Versions to a Model
	//----------------------------------------------------------------------------
func AddVersionsToModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	versionsIds,_ := vars["versionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.AddVersionsToModel(modelId, versionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more versionsIds as a Versions from a Model
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVersionsFromModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	versionsIds,_ := vars["versionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.RemoveVersionsFromModel(modelId, versionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more featureSetsIds as a FeatureSets to a Model
	//----------------------------------------------------------------------------
func AddFeatureSetsToModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.AddFeatureSetsToModel(modelId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featureSetsIds as a FeatureSets from a Model
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeatureSetsFromModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.RemoveFeatureSetsFromModel(modelId, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more experimentsIds as a Experiments to a Model
	//----------------------------------------------------------------------------
func AddExperimentsToModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.AddExperimentsToModel(modelId, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more experimentsIds as a Experiments from a Model
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExperimentsFromModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.RemoveExperimentsFromModel(modelId, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tagsIds as a Tags to a Model
	//----------------------------------------------------------------------------
func AddTagsToModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.AddTagsToModel(modelId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tagsIds as a Tags from a Model
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTagsFromModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	modelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model DAO
	//----------------------------------------------------------------------------
	requestResult := ModelDAO.RemoveTagsFromModel(modelId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
