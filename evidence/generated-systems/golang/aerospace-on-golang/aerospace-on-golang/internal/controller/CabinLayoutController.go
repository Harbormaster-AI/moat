package controller

import (
    CabinLayoutDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CabinLayoutDAO for database creation
//----------------------------------------------------------------------------
func CreateCabinLayout(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CabinLayout model
	//----------------------------------------------------------------------------
	data := model.CabinLayout{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CabinLayout model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout data access object to create
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.CreateCabinLayout( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CabinLayoutDAO to find the relevant CabinLayout
//----------------------------------------------------------------------------
func GetCabinLayout(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CabinLayout data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.GetCabinLayout(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CabinLayoutDAO for database read of all CabinLayouts
//----------------------------------------------------------------------------
func GetAllCabinLayout(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.GetAllCabinLayout()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CabinLayoutDAO for database save
//----------------------------------------------------------------------------
func UpdateCabinLayout(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CabinLayout model
	//----------------------------------------------------------------------------
	var data = model.CabinLayout{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CabinLayout model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.UpdateCabinLayout(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CabinLayoutDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCabinLayout(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CabinLayout data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CabinLayoutDAO.DeleteCabinLayout(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Variant on a CabinLayout
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignVariantToCabinLayout(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cabinLayoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantId,_ := strconv.ParseUint( vars["variantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout DAO
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.AssignVariantToCabinLayout(cabinLayoutId, variantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Variant on a CabinLayout
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignVariantFromCabinLayout( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	cabinLayoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout DAO
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.UnassignVariantFromCabinLayout(cabinLayoutId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more aircraftIds as a Aircraft to a CabinLayout
	//----------------------------------------------------------------------------
func AddAircraftToCabinLayout(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cabinLayoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftIds,_ := vars["aircraftIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout DAO
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.AddAircraftToCabinLayout(cabinLayoutId, aircraftIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more aircraftIds as a Aircraft from a CabinLayout
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAircraftFromCabinLayout(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cabinLayoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftIds,_ := vars["aircraftIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout DAO
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.RemoveAircraftFromCabinLayout(cabinLayoutId, aircraftIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more optionsIds as a Options to a CabinLayout
	//----------------------------------------------------------------------------
func AddOptionsToCabinLayout(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cabinLayoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	optionsIds,_ := vars["optionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout DAO
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.AddOptionsToCabinLayout(cabinLayoutId, optionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more optionsIds as a Options from a CabinLayout
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOptionsFromCabinLayout(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	cabinLayoutId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	optionsIds,_ := vars["optionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CabinLayout DAO
	//----------------------------------------------------------------------------
	requestResult := CabinLayoutDAO.RemoveOptionsFromCabinLayout(cabinLayoutId, optionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
