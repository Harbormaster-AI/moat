package controller

import (
    NetworkProfileDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to NetworkProfileDAO for database creation
//----------------------------------------------------------------------------
func CreateNetworkProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty NetworkProfile model
	//----------------------------------------------------------------------------
	data := model.NetworkProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a NetworkProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile data access object to create
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.CreateNetworkProfile( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to NetworkProfileDAO to find the relevant NetworkProfile
//----------------------------------------------------------------------------
func GetNetworkProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the NetworkProfile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.GetNetworkProfile(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to NetworkProfileDAO for database read of all NetworkProfiles
//----------------------------------------------------------------------------
func GetAllNetworkProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile data access object to get all
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.GetAllNetworkProfile()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to NetworkProfileDAO for database save
//----------------------------------------------------------------------------
func UpdateNetworkProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty NetworkProfile model
	//----------------------------------------------------------------------------
	var data = model.NetworkProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a NetworkProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.UpdateNetworkProfile(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to NetworkProfileDAO for database deletion
//----------------------------------------------------------------------------
func DeleteNetworkProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the NetworkProfile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := NetworkProfileDAO.DeleteNetworkProfile(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Device on a NetworkProfile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDeviceToNetworkProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	networkProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceId,_ := strconv.ParseUint( vars["deviceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile DAO
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.AssignDeviceToNetworkProfile(networkProfileId, deviceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Device on a NetworkProfile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDeviceFromNetworkProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	networkProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile DAO
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.UnassignDeviceFromNetworkProfile(networkProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Gateway on a NetworkProfile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignGatewayToNetworkProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	networkProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	gatewayId,_ := strconv.ParseUint( vars["gatewayId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile DAO
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.AssignGatewayToNetworkProfile(networkProfileId, gatewayId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Gateway on a NetworkProfile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignGatewayFromNetworkProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	networkProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile DAO
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.UnassignGatewayFromNetworkProfile(networkProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a SimCard on a NetworkProfile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSimCardToNetworkProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	networkProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	simCardId,_ := strconv.ParseUint( vars["simCardId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile DAO
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.AssignSimCardToNetworkProfile(networkProfileId, simCardId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SimCard on a NetworkProfile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSimCardFromNetworkProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	networkProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the NetworkProfile DAO
	//----------------------------------------------------------------------------
	requestResult := NetworkProfileDAO.UnassignSimCardFromNetworkProfile(networkProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


