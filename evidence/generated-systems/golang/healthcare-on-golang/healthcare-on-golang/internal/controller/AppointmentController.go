package controller

import (
    AppointmentDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AppointmentDAO for database creation
//----------------------------------------------------------------------------
func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Appointment model
	//----------------------------------------------------------------------------
	data := model.Appointment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Appointment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment data access object to create
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.CreateAppointment( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AppointmentDAO to find the relevant Appointment
//----------------------------------------------------------------------------
func GetAppointment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Appointment data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.GetAppointment(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AppointmentDAO for database read of all Appointments
//----------------------------------------------------------------------------
func GetAllAppointment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Appointment data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.GetAllAppointment()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AppointmentDAO for database save
//----------------------------------------------------------------------------
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Appointment model
	//----------------------------------------------------------------------------
	var data = model.Appointment{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Appointment model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.UpdateAppointment(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AppointmentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Appointment data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AppointmentDAO.DeleteAppointment(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a Appointment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToAppointment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.AssignPatientToAppointment(appointmentId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a Appointment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromAppointment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.UnassignPatientFromAppointment(appointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Clinician on a Appointment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClinicianToAppointment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	clinicianId,_ := strconv.ParseUint( vars["clinicianId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.AssignClinicianToAppointment(appointmentId, clinicianId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Clinician on a Appointment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClinicianFromAppointment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.UnassignClinicianFromAppointment(appointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Facility on a Appointment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToAppointment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.AssignFacilityToAppointment(appointmentId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a Appointment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromAppointment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.UnassignFacilityFromAppointment(appointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Encounter on a Appointment
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEncounterToAppointment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encounterId,_ := strconv.ParseUint( vars["encounterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.AssignEncounterToAppointment(appointmentId, encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Encounter on a Appointment
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEncounterFromAppointment( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	appointmentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Appointment DAO
	//----------------------------------------------------------------------------
	requestResult := AppointmentDAO.UnassignEncounterFromAppointment(appointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


