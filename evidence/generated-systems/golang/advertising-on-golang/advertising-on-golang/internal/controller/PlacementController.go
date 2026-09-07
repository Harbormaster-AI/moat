package controller

import (
    PlacementDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PlacementDAO for database creation
//----------------------------------------------------------------------------
func CreatePlacement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Placement model
	//----------------------------------------------------------------------------
	data := model.Placement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Placement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Placement data access object to create
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.CreatePlacement( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PlacementDAO to find the relevant Placement
//----------------------------------------------------------------------------
func GetPlacement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Placement data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.GetPlacement(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PlacementDAO for database read of all Placements
//----------------------------------------------------------------------------
func GetAllPlacement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Placement data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.GetAllPlacement()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PlacementDAO for database save
//----------------------------------------------------------------------------
func UpdatePlacement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Placement model
	//----------------------------------------------------------------------------
	var data = model.Placement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Placement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Placement data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.UpdatePlacement(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PlacementDAO for database deletion
//----------------------------------------------------------------------------
func DeletePlacement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Placement data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PlacementDAO.DeletePlacement(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a LineItem on a Placement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLineItemToPlacement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	placementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemId,_ := strconv.ParseUint( vars["lineItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Placement DAO
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.AssignLineItemToPlacement(placementId, lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LineItem on a Placement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLineItemFromPlacement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	placementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Placement DAO
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.UnassignLineItemFromPlacement(placementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AdSlot on a Placement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdSlotToPlacement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	placementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adSlotId,_ := strconv.ParseUint( vars["adSlotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Placement DAO
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.AssignAdSlotToPlacement(placementId, adSlotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AdSlot on a Placement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdSlotFromPlacement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	placementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Placement DAO
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.UnassignAdSlotFromPlacement(placementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Deal on a Placement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDealToPlacement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	placementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dealId,_ := strconv.ParseUint( vars["dealId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Placement DAO
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.AssignDealToPlacement(placementId, dealId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Deal on a Placement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDealFromPlacement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	placementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Placement DAO
	//----------------------------------------------------------------------------
	requestResult := PlacementDAO.UnassignDealFromPlacement(placementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


