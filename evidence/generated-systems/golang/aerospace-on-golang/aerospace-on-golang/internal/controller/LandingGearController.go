package controller

import (
    LandingGearDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to LandingGearDAO for database creation
//----------------------------------------------------------------------------
func CreateLandingGear(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LandingGear model
	//----------------------------------------------------------------------------
	data := model.LandingGear{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LandingGear model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LandingGear data access object to create
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.CreateLandingGear( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to LandingGearDAO to find the relevant LandingGear
//----------------------------------------------------------------------------
func GetLandingGear(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LandingGear data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.GetLandingGear(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to LandingGearDAO for database read of all LandingGears
//----------------------------------------------------------------------------
func GetAllLandingGear(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the LandingGear data access object to get all
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.GetAllLandingGear()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to LandingGearDAO for database save
//----------------------------------------------------------------------------
func UpdateLandingGear(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty LandingGear model
	//----------------------------------------------------------------------------
	var data = model.LandingGear{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a LandingGear model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the LandingGear data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.UpdateLandingGear(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to LandingGearDAO for database deletion
//----------------------------------------------------------------------------
func DeleteLandingGear(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the LandingGear data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := LandingGearDAO.DeleteLandingGear(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Supplier on a LandingGear
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSupplierToLandingGear(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	landingGearId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	supplierId,_ := strconv.ParseUint( vars["supplierId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LandingGear DAO
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.AssignSupplierToLandingGear(landingGearId, supplierId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Supplier on a LandingGear
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSupplierFromLandingGear( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	landingGearId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the LandingGear DAO
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.UnassignSupplierFromLandingGear(landingGearId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more variantsIds as a Variants to a LandingGear
	//----------------------------------------------------------------------------
func AddVariantsToLandingGear(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	landingGearId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LandingGear DAO
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.AddVariantsToLandingGear(landingGearId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variantsIds as a Variants from a LandingGear
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariantsFromLandingGear(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	landingGearId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the LandingGear DAO
	//----------------------------------------------------------------------------
	requestResult := LandingGearDAO.RemoveVariantsFromLandingGear(landingGearId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
