package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ContactDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateContact - creates a new db entry
//----------------------------------------------------------------------------
func CreateContact(obj model.Contact)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Contact with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Contact", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateContact", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetContact - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetContact(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Contact

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Contact with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Contact using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Contact using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetContact", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllContact - returns all
//----------------------------------------------------------------------------
func GetAllContact()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Contact

	//----------------------------------------------------------------------------
	// Request the ORM to find all Contact
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Contact" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Contact", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllContact", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateContact - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateContact(obj model.Contact)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Contact using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Contact using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateContact", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteContact - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteContact(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetContact(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Contact)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Contact using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Contact using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteContact", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Contact
//----------------------------------------------------------------------------
func AssignOrganizationToContact( contactId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
			// assign the Organization	to the Contact
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Contact
			//----------------------------------------------------------------------------
			return UpdateContact(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Contact
//----------------------------------------------------------------------------
func UnassignOrganizationFromContact(contactId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Contact
		//----------------------------------------------------------------------------
		return UpdateContact(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Contact
//----------------------------------------------------------------------------
func AssignAccountToContact( contactId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
			// assign the Account	to the Contact
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Contact
			//----------------------------------------------------------------------------
			return UpdateContact(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Contact
//----------------------------------------------------------------------------
func UnassignAccountFromContact(contactId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Contact
		//----------------------------------------------------------------------------
		return UpdateContact(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Contact
//----------------------------------------------------------------------------
func AssignOwnerToContact( contactId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
			// assign the Owner	to the Contact
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Contact
			//----------------------------------------------------------------------------
			return UpdateContact(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Contact
//----------------------------------------------------------------------------
func UnassignOwnerFromContact(contactId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Contact
		//----------------------------------------------------------------------------
		return UpdateContact(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more activitiesIds as a Activities to a Contact
//----------------------------------------------------------------------------
func AddActivitiesToContact ( contactId uint64, activitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more activitiesIds as a Activities from a Contact
//----------------------------------------------------------------------------
func RemoveActivitiesFromContact( contactId uint64, activitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more opportunitiesIds as a Opportunities to a Contact
//----------------------------------------------------------------------------
func AddOpportunitiesToContact ( contactId uint64, opportunitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

		// slice the ids on comma with no spaces
		ids := strings.Split( opportunitiesIds, ",")

		for _, opportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching opportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , opportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Opportunities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Opportunities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunities", opportunitiesId )
				return utils.RequestResult{false, msg, "unassignOpportunities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more opportunitiesIds as a Opportunities from a Contact
//----------------------------------------------------------------------------
func RemoveOpportunitiesFromContact( contactId uint64, opportunitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

		// slice the ids on comma with no spaces
		ids := strings.Split( opportunitiesIds, ",")

		for _, opportunitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Opportunity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Opportunity
			// with a matching opportunitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , opportunitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OpportunityObj from the Opportunities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Opportunities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunities", opportunitiesId )
				return utils.RequestResult{false, msg, "removeOpportunities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more casesIds as a Cases to a Contact
//----------------------------------------------------------------------------
func AddCasesToContact ( contactId uint64, casesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

		// slice the ids on comma with no spaces
		ids := strings.Split( casesIds, ",")

		for _, casesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Case_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Case_
			// with a matching casesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , casesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Cases using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Cases").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cases", casesId )
				return utils.RequestResult{false, msg, "unassignCases", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more casesIds as a Cases from a Contact
//----------------------------------------------------------------------------
func RemoveCasesFromContact( contactId uint64, casesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

		// slice the ids on comma with no spaces
		ids := strings.Split( casesIds, ",")

		for _, casesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Case_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Case_
			// with a matching casesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , casesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Case_Obj from the Cases array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Cases").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cases", casesId )
				return utils.RequestResult{false, msg, "removeCases", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a Contact
//----------------------------------------------------------------------------
func AddCampaignsToContact ( contactId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a Contact
//----------------------------------------------------------------------------
func RemoveCampaignsFromContact( contactId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more notesIds as a Notes to a Contact
//----------------------------------------------------------------------------
func AddNotesToContact ( contactId uint64, notesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more notesIds as a Notes from a Contact
//----------------------------------------------------------------------------
func RemoveNotesFromContact( contactId uint64, notesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more emailMessagesIds as a EmailMessages to a Contact
//----------------------------------------------------------------------------
func AddEmailMessagesToContact ( contactId uint64, emailMessagesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more emailMessagesIds as a EmailMessages from a Contact
//----------------------------------------------------------------------------
func RemoveEmailMessagesFromContact( contactId uint64, emailMessagesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contact with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContact(contactId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contact so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contact)

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
		// retrieve the modified Contact from the gorm
		//----------------------------------------------------------------------------
		return GetContact(contactId)

	} else {
		return parentRequestResult
	}
}

