package controller

import (
    EquityGrantDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EquityGrantDAO for database creation
//----------------------------------------------------------------------------
func CreateEquityGrant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EquityGrant model
	//----------------------------------------------------------------------------
	data := model.EquityGrant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EquityGrant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EquityGrant data access object to create
	//----------------------------------------------------------------------------
	requestResult := EquityGrantDAO.CreateEquityGrant( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EquityGrantDAO to find the relevant EquityGrant
//----------------------------------------------------------------------------
func GetEquityGrant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EquityGrant data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EquityGrantDAO.GetEquityGrant(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EquityGrantDAO for database read of all EquityGrants
//----------------------------------------------------------------------------
func GetAllEquityGrant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the EquityGrant data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EquityGrantDAO.GetAllEquityGrant()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EquityGrantDAO for database save
//----------------------------------------------------------------------------
func UpdateEquityGrant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EquityGrant model
	//----------------------------------------------------------------------------
	var data = model.EquityGrant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EquityGrant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EquityGrant data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EquityGrantDAO.UpdateEquityGrant(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EquityGrantDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEquityGrant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EquityGrant data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EquityGrantDAO.DeleteEquityGrant(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a CompensationPackage on a EquityGrant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCompensationPackageToEquityGrant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	equityGrantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compensationPackageId,_ := strconv.ParseUint( vars["compensationPackageId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EquityGrant DAO
	//----------------------------------------------------------------------------
	requestResult := EquityGrantDAO.AssignCompensationPackageToEquityGrant(equityGrantId, compensationPackageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CompensationPackage on a EquityGrant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCompensationPackageFromEquityGrant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	equityGrantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EquityGrant DAO
	//----------------------------------------------------------------------------
	requestResult := EquityGrantDAO.UnassignCompensationPackageFromEquityGrant(equityGrantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


