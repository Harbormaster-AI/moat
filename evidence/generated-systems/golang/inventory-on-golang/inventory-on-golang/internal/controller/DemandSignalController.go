package controller

import (
    DemandSignalDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DemandSignalDAO for database creation
//----------------------------------------------------------------------------
func CreateDemandSignal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DemandSignal model
	//----------------------------------------------------------------------------
	data := model.DemandSignal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DemandSignal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DemandSignal data access object to create
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.CreateDemandSignal( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DemandSignalDAO to find the relevant DemandSignal
//----------------------------------------------------------------------------
func GetDemandSignal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DemandSignal data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.GetDemandSignal(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DemandSignalDAO for database read of all DemandSignals
//----------------------------------------------------------------------------
func GetAllDemandSignal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DemandSignal data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.GetAllDemandSignal()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DemandSignalDAO for database save
//----------------------------------------------------------------------------
func UpdateDemandSignal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DemandSignal model
	//----------------------------------------------------------------------------
	var data = model.DemandSignal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DemandSignal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DemandSignal data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.UpdateDemandSignal(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DemandSignalDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDemandSignal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DemandSignal data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DemandSignalDAO.DeleteDemandSignal(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a DemandSignal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToDemandSignal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	demandSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DemandSignal DAO
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.AssignSkuToDemandSignal(demandSignalId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a DemandSignal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromDemandSignal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	demandSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DemandSignal DAO
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.UnassignSkuFromDemandSignal(demandSignalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more reservationsIds as a Reservations to a DemandSignal
	//----------------------------------------------------------------------------
func AddReservationsToDemandSignal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	demandSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservationsIds,_ := vars["reservationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DemandSignal DAO
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.AddReservationsToDemandSignal(demandSignalId, reservationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more reservationsIds as a Reservations from a DemandSignal
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveReservationsFromDemandSignal(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	demandSignalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	reservationsIds,_ := vars["reservationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DemandSignal DAO
	//----------------------------------------------------------------------------
	requestResult := DemandSignalDAO.RemoveReservationsFromDemandSignal(demandSignalId, reservationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
