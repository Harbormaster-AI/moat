package controller

import (
    PatientDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PatientDAO for database creation
//----------------------------------------------------------------------------
func CreatePatient(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Patient model
	//----------------------------------------------------------------------------
	data := model.Patient{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Patient model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Patient data access object to create
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.CreatePatient( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PatientDAO to find the relevant Patient
//----------------------------------------------------------------------------
func GetPatient(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Patient data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.GetPatient(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PatientDAO for database read of all Patients
//----------------------------------------------------------------------------
func GetAllPatient(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Patient data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.GetAllPatient()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PatientDAO for database save
//----------------------------------------------------------------------------
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Patient model
	//----------------------------------------------------------------------------
	var data = model.Patient{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Patient model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Patient data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.UpdatePatient(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PatientDAO for database deletion
//----------------------------------------------------------------------------
func DeletePatient(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Patient data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PatientDAO.DeletePatient(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more appointmentsIds as a Appointments to a Patient
	//----------------------------------------------------------------------------
func AddAppointmentsToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	appointmentsIds,_ := vars["appointmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddAppointmentsToPatient(patientId, appointmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more appointmentsIds as a Appointments from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAppointmentsFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	appointmentsIds,_ := vars["appointmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveAppointmentsFromPatient(patientId, appointmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more encountersIds as a Encounters to a Patient
	//----------------------------------------------------------------------------
func AddEncountersToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encountersIds,_ := vars["encountersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddEncountersToPatient(patientId, encountersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more encountersIds as a Encounters from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEncountersFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encountersIds,_ := vars["encountersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveEncountersFromPatient(patientId, encountersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more carePlansIds as a CarePlans to a Patient
	//----------------------------------------------------------------------------
func AddCarePlansToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	carePlansIds,_ := vars["carePlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddCarePlansToPatient(patientId, carePlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more carePlansIds as a CarePlans from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCarePlansFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	carePlansIds,_ := vars["carePlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveCarePlansFromPatient(patientId, carePlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more allergiesIds as a Allergies to a Patient
	//----------------------------------------------------------------------------
func AddAllergiesToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	allergiesIds,_ := vars["allergiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddAllergiesToPatient(patientId, allergiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more allergiesIds as a Allergies from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAllergiesFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	allergiesIds,_ := vars["allergiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveAllergiesFromPatient(patientId, allergiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more conditionsIds as a Conditions to a Patient
	//----------------------------------------------------------------------------
func AddConditionsToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	conditionsIds,_ := vars["conditionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddConditionsToPatient(patientId, conditionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more conditionsIds as a Conditions from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveConditionsFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	conditionsIds,_ := vars["conditionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveConditionsFromPatient(patientId, conditionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more medicationOrdersIds as a MedicationOrders to a Patient
	//----------------------------------------------------------------------------
func AddMedicationOrdersToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationOrdersIds,_ := vars["medicationOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddMedicationOrdersToPatient(patientId, medicationOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more medicationOrdersIds as a MedicationOrders from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMedicationOrdersFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationOrdersIds,_ := vars["medicationOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveMedicationOrdersFromPatient(patientId, medicationOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more labOrdersIds as a LabOrders to a Patient
	//----------------------------------------------------------------------------
func AddLabOrdersToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	labOrdersIds,_ := vars["labOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddLabOrdersToPatient(patientId, labOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more labOrdersIds as a LabOrders from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLabOrdersFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	labOrdersIds,_ := vars["labOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveLabOrdersFromPatient(patientId, labOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more imagingOrdersIds as a ImagingOrders to a Patient
	//----------------------------------------------------------------------------
func AddImagingOrdersToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingOrdersIds,_ := vars["imagingOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddImagingOrdersToPatient(patientId, imagingOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more imagingOrdersIds as a ImagingOrders from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveImagingOrdersFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingOrdersIds,_ := vars["imagingOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveImagingOrdersFromPatient(patientId, imagingOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more coveragesIds as a Coverages to a Patient
	//----------------------------------------------------------------------------
func AddCoveragesToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddCoveragesToPatient(patientId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more coveragesIds as a Coverages from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCoveragesFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveCoveragesFromPatient(patientId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a Patient
	//----------------------------------------------------------------------------
func AddClaimsToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddClaimsToPatient(patientId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveClaimsFromPatient(patientId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more devicesIds as a Devices to a Patient
	//----------------------------------------------------------------------------
func AddDevicesToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	devicesIds,_ := vars["devicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddDevicesToPatient(patientId, devicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more devicesIds as a Devices from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDevicesFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	devicesIds,_ := vars["devicesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveDevicesFromPatient(patientId, devicesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more observationsIds as a Observations to a Patient
	//----------------------------------------------------------------------------
func AddObservationsToPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.AddObservationsToPatient(patientId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more observationsIds as a Observations from a Patient
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObservationsFromPatient(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	patientId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	observationsIds,_ := vars["observationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Patient DAO
	//----------------------------------------------------------------------------
	requestResult := PatientDAO.RemoveObservationsFromPatient(patientId, observationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
