package controller

import (
    SupplierDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SupplierDAO for database creation
//----------------------------------------------------------------------------
func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Supplier model
	//----------------------------------------------------------------------------
	data := model.Supplier{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Supplier model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Supplier data access object to create
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.CreateSupplier( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SupplierDAO to find the relevant Supplier
//----------------------------------------------------------------------------
func GetSupplier(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Supplier data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.GetSupplier(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SupplierDAO for database read of all Suppliers
//----------------------------------------------------------------------------
func GetAllSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Supplier data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.GetAllSupplier()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SupplierDAO for database save
//----------------------------------------------------------------------------
func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Supplier model
	//----------------------------------------------------------------------------
	var data = model.Supplier{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Supplier model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Supplier data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.UpdateSupplier(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SupplierDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Supplier data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SupplierDAO.DeleteSupplier(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more manufacturersIds as a Manufacturers to a Supplier
	//----------------------------------------------------------------------------
func AddManufacturersToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	manufacturersIds,_ := vars["manufacturersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddManufacturersToSupplier(supplierId, manufacturersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more manufacturersIds as a Manufacturers from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveManufacturersFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	manufacturersIds,_ := vars["manufacturersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveManufacturersFromSupplier(supplierId, manufacturersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more componentsIds as a Components to a Supplier
	//----------------------------------------------------------------------------
func AddComponentsToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	componentsIds,_ := vars["componentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddComponentsToSupplier(supplierId, componentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more componentsIds as a Components from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveComponentsFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	componentsIds,_ := vars["componentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveComponentsFromSupplier(supplierId, componentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more engineTypesIds as a EngineTypes to a Supplier
	//----------------------------------------------------------------------------
func AddEngineTypesToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engineTypesIds,_ := vars["engineTypesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddEngineTypesToSupplier(supplierId, engineTypesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more engineTypesIds as a EngineTypes from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEngineTypesFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engineTypesIds,_ := vars["engineTypesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveEngineTypesFromSupplier(supplierId, engineTypesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more avionicsSuitesIds as a AvionicsSuites to a Supplier
	//----------------------------------------------------------------------------
func AddAvionicsSuitesToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	avionicsSuitesIds,_ := vars["avionicsSuitesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddAvionicsSuitesToSupplier(supplierId, avionicsSuitesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more avionicsSuitesIds as a AvionicsSuites from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAvionicsSuitesFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	avionicsSuitesIds,_ := vars["avionicsSuitesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveAvionicsSuitesFromSupplier(supplierId, avionicsSuitesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more apusIds as a Apus to a Supplier
	//----------------------------------------------------------------------------
func AddApusToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	apusIds,_ := vars["apusIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddApusToSupplier(supplierId, apusIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more apusIds as a Apus from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveApusFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	apusIds,_ := vars["apusIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveApusFromSupplier(supplierId, apusIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more landingGearsIds as a LandingGears to a Supplier
	//----------------------------------------------------------------------------
func AddLandingGearsToSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	landingGearsIds,_ := vars["landingGearsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.AddLandingGearsToSupplier(supplierId, landingGearsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more landingGearsIds as a LandingGears from a Supplier
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLandingGearsFromSupplier(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	supplierId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	landingGearsIds,_ := vars["landingGearsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Supplier DAO
	//----------------------------------------------------------------------------
	requestResult := SupplierDAO.RemoveLandingGearsFromSupplier(supplierId, landingGearsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
