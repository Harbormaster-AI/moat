package controller

import (
    UoMConversionDAO "inventory-on-golang/internal/dao"
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to UoMConversionDAO for database creation
//----------------------------------------------------------------------------
func CreateUoMConversion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UoMConversion model
	//----------------------------------------------------------------------------
	data := model.UoMConversion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UoMConversion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UoMConversion data access object to create
	//----------------------------------------------------------------------------
	requestResult := UoMConversionDAO.CreateUoMConversion( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to UoMConversionDAO to find the relevant UoMConversion
//----------------------------------------------------------------------------
func GetUoMConversion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UoMConversion data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UoMConversionDAO.GetUoMConversion(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to UoMConversionDAO for database read of all UoMConversions
//----------------------------------------------------------------------------
func GetAllUoMConversion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the UoMConversion data access object to get all
	//----------------------------------------------------------------------------
	requestResult := UoMConversionDAO.GetAllUoMConversion()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to UoMConversionDAO for database save
//----------------------------------------------------------------------------
func UpdateUoMConversion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UoMConversion model
	//----------------------------------------------------------------------------
	var data = model.UoMConversion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UoMConversion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UoMConversion data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UoMConversionDAO.UpdateUoMConversion(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to UoMConversionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteUoMConversion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UoMConversion data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := UoMConversionDAO.DeleteUoMConversion(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Sku on a UoMConversion
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSkuToUoMConversion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	uoMConversionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	skuId,_ := strconv.ParseUint( vars["skuId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UoMConversion DAO
	//----------------------------------------------------------------------------
	requestResult := UoMConversionDAO.AssignSkuToUoMConversion(uoMConversionId, skuId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Sku on a UoMConversion
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSkuFromUoMConversion( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	uoMConversionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UoMConversion DAO
	//----------------------------------------------------------------------------
	requestResult := UoMConversionDAO.UnassignSkuFromUoMConversion(uoMConversionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


