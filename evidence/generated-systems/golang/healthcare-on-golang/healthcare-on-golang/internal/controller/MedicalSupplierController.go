package controller

import (
    MedicalSupplierDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to MedicalSupplierDAO for database creation
//----------------------------------------------------------------------------
func CreateMedicalSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicalSupplier model
	//----------------------------------------------------------------------------
	data := model.MedicalSupplier{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicalSupplier model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicalSupplier data access object to create
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.CreateMedicalSupplier( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to MedicalSupplierDAO to find the relevant MedicalSupplier
//----------------------------------------------------------------------------
func GetMedicalSupplier(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicalSupplier data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.GetMedicalSupplier(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to MedicalSupplierDAO for database read of all MedicalSuppliers
//----------------------------------------------------------------------------
func GetAllMedicalSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the MedicalSupplier data access object to get all
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.GetAllMedicalSupplier()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to MedicalSupplierDAO for database save
//----------------------------------------------------------------------------
func UpdateMedicalSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty MedicalSupplier model
	//----------------------------------------------------------------------------
	var data = model.MedicalSupplier{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a MedicalSupplier model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the MedicalSupplier data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.UpdateMedicalSupplier(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to MedicalSupplierDAO for database deletion
//----------------------------------------------------------------------------
func DeleteMedicalSupplier(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the MedicalSupplier data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := MedicalSupplierDAO.DeleteMedicalSupplier(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more facilitiesIds as a Facilities to a MedicalSupplier
	//----------------------------------------------------------------------------
func AddFacilitiesToMedicalSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalSupplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilitiesIds,_ := vars["facilitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalSupplier DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.AddFacilitiesToMedicalSupplier(medicalSupplierId, facilitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more facilitiesIds as a Facilities from a MedicalSupplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFacilitiesFromMedicalSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalSupplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilitiesIds,_ := vars["facilitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalSupplier DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.RemoveFacilitiesFromMedicalSupplier(medicalSupplierId, facilitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more inventoryItemsIds as a InventoryItems to a MedicalSupplier
	//----------------------------------------------------------------------------
func AddInventoryItemsToMedicalSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalSupplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalSupplier DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.AddInventoryItemsToMedicalSupplier(medicalSupplierId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventoryItemsIds as a InventoryItems from a MedicalSupplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventoryItemsFromMedicalSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	medicalSupplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the MedicalSupplier DAO
	//----------------------------------------------------------------------------
	requestResult := MedicalSupplierDAO.RemoveInventoryItemsFromMedicalSupplier(medicalSupplierId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
