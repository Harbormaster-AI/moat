package controller

import (
    MedicationOrderDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MedicationOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateMedicationOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicationOrder model
	//----------------------------------------------------------------------------
	data := model.MedicationOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicationOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.CreateMedicationOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MedicationOrderDAO to find the relevant MedicationOrder
//----------------------------------------------------------------------------
func GetMedicationOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicationOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.GetMedicationOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MedicationOrderDAO for database read of all MedicationOrders
//----------------------------------------------------------------------------
func GetAllMedicationOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.GetAllMedicationOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MedicationOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateMedicationOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicationOrder model
	//----------------------------------------------------------------------------
	var data = model.MedicationOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicationOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.UpdateMedicationOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MedicationOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMedicationOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicationOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MedicationOrderDAO.DeleteMedicationOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Order on a MedicationOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrderToMedicationOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderId,_ := strconv.ParseUint( vars["orderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.AssignOrderToMedicationOrder(medicationOrderId, orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Order on a MedicationOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrderFromMedicationOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.UnassignOrderFromMedicationOrder(medicationOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Pharmacy on a MedicationOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPharmacyToMedicationOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pharmacyId,_ := strconv.ParseUint( vars["pharmacyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.AssignPharmacyToMedicationOrder(medicationOrderId, pharmacyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Pharmacy on a MedicationOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPharmacyFromMedicationOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	medicationOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.UnassignPharmacyFromMedicationOrder(medicationOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more dispensesIds as a Dispenses to a MedicationOrder
	//----------------------------------------------------------------------------
func AddDispensesToMedicationOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicationOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dispensesIds,_ := vars["dispensesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.AddDispensesToMedicationOrder(medicationOrderId, dispensesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dispensesIds as a Dispenses from a MedicationOrder
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDispensesFromMedicationOrder(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicationOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dispensesIds,_ := vars["dispensesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicationOrder DAO
	//----------------------------------------------------------------------------
	requestResult := MedicationOrderDAO.RemoveDispensesFromMedicationOrder(medicationOrderId, dispensesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
