package controller

import (
    OfferDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OfferDAO for database creation
//----------------------------------------------------------------------------
func CreateOffer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Offer model
	//----------------------------------------------------------------------------
	data := model.Offer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Offer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Offer data access object to create
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.CreateOffer( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OfferDAO to find the relevant Offer
//----------------------------------------------------------------------------
func GetOffer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Offer data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.GetOffer(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OfferDAO for database read of all Offers
//----------------------------------------------------------------------------
func GetAllOffer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Offer data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.GetAllOffer()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OfferDAO for database save
//----------------------------------------------------------------------------
func UpdateOffer(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Offer model
	//----------------------------------------------------------------------------
	var data = model.Offer{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Offer model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Offer data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.UpdateOffer(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OfferDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOffer(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Offer data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OfferDAO.DeleteOffer(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Requisition on a Offer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRequisitionToOffer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	requisitionId,_ := strconv.ParseUint( vars["requisitionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.AssignRequisitionToOffer(offerId, requisitionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Requisition on a Offer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRequisitionFromOffer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.UnassignRequisitionFromOffer(offerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Candidate on a Offer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCandidateToOffer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	candidateId,_ := strconv.ParseUint( vars["candidateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.AssignCandidateToOffer(offerId, candidateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Candidate on a Offer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCandidateFromOffer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.UnassignCandidateFromOffer(offerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ApprovedBy on a Offer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignApprovedByToOffer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approvedById,_ := strconv.ParseUint( vars["approvedById"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.AssignApprovedByToOffer(offerId, approvedById)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ApprovedBy on a Offer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignApprovedByFromOffer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.UnassignApprovedByFromOffer(offerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Contract on a Offer
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContractToOffer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractId,_ := strconv.ParseUint( vars["contractId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.AssignContractToOffer(offerId, contractId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contract on a Offer
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContractFromOffer( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	offerId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Offer DAO
	//----------------------------------------------------------------------------
	requestResult := OfferDAO.UnassignContractFromOffer(offerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


