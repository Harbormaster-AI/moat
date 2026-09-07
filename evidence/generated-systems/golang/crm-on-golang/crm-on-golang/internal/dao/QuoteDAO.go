package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing QuoteDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateQuote - creates a new db entry
//----------------------------------------------------------------------------
func CreateQuote(obj model.Quote)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Quote with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Quote", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateQuote", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetQuote - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetQuote(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Quote

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Quote with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Quote using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Quote using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetQuote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllQuote - returns all
//----------------------------------------------------------------------------
func GetAllQuote()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Quote

	//----------------------------------------------------------------------------
	// Request the ORM to find all Quote
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Quote" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Quote", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllQuote", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateQuote - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateQuote(obj model.Quote)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Quote using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Quote using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateQuote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteQuote - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteQuote(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetQuote(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Quote using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Quote using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteQuote", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Quote
//----------------------------------------------------------------------------
func AssignOrganizationToQuote( quoteId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

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
			// assign the Organization	to the Quote
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Quote
//----------------------------------------------------------------------------
func UnassignOrganizationFromQuote(quoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Quote
//----------------------------------------------------------------------------
func AssignAccountToQuote( quoteId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

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
			// assign the Account	to the Quote
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Quote
//----------------------------------------------------------------------------
func UnassignAccountFromQuote(quoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Opportunity on a Quote
//----------------------------------------------------------------------------
func AssignOpportunityToQuote( quoteId uint64, opportunityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

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
			// assign the Opportunity	to the Quote
			//----------------------------------------------------------------------------
			parentObj.Opportunity = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunity", opportunityId )
			return utils.RequestResult{false, msg, "assignOpportunity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Opportunity on a Quote
//----------------------------------------------------------------------------
func UnassignOpportunityFromQuote(quoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// assign an empty Opportunity to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.Opportunity = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.OpportunityId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Quote
//----------------------------------------------------------------------------
func AssignOwnerToQuote( quoteId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

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
			// assign the Owner	to the Quote
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Quote
//----------------------------------------------------------------------------
func UnassignOwnerFromQuote(quoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PriceBook on a Quote
//----------------------------------------------------------------------------
func AssignPriceBookToQuote( quoteId uint64, priceBookId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PriceBook

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PriceBook with a
		// matching priceBookId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, priceBookId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PriceBook	to the Quote
			//----------------------------------------------------------------------------
			parentObj.PriceBook = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBook", priceBookId )
			return utils.RequestResult{false, msg, "assignPriceBook", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PriceBook on a Quote
//----------------------------------------------------------------------------
func UnassignPriceBookFromQuote(quoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// assign an empty PriceBook to the PriceBook
		//----------------------------------------------------------------------------
		parentObj.PriceBook = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PriceBook
		//----------------------------------------------------------------------------
		parentObj.PriceBookId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Order on a Quote
//----------------------------------------------------------------------------
func AssignOrderToQuote( quoteId uint64, orderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Order

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Order with a
		// matching orderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, orderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Order	to the Quote
			//----------------------------------------------------------------------------
			parentObj.Order = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Order", orderId )
			return utils.RequestResult{false, msg, "assignOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Order on a Quote
//----------------------------------------------------------------------------
func UnassignOrderFromQuote(quoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// assign an empty Order to the Order
		//----------------------------------------------------------------------------
		parentObj.Order = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Order
		//----------------------------------------------------------------------------
		parentObj.OrderId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more lineItemsIds as a LineItems to a Quote
//----------------------------------------------------------------------------
func AddLineItemsToQuote ( quoteId uint64, lineItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QuoteLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QuoteLineItem
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
		// retrieve the modified Quote from the gorm
		//----------------------------------------------------------------------------
		return GetQuote(quoteId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more lineItemsIds as a LineItems from a Quote
//----------------------------------------------------------------------------
func RemoveLineItemsFromQuote( quoteId uint64, lineItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QuoteLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QuoteLineItem
			// with a matching lineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QuoteLineItemObj from the LineItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItems", lineItemsId )
				return utils.RequestResult{false, msg, "removeLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Quote from the gorm
		//----------------------------------------------------------------------------
		return GetQuote(quoteId)

	} else {
		return parentRequestResult
	}
}

