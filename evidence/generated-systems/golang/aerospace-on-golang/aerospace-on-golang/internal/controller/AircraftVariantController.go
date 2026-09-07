package controller

import (
    AircraftVariantDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftVariantDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraftVariant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftVariant model
	//----------------------------------------------------------------------------
	data := model.AircraftVariant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftVariant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.CreateAircraftVariant( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftVariantDAO to find the relevant AircraftVariant
//----------------------------------------------------------------------------
func GetAircraftVariant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftVariant data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.GetAircraftVariant(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftVariantDAO for database read of all AircraftVariants
//----------------------------------------------------------------------------
func GetAllAircraftVariant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.GetAllAircraftVariant()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftVariantDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraftVariant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftVariant model
	//----------------------------------------------------------------------------
	var data = model.AircraftVariant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftVariant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.UpdateAircraftVariant(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftVariantDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraftVariant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftVariant data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftVariantDAO.DeleteAircraftVariant(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Model_ on a AircraftVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignModel_ToAircraftVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	model_Id,_ := strconv.ParseUint( vars["model_Id"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AssignModel_ToAircraftVariant(aircraftVariantId, model_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Model_ on a AircraftVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignModel_FromAircraftVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.UnassignModel_FromAircraftVariant(aircraftVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a EngineType on a AircraftVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEngineTypeToAircraftVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engineTypeId,_ := strconv.ParseUint( vars["engineTypeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AssignEngineTypeToAircraftVariant(aircraftVariantId, engineTypeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a EngineType on a AircraftVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEngineTypeFromAircraftVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.UnassignEngineTypeFromAircraftVariant(aircraftVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AvionicsSuite on a AircraftVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAvionicsSuiteToAircraftVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	avionicsSuiteId,_ := strconv.ParseUint( vars["avionicsSuiteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AssignAvionicsSuiteToAircraftVariant(aircraftVariantId, avionicsSuiteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AvionicsSuite on a AircraftVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAvionicsSuiteFromAircraftVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.UnassignAvionicsSuiteFromAircraftVariant(aircraftVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Apu on a AircraftVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignApuToAircraftVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	apuId,_ := strconv.ParseUint( vars["apuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AssignApuToAircraftVariant(aircraftVariantId, apuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Apu on a AircraftVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignApuFromAircraftVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.UnassignApuFromAircraftVariant(aircraftVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LandingGear on a AircraftVariant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLandingGearToAircraftVariant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	landingGearId,_ := strconv.ParseUint( vars["landingGearId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AssignLandingGearToAircraftVariant(aircraftVariantId, landingGearId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LandingGear on a AircraftVariant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLandingGearFromAircraftVariant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.UnassignLandingGearFromAircraftVariant(aircraftVariantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more cabinLayoutsIds as a CabinLayouts to a AircraftVariant
	//----------------------------------------------------------------------------
func AddCabinLayoutsToAircraftVariant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cabinLayoutsIds,_ := vars["cabinLayoutsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AddCabinLayoutsToAircraftVariant(aircraftVariantId, cabinLayoutsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more cabinLayoutsIds as a CabinLayouts from a AircraftVariant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCabinLayoutsFromAircraftVariant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	cabinLayoutsIds,_ := vars["cabinLayoutsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.RemoveCabinLayoutsFromAircraftVariant(aircraftVariantId, cabinLayoutsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more optionsIds as a Options to a AircraftVariant
	//----------------------------------------------------------------------------
func AddOptionsToAircraftVariant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	optionsIds,_ := vars["optionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AddOptionsToAircraftVariant(aircraftVariantId, optionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more optionsIds as a Options from a AircraftVariant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOptionsFromAircraftVariant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	optionsIds,_ := vars["optionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.RemoveOptionsFromAircraftVariant(aircraftVariantId, optionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more packagesIds as a Packages to a AircraftVariant
	//----------------------------------------------------------------------------
func AddPackagesToAircraftVariant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	packagesIds,_ := vars["packagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.AddPackagesToAircraftVariant(aircraftVariantId, packagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more packagesIds as a Packages from a AircraftVariant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePackagesFromAircraftVariant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftVariantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	packagesIds,_ := vars["packagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftVariant DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftVariantDAO.RemovePackagesFromAircraftVariant(aircraftVariantId, packagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
