package controller

import (
    ExperimentVariantDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExperimentVariantDAO for database creation
//----------------------------------------------------------------------------
func CreateExperimentVariant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExperimentVariant model
	//----------------------------------------------------------------------------
	data := model.ExperimentVariant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExperimentVariant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.CreateExperimentVariant( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExperimentVariantDAO to find the relevant ExperimentVariant
//----------------------------------------------------------------------------
func GetExperimentVariant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExperimentVariant data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.GetExperimentVariant(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExperimentVariantDAO for database read of all ExperimentVariants
//----------------------------------------------------------------------------
func GetAllExperimentVariant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.GetAllExperimentVariant()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExperimentVariantDAO for database save
//----------------------------------------------------------------------------
func UpdateExperimentVariant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ExperimentVariant model
	//----------------------------------------------------------------------------
	var data = model.ExperimentVariant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ExperimentVariant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.UpdateExperimentVariant(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExperimentVariantDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExperimentVariant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ExperimentVariant data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExperimentVariantDAO.DeleteExperimentVariant(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Experiment on a ExperimentVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignExperimentToExperimentVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	experimentId,_ := strconv.ParseUint( vars["experimentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.AssignExperimentToExperimentVariant(experimentVariantId, experimentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Experiment on a ExperimentVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignExperimentFromExperimentVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.UnassignExperimentFromExperimentVariant(experimentVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a CreativeVariation on a ExperimentVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCreativeVariationToExperimentVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativeVariationId,_ := strconv.ParseUint( vars["creativeVariationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.AssignCreativeVariationToExperimentVariant(experimentVariantId, creativeVariationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CreativeVariation on a ExperimentVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCreativeVariationFromExperimentVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.UnassignCreativeVariationFromExperimentVariant(experimentVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LineItem on a ExperimentVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLineItemToExperimentVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemId,_ := strconv.ParseUint( vars["lineItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.AssignLineItemToExperimentVariant(experimentVariantId, lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LineItem on a ExperimentVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLineItemFromExperimentVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ExperimentVariant DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentVariantDAO.UnassignLineItemFromExperimentVariant(experimentVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


