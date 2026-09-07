package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CampaignDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCampaign - creates a new db entry
//----------------------------------------------------------------------------
func CreateCampaign(obj model.Campaign)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Campaign with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Campaign", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCampaign", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCampaign - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCampaign(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Campaign

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Campaign with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Campaign using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Campaign using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCampaign", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCampaign - returns all
//----------------------------------------------------------------------------
func GetAllCampaign()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Campaign

	//----------------------------------------------------------------------------
	// Request the ORM to find all Campaign
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Campaign" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Campaign", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCampaign", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCampaign - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCampaign(obj model.Campaign)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Campaign using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Campaign using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCampaign", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCampaign - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCampaign(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCampaign(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Campaign using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Campaign using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCampaign", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Campaign
//----------------------------------------------------------------------------
func AssignOrganizationToCampaign( campaignId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

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
			// assign the Organization	to the Campaign
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Campaign
			//----------------------------------------------------------------------------
			return UpdateCampaign(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Campaign
//----------------------------------------------------------------------------
func UnassignOrganizationFromCampaign(campaignId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Campaign
		//----------------------------------------------------------------------------
		return UpdateCampaign(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ParentCampaign on a Campaign
//----------------------------------------------------------------------------
func AssignParentCampaignToCampaign( campaignId uint64, parentCampaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Campaign

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Campaign with a
		// matching parentCampaignId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, parentCampaignId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ParentCampaign	to the Campaign
			//----------------------------------------------------------------------------
			parentObj.ParentCampaign = &childObj

			//----------------------------------------------------------------------------
			// save the Campaign
			//----------------------------------------------------------------------------
			return UpdateCampaign(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ParentCampaign", parentCampaignId )
			return utils.RequestResult{false, msg, "assignParentCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ParentCampaign on a Campaign
//----------------------------------------------------------------------------
func UnassignParentCampaignFromCampaign(campaignId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the ParentCampaign
		//----------------------------------------------------------------------------
		parentObj.ParentCampaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ParentCampaign
		//----------------------------------------------------------------------------
		parentObj.ParentCampaignId = nil;

		//----------------------------------------------------------------------------
		// save the Campaign
		//----------------------------------------------------------------------------
		return UpdateCampaign(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more childCampaignsIds as a ChildCampaigns to a Campaign
//----------------------------------------------------------------------------
func AddChildCampaignsToCampaign ( campaignId uint64, childCampaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( childCampaignsIds, ",")

		for _, childCampaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Campaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Campaign
			// with a matching childCampaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childCampaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ChildCampaigns using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildCampaigns").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildCampaigns", childCampaignsId )
				return utils.RequestResult{false, msg, "unassignChildCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more childCampaignsIds as a ChildCampaigns from a Campaign
//----------------------------------------------------------------------------
func RemoveChildCampaignsFromCampaign( campaignId uint64, childCampaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( childCampaignsIds, ",")

		for _, childCampaignsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Campaign

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Campaign
			// with a matching childCampaignsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childCampaignsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CampaignObj from the ChildCampaigns array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildCampaigns").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildCampaigns", childCampaignsId )
				return utils.RequestResult{false, msg, "removeChildCampaigns", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more membersIds as a Members to a Campaign
//----------------------------------------------------------------------------
func AddMembersToCampaign ( campaignId uint64, membersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( membersIds, ",")

		for _, membersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CampaignMember

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CampaignMember
			// with a matching membersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , membersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Members using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Members").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Members", membersId )
				return utils.RequestResult{false, msg, "unassignMembers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more membersIds as a Members from a Campaign
//----------------------------------------------------------------------------
func RemoveMembersFromCampaign( campaignId uint64, membersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( membersIds, ",")

		for _, membersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CampaignMember

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CampaignMember
			// with a matching membersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , membersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CampaignMemberObj from the Members array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Members").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Members", membersId )
				return utils.RequestResult{false, msg, "removeMembers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more opportunitiesIds as a Opportunities to a Campaign
//----------------------------------------------------------------------------
func AddOpportunitiesToCampaign ( campaignId uint64, opportunitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

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
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more opportunitiesIds as a Opportunities from a Campaign
//----------------------------------------------------------------------------
func RemoveOpportunitiesFromCampaign( campaignId uint64, opportunitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

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
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Campaign
//----------------------------------------------------------------------------
func AddAccountsToCampaign ( campaignId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Accounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a Campaign
//----------------------------------------------------------------------------
func RemoveAccountsFromCampaign( campaignId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the Accounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more leadsIds as a Leads to a Campaign
//----------------------------------------------------------------------------
func AddLeadsToCampaign ( campaignId uint64, leadsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( leadsIds, ",")

		for _, leadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Lead

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Lead
			// with a matching leadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , leadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Leads using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Leads").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Leads", leadsId )
				return utils.RequestResult{false, msg, "unassignLeads", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more leadsIds as a Leads from a Campaign
//----------------------------------------------------------------------------
func RemoveLeadsFromCampaign( campaignId uint64, leadsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( leadsIds, ",")

		for _, leadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Lead

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Lead
			// with a matching leadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , leadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LeadObj from the Leads array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Leads").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Leads", leadsId )
				return utils.RequestResult{false, msg, "removeLeads", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contactsIds as a Contacts to a Campaign
//----------------------------------------------------------------------------
func AddContactsToCampaign ( campaignId uint64, contactsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( contactsIds, ",")

		for _, contactsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contact

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contact
			// with a matching contactsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contactsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Contacts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contacts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contacts", contactsId )
				return utils.RequestResult{false, msg, "unassignContacts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contactsIds as a Contacts from a Campaign
//----------------------------------------------------------------------------
func RemoveContactsFromCampaign( campaignId uint64, contactsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( contactsIds, ",")

		for _, contactsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contact

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contact
			// with a matching contactsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contactsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ContactObj from the Contacts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contacts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contacts", contactsId )
				return utils.RequestResult{false, msg, "removeContacts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more teamsIds as a Teams to a Campaign
//----------------------------------------------------------------------------
func AddTeamsToCampaign ( campaignId uint64, teamsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( teamsIds, ",")

		for _, teamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Team

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Team
			// with a matching teamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , teamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Teams using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Teams").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Teams", teamsId )
				return utils.RequestResult{false, msg, "unassignTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more teamsIds as a Teams from a Campaign
//----------------------------------------------------------------------------
func RemoveTeamsFromCampaign( campaignId uint64, teamsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

		// slice the ids on comma with no spaces
		ids := strings.Split( teamsIds, ",")

		for _, teamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Team

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Team
			// with a matching teamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , teamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TeamObj from the Teams array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Teams").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Teams", teamsId )
				return utils.RequestResult{false, msg, "removeTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more activitiesIds as a Activities to a Campaign
//----------------------------------------------------------------------------
func AddActivitiesToCampaign ( campaignId uint64, activitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

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
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more activitiesIds as a Activities from a Campaign
//----------------------------------------------------------------------------
func RemoveActivitiesFromCampaign( campaignId uint64, activitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Campaign with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCampaign(campaignId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Campaign so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Campaign)

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
		// retrieve the modified Campaign from the gorm
		//----------------------------------------------------------------------------
		return GetCampaign(campaignId)

	} else {
		return parentRequestResult
	}
}

