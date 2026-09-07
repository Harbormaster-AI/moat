package controller

import (
    GeoRegionDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to GeoRegionDAO for database creation
//----------------------------------------------------------------------------
func CreateGeoRegion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GeoRegion model
	//----------------------------------------------------------------------------
	data := model.GeoRegion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GeoRegion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GeoRegion data access object to create
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.CreateGeoRegion( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to GeoRegionDAO to find the relevant GeoRegion
//----------------------------------------------------------------------------
func GetGeoRegion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GeoRegion data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.GetGeoRegion(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to GeoRegionDAO for database read of all GeoRegions
//----------------------------------------------------------------------------
func GetAllGeoRegion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the GeoRegion data access object to get all
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.GetAllGeoRegion()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to GeoRegionDAO for database save
//----------------------------------------------------------------------------
func UpdateGeoRegion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty GeoRegion model
	//----------------------------------------------------------------------------
	var data = model.GeoRegion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a GeoRegion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the GeoRegion data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.UpdateGeoRegion(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to GeoRegionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteGeoRegion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the GeoRegion data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := GeoRegionDAO.DeleteGeoRegion(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Parent on a GeoRegion
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignParentToGeoRegion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	geoRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	parentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GeoRegion DAO
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.AssignParentToGeoRegion(geoRegionId, parentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Parent on a GeoRegion
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignParentFromGeoRegion( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	geoRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the GeoRegion DAO
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.UnassignParentFromGeoRegion(geoRegionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more childrenIds as a Children to a GeoRegion
	//----------------------------------------------------------------------------
func AddChildrenToGeoRegion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	geoRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childrenIds,_ := vars["childrenIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GeoRegion DAO
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.AddChildrenToGeoRegion(geoRegionId, childrenIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more childrenIds as a Children from a GeoRegion
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveChildrenFromGeoRegion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	geoRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	childrenIds,_ := vars["childrenIds"]

	//----------------------------------------------------------------------------
	// Delegate to the GeoRegion DAO
	//----------------------------------------------------------------------------
	requestResult := GeoRegionDAO.RemoveChildrenFromGeoRegion(geoRegionId, childrenIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
