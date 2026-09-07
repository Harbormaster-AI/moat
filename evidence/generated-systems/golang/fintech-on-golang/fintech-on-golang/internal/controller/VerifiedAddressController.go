package controller

import (
    VerifiedAddressDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to VerifiedAddressDAO for database creation
//----------------------------------------------------------------------------
func CreateVerifiedAddress(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty VerifiedAddress model
	//----------------------------------------------------------------------------
	data := model.VerifiedAddress{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a VerifiedAddress model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the VerifiedAddress data access object to create
	//----------------------------------------------------------------------------
	requestResult := VerifiedAddressDAO.CreateVerifiedAddress( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to VerifiedAddressDAO to find the relevant VerifiedAddress
//----------------------------------------------------------------------------
func GetVerifiedAddress(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the VerifiedAddress data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := VerifiedAddressDAO.GetVerifiedAddress(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to VerifiedAddressDAO for database read of all VerifiedAddresss
//----------------------------------------------------------------------------
func GetAllVerifiedAddress(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the VerifiedAddress data access object to get all
	//----------------------------------------------------------------------------
	requestResult := VerifiedAddressDAO.GetAllVerifiedAddress()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to VerifiedAddressDAO for database save
//----------------------------------------------------------------------------
func UpdateVerifiedAddress(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty VerifiedAddress model
	//----------------------------------------------------------------------------
	var data = model.VerifiedAddress{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a VerifiedAddress model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the VerifiedAddress data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := VerifiedAddressDAO.UpdateVerifiedAddress(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to VerifiedAddressDAO for database deletion
//----------------------------------------------------------------------------
func DeleteVerifiedAddress(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the VerifiedAddress data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := VerifiedAddressDAO.DeleteVerifiedAddress(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a KycProfile on a VerifiedAddress
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignKycProfileToVerifiedAddress(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	verifiedAddressId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfileId,_ := strconv.ParseUint( vars["kycProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the VerifiedAddress DAO
	//----------------------------------------------------------------------------
	requestResult := VerifiedAddressDAO.AssignKycProfileToVerifiedAddress(verifiedAddressId, kycProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a KycProfile on a VerifiedAddress
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignKycProfileFromVerifiedAddress( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	verifiedAddressId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the VerifiedAddress DAO
	//----------------------------------------------------------------------------
	requestResult := VerifiedAddressDAO.UnassignKycProfileFromVerifiedAddress(verifiedAddressId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


