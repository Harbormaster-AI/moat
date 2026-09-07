package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LeadDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLead - creates a new db entry
//----------------------------------------------------------------------------
func CreateLead(obj model.Lead)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Lead with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Lead", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLead", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLead - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLead(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Lead

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Lead with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Lead using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Lead using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLead", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLead - returns all
//----------------------------------------------------------------------------
func GetAllLead()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Lead

	//----------------------------------------------------------------------------
	// Request the ORM to find all Lead
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Lead" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Lead", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLead", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLead - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLead(obj model.Lead)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Lead using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Lead using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLead", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLead - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLead(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLead(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Lead using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Lead using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLead", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Lead
//----------------------------------------------------------------------------
func AssignOrganizationToLead( leadId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

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
			// assign the Organization	to the Lead
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Lead
			//----------------------------------------------------------------------------
			return UpdateLead(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Lead
//----------------------------------------------------------------------------
func UnassignOrganizationFromLead(leadId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Lead
		//----------------------------------------------------------------------------
		return UpdateLead(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Lead
//----------------------------------------------------------------------------
func AssignOwnerToLead( leadId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

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
			// assign the Owner	to the Lead
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Lead
			//----------------------------------------------------------------------------
			return UpdateLead(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Lead
//----------------------------------------------------------------------------
func UnassignOwnerFromLead(leadId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Lead
		//----------------------------------------------------------------------------
		return UpdateLead(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ConvertedAccount on a Lead
//----------------------------------------------------------------------------
func AssignConvertedAccountToLead( leadId uint64, convertedAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Account

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Account with a
		// matching convertedAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, convertedAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ConvertedAccount	to the Lead
			//----------------------------------------------------------------------------
			parentObj.ConvertedAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Lead
			//----------------------------------------------------------------------------
			return UpdateLead(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConvertedAccount", convertedAccountId )
			return utils.RequestResult{false, msg, "assignConvertedAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ConvertedAccount on a Lead
//----------------------------------------------------------------------------
func UnassignConvertedAccountFromLead(leadId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// assign an empty Account to the ConvertedAccount
		//----------------------------------------------------------------------------
		parentObj.ConvertedAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ConvertedAccount
		//----------------------------------------------------------------------------
		parentObj.ConvertedAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Lead
		//----------------------------------------------------------------------------
		return UpdateLead(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ConvertedContact on a Lead
//----------------------------------------------------------------------------
func AssignConvertedContactToLead( leadId uint64, convertedContactId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Contact

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Contact with a
		// matching convertedContactId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, convertedContactId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ConvertedContact	to the Lead
			//----------------------------------------------------------------------------
			parentObj.ConvertedContact = &childObj

			//----------------------------------------------------------------------------
			// save the Lead
			//----------------------------------------------------------------------------
			return UpdateLead(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConvertedContact", convertedContactId )
			return utils.RequestResult{false, msg, "assignConvertedContact", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ConvertedContact on a Lead
//----------------------------------------------------------------------------
func UnassignConvertedContactFromLead(leadId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// assign an empty Contact to the ConvertedContact
		//----------------------------------------------------------------------------
		parentObj.ConvertedContact = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ConvertedContact
		//----------------------------------------------------------------------------
		parentObj.ConvertedContactId = nil;

		//----------------------------------------------------------------------------
		// save the Lead
		//----------------------------------------------------------------------------
		return UpdateLead(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ConvertedOpportunity on a Lead
//----------------------------------------------------------------------------
func AssignConvertedOpportunityToLead( leadId uint64, convertedOpportunityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Opportunity

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Opportunity with a
		// matching convertedOpportunityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, convertedOpportunityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ConvertedOpportunity	to the Lead
			//----------------------------------------------------------------------------
			parentObj.ConvertedOpportunity = &childObj

			//----------------------------------------------------------------------------
			// save the Lead
			//----------------------------------------------------------------------------
			return UpdateLead(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConvertedOpportunity", convertedOpportunityId )
			return utils.RequestResult{false, msg, "assignConvertedOpportunity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ConvertedOpportunity on a Lead
//----------------------------------------------------------------------------
func UnassignConvertedOpportunityFromLead(leadId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		//----------------------------------------------------------------------------
		// assign an empty Opportunity to the ConvertedOpportunity
		//----------------------------------------------------------------------------
		parentObj.ConvertedOpportunity = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ConvertedOpportunity
		//----------------------------------------------------------------------------
		parentObj.ConvertedOpportunityId = nil;

		//----------------------------------------------------------------------------
		// save the Lead
		//----------------------------------------------------------------------------
		return UpdateLead(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more activitiesIds as a Activities to a Lead
//----------------------------------------------------------------------------
func AddActivitiesToLead ( leadId uint64, activitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Activities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "unassignActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more activitiesIds as a Activities from a Lead
//----------------------------------------------------------------------------
func RemoveActivitiesFromLead( leadId uint64, activitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( activitiesIds, ",")

		for _, activitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Activity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Activity
			// with a matching activitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , activitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ActivityObj from the Activities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Activities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Activities", activitiesId )
				return utils.RequestResult{false, msg, "removeActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a Lead
//----------------------------------------------------------------------------
func AddCampaignsToLead ( leadId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( campaignsIds, ",")

		for _, campaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Campaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Campaign
			// with a matching campaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , campaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Campaigns using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Campaigns").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaigns", campaignsId )
				return utils.RequestResult{false, msg, "unassignCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a Lead
//----------------------------------------------------------------------------
func RemoveCampaignsFromLead( leadId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( campaignsIds, ",")

		for _, campaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Campaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Campaign
			// with a matching campaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , campaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CampaignObj from the Campaigns array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Campaigns").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaigns", campaignsId )
				return utils.RequestResult{false, msg, "removeCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more notesIds as a Notes to a Lead
//----------------------------------------------------------------------------
func AddNotesToLead ( leadId uint64, notesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( notesIds, ",")

		for _, notesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Note

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Note
			// with a matching notesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Notes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notes", notesId )
				return utils.RequestResult{false, msg, "unassignNotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more notesIds as a Notes from a Lead
//----------------------------------------------------------------------------
func RemoveNotesFromLead( leadId uint64, notesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( notesIds, ",")

		for _, notesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Note

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Note
			// with a matching notesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove NoteObj from the Notes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notes", notesId )
				return utils.RequestResult{false, msg, "removeNotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more emailMessagesIds as a EmailMessages to a Lead
//----------------------------------------------------------------------------
func AddEmailMessagesToLead ( leadId uint64, emailMessagesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( emailMessagesIds, ",")

		for _, emailMessagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailMessagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailMessagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EmailMessages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmailMessages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmailMessages", emailMessagesId )
				return utils.RequestResult{false, msg, "unassignEmailMessages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more emailMessagesIds as a EmailMessages from a Lead
//----------------------------------------------------------------------------
func RemoveEmailMessagesFromLead( leadId uint64, emailMessagesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Lead with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLead(leadId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lead so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lead)

		// slice the ids on comma with no spaces
		ids := strings.Split( emailMessagesIds, ",")

		for _, emailMessagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmailMessage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmailMessage
			// with a matching emailMessagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , emailMessagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmailMessageObj from the EmailMessages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EmailMessages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EmailMessages", emailMessagesId )
				return utils.RequestResult{false, msg, "removeEmailMessages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Lead from the gorm
		//----------------------------------------------------------------------------
		return GetLead(leadId)

	} else {
		return parentRequestResult
	}
}

