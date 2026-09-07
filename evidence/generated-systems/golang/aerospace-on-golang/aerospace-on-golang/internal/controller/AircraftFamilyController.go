package controller

import (
    AircraftFamilyDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AircraftFamilyDAO for database creation
//----------------------------------------------------------------------------
func CreateAircraftFamily(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftFamily model
	//----------------------------------------------------------------------------
	data := model.AircraftFamily{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftFamily model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftFamily data access object to create
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.CreateAircraftFamily( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AircraftFamilyDAO to find the relevant AircraftFamily
//----------------------------------------------------------------------------
func GetAircraftFamily(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftFamily data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.GetAircraftFamily(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AircraftFamilyDAO for database read of all AircraftFamilys
//----------------------------------------------------------------------------
func GetAllAircraftFamily(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AircraftFamily data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.GetAllAircraftFamily()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AircraftFamilyDAO for database save
//----------------------------------------------------------------------------
func UpdateAircraftFamily(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AircraftFamily model
	//----------------------------------------------------------------------------
	var data = model.AircraftFamily{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AircraftFamily model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftFamily data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.UpdateAircraftFamily(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AircraftFamilyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAircraftFamily(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AircraftFamily data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AircraftFamilyDAO.DeleteAircraftFamily(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Program on a AircraftFamily
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProgramToAircraftFamily(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	programId,_ := strconv.ParseUint( vars["programId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftFamily DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.AssignProgramToAircraftFamily(aircraftFamilyId, programId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Program on a AircraftFamily
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProgramFromAircraftFamily( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	aircraftFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AircraftFamily DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.UnassignProgramFromAircraftFamily(aircraftFamilyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more aircraftModelsIds as a AircraftModels to a AircraftFamily
	//----------------------------------------------------------------------------
func AddAircraftModelsToAircraftFamily(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftModelsIds,_ := vars["aircraftModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftFamily DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.AddAircraftModelsToAircraftFamily(aircraftFamilyId, aircraftModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more aircraftModelsIds as a AircraftModels from a AircraftFamily
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAircraftModelsFromAircraftFamily(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	aircraftFamilyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftModelsIds,_ := vars["aircraftModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AircraftFamily DAO
	//----------------------------------------------------------------------------
	requestResult := AircraftFamilyDAO.RemoveAircraftModelsFromAircraftFamily(aircraftFamilyId, aircraftModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
