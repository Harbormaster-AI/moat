package controller

import (
    ClinicianDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ClinicianDAO for database creation
//----------------------------------------------------------------------------
func CreateClinician(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Clinician model
	//----------------------------------------------------------------------------
	data := model.Clinician{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Clinician model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Clinician data access object to create
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.CreateClinician( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ClinicianDAO to find the relevant Clinician
//----------------------------------------------------------------------------
func GetClinician(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Clinician data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.GetClinician(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ClinicianDAO for database read of all Clinicians
//----------------------------------------------------------------------------
func GetAllClinician(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Clinician data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.GetAllClinician()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ClinicianDAO for database save
//----------------------------------------------------------------------------
func UpdateClinician(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Clinician model
	//----------------------------------------------------------------------------
	var data = model.Clinician{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Clinician model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Clinician data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.UpdateClinician(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ClinicianDAO for database deletion
//----------------------------------------------------------------------------
func DeleteClinician(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Clinician data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ClinicianDAO.DeleteClinician(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more careTeamsIds as a CareTeams to a Clinician
	//----------------------------------------------------------------------------
func AddCareTeamsToClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	careTeamsIds,_ := vars["careTeamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.AddCareTeamsToClinician(clinicianId, careTeamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more careTeamsIds as a CareTeams from a Clinician
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCareTeamsFromClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	careTeamsIds,_ := vars["careTeamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.RemoveCareTeamsFromClinician(clinicianId, careTeamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more appointmentsIds as a Appointments to a Clinician
	//----------------------------------------------------------------------------
func AddAppointmentsToClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	appointmentsIds,_ := vars["appointmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.AddAppointmentsToClinician(clinicianId, appointmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more appointmentsIds as a Appointments from a Clinician
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAppointmentsFromClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	appointmentsIds,_ := vars["appointmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.RemoveAppointmentsFromClinician(clinicianId, appointmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more encountersIds as a Encounters to a Clinician
	//----------------------------------------------------------------------------
func AddEncountersToClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encountersIds,_ := vars["encountersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.AddEncountersToClinician(clinicianId, encountersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more encountersIds as a Encounters from a Clinician
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEncountersFromClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encountersIds,_ := vars["encountersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.RemoveEncountersFromClinician(clinicianId, encountersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more proceduresIds as a Procedures to a Clinician
	//----------------------------------------------------------------------------
func AddProceduresToClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.AddProceduresToClinician(clinicianId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more proceduresIds as a Procedures from a Clinician
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProceduresFromClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.RemoveProceduresFromClinician(clinicianId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more imagingReportsIds as a ImagingReports to a Clinician
	//----------------------------------------------------------------------------
func AddImagingReportsToClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingReportsIds,_ := vars["imagingReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.AddImagingReportsToClinician(clinicianId, imagingReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more imagingReportsIds as a ImagingReports from a Clinician
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveImagingReportsFromClinician(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicianId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingReportsIds,_ := vars["imagingReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Clinician DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicianDAO.RemoveImagingReportsFromClinician(clinicianId, imagingReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
