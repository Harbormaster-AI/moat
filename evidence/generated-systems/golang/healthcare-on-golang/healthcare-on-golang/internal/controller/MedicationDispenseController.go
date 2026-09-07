package controller

import (
    MedicationDispenseDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MedicationDispenseDAO for database creation
//----------------------------------------------------------------------------
func CreateMedicationDispense(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicationDispense model
	//----------------------------------------------------------------------------
	data := model.MedicationDispense{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicationDispense model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense data access object to create
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.CreateMedicationDispense( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MedicationDispenseDAO to find the relevant MedicationDispense
//----------------------------------------------------------------------------
func GetMedicationDispense(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicationDispense data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.GetMedicationDispense(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MedicationDispenseDAO for database read of all MedicationDispenses
//----------------------------------------------------------------------------
func GetAllMedicationDispense(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.GetAllMedicationDispense()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MedicationDispenseDAO for database save
//----------------------------------------------------------------------------
func UpdateMedicationDispense(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicationDispense model
	//----------------------------------------------------------------------------
	var data = model.MedicationDispense{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicationDispense model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.UpdateMedicationDispense(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MedicationDispenseDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMedicationDispense(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicationDispense data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MedicationDispenseDAO.DeleteMedicationDispense(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a MedicationOrder on a MedicationDispense
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMedicationOrderToMedicationDispense(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationDispenseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationOrderId,_ := strconv.ParseUint( vars["medicationOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.AssignMedicationOrderToMedicationDispense(medicationDispenseId, medicationOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a MedicationOrder on a MedicationDispense
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMedicationOrderFromMedicationDispense( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationDispenseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.UnassignMedicationOrderFromMedicationDispense(medicationDispenseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Pharmacy on a MedicationDispense
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPharmacyToMedicationDispense(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationDispenseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pharmacyId,_ := strconv.ParseUint( vars["pharmacyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.AssignPharmacyToMedicationDispense(medicationDispenseId, pharmacyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Pharmacy on a MedicationDispense
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPharmacyFromMedicationDispense( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationDispenseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.UnassignPharmacyFromMedicationDispense(medicationDispenseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Patient on a MedicationDispense
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToMedicationDispense(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationDispenseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.AssignPatientToMedicationDispense(medicationDispenseId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a MedicationDispense
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromMedicationDispense( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationDispenseId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationDispense DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationDispenseDAO.UnassignPatientFromMedicationDispense(medicationDispenseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


