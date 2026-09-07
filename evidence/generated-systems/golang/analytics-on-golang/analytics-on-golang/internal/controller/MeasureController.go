package controller

import (
    MeasureDAO "analytics-on-golang/internal/dao"
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MeasureDAO for database creation
//----------------------------------------------------------------------------
func CreateMeasure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Measure model
	//----------------------------------------------------------------------------
	data := model.Measure{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Measure model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Measure data access object to create
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.CreateMeasure( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MeasureDAO to find the relevant Measure
//----------------------------------------------------------------------------
func GetMeasure(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Measure data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.GetMeasure(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MeasureDAO for database read of all Measures
//----------------------------------------------------------------------------
func GetAllMeasure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Measure data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.GetAllMeasure()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MeasureDAO for database save
//----------------------------------------------------------------------------
func UpdateMeasure(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Measure model
	//----------------------------------------------------------------------------
	var data = model.Measure{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Measure model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Measure data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.UpdateMeasure(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MeasureDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMeasure(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Measure data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MeasureDAO.DeleteMeasure(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a SemanticModel on a Measure
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSemanticModelToMeasure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	measureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	semanticModelId,_ := strconv.ParseUint( vars["semanticModelId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Measure DAO
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.AssignSemanticModelToMeasure(measureId, semanticModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SemanticModel on a Measure
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSemanticModelFromMeasure( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	measureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Measure DAO
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.UnassignSemanticModelFromMeasure(measureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more datasetsIds as a Datasets to a Measure
	//----------------------------------------------------------------------------
func AddDatasetsToMeasure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	measureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Measure DAO
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.AddDatasetsToMeasure(measureId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more datasetsIds as a Datasets from a Measure
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDatasetsFromMeasure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	measureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	datasetsIds,_ := vars["datasetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Measure DAO
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.RemoveDatasetsFromMeasure(measureId, datasetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more glossaryTermsIds as a GlossaryTerms to a Measure
	//----------------------------------------------------------------------------
func AddGlossaryTermsToMeasure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	measureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Measure DAO
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.AddGlossaryTermsToMeasure(measureId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more glossaryTermsIds as a GlossaryTerms from a Measure
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGlossaryTermsFromMeasure(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	measureId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	glossaryTermsIds,_ := vars["glossaryTermsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Measure DAO
	//----------------------------------------------------------------------------
	requestResult := MeasureDAO.RemoveGlossaryTermsFromMeasure(measureId, glossaryTermsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
