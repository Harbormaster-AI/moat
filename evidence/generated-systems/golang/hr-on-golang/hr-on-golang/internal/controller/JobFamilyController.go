package controller

import (
    JobFamilyDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to JobFamilyDAO for database creation
//----------------------------------------------------------------------------
func CreateJobFamily(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobFamily model
	//----------------------------------------------------------------------------
	data := model.JobFamily{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobFamily model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobFamily data access object to create
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.CreateJobFamily( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to JobFamilyDAO to find the relevant JobFamily
//----------------------------------------------------------------------------
func GetJobFamily(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobFamily data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.GetJobFamily(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to JobFamilyDAO for database read of all JobFamilys
//----------------------------------------------------------------------------
func GetAllJobFamily(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the JobFamily data access object to get all
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.GetAllJobFamily()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to JobFamilyDAO for database save
//----------------------------------------------------------------------------
func UpdateJobFamily(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty JobFamily model
	//----------------------------------------------------------------------------
	var data = model.JobFamily{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a JobFamily model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the JobFamily data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.UpdateJobFamily(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to JobFamilyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteJobFamily(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the JobFamily data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := JobFamilyDAO.DeleteJobFamily(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a JobFamily
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToJobFamily(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobFamily DAO
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.AssignOrganizationToJobFamily(jobFamilyId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a JobFamily
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromJobFamily( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	jobFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the JobFamily DAO
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.UnassignOrganizationFromJobFamily(jobFamilyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more jobProfilesIds as a JobProfiles to a JobFamily
	//----------------------------------------------------------------------------
func AddJobProfilesToJobFamily(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfilesIds,_ := vars["jobProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobFamily DAO
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.AddJobProfilesToJobFamily(jobFamilyId, jobProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more jobProfilesIds as a JobProfiles from a JobFamily
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveJobProfilesFromJobFamily(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	jobFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	jobProfilesIds,_ := vars["jobProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the JobFamily DAO
	//----------------------------------------------------------------------------
	requestResult := JobFamilyDAO.RemoveJobProfilesFromJobFamily(jobFamilyId, jobProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
