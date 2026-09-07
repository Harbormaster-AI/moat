package controller

import (
    Model_DAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to Model_DAO for database creation
//----------------------------------------------------------------------------
func CreateModel_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Model_ model
	//----------------------------------------------------------------------------
	data := model.Model_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Model_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Model_ data access object to create
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.CreateModel_( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to Model_DAO to find the relevant Model_
//----------------------------------------------------------------------------
func GetModel_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Model_ data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.GetModel_(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to Model_DAO for database read of all Model_s
//----------------------------------------------------------------------------
func GetAllModel_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Model_ data access object to get all
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.GetAllModel_()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to Model_DAO for database save
//----------------------------------------------------------------------------
func UpdateModel_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Model_ model
	//----------------------------------------------------------------------------
	var data = model.Model_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Model_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Model_ data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.UpdateModel_(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to Model_DAO for database deletion
//----------------------------------------------------------------------------
func DeleteModel_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Model_ data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := Model_DAO.DeleteModel_(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a Model_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToModel_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.AssignWorkspaceToModel_(model_Id, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a Model_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromModel_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.UnassignWorkspaceFromModel_(model_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more versionsIds as a Versions to a Model_
	//----------------------------------------------------------------------------
func AddVersionsToModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	versionsIds,_ := vars["versionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.AddVersionsToModel_(model_Id, versionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more versionsIds as a Versions from a Model_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVersionsFromModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	versionsIds,_ := vars["versionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.RemoveVersionsFromModel_(model_Id, versionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more featureSetsIds as a FeatureSets to a Model_
	//----------------------------------------------------------------------------
func AddFeatureSetsToModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.AddFeatureSetsToModel_(model_Id, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featureSetsIds as a FeatureSets from a Model_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeatureSetsFromModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featureSetsIds,_ := vars["featureSetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.RemoveFeatureSetsFromModel_(model_Id, featureSetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more experimentsIds as a Experiments to a Model_
	//----------------------------------------------------------------------------
func AddExperimentsToModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.AddExperimentsToModel_(model_Id, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more experimentsIds as a Experiments from a Model_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveExperimentsFromModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentsIds,_ := vars["experimentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.RemoveExperimentsFromModel_(model_Id, experimentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tagsIds as a Tags to a Model_
	//----------------------------------------------------------------------------
func AddTagsToModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.AddTagsToModel_(model_Id, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tagsIds as a Tags from a Model_
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTagsFromModel_(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	model_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Model_ DAO
	//----------------------------------------------------------------------------
	requestResult := Model_DAO.RemoveTagsFromModel_(model_Id, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
