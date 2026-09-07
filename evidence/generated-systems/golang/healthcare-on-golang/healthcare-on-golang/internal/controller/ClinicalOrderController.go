package controller

import (
    ClinicalOrderDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ClinicalOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateClinicalOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ClinicalOrder model
	//----------------------------------------------------------------------------
	data := model.ClinicalOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ClinicalOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.CreateClinicalOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ClinicalOrderDAO to find the relevant ClinicalOrder
//----------------------------------------------------------------------------
func GetClinicalOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ClinicalOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.GetClinicalOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ClinicalOrderDAO for database read of all ClinicalOrders
//----------------------------------------------------------------------------
func GetAllClinicalOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.GetAllClinicalOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ClinicalOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateClinicalOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ClinicalOrder model
	//----------------------------------------------------------------------------
	var data = model.ClinicalOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ClinicalOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.UpdateClinicalOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ClinicalOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteClinicalOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ClinicalOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ClinicalOrderDAO.DeleteClinicalOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a ClinicalOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToClinicalOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AssignPatientToClinicalOrder(clinicalOrderId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a ClinicalOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromClinicalOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.UnassignPatientFromClinicalOrder(clinicalOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Encounter on a ClinicalOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEncounterToClinicalOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encounterId,_ := strconv.ParseUint( vars["encounterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AssignEncounterToClinicalOrder(clinicalOrderId, encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Encounter on a ClinicalOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEncounterFromClinicalOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.UnassignEncounterFromClinicalOrder(clinicalOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a OrderingClinician on a ClinicalOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrderingClinicianToClinicalOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderingClinicianId,_ := strconv.ParseUint( vars["orderingClinicianId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AssignOrderingClinicianToClinicalOrder(clinicalOrderId, orderingClinicianId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a OrderingClinician on a ClinicalOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrderingClinicianFromClinicalOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.UnassignOrderingClinicianFromClinicalOrder(clinicalOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more medicationOrdersIds as a MedicationOrders to a ClinicalOrder
	//----------------------------------------------------------------------------
func AddMedicationOrdersToClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationOrdersIds,_ := vars["medicationOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AddMedicationOrdersToClinicalOrder(clinicalOrderId, medicationOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more medicationOrdersIds as a MedicationOrders from a ClinicalOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMedicationOrdersFromClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationOrdersIds,_ := vars["medicationOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.RemoveMedicationOrdersFromClinicalOrder(clinicalOrderId, medicationOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more laboratoryOrdersIds as a LaboratoryOrders to a ClinicalOrder
	//----------------------------------------------------------------------------
func AddLaboratoryOrdersToClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoryOrdersIds,_ := vars["laboratoryOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AddLaboratoryOrdersToClinicalOrder(clinicalOrderId, laboratoryOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more laboratoryOrdersIds as a LaboratoryOrders from a ClinicalOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLaboratoryOrdersFromClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoryOrdersIds,_ := vars["laboratoryOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.RemoveLaboratoryOrdersFromClinicalOrder(clinicalOrderId, laboratoryOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more imagingOrdersIds as a ImagingOrders to a ClinicalOrder
	//----------------------------------------------------------------------------
func AddImagingOrdersToClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingOrdersIds,_ := vars["imagingOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AddImagingOrdersToClinicalOrder(clinicalOrderId, imagingOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more imagingOrdersIds as a ImagingOrders from a ClinicalOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveImagingOrdersFromClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingOrdersIds,_ := vars["imagingOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.RemoveImagingOrdersFromClinicalOrder(clinicalOrderId, imagingOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more procedureOrdersIds as a ProcedureOrders to a ClinicalOrder
	//----------------------------------------------------------------------------
func AddProcedureOrdersToClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	procedureOrdersIds,_ := vars["procedureOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AddProcedureOrdersToClinicalOrder(clinicalOrderId, procedureOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more procedureOrdersIds as a ProcedureOrders from a ClinicalOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcedureOrdersFromClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	procedureOrdersIds,_ := vars["procedureOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.RemoveProcedureOrdersFromClinicalOrder(clinicalOrderId, procedureOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more authorizationsIds as a Authorizations to a ClinicalOrder
	//----------------------------------------------------------------------------
func AddAuthorizationsToClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	authorizationsIds,_ := vars["authorizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.AddAuthorizationsToClinicalOrder(clinicalOrderId, authorizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more authorizationsIds as a Authorizations from a ClinicalOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAuthorizationsFromClinicalOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	clinicalOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	authorizationsIds,_ := vars["authorizationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ClinicalOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ClinicalOrderDAO.RemoveAuthorizationsFromClinicalOrder(clinicalOrderId, authorizationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
