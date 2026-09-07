package controller

import (
    ProductionCertificateDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProductionCertificateDAO for database creation
//----------------------------------------------------------------------------
func CreateProductionCertificate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionCertificate model
	//----------------------------------------------------------------------------
	data := model.ProductionCertificate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionCertificate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionCertificate data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProductionCertificateDAO.CreateProductionCertificate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProductionCertificateDAO to find the relevant ProductionCertificate
//----------------------------------------------------------------------------
func GetProductionCertificate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionCertificate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionCertificateDAO.GetProductionCertificate(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProductionCertificateDAO for database read of all ProductionCertificates
//----------------------------------------------------------------------------
func GetAllProductionCertificate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ProductionCertificate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProductionCertificateDAO.GetAllProductionCertificate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProductionCertificateDAO for database save
//----------------------------------------------------------------------------
func UpdateProductionCertificate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionCertificate model
	//----------------------------------------------------------------------------
	var data = model.ProductionCertificate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionCertificate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionCertificate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionCertificateDAO.UpdateProductionCertificate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProductionCertificateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProductionCertificate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionCertificate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProductionCertificateDAO.DeleteProductionCertificate(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Manufacturer on a ProductionCertificate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignManufacturerToProductionCertificate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionCertificateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	manufacturerId,_ := strconv.ParseUint( vars["manufacturerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionCertificate DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionCertificateDAO.AssignManufacturerToProductionCertificate(productionCertificateId, manufacturerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Manufacturer on a ProductionCertificate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignManufacturerFromProductionCertificate( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionCertificateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionCertificate DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionCertificateDAO.UnassignManufacturerFromProductionCertificate(productionCertificateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


