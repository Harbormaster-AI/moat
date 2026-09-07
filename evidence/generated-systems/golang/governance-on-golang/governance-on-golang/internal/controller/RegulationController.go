package controller

import (
    RegulationDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RegulationDAO for database creation
//----------------------------------------------------------------------------
func CreateRegulation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Regulation model
	//----------------------------------------------------------------------------
	data := model.Regulation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Regulation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Regulation data access object to create
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.CreateRegulation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RegulationDAO to find the relevant Regulation
//----------------------------------------------------------------------------
func GetRegulation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Regulation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.GetRegulation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RegulationDAO for database read of all Regulations
//----------------------------------------------------------------------------
func GetAllRegulation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Regulation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.GetAllRegulation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RegulationDAO for database save
//----------------------------------------------------------------------------
func UpdateRegulation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Regulation model
	//----------------------------------------------------------------------------
	var data = model.Regulation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Regulation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Regulation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.UpdateRegulation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RegulationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRegulation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Regulation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RegulationDAO.DeleteRegulation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more obligationsIds as a Obligations to a Regulation
	//----------------------------------------------------------------------------
func AddObligationsToRegulation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	regulationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Regulation DAO
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.AddObligationsToRegulation(regulationId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more obligationsIds as a Obligations from a Regulation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObligationsFromRegulation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	regulationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Regulation DAO
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.RemoveObligationsFromRegulation(regulationId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more complianceProgramsIds as a CompliancePrograms to a Regulation
	//----------------------------------------------------------------------------
func AddComplianceProgramsToRegulation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	regulationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	complianceProgramsIds,_ := vars["complianceProgramsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Regulation DAO
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.AddComplianceProgramsToRegulation(regulationId, complianceProgramsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more complianceProgramsIds as a CompliancePrograms from a Regulation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveComplianceProgramsFromRegulation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	regulationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	complianceProgramsIds,_ := vars["complianceProgramsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Regulation DAO
	//----------------------------------------------------------------------------
	requestResult := RegulationDAO.RemoveComplianceProgramsFromRegulation(regulationId, complianceProgramsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
