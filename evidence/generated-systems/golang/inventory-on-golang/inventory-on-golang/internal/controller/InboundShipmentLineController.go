package controller

import (
    InboundShipmentLineDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InboundShipmentLineDAO for database creation
//----------------------------------------------------------------------------
func CreateInboundShipmentLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InboundShipmentLine model
	//----------------------------------------------------------------------------
	data := model.InboundShipmentLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InboundShipmentLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.CreateInboundShipmentLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InboundShipmentLineDAO to find the relevant InboundShipmentLine
//----------------------------------------------------------------------------
func GetInboundShipmentLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InboundShipmentLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.GetInboundShipmentLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InboundShipmentLineDAO for database read of all InboundShipmentLines
//----------------------------------------------------------------------------
func GetAllInboundShipmentLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.GetAllInboundShipmentLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InboundShipmentLineDAO for database save
//----------------------------------------------------------------------------
func UpdateInboundShipmentLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InboundShipmentLine model
	//----------------------------------------------------------------------------
	var data = model.InboundShipmentLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InboundShipmentLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.UpdateInboundShipmentLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InboundShipmentLineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInboundShipmentLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InboundShipmentLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InboundShipmentLineDAO.DeleteInboundShipmentLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a InboundShipment on a InboundShipmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInboundShipmentToInboundShipmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inboundShipmentId,_ := strconv.ParseUint( vars["inboundShipmentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.AssignInboundShipmentToInboundShipmentLine(inboundShipmentLineId, inboundShipmentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a InboundShipment on a InboundShipmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInboundShipmentFromInboundShipmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.UnassignInboundShipmentFromInboundShipmentLine(inboundShipmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Sku on a InboundShipmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToInboundShipmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.AssignSkuToInboundShipmentLine(inboundShipmentLineId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a InboundShipmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromInboundShipmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.UnassignSkuFromInboundShipmentLine(inboundShipmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lot on a InboundShipmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLotToInboundShipmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lotId,_ := strconv.ParseUint( vars["lotId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.AssignLotToInboundShipmentLine(inboundShipmentLineId, lotId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lot on a InboundShipmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLotFromInboundShipmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.UnassignLotFromInboundShipmentLine(inboundShipmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a DestinationLocation on a InboundShipmentLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDestinationLocationToInboundShipmentLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	destinationLocationId,_ := strconv.ParseUint( vars["destinationLocationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.AssignDestinationLocationToInboundShipmentLine(inboundShipmentLineId, destinationLocationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a DestinationLocation on a InboundShipmentLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDestinationLocationFromInboundShipmentLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.UnassignDestinationLocationFromInboundShipmentLine(inboundShipmentLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more serialNumbersIds as a SerialNumbers to a InboundShipmentLine
	//----------------------------------------------------------------------------
func AddSerialNumbersToInboundShipmentLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.AddSerialNumbersToInboundShipmentLine(inboundShipmentLineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more serialNumbersIds as a SerialNumbers from a InboundShipmentLine
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSerialNumbersFromInboundShipmentLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	inboundShipmentLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	serialNumbersIds,_ := vars["serialNumbersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InboundShipmentLine DAO
	//----------------------------------------------------------------------------
	requestResult := InboundShipmentLineDAO.RemoveSerialNumbersFromInboundShipmentLine(inboundShipmentLineId, serialNumbersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
