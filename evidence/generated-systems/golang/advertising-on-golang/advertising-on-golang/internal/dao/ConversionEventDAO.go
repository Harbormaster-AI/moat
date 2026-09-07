package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ConversionEventDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateConversionEvent - creates a new db entry
//----------------------------------------------------------------------------
func CreateConversionEvent(obj model.ConversionEvent)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ConversionEvent with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ConversionEvent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateConversionEvent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetConversionEvent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetConversionEvent(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ConversionEvent

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ConversionEvent with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ConversionEvent using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ConversionEvent using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetConversionEvent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllConversionEvent - returns all
//----------------------------------------------------------------------------
func GetAllConversionEvent()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ConversionEvent

	//----------------------------------------------------------------------------
	// Request the ORM to find all ConversionEvent
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ConversionEvent" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ConversionEvent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllConversionEvent", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateConversionEvent - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateConversionEvent(obj model.ConversionEvent)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ConversionEvent using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ConversionEvent using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateConversionEvent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteConversionEvent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteConversionEvent(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ConversionEvent with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetConversionEvent(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConversionEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ConversionEvent)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ConversionEvent using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ConversionEvent using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteConversionEvent", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Campaign on a ConversionEvent
//----------------------------------------------------------------------------
func AssignCampaignToConversionEvent( conversionEventId uint64, campaignId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ConversionEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConversionEvent(conversionEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConversionEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConversionEvent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Campaign

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Campaign with a
		// matching campaignId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, campaignId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Campaign	to the ConversionEvent
			//----------------------------------------------------------------------------
			parentObj.Campaign = &childObj

			//----------------------------------------------------------------------------
			// save the ConversionEvent
			//----------------------------------------------------------------------------
			return UpdateConversionEvent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Campaign", campaignId )
			return utils.RequestResult{false, msg, "assignCampaign", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Campaign on a ConversionEvent
//----------------------------------------------------------------------------
func UnassignCampaignFromConversionEvent(conversionEventId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ConversionEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConversionEvent(conversionEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConversionEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConversionEvent)

		//----------------------------------------------------------------------------
		// assign an empty Campaign to the Campaign
		//----------------------------------------------------------------------------
		parentObj.Campaign = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Campaign
		//----------------------------------------------------------------------------
		parentObj.CampaignId = nil;

		//----------------------------------------------------------------------------
		// save the ConversionEvent
		//----------------------------------------------------------------------------
		return UpdateConversionEvent(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LineItem on a ConversionEvent
//----------------------------------------------------------------------------
func AssignLineItemToConversionEvent( conversionEventId uint64, lineItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ConversionEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConversionEvent(conversionEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConversionEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConversionEvent)

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
			// assign the LineItem	to the ConversionEvent
			//----------------------------------------------------------------------------
			parentObj.LineItem = &childObj

			//----------------------------------------------------------------------------
			// save the ConversionEvent
			//----------------------------------------------------------------------------
			return UpdateConversionEvent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItem", lineItemId )
			return utils.RequestResult{false, msg, "assignLineItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LineItem on a ConversionEvent
//----------------------------------------------------------------------------
func UnassignLineItemFromConversionEvent(conversionEventId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ConversionEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConversionEvent(conversionEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConversionEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConversionEvent)

		//----------------------------------------------------------------------------
		// assign an empty LineItem to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LineItem
		//----------------------------------------------------------------------------
		parentObj.LineItemId = nil;

		//----------------------------------------------------------------------------
		// save the ConversionEvent
		//----------------------------------------------------------------------------
		return UpdateConversionEvent(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a TrackingPixel on a ConversionEvent
//----------------------------------------------------------------------------
func AssignTrackingPixelToConversionEvent( conversionEventId uint64, trackingPixelId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ConversionEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConversionEvent(conversionEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConversionEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConversionEvent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.TrackingPixel

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TrackingPixel with a
		// matching trackingPixelId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, trackingPixelId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the TrackingPixel	to the ConversionEvent
			//----------------------------------------------------------------------------
			parentObj.TrackingPixel = &childObj

			//----------------------------------------------------------------------------
			// save the ConversionEvent
			//----------------------------------------------------------------------------
			return UpdateConversionEvent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrackingPixel", trackingPixelId )
			return utils.RequestResult{false, msg, "assignTrackingPixel", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TrackingPixel on a ConversionEvent
//----------------------------------------------------------------------------
func UnassignTrackingPixelFromConversionEvent(conversionEventId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ConversionEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConversionEvent(conversionEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConversionEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConversionEvent)

		//----------------------------------------------------------------------------
		// assign an empty TrackingPixel to the TrackingPixel
		//----------------------------------------------------------------------------
		parentObj.TrackingPixel = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TrackingPixel
		//----------------------------------------------------------------------------
		parentObj.TrackingPixelId = nil;

		//----------------------------------------------------------------------------
		// save the ConversionEvent
		//----------------------------------------------------------------------------
		return UpdateConversionEvent(parentObj)

	} else {
		return parentRequestResult
	}

}


