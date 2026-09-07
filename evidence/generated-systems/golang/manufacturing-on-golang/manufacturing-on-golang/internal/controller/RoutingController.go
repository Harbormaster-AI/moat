package controller

import (
    RoutingDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RoutingDAO for database creation
//----------------------------------------------------------------------------
func CreateRouting(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Routing model
	//----------------------------------------------------------------------------
	data := model.Routing{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Routing model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Routing data access object to create
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.CreateRouting( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RoutingDAO to find the relevant Routing
//----------------------------------------------------------------------------
func GetRouting(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Routing data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.GetRouting(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RoutingDAO for database read of all Routings
//----------------------------------------------------------------------------
func GetAllRouting(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Routing data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.GetAllRouting()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RoutingDAO for database save
//----------------------------------------------------------------------------
func UpdateRouting(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Routing model
	//----------------------------------------------------------------------------
	var data = model.Routing{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Routing model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Routing data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.UpdateRouting(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RoutingDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRouting(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Routing data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RoutingDAO.DeleteRouting(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Item on a Routing
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToRouting(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	routingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Routing DAO
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.AssignItemToRouting(routingId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a Routing
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromRouting( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	routingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Routing DAO
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.UnassignItemFromRouting(routingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more operationsIds as a Operations to a Routing
	//----------------------------------------------------------------------------
func AddOperationsToRouting(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	routingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operationsIds,_ := vars["operationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Routing DAO
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.AddOperationsToRouting(routingId, operationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more operationsIds as a Operations from a Routing
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOperationsFromRouting(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	routingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operationsIds,_ := vars["operationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Routing DAO
	//----------------------------------------------------------------------------
	requestResult := RoutingDAO.RemoveOperationsFromRouting(routingId, operationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
