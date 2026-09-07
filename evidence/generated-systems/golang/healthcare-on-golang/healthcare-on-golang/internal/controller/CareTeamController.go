package controller

import (
    CareTeamDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CareTeamDAO for database creation
//----------------------------------------------------------------------------
func CreateCareTeam(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CareTeam model
	//----------------------------------------------------------------------------
	data := model.CareTeam{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CareTeam model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam data access object to create
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.CreateCareTeam( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CareTeamDAO to find the relevant CareTeam
//----------------------------------------------------------------------------
func GetCareTeam(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CareTeam data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.GetCareTeam(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CareTeamDAO for database read of all CareTeams
//----------------------------------------------------------------------------
func GetAllCareTeam(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CareTeam data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.GetAllCareTeam()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CareTeamDAO for database save
//----------------------------------------------------------------------------
func UpdateCareTeam(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CareTeam model
	//----------------------------------------------------------------------------
	var data = model.CareTeam{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CareTeam model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.UpdateCareTeam(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CareTeamDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCareTeam(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CareTeam data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CareTeamDAO.DeleteCareTeam(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Department on a CareTeam
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDepartmentToCareTeam(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTeamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentId,_ := strconv.ParseUint( vars["departmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam DAO
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.AssignDepartmentToCareTeam(careTeamId, departmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Department on a CareTeam
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDepartmentFromCareTeam( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	careTeamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam DAO
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.UnassignDepartmentFromCareTeam(careTeamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more cliniciansIds as a Clinicians to a CareTeam
	//----------------------------------------------------------------------------
func AddCliniciansToCareTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	careTeamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cliniciansIds,_ := vars["cliniciansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam DAO
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.AddCliniciansToCareTeam(careTeamId, cliniciansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more cliniciansIds as a Clinicians from a CareTeam
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCliniciansFromCareTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	careTeamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cliniciansIds,_ := vars["cliniciansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam DAO
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.RemoveCliniciansFromCareTeam(careTeamId, cliniciansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more patientsIds as a Patients to a CareTeam
	//----------------------------------------------------------------------------
func AddPatientsToCareTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	careTeamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientsIds,_ := vars["patientsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam DAO
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.AddPatientsToCareTeam(careTeamId, patientsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more patientsIds as a Patients from a CareTeam
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePatientsFromCareTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	careTeamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientsIds,_ := vars["patientsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CareTeam DAO
	//----------------------------------------------------------------------------
	requestResult := CareTeamDAO.RemovePatientsFromCareTeam(careTeamId, patientsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
