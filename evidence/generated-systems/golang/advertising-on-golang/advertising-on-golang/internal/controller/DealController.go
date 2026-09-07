package controller

import (
    DealDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DealDAO for database creation
//----------------------------------------------------------------------------
func CreateDeal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Deal model
	//----------------------------------------------------------------------------
	data := model.Deal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Deal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Deal data access object to create
	//----------------------------------------------------------------------------
	requestResult := DealDAO.CreateDeal( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DealDAO to find the relevant Deal
//----------------------------------------------------------------------------
func GetDeal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Deal data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DealDAO.GetDeal(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DealDAO for database read of all Deals
//----------------------------------------------------------------------------
func GetAllDeal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Deal data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DealDAO.GetAllDeal()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DealDAO for database save
//----------------------------------------------------------------------------
func UpdateDeal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Deal model
	//----------------------------------------------------------------------------
	var data = model.Deal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Deal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Deal data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DealDAO.UpdateDeal(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DealDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDeal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Deal data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DealDAO.DeleteDeal(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Publisher on a Deal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPublisherToDeal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	publisherId,_ := strconv.ParseUint( vars["publisherId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Deal DAO
	//----------------------------------------------------------------------------
	requestResult := DealDAO.AssignPublisherToDeal(dealId, publisherId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Publisher on a Deal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPublisherFromDeal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Deal DAO
	//----------------------------------------------------------------------------
	requestResult := DealDAO.UnassignPublisherFromDeal(dealId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more inventorySourcesIds as a InventorySources to a Deal
	//----------------------------------------------------------------------------
func AddInventorySourcesToDeal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventorySourcesIds,_ := vars["inventorySourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Deal DAO
	//----------------------------------------------------------------------------
	requestResult := DealDAO.AddInventorySourcesToDeal(dealId, inventorySourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventorySourcesIds as a InventorySources from a Deal
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventorySourcesFromDeal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventorySourcesIds,_ := vars["inventorySourcesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Deal DAO
	//----------------------------------------------------------------------------
	requestResult := DealDAO.RemoveInventorySourcesFromDeal(dealId, inventorySourcesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more placementsIds as a Placements to a Deal
	//----------------------------------------------------------------------------
func AddPlacementsToDeal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	placementsIds,_ := vars["placementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Deal DAO
	//----------------------------------------------------------------------------
	requestResult := DealDAO.AddPlacementsToDeal(dealId, placementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more placementsIds as a Placements from a Deal
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlacementsFromDeal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dealId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	placementsIds,_ := vars["placementsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Deal DAO
	//----------------------------------------------------------------------------
	requestResult := DealDAO.RemovePlacementsFromDeal(dealId, placementsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
