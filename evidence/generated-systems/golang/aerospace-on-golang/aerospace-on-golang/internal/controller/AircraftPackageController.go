package controller

import (
    AircraftPackageDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftPackageDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraftPackage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftPackage model
	//----------------------------------------------------------------------------
	data := model.AircraftPackage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftPackage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftPackage data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.CreateAircraftPackage( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftPackageDAO to find the relevant AircraftPackage
//----------------------------------------------------------------------------
func GetAircraftPackage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftPackage data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.GetAircraftPackage(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftPackageDAO for database read of all AircraftPackages
//----------------------------------------------------------------------------
func GetAllAircraftPackage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AircraftPackage data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.GetAllAircraftPackage()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftPackageDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraftPackage(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftPackage model
	//----------------------------------------------------------------------------
	var data = model.AircraftPackage{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftPackage model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftPackage data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.UpdateAircraftPackage(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftPackageDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraftPackage(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftPackage data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftPackageDAO.DeleteAircraftPackage(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more optionsIds as a Options to a AircraftPackage
	//----------------------------------------------------------------------------
func AddOptionsToAircraftPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	optionsIds,_ := vars["optionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftPackage DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.AddOptionsToAircraftPackage(aircraftPackageId, optionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more optionsIds as a Options from a AircraftPackage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOptionsFromAircraftPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	optionsIds,_ := vars["optionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftPackage DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.RemoveOptionsFromAircraftPackage(aircraftPackageId, optionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more variantsIds as a Variants to a AircraftPackage
	//----------------------------------------------------------------------------
func AddVariantsToAircraftPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftPackage DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.AddVariantsToAircraftPackage(aircraftPackageId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variantsIds as a Variants from a AircraftPackage
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariantsFromAircraftPackage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftPackageId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftPackage DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftPackageDAO.RemoveVariantsFromAircraftPackage(aircraftPackageId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
