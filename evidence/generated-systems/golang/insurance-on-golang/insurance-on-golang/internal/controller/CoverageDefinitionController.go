package controller

import (
    CoverageDefinitionDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CoverageDefinitionDAO for database creation
//----------------------------------------------------------------------------
func CreateCoverageDefinition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CoverageDefinition model
	//----------------------------------------------------------------------------
	data := model.CoverageDefinition{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CoverageDefinition model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CoverageDefinition data access object to create
	//----------------------------------------------------------------------------
	requestResult := CoverageDefinitionDAO.CreateCoverageDefinition( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CoverageDefinitionDAO to find the relevant CoverageDefinition
//----------------------------------------------------------------------------
func GetCoverageDefinition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CoverageDefinition data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CoverageDefinitionDAO.GetCoverageDefinition(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CoverageDefinitionDAO for database read of all CoverageDefinitions
//----------------------------------------------------------------------------
func GetAllCoverageDefinition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CoverageDefinition data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CoverageDefinitionDAO.GetAllCoverageDefinition()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CoverageDefinitionDAO for database save
//----------------------------------------------------------------------------
func UpdateCoverageDefinition(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CoverageDefinition model
	//----------------------------------------------------------------------------
	var data = model.CoverageDefinition{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CoverageDefinition model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CoverageDefinition data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CoverageDefinitionDAO.UpdateCoverageDefinition(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CoverageDefinitionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCoverageDefinition(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CoverageDefinition data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CoverageDefinitionDAO.DeleteCoverageDefinition(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Product on a CoverageDefinition
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToCoverageDefinition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	coverageDefinitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CoverageDefinition DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDefinitionDAO.AssignProductToCoverageDefinition(coverageDefinitionId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a CoverageDefinition
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromCoverageDefinition( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	coverageDefinitionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CoverageDefinition DAO
	//----------------------------------------------------------------------------
	requestResult := CoverageDefinitionDAO.UnassignProductFromCoverageDefinition(coverageDefinitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


