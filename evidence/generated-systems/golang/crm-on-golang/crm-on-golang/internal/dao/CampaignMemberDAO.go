package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CampaignMemberDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCampaignMember - creates a new db entry
//----------------------------------------------------------------------------
func CreateCampaignMember(obj model.CampaignMember)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a CampaignMember with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CampaignMember", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCampaignMember", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCampaignMember - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCampaignMember(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CampaignMember

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CampaignMember with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CampaignMember using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CampaignMember using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCampaignMember", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCampaignMember - returns all
//----------------------------------------------------------------------------
func GetAllCampaignMember()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CampaignMember

	//----------------------------------------------------------------------------
	// Request the ORM to find all CampaignMember
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CampaignMember" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CampaignMember", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCampaignMember", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCampaignMember - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCampaignMember(obj model.CampaignMember)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a CampaignMember using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CampaignMember using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCampaignMember", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCampaignMember - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCampaignMember(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CampaignMember with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCampaignMember(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CampaignMember so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CampaignMember)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CampaignMember using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CampaignMember using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCampaignMember", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Campaign on a CampaignMember
//----------------------------------------------------------------------------
func AssignCampaignToCampaignMember( campaignMemberId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CampaignMember with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaignMember(campaignMemberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CampaignMember so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CampaignMember)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Campaign

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Campaign with a
		// matching campaignId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, campaignId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Campaign	to the CampaignMember
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the CampaignMember
			//----------------------------------------------------------------------------
			return UpdateCampaignMember(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a CampaignMember
//----------------------------------------------------------------------------
func UnassignCampaignFromCampaignMember(campaignMemberId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CampaignMember with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaignMember(campaignMemberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CampaignMember so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CampaignMember)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the CampaignMember
		//----------------------------------------------------------------------------
		return UpdateCampaignMember(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lead on a CampaignMember
//----------------------------------------------------------------------------
func AssignLeadToCampaignMember( campaignMemberId uint64, leadId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CampaignMember with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaignMember(campaignMemberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CampaignMember so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CampaignMember)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Lead

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Lead with a
		// matching leadId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, leadId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Lead	to the CampaignMember
			//----------------------------------------------------------------------------
			parentObj.Lead = &childObj

			//----------------------------------------------------------------------------
			// save the CampaignMember
			//----------------------------------------------------------------------------
			return UpdateCampaignMember(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lead", leadId )
			return utils.RequestResult{false, msg, "assignLead", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lead on a CampaignMember
//----------------------------------------------------------------------------
func UnassignLeadFromCampaignMember(campaignMemberId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CampaignMember with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaignMember(campaignMemberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CampaignMember so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CampaignMember)

		//----------------------------------------------------------------------------
		// assign an empty Lead to the Lead
		//----------------------------------------------------------------------------
		parentObj.Lead = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lead
		//----------------------------------------------------------------------------
		parentObj.LeadId = nil;

		//----------------------------------------------------------------------------
		// save the CampaignMember
		//----------------------------------------------------------------------------
		return UpdateCampaignMember(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Contact on a CampaignMember
//----------------------------------------------------------------------------
func AssignContactToCampaignMember( campaignMemberId uint64, contactId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CampaignMember with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaignMember(campaignMemberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CampaignMember so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CampaignMember)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Contact

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Contact with a
		// matching contactId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, contactId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Contact	to the CampaignMember
			//----------------------------------------------------------------------------
			parentObj.Contact = &childObj

			//----------------------------------------------------------------------------
			// save the CampaignMember
			//----------------------------------------------------------------------------
			return UpdateCampaignMember(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contact", contactId )
			return utils.RequestResult{false, msg, "assignContact", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Contact on a CampaignMember
//----------------------------------------------------------------------------
func UnassignContactFromCampaignMember(campaignMemberId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CampaignMember with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaignMember(campaignMemberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CampaignMember so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CampaignMember)

		//----------------------------------------------------------------------------
		// assign an empty Contact to the Contact
		//----------------------------------------------------------------------------
		parentObj.Contact = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Contact
		//----------------------------------------------------------------------------
		parentObj.ContactId = nil;

		//----------------------------------------------------------------------------
		// save the CampaignMember
		//----------------------------------------------------------------------------
		return UpdateCampaignMember(parentObj)

	} else {
		return parentRequestResult
	}

}


