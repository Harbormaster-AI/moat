package controller

import (
    CompetencyDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CompetencyDAO for database creation
//----------------------------------------------------------------------------
func CreateCompetency(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Competency model
	//----------------------------------------------------------------------------
	data := model.Competency{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Competency model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Competency data access object to create
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.CreateCompetency( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CompetencyDAO to find the relevant Competency
//----------------------------------------------------------------------------
func GetCompetency(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Competency data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.GetCompetency(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CompetencyDAO for database read of all Competencys
//----------------------------------------------------------------------------
func GetAllCompetency(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Competency data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.GetAllCompetency()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CompetencyDAO for database save
//----------------------------------------------------------------------------
func UpdateCompetency(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Competency model
	//----------------------------------------------------------------------------
	var data = model.Competency{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Competency model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Competency data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.UpdateCompetency(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CompetencyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCompetency(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Competency data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CompetencyDAO.DeleteCompetency(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more jobProfilesIds as a JobProfiles to a Competency
	//----------------------------------------------------------------------------
func AddJobProfilesToCompetency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	competencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfilesIds,_ := vars["jobProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Competency DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.AddJobProfilesToCompetency(competencyId, jobProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more jobProfilesIds as a JobProfiles from a Competency
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveJobProfilesFromCompetency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	competencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfilesIds,_ := vars["jobProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Competency DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.RemoveJobProfilesFromCompetency(competencyId, jobProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more competencyRatingsIds as a CompetencyRatings to a Competency
	//----------------------------------------------------------------------------
func AddCompetencyRatingsToCompetency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	competencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	competencyRatingsIds,_ := vars["competencyRatingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Competency DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.AddCompetencyRatingsToCompetency(competencyId, competencyRatingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more competencyRatingsIds as a CompetencyRatings from a Competency
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCompetencyRatingsFromCompetency(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	competencyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	competencyRatingsIds,_ := vars["competencyRatingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Competency DAO
	//----------------------------------------------------------------------------
	requestResult := CompetencyDAO.RemoveCompetencyRatingsFromCompetency(competencyId, competencyRatingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
