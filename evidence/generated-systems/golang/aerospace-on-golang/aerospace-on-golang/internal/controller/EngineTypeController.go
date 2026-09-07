package controller

import (
    EngineTypeDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EngineTypeDAO for database creation
//----------------------------------------------------------------------------
func CreateEngineType(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EngineType model
	//----------------------------------------------------------------------------
	data := model.EngineType{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EngineType model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EngineType data access object to create
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.CreateEngineType( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EngineTypeDAO to find the relevant EngineType
//----------------------------------------------------------------------------
func GetEngineType(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EngineType data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.GetEngineType(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EngineTypeDAO for database read of all EngineTypes
//----------------------------------------------------------------------------
func GetAllEngineType(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the EngineType data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.GetAllEngineType()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EngineTypeDAO for database save
//----------------------------------------------------------------------------
func UpdateEngineType(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty EngineType model
	//----------------------------------------------------------------------------
	var data = model.EngineType{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a EngineType model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the EngineType data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.UpdateEngineType(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EngineTypeDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEngineType(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the EngineType data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EngineTypeDAO.DeleteEngineType(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Supplier on a EngineType
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSupplierToEngineType(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	engineTypeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	supplierId,_ := strconv.ParseUint( vars["supplierId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EngineType DAO
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.AssignSupplierToEngineType(engineTypeId, supplierId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Supplier on a EngineType
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSupplierFromEngineType( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	engineTypeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the EngineType DAO
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.UnassignSupplierFromEngineType(engineTypeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more compatibleModelsIds as a CompatibleModels to a EngineType
	//----------------------------------------------------------------------------
func AddCompatibleModelsToEngineType(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	engineTypeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compatibleModelsIds,_ := vars["compatibleModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the EngineType DAO
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.AddCompatibleModelsToEngineType(engineTypeId, compatibleModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more compatibleModelsIds as a CompatibleModels from a EngineType
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCompatibleModelsFromEngineType(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	engineTypeId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	compatibleModelsIds,_ := vars["compatibleModelsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the EngineType DAO
	//----------------------------------------------------------------------------
	requestResult := EngineTypeDAO.RemoveCompatibleModelsFromEngineType(engineTypeId, compatibleModelsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
