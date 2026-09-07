package controller

import (
    AvionicsSuiteDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AvionicsSuiteDAO for database creation
//----------------------------------------------------------------------------
func CreateAvionicsSuite(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AvionicsSuite model
	//----------------------------------------------------------------------------
	data := model.AvionicsSuite{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AvionicsSuite model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite data access object to create
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.CreateAvionicsSuite( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AvionicsSuiteDAO to find the relevant AvionicsSuite
//----------------------------------------------------------------------------
func GetAvionicsSuite(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AvionicsSuite data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.GetAvionicsSuite(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AvionicsSuiteDAO for database read of all AvionicsSuites
//----------------------------------------------------------------------------
func GetAllAvionicsSuite(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.GetAllAvionicsSuite()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AvionicsSuiteDAO for database save
//----------------------------------------------------------------------------
func UpdateAvionicsSuite(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty AvionicsSuite model
	//----------------------------------------------------------------------------
	var data = model.AvionicsSuite{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a AvionicsSuite model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.UpdateAvionicsSuite(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AvionicsSuiteDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAvionicsSuite(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the AvionicsSuite data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AvionicsSuiteDAO.DeleteAvionicsSuite(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Supplier on a AvionicsSuite
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSupplierToAvionicsSuite(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	avionicsSuiteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	supplierId,_ := strconv.ParseUint( vars["supplierId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite DAO
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.AssignSupplierToAvionicsSuite(avionicsSuiteId, supplierId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Supplier on a AvionicsSuite
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSupplierFromAvionicsSuite( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	avionicsSuiteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite DAO
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.UnassignSupplierFromAvionicsSuite(avionicsSuiteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more variantsIds as a Variants to a AvionicsSuite
	//----------------------------------------------------------------------------
func AddVariantsToAvionicsSuite(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	avionicsSuiteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite DAO
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.AddVariantsToAvionicsSuite(avionicsSuiteId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variantsIds as a Variants from a AvionicsSuite
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariantsFromAvionicsSuite(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	avionicsSuiteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantsIds,_ := vars["variantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite DAO
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.RemoveVariantsFromAvionicsSuite(avionicsSuiteId, variantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more softwareLoadsIds as a SoftwareLoads to a AvionicsSuite
	//----------------------------------------------------------------------------
func AddSoftwareLoadsToAvionicsSuite(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	avionicsSuiteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	softwareLoadsIds,_ := vars["softwareLoadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite DAO
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.AddSoftwareLoadsToAvionicsSuite(avionicsSuiteId, softwareLoadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more softwareLoadsIds as a SoftwareLoads from a AvionicsSuite
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSoftwareLoadsFromAvionicsSuite(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	avionicsSuiteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	softwareLoadsIds,_ := vars["softwareLoadsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the AvionicsSuite DAO
	//----------------------------------------------------------------------------
	requestResult := AvionicsSuiteDAO.RemoveSoftwareLoadsFromAvionicsSuite(avionicsSuiteId, softwareLoadsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
