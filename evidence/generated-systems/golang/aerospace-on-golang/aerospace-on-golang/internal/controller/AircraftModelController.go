package controller

import (
    AircraftModelDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftModelDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraftModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftModel model
	//----------------------------------------------------------------------------
	data := model.AircraftModel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftModel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.CreateAircraftModel( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftModelDAO to find the relevant AircraftModel
//----------------------------------------------------------------------------
func GetAircraftModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftModel data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.GetAircraftModel(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftModelDAO for database read of all AircraftModels
//----------------------------------------------------------------------------
func GetAllAircraftModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.GetAllAircraftModel()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftModelDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraftModel(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftModel model
	//----------------------------------------------------------------------------
	var data = model.AircraftModel{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftModel model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.UpdateAircraftModel(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftModelDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraftModel(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftModel data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftModelDAO.DeleteAircraftModel(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Family on a AircraftModel
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFamilyToAircraftModel(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	familyId,_ := strconv.ParseUint( vars["familyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.AssignFamilyToAircraftModel(aircraftModelId, familyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Family on a AircraftModel
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFamilyFromAircraftModel( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.UnassignFamilyFromAircraftModel(aircraftModelId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more variantsIds as a Variants to a AircraftModel
	//----------------------------------------------------------------------------
func AddVariantsToAircraftModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.AddVariantsToAircraftModel(aircraftModelId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variantsIds as a Variants from a AircraftModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariantsFromAircraftModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.RemoveVariantsFromAircraftModel(aircraftModelId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more engineTypesIds as a EngineTypes to a AircraftModel
	//----------------------------------------------------------------------------
func AddEngineTypesToAircraftModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engineTypesIds,_ := vars["engineTypesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.AddEngineTypesToAircraftModel(aircraftModelId, engineTypesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more engineTypesIds as a EngineTypes from a AircraftModel
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveEngineTypesFromAircraftModel(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftModelId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	engineTypesIds,_ := vars["engineTypesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftModel DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftModelDAO.RemoveEngineTypesFromAircraftModel(aircraftModelId, engineTypesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
