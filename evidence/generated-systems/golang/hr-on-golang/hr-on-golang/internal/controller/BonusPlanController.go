package controller

import (
    BonusPlanDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BonusPlanDAO for database creation
//----------------------------------------------------------------------------
func CreateBonusPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BonusPlan model
	//----------------------------------------------------------------------------
	data := model.BonusPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BonusPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BonusPlan data access object to create
	//----------------------------------------------------------------------------
	requestResult := BonusPlanDAO.CreateBonusPlan( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BonusPlanDAO to find the relevant BonusPlan
//----------------------------------------------------------------------------
func GetBonusPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BonusPlan data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BonusPlanDAO.GetBonusPlan(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BonusPlanDAO for database read of all BonusPlans
//----------------------------------------------------------------------------
func GetAllBonusPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BonusPlan data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BonusPlanDAO.GetAllBonusPlan()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BonusPlanDAO for database save
//----------------------------------------------------------------------------
func UpdateBonusPlan(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BonusPlan model
	//----------------------------------------------------------------------------
	var data = model.BonusPlan{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BonusPlan model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BonusPlan data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BonusPlanDAO.UpdateBonusPlan(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BonusPlanDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBonusPlan(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BonusPlan data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BonusPlanDAO.DeleteBonusPlan(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more compensationPackagesIds as a CompensationPackages to a BonusPlan
	//----------------------------------------------------------------------------
func AddCompensationPackagesToBonusPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bonusPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compensationPackagesIds,_ := vars["compensationPackagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BonusPlan DAO
	//----------------------------------------------------------------------------
	requestResult := BonusPlanDAO.AddCompensationPackagesToBonusPlan(bonusPlanId, compensationPackagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more compensationPackagesIds as a CompensationPackages from a BonusPlan
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCompensationPackagesFromBonusPlan(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bonusPlanId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compensationPackagesIds,_ := vars["compensationPackagesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BonusPlan DAO
	//----------------------------------------------------------------------------
	requestResult := BonusPlanDAO.RemoveCompensationPackagesFromBonusPlan(bonusPlanId, compensationPackagesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
