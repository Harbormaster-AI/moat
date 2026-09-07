package controller

import (
    TelemetrySchemaDAO "iot-on-golang/internal/dao"
    "iot-on-golang/internal/model"
    "iot-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TelemetrySchemaDAO for database creation
//----------------------------------------------------------------------------
func CreateTelemetrySchema(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TelemetrySchema model
	//----------------------------------------------------------------------------
	data := model.TelemetrySchema{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TelemetrySchema model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetrySchema data access object to create
	//----------------------------------------------------------------------------
	requestResult := TelemetrySchemaDAO.CreateTelemetrySchema( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TelemetrySchemaDAO to find the relevant TelemetrySchema
//----------------------------------------------------------------------------
func GetTelemetrySchema(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TelemetrySchema data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TelemetrySchemaDAO.GetTelemetrySchema(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TelemetrySchemaDAO for database read of all TelemetrySchemas
//----------------------------------------------------------------------------
func GetAllTelemetrySchema(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TelemetrySchema data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TelemetrySchemaDAO.GetAllTelemetrySchema()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TelemetrySchemaDAO for database save
//----------------------------------------------------------------------------
func UpdateTelemetrySchema(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TelemetrySchema model
	//----------------------------------------------------------------------------
	var data = model.TelemetrySchema{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TelemetrySchema model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TelemetrySchema data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TelemetrySchemaDAO.UpdateTelemetrySchema(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TelemetrySchemaDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTelemetrySchema(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TelemetrySchema data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TelemetrySchemaDAO.DeleteTelemetrySchema(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more streamsIds as a Streams to a TelemetrySchema
	//----------------------------------------------------------------------------
func AddStreamsToTelemetrySchema(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	telemetrySchemaId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	streamsIds,_ := vars["streamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TelemetrySchema DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetrySchemaDAO.AddStreamsToTelemetrySchema(telemetrySchemaId, streamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more streamsIds as a Streams from a TelemetrySchema
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveStreamsFromTelemetrySchema(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	telemetrySchemaId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	streamsIds,_ := vars["streamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TelemetrySchema DAO
	//----------------------------------------------------------------------------
	requestResult := TelemetrySchemaDAO.RemoveStreamsFromTelemetrySchema(telemetrySchemaId, streamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
