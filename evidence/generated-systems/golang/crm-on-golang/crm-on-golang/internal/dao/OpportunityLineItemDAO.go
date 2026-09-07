package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OpportunityLineItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOpportunityLineItem - creates a new db entry
//----------------------------------------------------------------------------
func CreateOpportunityLineItem(obj model.OpportunityLineItem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a OpportunityLineItem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a OpportunityLineItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOpportunityLineItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOpportunityLineItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOpportunityLineItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.OpportunityLineItem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a OpportunityLineItem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a OpportunityLineItem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a OpportunityLineItem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOpportunityLineItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOpportunityLineItem - returns all
//----------------------------------------------------------------------------
func GetAllOpportunityLineItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.OpportunityLineItem

	//----------------------------------------------------------------------------
	// Request the ORM to find all OpportunityLineItem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all OpportunityLineItem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all OpportunityLineItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOpportunityLineItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOpportunityLineItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOpportunityLineItem(obj model.OpportunityLineItem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a OpportunityLineItem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a OpportunityLineItem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOpportunityLineItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOpportunityLineItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOpportunityLineItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the OpportunityLineItem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOpportunityLineItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.OpportunityLineItem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a OpportunityLineItem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a OpportunityLineItem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOpportunityLineItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Opportunity on a OpportunityLineItem
//----------------------------------------------------------------------------
func AssignOpportunityToOpportunityLineItem( opportunityLineItemId uint64, opportunityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OpportunityLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityLineItem(opportunityLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityLineItem)

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
			// assign the Opportunity	to the OpportunityLineItem
			//----------------------------------------------------------------------------
			parentObj.Opportunity = &childObj

			//----------------------------------------------------------------------------
			// save the OpportunityLineItem
			//----------------------------------------------------------------------------
			return UpdateOpportunityLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunity", opportunityId )
			return utils.RequestResult{false, msg, "assignOpportunity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Opportunity on a OpportunityLineItem
//----------------------------------------------------------------------------
func UnassignOpportunityFromOpportunityLineItem(opportunityLineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OpportunityLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityLineItem(opportunityLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityLineItem)

		//----------------------------------------------------------------------------
		// assign an empty Opportunity to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.Opportunity = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.OpportunityId = nil;

		//----------------------------------------------------------------------------
		// save the OpportunityLineItem
		//----------------------------------------------------------------------------
		return UpdateOpportunityLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a OpportunityLineItem
//----------------------------------------------------------------------------
func AssignProductToOpportunityLineItem( opportunityLineItemId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OpportunityLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityLineItem(opportunityLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityLineItem)

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
			// assign the Product	to the OpportunityLineItem
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the OpportunityLineItem
			//----------------------------------------------------------------------------
			return UpdateOpportunityLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a OpportunityLineItem
//----------------------------------------------------------------------------
func UnassignProductFromOpportunityLineItem(opportunityLineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OpportunityLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityLineItem(opportunityLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityLineItem)

		//----------------------------------------------------------------------------
		// assign an empty Product to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the OpportunityLineItem
		//----------------------------------------------------------------------------
		return UpdateOpportunityLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PriceBookEntry on a OpportunityLineItem
//----------------------------------------------------------------------------
func AssignPriceBookEntryToOpportunityLineItem( opportunityLineItemId uint64, priceBookEntryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OpportunityLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityLineItem(opportunityLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityLineItem)

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
			// assign the PriceBookEntry	to the OpportunityLineItem
			//----------------------------------------------------------------------------
			parentObj.PriceBookEntry = &childObj

			//----------------------------------------------------------------------------
			// save the OpportunityLineItem
			//----------------------------------------------------------------------------
			return UpdateOpportunityLineItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBookEntry", priceBookEntryId )
			return utils.RequestResult{false, msg, "assignPriceBookEntry", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PriceBookEntry on a OpportunityLineItem
//----------------------------------------------------------------------------
func UnassignPriceBookEntryFromOpportunityLineItem(opportunityLineItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OpportunityLineItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityLineItem(opportunityLineItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityLineItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityLineItem)

		//----------------------------------------------------------------------------
		// assign an empty PriceBookEntry to the PriceBookEntry
		//----------------------------------------------------------------------------
		parentObj.PriceBookEntry = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PriceBookEntry
		//----------------------------------------------------------------------------
		parentObj.PriceBookEntryId = nil;

		//----------------------------------------------------------------------------
		// save the OpportunityLineItem
		//----------------------------------------------------------------------------
		return UpdateOpportunityLineItem(parentObj)

	} else {
		return parentRequestResult
	}

}


