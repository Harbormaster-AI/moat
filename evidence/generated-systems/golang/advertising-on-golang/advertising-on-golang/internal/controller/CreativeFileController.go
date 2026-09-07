package controller

import (
    CreativeFileDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CreativeFileDAO for database creation
//----------------------------------------------------------------------------
func CreateCreativeFile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeFile model
	//----------------------------------------------------------------------------
	data := model.CreativeFile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeFile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeFile data access object to create
	//----------------------------------------------------------------------------
	requestResult := CreativeFileDAO.CreateCreativeFile( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CreativeFileDAO to find the relevant CreativeFile
//----------------------------------------------------------------------------
func GetCreativeFile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeFile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeFileDAO.GetCreativeFile(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CreativeFileDAO for database read of all CreativeFiles
//----------------------------------------------------------------------------
func GetAllCreativeFile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CreativeFile data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CreativeFileDAO.GetAllCreativeFile()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CreativeFileDAO for database save
//----------------------------------------------------------------------------
func UpdateCreativeFile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeFile model
	//----------------------------------------------------------------------------
	var data = model.CreativeFile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeFile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeFile data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeFileDAO.UpdateCreativeFile(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CreativeFileDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCreativeFile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeFile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CreativeFileDAO.DeleteCreativeFile(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a CreativeAsset on a CreativeFile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCreativeAssetToCreativeFile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeFileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativeAssetId,_ := strconv.ParseUint( vars["creativeAssetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeFile DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeFileDAO.AssignCreativeAssetToCreativeFile(creativeFileId, creativeAssetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CreativeAsset on a CreativeFile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCreativeAssetFromCreativeFile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeFileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeFile DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeFileDAO.UnassignCreativeAssetFromCreativeFile(creativeFileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


