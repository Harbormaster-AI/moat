package controller

import (
    CampaignMemberDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CampaignMemberDAO for database creation
//----------------------------------------------------------------------------
func CreateCampaignMember(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CampaignMember model
	//----------------------------------------------------------------------------
	data := model.CampaignMember{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CampaignMember model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember data access object to create
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.CreateCampaignMember( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CampaignMemberDAO to find the relevant CampaignMember
//----------------------------------------------------------------------------
func GetCampaignMember(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CampaignMember data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.GetCampaignMember(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CampaignMemberDAO for database read of all CampaignMembers
//----------------------------------------------------------------------------
func GetAllCampaignMember(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.GetAllCampaignMember()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CampaignMemberDAO for database save
//----------------------------------------------------------------------------
func UpdateCampaignMember(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CampaignMember model
	//----------------------------------------------------------------------------
	var data = model.CampaignMember{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CampaignMember model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.UpdateCampaignMember(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CampaignMemberDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCampaignMember(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CampaignMember data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CampaignMemberDAO.DeleteCampaignMember(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a CampaignMember
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToCampaignMember(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignMemberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.AssignCampaignToCampaignMember(campaignMemberId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a CampaignMember
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromCampaignMember( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignMemberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.UnassignCampaignFromCampaignMember(campaignMemberId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Lead on a CampaignMember
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLeadToCampaignMember(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignMemberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	leadId,_ := strconv.ParseUint( vars["leadId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.AssignLeadToCampaignMember(campaignMemberId, leadId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Lead on a CampaignMember
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLeadFromCampaignMember( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignMemberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.UnassignLeadFromCampaignMember(campaignMemberId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Contact on a CampaignMember
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignContactToCampaignMember(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignMemberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contactId,_ := strconv.ParseUint( vars["contactId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.AssignContactToCampaignMember(campaignMemberId, contactId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Contact on a CampaignMember
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignContactFromCampaignMember( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	campaignMemberId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CampaignMember DAO
	//----------------------------------------------------------------------------
	requestResult := CampaignMemberDAO.UnassignContactFromCampaignMember(campaignMemberId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


