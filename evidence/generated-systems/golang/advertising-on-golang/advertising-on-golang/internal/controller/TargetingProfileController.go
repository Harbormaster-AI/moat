package controller

import (
    TargetingProfileDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TargetingProfileDAO for database creation
//----------------------------------------------------------------------------
func CreateTargetingProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TargetingProfile model
	//----------------------------------------------------------------------------
	data := model.TargetingProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TargetingProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile data access object to create
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.CreateTargetingProfile( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TargetingProfileDAO to find the relevant TargetingProfile
//----------------------------------------------------------------------------
func GetTargetingProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TargetingProfile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.GetTargetingProfile(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TargetingProfileDAO for database read of all TargetingProfiles
//----------------------------------------------------------------------------
func GetAllTargetingProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.GetAllTargetingProfile()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TargetingProfileDAO for database save
//----------------------------------------------------------------------------
func UpdateTargetingProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TargetingProfile model
	//----------------------------------------------------------------------------
	var data = model.TargetingProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TargetingProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.UpdateTargetingProfile(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TargetingProfileDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTargetingProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TargetingProfile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TargetingProfileDAO.DeleteTargetingProfile(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a BrandSafetyPolicy on a TargetingProfile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBrandSafetyPolicyToTargetingProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	brandSafetyPolicyId,_ := strconv.ParseUint( vars["brandSafetyPolicyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.AssignBrandSafetyPolicyToTargetingProfile(targetingProfileId, brandSafetyPolicyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BrandSafetyPolicy on a TargetingProfile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBrandSafetyPolicyFromTargetingProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.UnassignBrandSafetyPolicyFromTargetingProfile(targetingProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more audienceSegmentsIds as a AudienceSegments to a TargetingProfile
	//----------------------------------------------------------------------------
func AddAudienceSegmentsToTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	audienceSegmentsIds,_ := vars["audienceSegmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.AddAudienceSegmentsToTargetingProfile(targetingProfileId, audienceSegmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more audienceSegmentsIds as a AudienceSegments from a TargetingProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAudienceSegmentsFromTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	audienceSegmentsIds,_ := vars["audienceSegmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.RemoveAudienceSegmentsFromTargetingProfile(targetingProfileId, audienceSegmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more geoRegionsIds as a GeoRegions to a TargetingProfile
	//----------------------------------------------------------------------------
func AddGeoRegionsToTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	geoRegionsIds,_ := vars["geoRegionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.AddGeoRegionsToTargetingProfile(targetingProfileId, geoRegionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more geoRegionsIds as a GeoRegions from a TargetingProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveGeoRegionsFromTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	geoRegionsIds,_ := vars["geoRegionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.RemoveGeoRegionsFromTargetingProfile(targetingProfileId, geoRegionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contentCategoriesIds as a ContentCategories to a TargetingProfile
	//----------------------------------------------------------------------------
func AddContentCategoriesToTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contentCategoriesIds,_ := vars["contentCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.AddContentCategoriesToTargetingProfile(targetingProfileId, contentCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contentCategoriesIds as a ContentCategories from a TargetingProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContentCategoriesFromTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contentCategoriesIds,_ := vars["contentCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.RemoveContentCategoriesFromTargetingProfile(targetingProfileId, contentCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more deviceCriteriaIds as a DeviceCriteria to a TargetingProfile
	//----------------------------------------------------------------------------
func AddDeviceCriteriaToTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceCriteriaIds,_ := vars["deviceCriteriaIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.AddDeviceCriteriaToTargetingProfile(targetingProfileId, deviceCriteriaIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more deviceCriteriaIds as a DeviceCriteria from a TargetingProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDeviceCriteriaFromTargetingProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	targetingProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	deviceCriteriaIds,_ := vars["deviceCriteriaIds"]

	//----------------------------------------------------------------------------
	// Delegate to the TargetingProfile DAO
	//----------------------------------------------------------------------------
	requestResult := TargetingProfileDAO.RemoveDeviceCriteriaFromTargetingProfile(targetingProfileId, deviceCriteriaIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
