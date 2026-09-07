package controller

import (
    BrandSafetyPolicyDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BrandSafetyPolicyDAO for database creation
//----------------------------------------------------------------------------
func CreateBrandSafetyPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BrandSafetyPolicy model
	//----------------------------------------------------------------------------
	data := model.BrandSafetyPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BrandSafetyPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BrandSafetyPolicy data access object to create
	//----------------------------------------------------------------------------
	requestResult := BrandSafetyPolicyDAO.CreateBrandSafetyPolicy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BrandSafetyPolicyDAO to find the relevant BrandSafetyPolicy
//----------------------------------------------------------------------------
func GetBrandSafetyPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BrandSafetyPolicy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BrandSafetyPolicyDAO.GetBrandSafetyPolicy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BrandSafetyPolicyDAO for database read of all BrandSafetyPolicys
//----------------------------------------------------------------------------
func GetAllBrandSafetyPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BrandSafetyPolicy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BrandSafetyPolicyDAO.GetAllBrandSafetyPolicy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BrandSafetyPolicyDAO for database save
//----------------------------------------------------------------------------
func UpdateBrandSafetyPolicy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BrandSafetyPolicy model
	//----------------------------------------------------------------------------
	var data = model.BrandSafetyPolicy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BrandSafetyPolicy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BrandSafetyPolicy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BrandSafetyPolicyDAO.UpdateBrandSafetyPolicy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BrandSafetyPolicyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBrandSafetyPolicy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BrandSafetyPolicy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BrandSafetyPolicyDAO.DeleteBrandSafetyPolicy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more targetingProfilesIds as a TargetingProfiles to a BrandSafetyPolicy
	//----------------------------------------------------------------------------
func AddTargetingProfilesToBrandSafetyPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	brandSafetyPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	targetingProfilesIds,_ := vars["targetingProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BrandSafetyPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := BrandSafetyPolicyDAO.AddTargetingProfilesToBrandSafetyPolicy(brandSafetyPolicyId, targetingProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more targetingProfilesIds as a TargetingProfiles from a BrandSafetyPolicy
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTargetingProfilesFromBrandSafetyPolicy(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	brandSafetyPolicyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	targetingProfilesIds,_ := vars["targetingProfilesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BrandSafetyPolicy DAO
	//----------------------------------------------------------------------------
	requestResult := BrandSafetyPolicyDAO.RemoveTargetingProfilesFromBrandSafetyPolicy(brandSafetyPolicyId, targetingProfilesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
