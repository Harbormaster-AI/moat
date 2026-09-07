package controller

import (
    Component_DAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to Component_DAO for database creation
//----------------------------------------------------------------------------
func CreateComponent_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Component_ model
	//----------------------------------------------------------------------------
	data := model.Component_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Component_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Component_ data access object to create
	//----------------------------------------------------------------------------
	requestResult := Component_DAO.CreateComponent_( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to Component_DAO to find the relevant Component_
//----------------------------------------------------------------------------
func GetComponent_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Component_ data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Component_DAO.GetComponent_(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to Component_DAO for database read of all Component_s
//----------------------------------------------------------------------------
func GetAllComponent_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Component_ data access object to get all
	//----------------------------------------------------------------------------
	requestResult := Component_DAO.GetAllComponent_()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to Component_DAO for database save
//----------------------------------------------------------------------------
func UpdateComponent_(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Component_ model
	//----------------------------------------------------------------------------
	var data = model.Component_{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Component_ model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Component_ data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := Component_DAO.UpdateComponent_(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to Component_DAO for database deletion
//----------------------------------------------------------------------------
func DeleteComponent_(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Component_ data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := Component_DAO.DeleteComponent_(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Supplier on a Component_
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSupplierToComponent_(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	component_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	supplierId,_ := strconv.ParseUint( vars["supplierId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Component_ DAO
	//----------------------------------------------------------------------------
	requestResult := Component_DAO.AssignSupplierToComponent_(component_Id, supplierId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Supplier on a Component_
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSupplierFromComponent_( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	component_Id,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Component_ DAO
	//----------------------------------------------------------------------------
	requestResult := Component_DAO.UnassignSupplierFromComponent_(component_Id)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


