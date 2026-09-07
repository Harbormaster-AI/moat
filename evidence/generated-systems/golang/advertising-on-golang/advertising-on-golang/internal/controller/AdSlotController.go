package controller

import (
    AdSlotDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AdSlotDAO for database creation
//----------------------------------------------------------------------------
func CreateAdSlot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AdSlot model
	//----------------------------------------------------------------------------
	data := model.AdSlot{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AdSlot model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot data access object to create
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.CreateAdSlot( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AdSlotDAO to find the relevant AdSlot
//----------------------------------------------------------------------------
func GetAdSlot(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AdSlot data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.GetAdSlot(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AdSlotDAO for database read of all AdSlots
//----------------------------------------------------------------------------
func GetAllAdSlot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AdSlot data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.GetAllAdSlot()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AdSlotDAO for database save
//----------------------------------------------------------------------------
func UpdateAdSlot(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AdSlot model
	//----------------------------------------------------------------------------
	var data = model.AdSlot{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AdSlot model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.UpdateAdSlot(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AdSlotDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAdSlot(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AdSlot data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AdSlotDAO.DeleteAdSlot(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a InventorySource on a AdSlot
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInventorySourceToAdSlot(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adSlotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventorySourceId,_ := strconv.ParseUint( vars["inventorySourceId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot DAO
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.AssignInventorySourceToAdSlot(adSlotId, inventorySourceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InventorySource on a AdSlot
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInventorySourceFromAdSlot( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	adSlotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot DAO
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.UnassignInventorySourceFromAdSlot(adSlotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more placementsIds as a Placements to a AdSlot
	//----------------------------------------------------------------------------
func AddPlacementsToAdSlot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adSlotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	placementsIds,_ := vars["placementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot DAO
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.AddPlacementsToAdSlot(adSlotId, placementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more placementsIds as a Placements from a AdSlot
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlacementsFromAdSlot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adSlotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	placementsIds,_ := vars["placementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot DAO
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.RemovePlacementsFromAdSlot(adSlotId, placementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ratesIds as a Rates to a AdSlot
	//----------------------------------------------------------------------------
func AddRatesToAdSlot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adSlotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ratesIds,_ := vars["ratesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot DAO
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.AddRatesToAdSlot(adSlotId, ratesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ratesIds as a Rates from a AdSlot
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRatesFromAdSlot(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	adSlotId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ratesIds,_ := vars["ratesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AdSlot DAO
	//----------------------------------------------------------------------------
	requestResult := AdSlotDAO.RemoveRatesFromAdSlot(adSlotId, ratesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
