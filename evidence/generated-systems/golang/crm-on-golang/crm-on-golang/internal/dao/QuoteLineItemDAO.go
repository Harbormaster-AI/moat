package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing QuoteLineItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateQuoteLineItem - creates a new db entry
//----------------------------------------------------------------------------
func CreateQuoteLineItem(obj model.QuoteLineItem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a QuoteLineItem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a QuoteLineItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateQuoteLineItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetQuoteLineItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetQuoteLineItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.QuoteLineItem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a QuoteLineItem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a QuoteLineItem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a QuoteLineItem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetQuoteLineItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllQuoteLineItem - returns all
//----------------------------------------------------------------------------
func GetAllQuoteLineItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.QuoteLineItem

	//----------------------------------------------------------------------------
	// Request the ORM to find all QuoteLineItem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all QuoteLineItem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all QuoteLineItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllQuoteLineItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateQuoteLineItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateQuoteLineItem(obj model.QuoteLineItem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a QuoteLineItem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a QuoteLineItem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateQuoteLineItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteQuoteLineItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteQuoteLineItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetQuoteLineItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a QuoteLineItem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a QuoteLineItem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteQuoteLineItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Quote on a QuoteLineItem
//----------------------------------------------------------------------------
func AssignQuoteToQuoteLineItem( quoteLineItemId uint64, quoteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

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
			// assign the Quote	to the QuoteLineItem
			//----------------------------------------------------------------------------
			parentObj.Quote = &childObj

			//----------------------------------------------------------------------------
			// save the QuoteLineItem
			//----------------------------------------------------------------------------
			return UpdateQuoteLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quote", quoteId )
			return utils.RequestResult{false, msg, "assignQuote", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Quote on a QuoteLineItem
//----------------------------------------------------------------------------
func UnassignQuoteFromQuoteLineItem(quoteLineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// assign an empty Quote to the Quote
		//----------------------------------------------------------------------------
		parentObj.Quote = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Quote
		//----------------------------------------------------------------------------
		parentObj.QuoteId = nil;

		//----------------------------------------------------------------------------
		// save the QuoteLineItem
		//----------------------------------------------------------------------------
		return UpdateQuoteLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a QuoteLineItem
//----------------------------------------------------------------------------
func AssignProductToQuoteLineItem( quoteLineItemId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Product

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Product with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the QuoteLineItem
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the QuoteLineItem
			//----------------------------------------------------------------------------
			return UpdateQuoteLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a QuoteLineItem
//----------------------------------------------------------------------------
func UnassignProductFromQuoteLineItem(quoteLineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// assign an empty Product to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the QuoteLineItem
		//----------------------------------------------------------------------------
		return UpdateQuoteLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PriceBookEntry on a QuoteLineItem
//----------------------------------------------------------------------------
func AssignPriceBookEntryToQuoteLineItem( quoteLineItemId uint64, priceBookEntryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PriceBookEntry

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PriceBookEntry with a
		// matching priceBookEntryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, priceBookEntryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PriceBookEntry	to the QuoteLineItem
			//----------------------------------------------------------------------------
			parentObj.PriceBookEntry = &childObj

			//----------------------------------------------------------------------------
			// save the QuoteLineItem
			//----------------------------------------------------------------------------
			return UpdateQuoteLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBookEntry", priceBookEntryId )
			return utils.RequestResult{false, msg, "assignPriceBookEntry", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PriceBookEntry on a QuoteLineItem
//----------------------------------------------------------------------------
func UnassignPriceBookEntryFromQuoteLineItem(quoteLineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// assign an empty PriceBookEntry to the PriceBookEntry
		//----------------------------------------------------------------------------
		parentObj.PriceBookEntry = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PriceBookEntry
		//----------------------------------------------------------------------------
		parentObj.PriceBookEntryId = nil;

		//----------------------------------------------------------------------------
		// save the QuoteLineItem
		//----------------------------------------------------------------------------
		return UpdateQuoteLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a OpportunityLineItem on a QuoteLineItem
//----------------------------------------------------------------------------
func AssignOpportunityLineItemToQuoteLineItem( quoteLineItemId uint64, opportunityLineItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.OpportunityLineItem

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a OpportunityLineItem with a
		// matching opportunityLineItemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, opportunityLineItemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the OpportunityLineItem	to the QuoteLineItem
			//----------------------------------------------------------------------------
			parentObj.OpportunityLineItem = &childObj

			//----------------------------------------------------------------------------
			// save the QuoteLineItem
			//----------------------------------------------------------------------------
			return UpdateQuoteLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OpportunityLineItem", opportunityLineItemId )
			return utils.RequestResult{false, msg, "assignOpportunityLineItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a OpportunityLineItem on a QuoteLineItem
//----------------------------------------------------------------------------
func UnassignOpportunityLineItemFromQuoteLineItem(quoteLineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the QuoteLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuoteLineItem(quoteLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.QuoteLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.QuoteLineItem)

		//----------------------------------------------------------------------------
		// assign an empty OpportunityLineItem to the OpportunityLineItem
		//----------------------------------------------------------------------------
		parentObj.OpportunityLineItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the OpportunityLineItem
		//----------------------------------------------------------------------------
		parentObj.OpportunityLineItemId = nil;

		//----------------------------------------------------------------------------
		// save the QuoteLineItem
		//----------------------------------------------------------------------------
		return UpdateQuoteLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}


