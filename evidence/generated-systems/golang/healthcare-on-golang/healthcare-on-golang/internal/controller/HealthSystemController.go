package controller

import (
    HealthSystemDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to HealthSystemDAO for database creation
//----------------------------------------------------------------------------
func CreateHealthSystem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty HealthSystem model
	//----------------------------------------------------------------------------
	data := model.HealthSystem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a HealthSystem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the HealthSystem data access object to create
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.CreateHealthSystem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to HealthSystemDAO to find the relevant HealthSystem
//----------------------------------------------------------------------------
func GetHealthSystem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the HealthSystem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.GetHealthSystem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to HealthSystemDAO for database read of all HealthSystems
//----------------------------------------------------------------------------
func GetAllHealthSystem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the HealthSystem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.GetAllHealthSystem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to HealthSystemDAO for database save
//----------------------------------------------------------------------------
func UpdateHealthSystem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty HealthSystem model
	//----------------------------------------------------------------------------
	var data = model.HealthSystem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a HealthSystem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the HealthSystem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.UpdateHealthSystem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to HealthSystemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteHealthSystem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the HealthSystem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := HealthSystemDAO.DeleteHealthSystem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more facilitiesIds as a Facilities to a HealthSystem
	//----------------------------------------------------------------------------
func AddFacilitiesToHealthSystem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	healthSystemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilitiesIds,_ := vars["facilitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the HealthSystem DAO
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.AddFacilitiesToHealthSystem(healthSystemId, facilitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more facilitiesIds as a Facilities from a HealthSystem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFacilitiesFromHealthSystem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	healthSystemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilitiesIds,_ := vars["facilitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the HealthSystem DAO
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.RemoveFacilitiesFromHealthSystem(healthSystemId, facilitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more suppliersIds as a Suppliers to a HealthSystem
	//----------------------------------------------------------------------------
func AddSuppliersToHealthSystem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	healthSystemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the HealthSystem DAO
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.AddSuppliersToHealthSystem(healthSystemId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more suppliersIds as a Suppliers from a HealthSystem
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSuppliersFromHealthSystem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	healthSystemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the HealthSystem DAO
	//----------------------------------------------------------------------------
	requestResult := HealthSystemDAO.RemoveSuppliersFromHealthSystem(healthSystemId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
