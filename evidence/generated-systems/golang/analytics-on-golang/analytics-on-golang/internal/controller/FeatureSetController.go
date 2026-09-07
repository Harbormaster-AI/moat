package controller

import (
    FeatureSetDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FeatureSetDAO for database creation
//----------------------------------------------------------------------------
func CreateFeatureSet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FeatureSet model
	//----------------------------------------------------------------------------
	data := model.FeatureSet{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FeatureSet model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet data access object to create
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.CreateFeatureSet( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FeatureSetDAO to find the relevant FeatureSet
//----------------------------------------------------------------------------
func GetFeatureSet(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FeatureSet data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.GetFeatureSet(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FeatureSetDAO for database read of all FeatureSets
//----------------------------------------------------------------------------
func GetAllFeatureSet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.GetAllFeatureSet()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FeatureSetDAO for database save
//----------------------------------------------------------------------------
func UpdateFeatureSet(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty FeatureSet model
	//----------------------------------------------------------------------------
	var data = model.FeatureSet{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a FeatureSet model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.UpdateFeatureSet(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FeatureSetDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFeatureSet(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the FeatureSet data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FeatureSetDAO.DeleteFeatureSet(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Workspace on a FeatureSet
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkspaceToFeatureSet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workspaceId,_ := strconv.ParseUint( vars["workspaceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.AssignWorkspaceToFeatureSet(featureSetId, workspaceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workspace on a FeatureSet
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkspaceFromFeatureSet( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.UnassignWorkspaceFromFeatureSet(featureSetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more featuresIds as a Features to a FeatureSet
	//----------------------------------------------------------------------------
func AddFeaturesToFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featuresIds,_ := vars["featuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.AddFeaturesToFeatureSet(featureSetId, featuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more featuresIds as a Features from a FeatureSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFeaturesFromFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	featuresIds,_ := vars["featuresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.RemoveFeaturesFromFeatureSet(featureSetId, featuresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a FeatureSet
	//----------------------------------------------------------------------------
func AddDatasetsToFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.AddDatasetsToFeatureSet(featureSetId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a FeatureSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.RemoveDatasetsFromFeatureSet(featureSetId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelsIds as a Models to a FeatureSet
	//----------------------------------------------------------------------------
func AddModelsToFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.AddModelsToFeatureSet(featureSetId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelsIds as a Models from a FeatureSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelsFromFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelsIds,_ := vars["modelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.RemoveModelsFromFeatureSet(featureSetId, modelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more modelVersionsIds as a ModelVersions to a FeatureSet
	//----------------------------------------------------------------------------
func AddModelVersionsToFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionsIds,_ := vars["modelVersionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.AddModelVersionsToFeatureSet(featureSetId, modelVersionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more modelVersionsIds as a ModelVersions from a FeatureSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveModelVersionsFromFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	modelVersionsIds,_ := vars["modelVersionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.RemoveModelVersionsFromFeatureSet(featureSetId, modelVersionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tagsIds as a Tags to a FeatureSet
	//----------------------------------------------------------------------------
func AddTagsToFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.AddTagsToFeatureSet(featureSetId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tagsIds as a Tags from a FeatureSet
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTagsFromFeatureSet(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	featureSetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tagsIds,_ := vars["tagsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the FeatureSet DAO
	//----------------------------------------------------------------------------
	requestResult := FeatureSetDAO.RemoveTagsFromFeatureSet(featureSetId, tagsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
