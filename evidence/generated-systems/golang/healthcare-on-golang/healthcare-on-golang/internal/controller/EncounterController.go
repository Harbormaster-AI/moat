package controller

import (
    EncounterDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EncounterDAO for database creation
//----------------------------------------------------------------------------
func CreateEncounter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Encounter model
	//----------------------------------------------------------------------------
	data := model.Encounter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Encounter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter data access object to create
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.CreateEncounter( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EncounterDAO to find the relevant Encounter
//----------------------------------------------------------------------------
func GetEncounter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Encounter data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.GetEncounter(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EncounterDAO for database read of all Encounters
//----------------------------------------------------------------------------
func GetAllEncounter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Encounter data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.GetAllEncounter()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EncounterDAO for database save
//----------------------------------------------------------------------------
func UpdateEncounter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Encounter model
	//----------------------------------------------------------------------------
	var data = model.Encounter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Encounter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.UpdateEncounter(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EncounterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEncounter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Encounter data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EncounterDAO.DeleteEncounter(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a Encounter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToEncounter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AssignPatientToEncounter(encounterId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a Encounter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromEncounter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.UnassignPatientFromEncounter(encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Clinician on a Encounter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClinicianToEncounter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	clinicianId,_ := strconv.ParseUint( vars["clinicianId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AssignClinicianToEncounter(encounterId, clinicianId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Clinician on a Encounter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClinicianFromEncounter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.UnassignClinicianFromEncounter(encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Facility on a Encounter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToEncounter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AssignFacilityToEncounter(encounterId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a Encounter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromEncounter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.UnassignFacilityFromEncounter(encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Appointment on a Encounter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAppointmentToEncounter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	appointmentId,_ := strconv.ParseUint( vars["appointmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AssignAppointmentToEncounter(encounterId, appointmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Appointment on a Encounter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAppointmentFromEncounter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.UnassignAppointmentFromEncounter(encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Admission on a Encounter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdmissionToEncounter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	admissionId,_ := strconv.ParseUint( vars["admissionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AssignAdmissionToEncounter(encounterId, admissionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Admission on a Encounter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdmissionFromEncounter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.UnassignAdmissionFromEncounter(encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Discharge on a Encounter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDischargeToEncounter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dischargeId,_ := strconv.ParseUint( vars["dischargeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AssignDischargeToEncounter(encounterId, dischargeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Discharge on a Encounter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDischargeFromEncounter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.UnassignDischargeFromEncounter(encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more diagnosesIds as a Diagnoses to a Encounter
	//----------------------------------------------------------------------------
func AddDiagnosesToEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	diagnosesIds,_ := vars["diagnosesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AddDiagnosesToEncounter(encounterId, diagnosesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more diagnosesIds as a Diagnoses from a Encounter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDiagnosesFromEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	diagnosesIds,_ := vars["diagnosesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.RemoveDiagnosesFromEncounter(encounterId, diagnosesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more proceduresIds as a Procedures to a Encounter
	//----------------------------------------------------------------------------
func AddProceduresToEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AddProceduresToEncounter(encounterId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more proceduresIds as a Procedures from a Encounter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProceduresFromEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	proceduresIds,_ := vars["proceduresIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.RemoveProceduresFromEncounter(encounterId, proceduresIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more observationsIds as a Observations to a Encounter
	//----------------------------------------------------------------------------
func AddObservationsToEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AddObservationsToEncounter(encounterId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more observationsIds as a Observations from a Encounter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObservationsFromEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.RemoveObservationsFromEncounter(encounterId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a Encounter
	//----------------------------------------------------------------------------
func AddOrdersToEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.AddOrdersToEncounter(encounterId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a Encounter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromEncounter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	encounterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Encounter DAO
	//----------------------------------------------------------------------------
	requestResult := EncounterDAO.RemoveOrdersFromEncounter(encounterId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
