package controller

import (
    SimCardDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SimCardDAO for database creation
//----------------------------------------------------------------------------
func CreateSimCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SimCard model
	//----------------------------------------------------------------------------
	data := model.SimCard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SimCard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SimCard data access object to create
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.CreateSimCard( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SimCardDAO to find the relevant SimCard
//----------------------------------------------------------------------------
func GetSimCard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SimCard data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.GetSimCard(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SimCardDAO for database read of all SimCards
//----------------------------------------------------------------------------
func GetAllSimCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SimCard data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.GetAllSimCard()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SimCardDAO for database save
//----------------------------------------------------------------------------
func UpdateSimCard(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SimCard model
	//----------------------------------------------------------------------------
	var data = model.SimCard{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SimCard model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SimCard data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.UpdateSimCard(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SimCardDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSimCard(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SimCard data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SimCardDAO.DeleteSimCard(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Tenant on a SimCard
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignTenantToSimCard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	simCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tenantId,_ := strconv.ParseUint( vars["tenantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SimCard DAO
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.AssignTenantToSimCard(simCardId, tenantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Tenant on a SimCard
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignTenantFromSimCard( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	simCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SimCard DAO
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.UnassignTenantFromSimCard(simCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ConnectivityPlan on a SimCard
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignConnectivityPlanToSimCard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	simCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	connectivityPlanId,_ := strconv.ParseUint( vars["connectivityPlanId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SimCard DAO
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.AssignConnectivityPlanToSimCard(simCardId, connectivityPlanId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ConnectivityPlan on a SimCard
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignConnectivityPlanFromSimCard( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	simCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SimCard DAO
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.UnassignConnectivityPlanFromSimCard(simCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more networkProfilesIds as a NetworkProfiles to a SimCard
	//----------------------------------------------------------------------------
func AddNetworkProfilesToSimCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	simCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	networkProfilesIds,_ := vars["networkProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SimCard DAO
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.AddNetworkProfilesToSimCard(simCardId, networkProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more networkProfilesIds as a NetworkProfiles from a SimCard
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveNetworkProfilesFromSimCard(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	simCardId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	networkProfilesIds,_ := vars["networkProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SimCard DAO
	//----------------------------------------------------------------------------
	requestResult := SimCardDAO.RemoveNetworkProfilesFromSimCard(simCardId, networkProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
