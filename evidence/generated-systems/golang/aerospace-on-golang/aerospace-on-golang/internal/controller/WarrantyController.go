package controller

import (
    WarrantyDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to WarrantyDAO for database creation
//----------------------------------------------------------------------------
func CreateWarranty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Warranty model
	//----------------------------------------------------------------------------
	data := model.Warranty{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Warranty model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Warranty data access object to create
	//----------------------------------------------------------------------------
	requestResult := WarrantyDAO.CreateWarranty( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to WarrantyDAO to find the relevant Warranty
//----------------------------------------------------------------------------
func GetWarranty(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Warranty data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WarrantyDAO.GetWarranty(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to WarrantyDAO for database read of all Warrantys
//----------------------------------------------------------------------------
func GetAllWarranty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Warranty data access object to get all
	//----------------------------------------------------------------------------
	requestResult := WarrantyDAO.GetAllWarranty()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to WarrantyDAO for database save
//----------------------------------------------------------------------------
func UpdateWarranty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Warranty model
	//----------------------------------------------------------------------------
	var data = model.Warranty{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Warranty model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Warranty data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WarrantyDAO.UpdateWarranty(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to WarrantyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteWarranty(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Warranty data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := WarrantyDAO.DeleteWarranty(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Aircraft on a Warranty
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAircraftToWarranty(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	warrantyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftId,_ := strconv.ParseUint( vars["aircraftId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Warranty DAO
	//----------------------------------------------------------------------------
	requestResult := WarrantyDAO.AssignAircraftToWarranty(warrantyId, aircraftId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Aircraft on a Warranty
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAircraftFromWarranty( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	warrantyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Warranty DAO
	//----------------------------------------------------------------------------
	requestResult := WarrantyDAO.UnassignAircraftFromWarranty(warrantyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


