package controller

import (
    KYCDocumentDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to KYCDocumentDAO for database creation
//----------------------------------------------------------------------------
func CreateKYCDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty KYCDocument model
	//----------------------------------------------------------------------------
	data := model.KYCDocument{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a KYCDocument model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KYCDocument data access object to create
	//----------------------------------------------------------------------------
	requestResult := KYCDocumentDAO.CreateKYCDocument( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to KYCDocumentDAO to find the relevant KYCDocument
//----------------------------------------------------------------------------
func GetKYCDocument(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the KYCDocument data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := KYCDocumentDAO.GetKYCDocument(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to KYCDocumentDAO for database read of all KYCDocuments
//----------------------------------------------------------------------------
func GetAllKYCDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the KYCDocument data access object to get all
	//----------------------------------------------------------------------------
	requestResult := KYCDocumentDAO.GetAllKYCDocument()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to KYCDocumentDAO for database save
//----------------------------------------------------------------------------
func UpdateKYCDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty KYCDocument model
	//----------------------------------------------------------------------------
	var data = model.KYCDocument{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a KYCDocument model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KYCDocument data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := KYCDocumentDAO.UpdateKYCDocument(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to KYCDocumentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteKYCDocument(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the KYCDocument data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := KYCDocumentDAO.DeleteKYCDocument(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a KycProfile on a KYCDocument
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignKycProfileToKYCDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	kYCDocumentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	kycProfileId,_ := strconv.ParseUint( vars["kycProfileId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the KYCDocument DAO
	//----------------------------------------------------------------------------
	requestResult := KYCDocumentDAO.AssignKycProfileToKYCDocument(kYCDocumentId, kycProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a KycProfile on a KYCDocument
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignKycProfileFromKYCDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	kYCDocumentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the KYCDocument DAO
	//----------------------------------------------------------------------------
	requestResult := KYCDocumentDAO.UnassignKycProfileFromKYCDocument(kYCDocumentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


