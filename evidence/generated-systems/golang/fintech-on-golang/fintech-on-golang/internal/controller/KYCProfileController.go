package controller

import (
    KYCProfileDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to KYCProfileDAO for database creation
//----------------------------------------------------------------------------
func CreateKYCProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty KYCProfile model
	//----------------------------------------------------------------------------
	data := model.KYCProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a KYCProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile data access object to create
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.CreateKYCProfile( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to KYCProfileDAO to find the relevant KYCProfile
//----------------------------------------------------------------------------
func GetKYCProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the KYCProfile data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.GetKYCProfile(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to KYCProfileDAO for database read of all KYCProfiles
//----------------------------------------------------------------------------
func GetAllKYCProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile data access object to get all
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.GetAllKYCProfile()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to KYCProfileDAO for database save
//----------------------------------------------------------------------------
func UpdateKYCProfile(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty KYCProfile model
	//----------------------------------------------------------------------------
	var data = model.KYCProfile{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a KYCProfile model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.UpdateKYCProfile(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to KYCProfileDAO for database deletion
//----------------------------------------------------------------------------
func DeleteKYCProfile(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the KYCProfile data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := KYCProfileDAO.DeleteKYCProfile(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a KYCProfile
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToKYCProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.AssignCustomerToKYCProfile(kYCProfileId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a KYCProfile
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromKYCProfile( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.UnassignCustomerFromKYCProfile(kYCProfileId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more documentsIds as a Documents to a KYCProfile
	//----------------------------------------------------------------------------
func AddDocumentsToKYCProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	documentsIds,_ := vars["documentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.AddDocumentsToKYCProfile(kYCProfileId, documentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more documentsIds as a Documents from a KYCProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDocumentsFromKYCProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	documentsIds,_ := vars["documentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.RemoveDocumentsFromKYCProfile(kYCProfileId, documentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more screeningsIds as a Screenings to a KYCProfile
	//----------------------------------------------------------------------------
func AddScreeningsToKYCProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	screeningsIds,_ := vars["screeningsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.AddScreeningsToKYCProfile(kYCProfileId, screeningsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more screeningsIds as a Screenings from a KYCProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveScreeningsFromKYCProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	screeningsIds,_ := vars["screeningsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.RemoveScreeningsFromKYCProfile(kYCProfileId, screeningsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more addressesIds as a Addresses to a KYCProfile
	//----------------------------------------------------------------------------
func AddAddressesToKYCProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	addressesIds,_ := vars["addressesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.AddAddressesToKYCProfile(kYCProfileId, addressesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more addressesIds as a Addresses from a KYCProfile
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAddressesFromKYCProfile(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	kYCProfileId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	addressesIds,_ := vars["addressesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the KYCProfile DAO
	//----------------------------------------------------------------------------
	requestResult := KYCProfileDAO.RemoveAddressesFromKYCProfile(kYCProfileId, addressesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
