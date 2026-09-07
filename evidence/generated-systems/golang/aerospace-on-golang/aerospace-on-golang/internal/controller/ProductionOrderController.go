package controller

import (
    ProductionOrderDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProductionOrderDAO for database creation
//----------------------------------------------------------------------------
func CreateProductionOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionOrder model
	//----------------------------------------------------------------------------
	data := model.ProductionOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.CreateProductionOrder( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProductionOrderDAO to find the relevant ProductionOrder
//----------------------------------------------------------------------------
func GetProductionOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionOrder data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.GetProductionOrder(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProductionOrderDAO for database read of all ProductionOrders
//----------------------------------------------------------------------------
func GetAllProductionOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.GetAllProductionOrder()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProductionOrderDAO for database save
//----------------------------------------------------------------------------
func UpdateProductionOrder(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionOrder model
	//----------------------------------------------------------------------------
	var data = model.ProductionOrder{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionOrder model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.UpdateProductionOrder(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProductionOrderDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProductionOrder(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionOrder data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProductionOrderDAO.DeleteProductionOrder(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Variant on a ProductionOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignVariantToProductionOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variantId,_ := strconv.ParseUint( vars["variantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.AssignVariantToProductionOrder(productionOrderId, variantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Variant on a ProductionOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignVariantFromProductionOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.UnassignVariantFromProductionOrder(productionOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Plant on a ProductionOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToProductionOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.AssignPlantToProductionOrder(productionOrderId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a ProductionOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromProductionOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.UnassignPlantFromProductionOrder(productionOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a AircraftOrder on a ProductionOrder
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAircraftOrderToProductionOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	aircraftOrderId,_ := strconv.ParseUint( vars["aircraftOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.AssignAircraftOrderToProductionOrder(productionOrderId, aircraftOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AircraftOrder on a ProductionOrder
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAircraftOrderFromProductionOrder( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionOrderId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionOrder DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionOrderDAO.UnassignAircraftOrderFromProductionOrder(productionOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


