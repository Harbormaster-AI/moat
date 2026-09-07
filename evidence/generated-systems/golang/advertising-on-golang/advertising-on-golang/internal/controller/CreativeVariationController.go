package controller

import (
    CreativeVariationDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CreativeVariationDAO for database creation
//----------------------------------------------------------------------------
func CreateCreativeVariation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeVariation model
	//----------------------------------------------------------------------------
	data := model.CreativeVariation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeVariation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeVariation data access object to create
	//----------------------------------------------------------------------------
	requestResult := CreativeVariationDAO.CreateCreativeVariation( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CreativeVariationDAO to find the relevant CreativeVariation
//----------------------------------------------------------------------------
func GetCreativeVariation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeVariation data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeVariationDAO.GetCreativeVariation(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CreativeVariationDAO for database read of all CreativeVariations
//----------------------------------------------------------------------------
func GetAllCreativeVariation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CreativeVariation data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CreativeVariationDAO.GetAllCreativeVariation()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CreativeVariationDAO for database save
//----------------------------------------------------------------------------
func UpdateCreativeVariation(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeVariation model
	//----------------------------------------------------------------------------
	var data = model.CreativeVariation{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeVariation model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeVariation data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeVariationDAO.UpdateCreativeVariation(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CreativeVariationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCreativeVariation(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeVariation data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CreativeVariationDAO.DeleteCreativeVariation(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a CreativeAsset on a CreativeVariation
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCreativeAssetToCreativeVariation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeVariationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativeAssetId,_ := strconv.ParseUint( vars["creativeAssetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeVariation DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeVariationDAO.AssignCreativeAssetToCreativeVariation(creativeVariationId, creativeAssetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CreativeAsset on a CreativeVariation
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCreativeAssetFromCreativeVariation( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeVariationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeVariation DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeVariationDAO.UnassignCreativeAssetFromCreativeVariation(creativeVariationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


