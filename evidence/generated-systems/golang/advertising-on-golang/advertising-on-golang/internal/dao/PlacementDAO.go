package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PlacementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePlacement - creates a new db entry
//----------------------------------------------------------------------------
func CreatePlacement(obj model.Placement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Placement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Placement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePlacement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPlacement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPlacement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Placement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Placement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Placement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Placement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPlacement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPlacement - returns all
//----------------------------------------------------------------------------
func GetAllPlacement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Placement

	//----------------------------------------------------------------------------
	// Request the ORM to find all Placement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Placement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Placement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPlacement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePlacement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePlacement(obj model.Placement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Placement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Placement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePlacement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePlacement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePlacement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Placement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPlacement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Placement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Placement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Placement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Placement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePlacement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a LineItem on a Placement
//----------------------------------------------------------------------------
func AssignLineItemToPlacement( placementId uint64, lineItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Placement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlacement(placementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Placement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Placement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LineItem

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LineItem with a
		// matching lineItemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lineItemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LineItem	to the Placement
			//----------------------------------------------------------------------------
			parentObj.LineItem = &childObj

			//----------------------------------------------------------------------------
			// save the Placement
			//----------------------------------------------------------------------------
			return UpdatePlacement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItem", lineItemId )
			return utils.RequestResult{false, msg, "assignLineItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LineItem on a Placement
//----------------------------------------------------------------------------
func UnassignLineItemFromPlacement(placementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Placement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlacement(placementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Placement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Placement)

		//----------------------------------------------------------------------------
		// assign an empty LineItem to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItemId = nil;

		//----------------------------------------------------------------------------
		// save the Placement
		//----------------------------------------------------------------------------
		return UpdatePlacement(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AdSlot on a Placement
//----------------------------------------------------------------------------
func AssignAdSlotToPlacement( placementId uint64, adSlotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Placement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlacement(placementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Placement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Placement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AdSlot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AdSlot with a
		// matching adSlotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, adSlotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AdSlot	to the Placement
			//----------------------------------------------------------------------------
			parentObj.AdSlot = &childObj

			//----------------------------------------------------------------------------
			// save the Placement
			//----------------------------------------------------------------------------
			return UpdatePlacement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdSlot", adSlotId )
			return utils.RequestResult{false, msg, "assignAdSlot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AdSlot on a Placement
//----------------------------------------------------------------------------
func UnassignAdSlotFromPlacement(placementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Placement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlacement(placementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Placement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Placement)

		//----------------------------------------------------------------------------
		// assign an empty AdSlot to the AdSlot
		//----------------------------------------------------------------------------
		parentObj.AdSlot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AdSlot
		//----------------------------------------------------------------------------
		parentObj.AdSlotId = nil;

		//----------------------------------------------------------------------------
		// save the Placement
		//----------------------------------------------------------------------------
		return UpdatePlacement(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Deal on a Placement
//----------------------------------------------------------------------------
func AssignDealToPlacement( placementId uint64, dealId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Placement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlacement(placementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Placement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Placement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Deal

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Deal with a
		// matching dealId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, dealId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Deal	to the Placement
			//----------------------------------------------------------------------------
			parentObj.Deal = &childObj

			//----------------------------------------------------------------------------
			// save the Placement
			//----------------------------------------------------------------------------
			return UpdatePlacement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deal", dealId )
			return utils.RequestResult{false, msg, "assignDeal", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Deal on a Placement
//----------------------------------------------------------------------------
func UnassignDealFromPlacement(placementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Placement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlacement(placementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Placement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Placement)

		//----------------------------------------------------------------------------
		// assign an empty Deal to the Deal
		//----------------------------------------------------------------------------
		parentObj.Deal = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Deal
		//----------------------------------------------------------------------------
		parentObj.DealId = nil;

		//----------------------------------------------------------------------------
		// save the Placement
		//----------------------------------------------------------------------------
		return UpdatePlacement(parentObj)

	} else {
		return parentRequestResult
	}

}


