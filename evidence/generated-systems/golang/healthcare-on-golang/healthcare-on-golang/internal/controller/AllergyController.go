package controller

import (
    AllergyDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AllergyDAO for database creation
//----------------------------------------------------------------------------
func CreateAllergy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Allergy model
	//----------------------------------------------------------------------------
	data := model.Allergy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Allergy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Allergy data access object to create
	//----------------------------------------------------------------------------
	requestResult := AllergyDAO.CreateAllergy( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AllergyDAO to find the relevant Allergy
//----------------------------------------------------------------------------
func GetAllergy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Allergy data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AllergyDAO.GetAllergy(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AllergyDAO for database read of all Allergys
//----------------------------------------------------------------------------
func GetAllAllergy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Allergy data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AllergyDAO.GetAllAllergy()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AllergyDAO for database save
//----------------------------------------------------------------------------
func UpdateAllergy(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Allergy model
	//----------------------------------------------------------------------------
	var data = model.Allergy{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Allergy model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Allergy data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AllergyDAO.UpdateAllergy(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AllergyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAllergy(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Allergy data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AllergyDAO.DeleteAllergy(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Patient on a Allergy
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPatientToAllergy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	allergyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	patientId,_ := strconv.ParseUint( vars["patientId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Allergy DAO
	//----------------------------------------------------------------------------
	requestResult := AllergyDAO.AssignPatientToAllergy(allergyId, patientId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Patient on a Allergy
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPatientFromAllergy( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	allergyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Allergy DAO
	//----------------------------------------------------------------------------
	requestResult := AllergyDAO.UnassignPatientFromAllergy(allergyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


