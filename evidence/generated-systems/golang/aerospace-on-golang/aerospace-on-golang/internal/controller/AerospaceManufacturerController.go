package controller

import (
    AerospaceManufacturerDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AerospaceManufacturerDAO for database creation
//----------------------------------------------------------------------------
func CreateAerospaceManufacturer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AerospaceManufacturer model
	//----------------------------------------------------------------------------
	data := model.AerospaceManufacturer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AerospaceManufacturer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer data access object to create
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.CreateAerospaceManufacturer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AerospaceManufacturerDAO to find the relevant AerospaceManufacturer
//----------------------------------------------------------------------------
func GetAerospaceManufacturer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AerospaceManufacturer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.GetAerospaceManufacturer(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AerospaceManufacturerDAO for database read of all AerospaceManufacturers
//----------------------------------------------------------------------------
func GetAllAerospaceManufacturer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.GetAllAerospaceManufacturer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AerospaceManufacturerDAO for database save
//----------------------------------------------------------------------------
func UpdateAerospaceManufacturer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AerospaceManufacturer model
	//----------------------------------------------------------------------------
	var data = model.AerospaceManufacturer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AerospaceManufacturer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.UpdateAerospaceManufacturer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AerospaceManufacturerDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAerospaceManufacturer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AerospaceManufacturer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AerospaceManufacturerDAO.DeleteAerospaceManufacturer(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more programsIds as a Programs to a AerospaceManufacturer
	//----------------------------------------------------------------------------
func AddProgramsToAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	programsIds,_ := vars["programsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.AddProgramsToAerospaceManufacturer(aerospaceManufacturerId, programsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more programsIds as a Programs from a AerospaceManufacturer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProgramsFromAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	programsIds,_ := vars["programsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.RemoveProgramsFromAerospaceManufacturer(aerospaceManufacturerId, programsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more plantsIds as a Plants to a AerospaceManufacturer
	//----------------------------------------------------------------------------
func AddPlantsToAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantsIds,_ := vars["plantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.AddPlantsToAerospaceManufacturer(aerospaceManufacturerId, plantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more plantsIds as a Plants from a AerospaceManufacturer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlantsFromAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantsIds,_ := vars["plantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.RemovePlantsFromAerospaceManufacturer(aerospaceManufacturerId, plantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more suppliersIds as a Suppliers to a AerospaceManufacturer
	//----------------------------------------------------------------------------
func AddSuppliersToAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.AddSuppliersToAerospaceManufacturer(aerospaceManufacturerId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more suppliersIds as a Suppliers from a AerospaceManufacturer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSuppliersFromAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.RemoveSuppliersFromAerospaceManufacturer(aerospaceManufacturerId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more productionCertificatesIds as a ProductionCertificates to a AerospaceManufacturer
	//----------------------------------------------------------------------------
func AddProductionCertificatesToAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionCertificatesIds,_ := vars["productionCertificatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.AddProductionCertificatesToAerospaceManufacturer(aerospaceManufacturerId, productionCertificatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productionCertificatesIds as a ProductionCertificates from a AerospaceManufacturer
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductionCertificatesFromAerospaceManufacturer(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aerospaceManufacturerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionCertificatesIds,_ := vars["productionCertificatesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AerospaceManufacturer DAO
	//----------------------------------------------------------------------------
	requestResult := AerospaceManufacturerDAO.RemoveProductionCertificatesFromAerospaceManufacturer(aerospaceManufacturerId, productionCertificatesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
