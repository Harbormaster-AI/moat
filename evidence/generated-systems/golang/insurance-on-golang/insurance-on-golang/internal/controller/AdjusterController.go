package controller

import (
    AdjusterDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AdjusterDAO for database creation
//----------------------------------------------------------------------------
func CreateAdjuster(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Adjuster model
	//----------------------------------------------------------------------------
	data := model.Adjuster{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Adjuster model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Adjuster data access object to create
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.CreateAdjuster( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AdjusterDAO to find the relevant Adjuster
//----------------------------------------------------------------------------
func GetAdjuster(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Adjuster data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.GetAdjuster(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AdjusterDAO for database read of all Adjusters
//----------------------------------------------------------------------------
func GetAllAdjuster(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Adjuster data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.GetAllAdjuster()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AdjusterDAO for database save
//----------------------------------------------------------------------------
func UpdateAdjuster(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Adjuster model
	//----------------------------------------------------------------------------
	var data = model.Adjuster{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Adjuster model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Adjuster data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.UpdateAdjuster(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AdjusterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAdjuster(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Adjuster data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AdjusterDAO.DeleteAdjuster(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more claimsIds as a Claims to a Adjuster
	//----------------------------------------------------------------------------
func AddClaimsToAdjuster(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adjusterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Adjuster DAO
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.AddClaimsToAdjuster(adjusterId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more claimsIds as a Claims from a Adjuster
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveClaimsFromAdjuster(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adjusterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimsIds,_ := vars["claimsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Adjuster DAO
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.RemoveClaimsFromAdjuster(adjusterId, claimsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more serviceProvidersIds as a ServiceProviders to a Adjuster
	//----------------------------------------------------------------------------
func AddServiceProvidersToAdjuster(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adjusterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serviceProvidersIds,_ := vars["serviceProvidersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Adjuster DAO
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.AddServiceProvidersToAdjuster(adjusterId, serviceProvidersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serviceProvidersIds as a ServiceProviders from a Adjuster
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveServiceProvidersFromAdjuster(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adjusterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serviceProvidersIds,_ := vars["serviceProvidersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Adjuster DAO
	//----------------------------------------------------------------------------
	requestResult := AdjusterDAO.RemoveServiceProvidersFromAdjuster(adjusterId, serviceProvidersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
