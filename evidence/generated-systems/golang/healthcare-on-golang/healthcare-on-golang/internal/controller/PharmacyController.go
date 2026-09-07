package controller

import (
    PharmacyDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PharmacyDAO for database creation
//----------------------------------------------------------------------------
func CreatePharmacy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Pharmacy model
	//----------------------------------------------------------------------------
	data := model.Pharmacy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Pharmacy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy data access object to create
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.CreatePharmacy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PharmacyDAO to find the relevant Pharmacy
//----------------------------------------------------------------------------
func GetPharmacy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Pharmacy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.GetPharmacy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PharmacyDAO for database read of all Pharmacys
//----------------------------------------------------------------------------
func GetAllPharmacy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.GetAllPharmacy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PharmacyDAO for database save
//----------------------------------------------------------------------------
func UpdatePharmacy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Pharmacy model
	//----------------------------------------------------------------------------
	var data = model.Pharmacy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Pharmacy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.UpdatePharmacy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PharmacyDAO for database deletion
//----------------------------------------------------------------------------
func DeletePharmacy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Pharmacy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PharmacyDAO.DeletePharmacy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Facility on a Pharmacy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToPharmacy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	pharmacyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy DAO
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.AssignFacilityToPharmacy(pharmacyId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a Pharmacy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromPharmacy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	pharmacyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy DAO
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.UnassignFacilityFromPharmacy(pharmacyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more medicationDispensesIds as a MedicationDispenses to a Pharmacy
	//----------------------------------------------------------------------------
func AddMedicationDispensesToPharmacy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pharmacyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationDispensesIds,_ := vars["medicationDispensesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy DAO
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.AddMedicationDispensesToPharmacy(pharmacyId, medicationDispensesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more medicationDispensesIds as a MedicationDispenses from a Pharmacy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMedicationDispensesFromPharmacy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pharmacyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationDispensesIds,_ := vars["medicationDispensesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy DAO
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.RemoveMedicationDispensesFromPharmacy(pharmacyId, medicationDispensesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more medicationOrdersIds as a MedicationOrders to a Pharmacy
	//----------------------------------------------------------------------------
func AddMedicationOrdersToPharmacy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pharmacyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationOrdersIds,_ := vars["medicationOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy DAO
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.AddMedicationOrdersToPharmacy(pharmacyId, medicationOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more medicationOrdersIds as a MedicationOrders from a Pharmacy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveMedicationOrdersFromPharmacy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	pharmacyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	medicationOrdersIds,_ := vars["medicationOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Pharmacy DAO
	//----------------------------------------------------------------------------
	requestResult := PharmacyDAO.RemoveMedicationOrdersFromPharmacy(pharmacyId, medicationOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
