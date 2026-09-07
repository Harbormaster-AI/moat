package controller

import (
    ThirdPartyDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ThirdPartyDAO for database creation
//----------------------------------------------------------------------------
func CreateThirdParty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdParty model
	//----------------------------------------------------------------------------
	data := model.ThirdParty{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdParty model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty data access object to create
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.CreateThirdParty( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ThirdPartyDAO to find the relevant ThirdParty
//----------------------------------------------------------------------------
func GetThirdParty(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdParty data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.GetThirdParty(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ThirdPartyDAO for database read of all ThirdPartys
//----------------------------------------------------------------------------
func GetAllThirdParty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.GetAllThirdParty()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ThirdPartyDAO for database save
//----------------------------------------------------------------------------
func UpdateThirdParty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdParty model
	//----------------------------------------------------------------------------
	var data = model.ThirdParty{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdParty model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.UpdateThirdParty(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ThirdPartyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteThirdParty(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdParty data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ThirdPartyDAO.DeleteThirdParty(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more subrogationsIds as a Subrogations to a ThirdParty
	//----------------------------------------------------------------------------
func AddSubrogationsToThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	subrogationsIds,_ := vars["subrogationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.AddSubrogationsToThirdParty(thirdPartyId, subrogationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more subrogationsIds as a Subrogations from a ThirdParty
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSubrogationsFromThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	subrogationsIds,_ := vars["subrogationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.RemoveSubrogationsFromThirdParty(thirdPartyId, subrogationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
