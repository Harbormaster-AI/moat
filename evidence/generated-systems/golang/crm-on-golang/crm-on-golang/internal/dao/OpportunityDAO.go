package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OpportunityDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOpportunity - creates a new db entry
//----------------------------------------------------------------------------
func CreateOpportunity(obj model.Opportunity)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Opportunity with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Opportunity", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOpportunity", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOpportunity - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOpportunity(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Opportunity

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Opportunity with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Opportunity using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Opportunity using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOpportunity", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOpportunity - returns all
//----------------------------------------------------------------------------
func GetAllOpportunity()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Opportunity

	//----------------------------------------------------------------------------
	// Request the ORM to find all Opportunity
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Opportunity" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Opportunity", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOpportunity", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOpportunity - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOpportunity(obj model.Opportunity)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Opportunity using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Opportunity using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOpportunity", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOpportunity - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOpportunity(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOpportunity(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Opportunity)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Opportunity using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Opportunity using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOpportunity", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Opportunity
//----------------------------------------------------------------------------
func AssignOrganizationToOpportunity( opportunityId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
			// assign the Organization	to the Opportunity
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Opportunity
			//----------------------------------------------------------------------------
			return UpdateOpportunity(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Opportunity
//----------------------------------------------------------------------------
func UnassignOrganizationFromOpportunity(opportunityId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Opportunity
		//----------------------------------------------------------------------------
		return UpdateOpportunity(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Opportunity
//----------------------------------------------------------------------------
func AssignAccountToOpportunity( opportunityId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
			// assign the Account	to the Opportunity
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Opportunity
			//----------------------------------------------------------------------------
			return UpdateOpportunity(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Opportunity
//----------------------------------------------------------------------------
func UnassignAccountFromOpportunity(opportunityId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Opportunity
		//----------------------------------------------------------------------------
		return UpdateOpportunity(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Opportunity
//----------------------------------------------------------------------------
func AssignOwnerToOpportunity( opportunityId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
			// assign the Owner	to the Opportunity
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Opportunity
			//----------------------------------------------------------------------------
			return UpdateOpportunity(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Opportunity
//----------------------------------------------------------------------------
func UnassignOwnerFromOpportunity(opportunityId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Opportunity
		//----------------------------------------------------------------------------
		return UpdateOpportunity(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more contactsIds as a Contacts to a Opportunity
//----------------------------------------------------------------------------
func AddContactsToOpportunity ( opportunityId uint64, contactsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contactsIds as a Contacts from a Opportunity
//----------------------------------------------------------------------------
func RemoveContactsFromOpportunity( opportunityId uint64, contactsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more lineItemsIds as a LineItems to a Opportunity
//----------------------------------------------------------------------------
func AddLineItemsToOpportunity ( opportunityId uint64, lineItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OpportunityLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OpportunityLineItem
			// with a matching lineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LineItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItems", lineItemsId )
				return utils.RequestResult{false, msg, "unassignLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more lineItemsIds as a LineItems from a Opportunity
//----------------------------------------------------------------------------
func RemoveLineItemsFromOpportunity( opportunityId uint64, lineItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OpportunityLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OpportunityLineItem
			// with a matching lineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OpportunityLineItemObj from the LineItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItems", lineItemsId )
				return utils.RequestResult{false, msg, "removeLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more stageHistoryIds as a StageHistory to a Opportunity
//----------------------------------------------------------------------------
func AddStageHistoryToOpportunity ( opportunityId uint64, stageHistoryIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( stageHistoryIds, ",")

		for _, stageHistoryId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OpportunityStageHistory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OpportunityStageHistory
			// with a matching stageHistoryId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , stageHistoryId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the StageHistory using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("StageHistory").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "StageHistory", stageHistoryId )
				return utils.RequestResult{false, msg, "unassignStageHistory", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more stageHistoryIds as a StageHistory from a Opportunity
//----------------------------------------------------------------------------
func RemoveStageHistoryFromOpportunity( opportunityId uint64, stageHistoryIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( stageHistoryIds, ",")

		for _, stageHistoryId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OpportunityStageHistory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OpportunityStageHistory
			// with a matching stageHistoryId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , stageHistoryId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OpportunityStageHistoryObj from the StageHistory array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("StageHistory").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "StageHistory", stageHistoryId )
				return utils.RequestResult{false, msg, "removeStageHistory", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more quotesIds as a Quotes to a Opportunity
//----------------------------------------------------------------------------
func AddQuotesToOpportunity ( opportunityId uint64, quotesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Quotes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "unassignQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more quotesIds as a Quotes from a Opportunity
//----------------------------------------------------------------------------
func RemoveQuotesFromOpportunity( opportunityId uint64, quotesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( quotesIds, ",")

		for _, quotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Quote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Quote
			// with a matching quotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QuoteObj from the Quotes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Quotes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quotes", quotesId )
				return utils.RequestResult{false, msg, "removeQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ordersIds as a Orders to a Opportunity
//----------------------------------------------------------------------------
func AddOrdersToOpportunity ( opportunityId uint64, ordersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Order

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Order
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Orders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "unassignOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ordersIds as a Orders from a Opportunity
//----------------------------------------------------------------------------
func RemoveOrdersFromOpportunity( opportunityId uint64, ordersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Order

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Order
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OrderObj from the Orders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "removeOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more campaignsIds as a Campaigns to a Opportunity
//----------------------------------------------------------------------------
func AddCampaignsToOpportunity ( opportunityId uint64, campaignsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more campaignsIds as a Campaigns from a Opportunity
//----------------------------------------------------------------------------
func RemoveCampaignsFromOpportunity( opportunityId uint64, campaignsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more activitiesIds as a Activities to a Opportunity
//----------------------------------------------------------------------------
func AddActivitiesToOpportunity ( opportunityId uint64, activitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more activitiesIds as a Activities from a Opportunity
//----------------------------------------------------------------------------
func RemoveActivitiesFromOpportunity( opportunityId uint64, activitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more teamsIds as a Teams to a Opportunity
//----------------------------------------------------------------------------
func AddTeamsToOpportunity ( opportunityId uint64, teamsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more teamsIds as a Teams from a Opportunity
//----------------------------------------------------------------------------
func RemoveTeamsFromOpportunity( opportunityId uint64, teamsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Opportunity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunity(opportunityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Opportunity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Opportunity)

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
		// retrieve the modified Opportunity from the gorm
		//----------------------------------------------------------------------------
		return GetOpportunity(opportunityId)

	} else {
		return parentRequestResult
	}
}

