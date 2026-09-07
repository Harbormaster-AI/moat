package controller

import (
    JobProfileDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to JobProfileDAO for database creation
//----------------------------------------------------------------------------
func CreateJobProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobProfile model
	//----------------------------------------------------------------------------
	data := model.JobProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile data access object to create
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.CreateJobProfile( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to JobProfileDAO to find the relevant JobProfile
//----------------------------------------------------------------------------
func GetJobProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobProfile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.GetJobProfile(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to JobProfileDAO for database read of all JobProfiles
//----------------------------------------------------------------------------
func GetAllJobProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the JobProfile data access object to get all
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.GetAllJobProfile()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to JobProfileDAO for database save
//----------------------------------------------------------------------------
func UpdateJobProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobProfile model
	//----------------------------------------------------------------------------
	var data = model.JobProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.UpdateJobProfile(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to JobProfileDAO for database deletion
//----------------------------------------------------------------------------
func DeleteJobProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobProfile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := JobProfileDAO.DeleteJobProfile(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a JobFamily on a JobProfile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignJobFamilyToJobProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobFamilyId,_ := strconv.ParseUint( vars["jobFamilyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.AssignJobFamilyToJobProfile(jobProfileId, jobFamilyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a JobFamily on a JobProfile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignJobFamilyFromJobProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.UnassignJobFamilyFromJobProfile(jobProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more competenciesIds as a Competencies to a JobProfile
	//----------------------------------------------------------------------------
func AddCompetenciesToJobProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	competenciesIds,_ := vars["competenciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.AddCompetenciesToJobProfile(jobProfileId, competenciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more competenciesIds as a Competencies from a JobProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCompetenciesFromJobProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	competenciesIds,_ := vars["competenciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.RemoveCompetenciesFromJobProfile(jobProfileId, competenciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more trainingRecommendationsIds as a TrainingRecommendations to a JobProfile
	//----------------------------------------------------------------------------
func AddTrainingRecommendationsToJobProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRecommendationsIds,_ := vars["trainingRecommendationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.AddTrainingRecommendationsToJobProfile(jobProfileId, trainingRecommendationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more trainingRecommendationsIds as a TrainingRecommendations from a JobProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTrainingRecommendationsFromJobProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	trainingRecommendationsIds,_ := vars["trainingRecommendationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.RemoveTrainingRecommendationsFromJobProfile(jobProfileId, trainingRecommendationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more positionsIds as a Positions to a JobProfile
	//----------------------------------------------------------------------------
func AddPositionsToJobProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.AddPositionsToJobProfile(jobProfileId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more positionsIds as a Positions from a JobProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePositionsFromJobProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobProfile DAO
	//----------------------------------------------------------------------------
	requestResult := JobProfileDAO.RemovePositionsFromJobProfile(jobProfileId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
