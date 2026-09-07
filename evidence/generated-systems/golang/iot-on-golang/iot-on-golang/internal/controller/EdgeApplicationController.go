package controller

import (
    EdgeApplicationDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EdgeApplicationDAO for database creation
//----------------------------------------------------------------------------
func CreateEdgeApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EdgeApplication model
	//----------------------------------------------------------------------------
	data := model.EdgeApplication{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EdgeApplication model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EdgeApplication data access object to create
	//----------------------------------------------------------------------------
	requestResult := EdgeApplicationDAO.CreateEdgeApplication( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EdgeApplicationDAO to find the relevant EdgeApplication
//----------------------------------------------------------------------------
func GetEdgeApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EdgeApplication data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EdgeApplicationDAO.GetEdgeApplication(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EdgeApplicationDAO for database read of all EdgeApplications
//----------------------------------------------------------------------------
func GetAllEdgeApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the EdgeApplication data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EdgeApplicationDAO.GetAllEdgeApplication()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EdgeApplicationDAO for database save
//----------------------------------------------------------------------------
func UpdateEdgeApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EdgeApplication model
	//----------------------------------------------------------------------------
	var data = model.EdgeApplication{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EdgeApplication model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EdgeApplication data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EdgeApplicationDAO.UpdateEdgeApplication(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EdgeApplicationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEdgeApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EdgeApplication data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EdgeApplicationDAO.DeleteEdgeApplication(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Gateway on a EdgeApplication
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignGatewayToEdgeApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	edgeApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	gatewayId,_ := strconv.ParseUint( vars["gatewayId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EdgeApplication DAO
	//----------------------------------------------------------------------------
	requestResult := EdgeApplicationDAO.AssignGatewayToEdgeApplication(edgeApplicationId, gatewayId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Gateway on a EdgeApplication
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignGatewayFromEdgeApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	edgeApplicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EdgeApplication DAO
	//----------------------------------------------------------------------------
	requestResult := EdgeApplicationDAO.UnassignGatewayFromEdgeApplication(edgeApplicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


