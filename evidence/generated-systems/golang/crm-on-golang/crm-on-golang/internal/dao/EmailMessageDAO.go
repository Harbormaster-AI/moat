package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EmailMessageDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEmailMessage - creates a new db entry
//----------------------------------------------------------------------------
func CreateEmailMessage(obj model.EmailMessage)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a EmailMessage with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a EmailMessage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEmailMessage", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEmailMessage - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEmailMessage(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.EmailMessage

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a EmailMessage with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a EmailMessage using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a EmailMessage using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEmailMessage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEmailMessage - returns all
//----------------------------------------------------------------------------
func GetAllEmailMessage()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.EmailMessage

	//----------------------------------------------------------------------------
	// Request the ORM to find all EmailMessage
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all EmailMessage" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all EmailMessage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEmailMessage", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEmailMessage - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEmailMessage(obj model.EmailMessage)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a EmailMessage using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a EmailMessage using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEmailMessage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEmailMessage - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEmailMessage(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEmailMessage(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a EmailMessage using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a EmailMessage using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEmailMessage", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a EmailMessage
//----------------------------------------------------------------------------
func AssignOrganizationToEmailMessage( emailMessageId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a EmailMessage
//----------------------------------------------------------------------------
func UnassignOrganizationFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a EmailMessage
//----------------------------------------------------------------------------
func AssignOwnerToEmailMessage( emailMessageId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.User

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a User with a
		// matching ownerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, ownerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Owner	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a EmailMessage
//----------------------------------------------------------------------------
func UnassignOwnerFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a EmailMessage
//----------------------------------------------------------------------------
func AssignAccountToEmailMessage( emailMessageId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching accountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, accountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Account	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a EmailMessage
//----------------------------------------------------------------------------
func UnassignAccountFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Contact on a EmailMessage
//----------------------------------------------------------------------------
func AssignContactToEmailMessage( emailMessageId uint64, contactId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

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
			// assign the Contact	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Contact = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contact", contactId )
			return utils.RequestResult{false, msg, "assignContact", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Contact on a EmailMessage
//----------------------------------------------------------------------------
func UnassignContactFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty Contact to the Contact
		//----------------------------------------------------------------------------
		parentObj.Contact = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Contact
		//----------------------------------------------------------------------------
		parentObj.ContactId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lead on a EmailMessage
//----------------------------------------------------------------------------
func AssignLeadToEmailMessage( emailMessageId uint64, leadId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

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
			// assign the Lead	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Lead = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lead", leadId )
			return utils.RequestResult{false, msg, "assignLead", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lead on a EmailMessage
//----------------------------------------------------------------------------
func UnassignLeadFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty Lead to the Lead
		//----------------------------------------------------------------------------
		parentObj.Lead = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lead
		//----------------------------------------------------------------------------
		parentObj.LeadId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Case on a EmailMessage
//----------------------------------------------------------------------------
func AssignCaseToEmailMessage( emailMessageId uint64, caseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Case_

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Case_ with a
		// matching caseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, caseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Case	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Case = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Case", caseId )
			return utils.RequestResult{false, msg, "assignCase", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Case on a EmailMessage
//----------------------------------------------------------------------------
func UnassignCaseFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty Case_ to the Case
		//----------------------------------------------------------------------------
		parentObj.Case = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Case
		//----------------------------------------------------------------------------
		parentObj.CaseId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Opportunity on a EmailMessage
//----------------------------------------------------------------------------
func AssignOpportunityToEmailMessage( emailMessageId uint64, opportunityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Opportunity

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Opportunity with a
		// matching opportunityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, opportunityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Opportunity	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Opportunity = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunity", opportunityId )
			return utils.RequestResult{false, msg, "assignOpportunity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Opportunity on a EmailMessage
//----------------------------------------------------------------------------
func UnassignOpportunityFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty Opportunity to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.Opportunity = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.OpportunityId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Campaign on a EmailMessage
//----------------------------------------------------------------------------
func AssignCampaignToEmailMessage( emailMessageId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

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
			// assign the Campaign	to the EmailMessage
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the EmailMessage
			//----------------------------------------------------------------------------
			return UpdateEmailMessage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a EmailMessage
//----------------------------------------------------------------------------
func UnassignCampaignFromEmailMessage(emailMessageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmailMessage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmailMessage(emailMessageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmailMessage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmailMessage)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the EmailMessage
		//----------------------------------------------------------------------------
		return UpdateEmailMessage(parentObj)

	} else {
		return parentRequestResult
	}

}


