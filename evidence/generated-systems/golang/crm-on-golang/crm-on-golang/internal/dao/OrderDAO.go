package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateOrder(obj model.Order)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Order with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Order", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Order

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Order with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Order using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Order using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOrder - returns all
//----------------------------------------------------------------------------
func GetAllOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Order

	//----------------------------------------------------------------------------
	// Request the ORM to find all Order
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Order" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Order", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOrder(obj model.Order)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Order using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Order using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Order using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Order using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Order
//----------------------------------------------------------------------------
func AssignOrganizationToOrder( orderId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

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
			// assign the Organization	to the Order
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Order
			//----------------------------------------------------------------------------
			return UpdateOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Order
//----------------------------------------------------------------------------
func UnassignOrganizationFromOrder(orderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Order
		//----------------------------------------------------------------------------
		return UpdateOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Account on a Order
//----------------------------------------------------------------------------
func AssignAccountToOrder( orderId uint64, accountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

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
			// assign the Account	to the Order
			//----------------------------------------------------------------------------
			parentObj.Account = &childObj

			//----------------------------------------------------------------------------
			// save the Order
			//----------------------------------------------------------------------------
			return UpdateOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Account", accountId )
			return utils.RequestResult{false, msg, "assignAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Account on a Order
//----------------------------------------------------------------------------
func UnassignAccountFromOrder(orderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// assign an empty Account to the Account
		//----------------------------------------------------------------------------
		parentObj.Account = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Account
		//----------------------------------------------------------------------------
		parentObj.AccountId = nil;

		//----------------------------------------------------------------------------
		// save the Order
		//----------------------------------------------------------------------------
		return UpdateOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Opportunity on a Order
//----------------------------------------------------------------------------
func AssignOpportunityToOrder( orderId uint64, opportunityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

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
			// assign the Opportunity	to the Order
			//----------------------------------------------------------------------------
			parentObj.Opportunity = &childObj

			//----------------------------------------------------------------------------
			// save the Order
			//----------------------------------------------------------------------------
			return UpdateOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunity", opportunityId )
			return utils.RequestResult{false, msg, "assignOpportunity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Opportunity on a Order
//----------------------------------------------------------------------------
func UnassignOpportunityFromOrder(orderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// assign an empty Opportunity to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.Opportunity = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.OpportunityId = nil;

		//----------------------------------------------------------------------------
		// save the Order
		//----------------------------------------------------------------------------
		return UpdateOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Quote on a Order
//----------------------------------------------------------------------------
func AssignQuoteToOrder( orderId uint64, quoteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Quote

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Quote with a
		// matching quoteId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, quoteId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Quote	to the Order
			//----------------------------------------------------------------------------
			parentObj.Quote = &childObj

			//----------------------------------------------------------------------------
			// save the Order
			//----------------------------------------------------------------------------
			return UpdateOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quote", quoteId )
			return utils.RequestResult{false, msg, "assignQuote", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Quote on a Order
//----------------------------------------------------------------------------
func UnassignQuoteFromOrder(orderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// assign an empty Quote to the Quote
		//----------------------------------------------------------------------------
		parentObj.Quote = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Quote
		//----------------------------------------------------------------------------
		parentObj.QuoteId = nil;

		//----------------------------------------------------------------------------
		// save the Order
		//----------------------------------------------------------------------------
		return UpdateOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a Order
//----------------------------------------------------------------------------
func AssignOwnerToOrder( orderId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

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
			// assign the Owner	to the Order
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the Order
			//----------------------------------------------------------------------------
			return UpdateOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a Order
//----------------------------------------------------------------------------
func UnassignOwnerFromOrder(orderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// assign an empty User to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the Order
		//----------------------------------------------------------------------------
		return UpdateOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Contract on a Order
//----------------------------------------------------------------------------
func AssignContractToOrder( orderId uint64, contractId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Contract

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Contract with a
		// matching contractId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, contractId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Contract	to the Order
			//----------------------------------------------------------------------------
			parentObj.Contract = &childObj

			//----------------------------------------------------------------------------
			// save the Order
			//----------------------------------------------------------------------------
			return UpdateOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contract", contractId )
			return utils.RequestResult{false, msg, "assignContract", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Contract on a Order
//----------------------------------------------------------------------------
func UnassignContractFromOrder(orderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// assign an empty Contract to the Contract
		//----------------------------------------------------------------------------
		parentObj.Contract = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Contract
		//----------------------------------------------------------------------------
		parentObj.ContractId = nil;

		//----------------------------------------------------------------------------
		// save the Order
		//----------------------------------------------------------------------------
		return UpdateOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PriceBook on a Order
//----------------------------------------------------------------------------
func AssignPriceBookToOrder( orderId uint64, priceBookId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

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
			// assign the PriceBook	to the Order
			//----------------------------------------------------------------------------
			parentObj.PriceBook = &childObj

			//----------------------------------------------------------------------------
			// save the Order
			//----------------------------------------------------------------------------
			return UpdateOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBook", priceBookId )
			return utils.RequestResult{false, msg, "assignPriceBook", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PriceBook on a Order
//----------------------------------------------------------------------------
func UnassignPriceBookFromOrder(orderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		//----------------------------------------------------------------------------
		// assign an empty PriceBook to the PriceBook
		//----------------------------------------------------------------------------
		parentObj.PriceBook = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PriceBook
		//----------------------------------------------------------------------------
		parentObj.PriceBookId = nil;

		//----------------------------------------------------------------------------
		// save the Order
		//----------------------------------------------------------------------------
		return UpdateOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more itemsIds as a Items to a Order
//----------------------------------------------------------------------------
func AddItemsToOrder ( orderId uint64, itemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OrderItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OrderItem
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Items using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "unassignItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Order from the gorm
		//----------------------------------------------------------------------------
		return GetOrder(orderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more itemsIds as a Items from a Order
//----------------------------------------------------------------------------
func RemoveItemsFromOrder( orderId uint64, itemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Order with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrder(orderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Order so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Order)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OrderItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OrderItem
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OrderItemObj from the Items array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "removeItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Order from the gorm
		//----------------------------------------------------------------------------
		return GetOrder(orderId)

	} else {
		return parentRequestResult
	}
}

