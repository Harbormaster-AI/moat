package controller

import (
    ReservationDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ReservationDAO for database creation
//----------------------------------------------------------------------------
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Reservation model
	//----------------------------------------------------------------------------
	data := model.Reservation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Reservation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation data access object to create
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.CreateReservation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ReservationDAO to find the relevant Reservation
//----------------------------------------------------------------------------
func GetReservation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Reservation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.GetReservation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ReservationDAO for database read of all Reservations
//----------------------------------------------------------------------------
func GetAllReservation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Reservation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.GetAllReservation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ReservationDAO for database save
//----------------------------------------------------------------------------
func UpdateReservation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Reservation model
	//----------------------------------------------------------------------------
	var data = model.Reservation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Reservation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.UpdateReservation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ReservationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteReservation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Reservation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ReservationDAO.DeleteReservation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a Reservation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToReservation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.AssignSkuToReservation(reservationId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a Reservation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromReservation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.UnassignSkuFromReservation(reservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Warehouse on a Reservation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWarehouseToReservation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehouseId,_ := strconv.ParseUint( vars["warehouseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.AssignWarehouseToReservation(reservationId, warehouseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Warehouse on a Reservation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWarehouseFromReservation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.UnassignWarehouseFromReservation(reservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Location on a Reservation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLocationToReservation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	locationId,_ := strconv.ParseUint( vars["locationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.AssignLocationToReservation(reservationId, locationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Location on a Reservation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLocationFromReservation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.UnassignLocationFromReservation(reservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a InventoryItem on a Reservation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInventoryItemToReservation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemId,_ := strconv.ParseUint( vars["inventoryItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.AssignInventoryItemToReservation(reservationId, inventoryItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InventoryItem on a Reservation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInventoryItemFromReservation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.UnassignInventoryItemFromReservation(reservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a Reservation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToReservation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.AssignLotToReservation(reservationId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a Reservation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromReservation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.UnassignLotFromReservation(reservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DemandSignal on a Reservation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDemandSignalToReservation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	demandSignalId,_ := strconv.ParseUint( vars["demandSignalId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.AssignDemandSignalToReservation(reservationId, demandSignalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DemandSignal on a Reservation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDemandSignalFromReservation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.UnassignDemandSignalFromReservation(reservationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a Reservation
	//----------------------------------------------------------------------------
func AddSerialNumbersToReservation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.AddSerialNumbersToReservation(reservationId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a Reservation
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromReservation(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reservationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Reservation DAO
	//----------------------------------------------------------------------------
	requestResult := ReservationDAO.RemoveSerialNumbersFromReservation(reservationId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
