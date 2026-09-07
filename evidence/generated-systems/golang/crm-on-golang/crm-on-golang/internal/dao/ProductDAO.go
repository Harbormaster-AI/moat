package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProductDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProduct - creates a new db entry
//----------------------------------------------------------------------------
func CreateProduct(obj model.Product)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Product with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Product", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProduct", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProduct - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProduct(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Product

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Product with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Product using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Product using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProduct", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProduct - returns all
//----------------------------------------------------------------------------
func GetAllProduct()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Product

	//----------------------------------------------------------------------------
	// Request the ORM to find all Product
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Product" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Product", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProduct", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProduct - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProduct(obj model.Product)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Product using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Product using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProduct", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProduct - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProduct(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProduct(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Product)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Product using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Product using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProduct", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Product
//----------------------------------------------------------------------------
func AssignOrganizationToProduct( productId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

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
			// assign the Organization	to the Product
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Product
			//----------------------------------------------------------------------------
			return UpdateProduct(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Product
//----------------------------------------------------------------------------
func UnassignOrganizationFromProduct(productId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Product
		//----------------------------------------------------------------------------
		return UpdateProduct(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more priceBookEntriesIds as a PriceBookEntries to a Product
//----------------------------------------------------------------------------
func AddPriceBookEntriesToProduct ( productId uint64, priceBookEntriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( priceBookEntriesIds, ",")

		for _, priceBookEntriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PriceBookEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PriceBookEntry
			// with a matching priceBookEntriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , priceBookEntriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PriceBookEntries using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PriceBookEntries").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBookEntries", priceBookEntriesId )
				return utils.RequestResult{false, msg, "unassignPriceBookEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more priceBookEntriesIds as a PriceBookEntries from a Product
//----------------------------------------------------------------------------
func RemovePriceBookEntriesFromProduct( productId uint64, priceBookEntriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( priceBookEntriesIds, ",")

		for _, priceBookEntriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PriceBookEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PriceBookEntry
			// with a matching priceBookEntriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , priceBookEntriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PriceBookEntryObj from the PriceBookEntries array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PriceBookEntries").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBookEntries", priceBookEntriesId )
				return utils.RequestResult{false, msg, "removePriceBookEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more opportunityLineItemsIds as a OpportunityLineItems to a Product
//----------------------------------------------------------------------------
func AddOpportunityLineItemsToProduct ( productId uint64, opportunityLineItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( opportunityLineItemsIds, ",")

		for _, opportunityLineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OpportunityLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OpportunityLineItem
			// with a matching opportunityLineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , opportunityLineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OpportunityLineItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OpportunityLineItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OpportunityLineItems", opportunityLineItemsId )
				return utils.RequestResult{false, msg, "unassignOpportunityLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more opportunityLineItemsIds as a OpportunityLineItems from a Product
//----------------------------------------------------------------------------
func RemoveOpportunityLineItemsFromProduct( productId uint64, opportunityLineItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( opportunityLineItemsIds, ",")

		for _, opportunityLineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OpportunityLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OpportunityLineItem
			// with a matching opportunityLineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , opportunityLineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OpportunityLineItemObj from the OpportunityLineItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OpportunityLineItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OpportunityLineItems", opportunityLineItemsId )
				return utils.RequestResult{false, msg, "removeOpportunityLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more quoteLineItemsIds as a QuoteLineItems to a Product
//----------------------------------------------------------------------------
func AddQuoteLineItemsToProduct ( productId uint64, quoteLineItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( quoteLineItemsIds, ",")

		for _, quoteLineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QuoteLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QuoteLineItem
			// with a matching quoteLineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quoteLineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the QuoteLineItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("QuoteLineItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "QuoteLineItems", quoteLineItemsId )
				return utils.RequestResult{false, msg, "unassignQuoteLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more quoteLineItemsIds as a QuoteLineItems from a Product
//----------------------------------------------------------------------------
func RemoveQuoteLineItemsFromProduct( productId uint64, quoteLineItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( quoteLineItemsIds, ",")

		for _, quoteLineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QuoteLineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QuoteLineItem
			// with a matching quoteLineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , quoteLineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QuoteLineItemObj from the QuoteLineItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("QuoteLineItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "QuoteLineItems", quoteLineItemsId )
				return utils.RequestResult{false, msg, "removeQuoteLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more orderItemsIds as a OrderItems to a Product
//----------------------------------------------------------------------------
func AddOrderItemsToProduct ( productId uint64, orderItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( orderItemsIds, ",")

		for _, orderItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OrderItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OrderItem
			// with a matching orderItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , orderItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OrderItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OrderItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OrderItems", orderItemsId )
				return utils.RequestResult{false, msg, "unassignOrderItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more orderItemsIds as a OrderItems from a Product
//----------------------------------------------------------------------------
func RemoveOrderItemsFromProduct( productId uint64, orderItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Product with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProduct(productId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Product so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Product)

		// slice the ids on comma with no spaces
		ids := strings.Split( orderItemsIds, ",")

		for _, orderItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OrderItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OrderItem
			// with a matching orderItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , orderItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OrderItemObj from the OrderItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OrderItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OrderItems", orderItemsId )
				return utils.RequestResult{false, msg, "removeOrderItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Product from the gorm
		//----------------------------------------------------------------------------
		return GetProduct(productId)

	} else {
		return parentRequestResult
	}
}

