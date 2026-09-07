package controller

import (
    AircraftOptionDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftOptionDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraftOption(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftOption model
	//----------------------------------------------------------------------------
	data := model.AircraftOption{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftOption model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOption data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.CreateAircraftOption( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftOptionDAO to find the relevant AircraftOption
//----------------------------------------------------------------------------
func GetAircraftOption(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftOption data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.GetAircraftOption(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftOptionDAO for database read of all AircraftOptions
//----------------------------------------------------------------------------
func GetAllAircraftOption(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AircraftOption data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.GetAllAircraftOption()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftOptionDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraftOption(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftOption model
	//----------------------------------------------------------------------------
	var data = model.AircraftOption{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftOption model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOption data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.UpdateAircraftOption(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftOptionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraftOption(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftOption data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftOptionDAO.DeleteAircraftOption(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more variantsIds as a Variants to a AircraftOption
	//----------------------------------------------------------------------------
func AddVariantsToAircraftOption(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftOptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOption DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.AddVariantsToAircraftOption(aircraftOptionId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variantsIds as a Variants from a AircraftOption
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariantsFromAircraftOption(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftOptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOption DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.RemoveVariantsFromAircraftOption(aircraftOptionId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more packagesIds as a Packages to a AircraftOption
	//----------------------------------------------------------------------------
func AddPackagesToAircraftOption(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftOptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	packagesIds,_ := vars["packagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOption DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.AddPackagesToAircraftOption(aircraftOptionId, packagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more packagesIds as a Packages from a AircraftOption
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePackagesFromAircraftOption(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftOptionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	packagesIds,_ := vars["packagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftOption DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftOptionDAO.RemovePackagesFromAircraftOption(aircraftOptionId, packagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
