package controller

import (
    ExperimentDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ExperimentDAO for database creation
//----------------------------------------------------------------------------
func CreateExperiment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Experiment model
	//----------------------------------------------------------------------------
	data := model.Experiment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Experiment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment data access object to create
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.CreateExperiment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ExperimentDAO to find the relevant Experiment
//----------------------------------------------------------------------------
func GetExperiment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Experiment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.GetExperiment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ExperimentDAO for database read of all Experiments
//----------------------------------------------------------------------------
func GetAllExperiment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Experiment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.GetAllExperiment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ExperimentDAO for database save
//----------------------------------------------------------------------------
func UpdateExperiment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Experiment model
	//----------------------------------------------------------------------------
	var data = model.Experiment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Experiment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.UpdateExperiment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ExperimentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteExperiment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Experiment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ExperimentDAO.DeleteExperiment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a Experiment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToExperiment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.AssignCampaignToExperiment(experimentId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a Experiment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromExperiment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.UnassignCampaignFromExperiment(experimentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more variantsIds as a Variants to a Experiment
	//----------------------------------------------------------------------------
func AddVariantsToExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.AddVariantsToExperiment(experimentId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variantsIds as a Variants from a Experiment
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariantsFromExperiment(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	experimentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Experiment DAO
	//----------------------------------------------------------------------------
	requestResult := ExperimentDAO.RemoveVariantsFromExperiment(experimentId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
