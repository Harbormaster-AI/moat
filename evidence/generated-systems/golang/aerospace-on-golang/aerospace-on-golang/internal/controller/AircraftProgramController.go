package controller

import (
    AircraftProgramDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftProgramDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraftProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftProgram model
	//----------------------------------------------------------------------------
	data := model.AircraftProgram{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftProgram model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.CreateAircraftProgram( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftProgramDAO to find the relevant AircraftProgram
//----------------------------------------------------------------------------
func GetAircraftProgram(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftProgram data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.GetAircraftProgram(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftProgramDAO for database read of all AircraftPrograms
//----------------------------------------------------------------------------
func GetAllAircraftProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.GetAllAircraftProgram()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftProgramDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraftProgram(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftProgram model
	//----------------------------------------------------------------------------
	var data = model.AircraftProgram{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftProgram model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.UpdateAircraftProgram(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftProgramDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraftProgram(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftProgram data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftProgramDAO.DeleteAircraftProgram(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Manufacturer on a AircraftProgram
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignManufacturerToAircraftProgram(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	manufacturerId,_ := strconv.ParseUint( vars["manufacturerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.AssignManufacturerToAircraftProgram(aircraftProgramId, manufacturerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Manufacturer on a AircraftProgram
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignManufacturerFromAircraftProgram( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.UnassignManufacturerFromAircraftProgram(aircraftProgramId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a TypeCertificate on a AircraftProgram
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTypeCertificateToAircraftProgram(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	typeCertificateId,_ := strconv.ParseUint( vars["typeCertificateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.AssignTypeCertificateToAircraftProgram(aircraftProgramId, typeCertificateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a TypeCertificate on a AircraftProgram
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTypeCertificateFromAircraftProgram( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.UnassignTypeCertificateFromAircraftProgram(aircraftProgramId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more aircraftFamiliesIds as a AircraftFamilies to a AircraftProgram
	//----------------------------------------------------------------------------
func AddAircraftFamiliesToAircraftProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftFamiliesIds,_ := vars["aircraftFamiliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.AddAircraftFamiliesToAircraftProgram(aircraftProgramId, aircraftFamiliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more aircraftFamiliesIds as a AircraftFamilies from a AircraftProgram
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAircraftFamiliesFromAircraftProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftFamiliesIds,_ := vars["aircraftFamiliesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.RemoveAircraftFamiliesFromAircraftProgram(aircraftProgramId, aircraftFamiliesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more keySuppliersIds as a KeySuppliers to a AircraftProgram
	//----------------------------------------------------------------------------
func AddKeySuppliersToAircraftProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	keySuppliersIds,_ := vars["keySuppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.AddKeySuppliersToAircraftProgram(aircraftProgramId, keySuppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more keySuppliersIds as a KeySuppliers from a AircraftProgram
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveKeySuppliersFromAircraftProgram(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftProgramId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	keySuppliersIds,_ := vars["keySuppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftProgram DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftProgramDAO.RemoveKeySuppliersFromAircraftProgram(aircraftProgramId, keySuppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
