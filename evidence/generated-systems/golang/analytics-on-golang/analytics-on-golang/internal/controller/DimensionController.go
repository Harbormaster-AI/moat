package controller

import (
    DimensionDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DimensionDAO for database creation
//----------------------------------------------------------------------------
func CreateDimension(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dimension model
	//----------------------------------------------------------------------------
	data := model.Dimension{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dimension model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dimension data access object to create
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.CreateDimension( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DimensionDAO to find the relevant Dimension
//----------------------------------------------------------------------------
func GetDimension(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dimension data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.GetDimension(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DimensionDAO for database read of all Dimensions
//----------------------------------------------------------------------------
func GetAllDimension(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Dimension data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.GetAllDimension()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DimensionDAO for database save
//----------------------------------------------------------------------------
func UpdateDimension(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Dimension model
	//----------------------------------------------------------------------------
	var data = model.Dimension{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Dimension model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Dimension data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.UpdateDimension(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DimensionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDimension(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Dimension data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DimensionDAO.DeleteDimension(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a SemanticModel on a Dimension
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSemanticModelToDimension(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dimensionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	semanticModelId,_ := strconv.ParseUint( vars["semanticModelId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dimension DAO
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.AssignSemanticModelToDimension(dimensionId, semanticModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SemanticModel on a Dimension
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSemanticModelFromDimension( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dimensionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Dimension DAO
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.UnassignSemanticModelFromDimension(dimensionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Dimension
	//----------------------------------------------------------------------------
func AddDatasetsToDimension(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dimensionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dimension DAO
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.AddDatasetsToDimension(dimensionId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Dimension
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromDimension(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dimensionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dimension DAO
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.RemoveDatasetsFromDimension(dimensionId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more glossaryTermsIds as a GlossaryTerms to a Dimension
	//----------------------------------------------------------------------------
func AddGlossaryTermsToDimension(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dimensionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dimension DAO
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.AddGlossaryTermsToDimension(dimensionId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more glossaryTermsIds as a GlossaryTerms from a Dimension
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGlossaryTermsFromDimension(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dimensionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Dimension DAO
	//----------------------------------------------------------------------------
	requestResult := DimensionDAO.RemoveGlossaryTermsFromDimension(dimensionId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
