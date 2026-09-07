package controller

import (
    ProductOfferingDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProductOfferingDAO for database creation
//----------------------------------------------------------------------------
func CreateProductOffering(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductOffering model
	//----------------------------------------------------------------------------
	data := model.ProductOffering{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductOffering model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductOffering data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.CreateProductOffering( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProductOfferingDAO to find the relevant ProductOffering
//----------------------------------------------------------------------------
func GetProductOffering(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductOffering data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.GetProductOffering(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProductOfferingDAO for database read of all ProductOfferings
//----------------------------------------------------------------------------
func GetAllProductOffering(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ProductOffering data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.GetAllProductOffering()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProductOfferingDAO for database save
//----------------------------------------------------------------------------
func UpdateProductOffering(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductOffering model
	//----------------------------------------------------------------------------
	var data = model.ProductOffering{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductOffering model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductOffering data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.UpdateProductOffering(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProductOfferingDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProductOffering(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductOffering data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProductOfferingDAO.DeleteProductOffering(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Institution on a ProductOffering
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInstitutionToProductOffering(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productOfferingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	institutionId,_ := strconv.ParseUint( vars["institutionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductOffering DAO
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.AssignInstitutionToProductOffering(productOfferingId, institutionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Institution on a ProductOffering
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInstitutionFromProductOffering( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productOfferingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductOffering DAO
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.UnassignInstitutionFromProductOffering(productOfferingId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more pricingPlansIds as a PricingPlans to a ProductOffering
	//----------------------------------------------------------------------------
func AddPricingPlansToProductOffering(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productOfferingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pricingPlansIds,_ := vars["pricingPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ProductOffering DAO
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.AddPricingPlansToProductOffering(productOfferingId, pricingPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more pricingPlansIds as a PricingPlans from a ProductOffering
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePricingPlansFromProductOffering(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productOfferingId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pricingPlansIds,_ := vars["pricingPlansIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ProductOffering DAO
	//----------------------------------------------------------------------------
	requestResult := ProductOfferingDAO.RemovePricingPlansFromProductOffering(productOfferingId, pricingPlansIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
